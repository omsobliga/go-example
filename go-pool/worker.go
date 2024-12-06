package go_pool

import (
	"time"
)

type Worker struct {
	// pool who owns this worker
	pool *Pool

	// task is a job should be done
	task chan f

	// recycleTime will be updated when putting a worker back into queue
	recycleTime time.Time
}

func NewWorker(pool *Pool) *Worker {
	return &Worker{
		pool: pool,
		task: make(chan f, 1),
	}
}

func (w *Worker) run() {
	go func() {
		//循环监听任务，一旦有任务立即取出执行
		for f := range w.task {
			if f == nil {
				w.pool.decrRunning()
				return
			}
			f()
			w.pool.putWorker(w)
		}
	}()
}
