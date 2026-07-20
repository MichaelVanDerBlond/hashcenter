package jobs

import (
	"context"
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/database"
	"github.com/MichaelVanDerBlond/hashcenter/internal/hashcat"
	"github.com/MichaelVanDerBlond/hashcenter/internal/repository"
	"github.com/MichaelVanDerBlond/hashcenter/internal/session"
)

func (m *Manager) Run(ctx context.Context, job Job) (*Result, error) {

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	repo := repository.NewSQLiteSessionRepository(db)

	s := session.New()
	s.Backend = "hashcat"
	s.CaptureFile = job.HashFile

	if err := repo.Save(s); err != nil {
		return nil, err
	}

	s.Start()

	if err := repo.Update(s); err != nil {
		return nil, err
	}

	runner := hashcat.New()

	version, err := runner.Version(ctx)
	if err != nil {
		s.Fail(err)
		_ = repo.Update(s)
		return nil, err
	}

	args := []string{
		"-m", fmt.Sprint(job.HashMode),
		"-a", fmt.Sprint(job.AttackMode),
		job.HashFile,
		job.Dictionary,
	}

	args = append(args, job.Extra...)

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
	defer m.Remove(s.ID)

	result, err = runner.Wait(result)

	if err != nil {
		s.Fail(err)
	} else {
		s.Finish()
	}

	_ = repo.Update(s)

	_ = version

	return &Result{
		SessionID: s.ID,
		Result:    result,
		Status:    status,
	}, err
}
