package services

import (
	"context"
	"log"
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
	"github.com/MichaelVanDerBlond/hashcenter/internal/task"
)

type TaskWorker struct {
	manager *task.Manager
}

func NewTaskWorker(manager *task.Manager) *TaskWorker {
	if manager == nil {
		manager = task.DefaultManager()
	}

	return &TaskWorker{
		manager: manager,
	}
}

func (w *TaskWorker) Run(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		t, ok := w.manager.NextQueued()
		if !ok {
			time.Sleep(250 * time.Millisecond)
			continue
		}

		job := attack.Job{
			HashFile:    t.HashFile,
			Dictionary:  t.Dictionary,
			HashMode:    t.HashMode,
			AttackMode:  t.AttackMode,
			Rule:        t.Rule,
			Mask:        t.Mask,
			SessionName: t.SessionName,
			Device:      t.Device,
			Workload:    t.Workload,
		}

		if err := attack.Execute(ctx, job); err != nil {
			w.manager.SetError(t.ID, err.Error())
			log.Printf("task %s failed: %v", t.ID, err)
			continue
		}

		w.manager.SetState(t.ID, task.Finished)
	}
}
