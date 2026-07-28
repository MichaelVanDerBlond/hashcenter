package web

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/attack"
)

type queuedAttack struct {
	meta *QueueJob
	job  attack.Job
}

var (
	queue chan queuedAttack
	once  sync.Once
)

func startQueue() {
	once.Do(func() {
		queue = make(chan queuedAttack, 128)

		go func() {
			for item := range queue {
				queueMu.Lock()
				item.meta.State = "running"
				item.meta.Started = time.Now()
				queueMu.Unlock()

				err := attack.Execute(context.Background(), item.job)

				queueMu.Lock()
				item.meta.Finished = time.Now()

				if err != nil {
					item.meta.State = "failed"
					queueMu.Unlock()

					log.Println(err)
					continue
				}

				item.meta.State = "finished"
				queueMu.Unlock()
			}
		}()
	})
}

func enqueue(job attack.Job) {
	startQueue()

	meta := addQueueJob(
		job.HashFile,
		job.Dictionary,
	)

	queue <- queuedAttack{
		meta: meta,
		job:  job,
	}
}
