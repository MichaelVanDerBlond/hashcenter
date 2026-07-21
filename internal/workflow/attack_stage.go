package workflow

import (
	"context"
	"errors"

	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
)

type AttackStage struct{}

func NewAttackStage() *AttackStage {
	return &AttackStage{}
}

func (a *AttackStage) Name() string {
	return "attack"
}

func (a *AttackStage) ShouldRun(wf *Context) bool {
	if wf == nil {
		return false
	}

	if wf.Analyze == nil {
		return false
	}

	return wf.Analyze.HashcatReady
}

func (a *AttackStage) Run(ctx context.Context, wf *Context) error {

	if wf == nil {
		return errors.New("workflow context is nil")
	}

	job := jobs.Job{
		HashFile:    wf.HashFile,
		Dictionary:  wf.Dictionary,
		AttackMode:  wf.AttackMode,
		HashMode:    wf.HashMode,
		Rule:        wf.Rule,
		Mask:        wf.Mask,
		SessionName: wf.SessionName,
		Device:      wf.Device,
		Workload:    wf.Workload,
		Extra:       wf.Extra,
	}

	result, err := jobs.DefaultManager().Run(ctx, job)
	if err != nil {
		return err
	}

	wf.SessionID = result.SessionID
	wf.Result = result

	return nil
}
