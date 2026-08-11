package web

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
	"github.com/MichaelVanDerBlond/hashcenter/internal/task"
)

var queueOnce sync.Once

func startQueue() {
	queueOnce.Do(func() {
		go taskWorker()
	})
}

func taskWorker() {
	manager := task.DefaultManager()

	for {
		t, ok := manager.NextQueued()
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

		err := attack.Execute(context.Background(), job)

		if err != nil {
			manager.SetError(t.ID, err.Error())
			log.Printf("task %s failed: %v", t.ID, err)
			continue
		}

		manager.SetState(t.ID, task.Finished)
	}
}

func enqueue(job attack.Job) {
	startQueue()

	t := task.NewAttack("", job)

	created, err := task.DefaultManager().Create(t)
	if err != nil {
		log.Printf("failed to create task: %v", err)
		return
	}

	if created == nil {
		log.Printf("failed to create task")
		return
	}

	log.Printf("task queued: %s", created.ID)
}
