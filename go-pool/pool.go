package go_pool

import (
	"errors"
	"sync"
	"time"
)

type sig struct{}

type f func()

type Pool struct {
	// capacity of the pool
	capacity int

	// running is the number of the currently running goroutines
	running int

	// expiryDuration set the expired time(second) of every worker
	expiryDuration time.Duration

	// workers is a slice the store the available workers
	workers []*Worker

	// release is used to notice the pool to closed itself
	release chan sig

	// lock for synchronous operation
	lock sync.Mutex

	once sync.Once
}

const DefaultExpiryDuration = time.Second * time.Duration(30)

func NewPool(capacity int) (*Pool, error) {
	return NewTimingPool(capacity, DefaultExpiryDuration)
}

func NewTimingPool(capacity int, expiryDuration time.Duration) (*Pool, error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than zero")
	}
	if expiryDuration <= 0 {
		return nil, errors.New("expiryDuration must be greater than zero")
	}
	p := &Pool{
		capacity:       capacity,
		expiryDuration: expiryDuration,
		release:        make(chan sig, 1),
	}
	p.monitorAndClear()
	return p, nil
}

// clear expired workers regularly
func (p *Pool) monitorAndClear() {
	//
}

// submit a task to pool
func (p *Pool) submit(task f) error {

	if len(p.workers) >= p.capacity {
		return errors.New("pool is full")
	}
	taskCh := make(chan f, 1)
	taskCh <- task

	if p.running < len(p.workers) {
		// exist empty worker
		n := len(p.workers) - 1
		worker := p.workers[n]
		p.workers = p.workers[:n]
		worker.task = taskCh
		worker.recycleTime = time.Now()
	} else if p.running == len(p.workers) {
		// not exist empty worker
		worker := &Worker{
			pool:        p,
			task:        taskCh,
			recycleTime: time.Now(),
		}
		p.workers = append(p.workers, worker)
	}
	return nil
}

// getWorker returns a available worker to run the tasks
func (p *Pool) getWorker() *Worker {
	p.lock.Lock()
	defer p.lock.Unlock()
}
