package jobs

import (
	"context"

	"github.com/MichaelVanDerBlond/hashcenter/internal/database"
	"github.com/MichaelVanDerBlond/hashcenter/internal/hashcat"
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
	"github.com/MichaelVanDerBlond/hashcenter/internal/session"
)

func (m *Manager) Run(ctx context.Context, job Job) (*Result, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := repository.NewSQLiteSessionRepository(db)

	s := session.New()
	if job.SessionID != "" {
		s.ID = job.SessionID
	}

	s.Backend = "hashcat"
	s.CaptureFile = job.HashFile
	s.Dictionary = job.Dictionary
	s.AttackMode = job.AttackMode
	s.HashMode = job.HashMode
	s.Rule = job.Rule
	s.Mask = job.Mask
	s.Device = job.Device
	s.Workload = job.Workload
	s.SessionName = job.SessionName

	if err := repo.Save(s); err != nil {
		return nil, err
	}

	s.Start()

	if err := repo.Update(s); err != nil {
		return nil, err
	}

	runner := hashcat.New()

	if _, err := runner.Version(ctx); err != nil {
		s.Fail(err)
		_ = repo.Update(s)
		return nil, err
	}

	args := buildArgs(job)

	result, status, err := runner.Start(ctx, args...)
	if err != nil {
		s.Fail(err)
		_ = repo.Update(s)
		return nil, err
	}

	runtime := &Runtime{
		SessionID: s.ID,
		Runner:    runner,
		Status:    status,
		Result:    result,
	}

	m.Register(runtime)
	defer m.Remove(runtime.SessionID)

	result, err = runner.Wait(result)

	if err != nil {
		s.Fail(err)
	} else {
		s.Finish()
	}

	_ = repo.Update(s)

	return &Result{
		SessionID: s.ID,
		Result:    result,
		Status:    status,
	}, err
}
