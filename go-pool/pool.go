package go_pool

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type sig struct{}

type f func()

type Pool struct {
	// capacity of the pool
	capacity int32

	// running is the number of the currently running goroutines
	running int32

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
		capacity:       int32(capacity),
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

// Submit a task to pool
func (p *Pool) Submit(task f) error {
	if len(p.release) > 0 {
		return errors.New("pool is closed")
	}

	w := p.getWorker()
	w.task <- task
	return nil
}

func (p *Pool) Running() int {
	return int(atomic.LoadInt32(&p.running))
}

func (p *Pool) incrRunning() {
	atomic.AddInt32(&p.running, 1)
}

func (p *Pool) decrRunning() {
	atomic.AddInt32(&p.running, -1)
}

func (p *Pool) Cap() int {
	return int(atomic.LoadInt32(&p.capacity))
}

// getWorker returns an available worker to run the tasks
func (p *Pool) getWorker() *Worker {
	var w *Worker
	waiting := false

	p.lock.Lock()
	// 有可用 worker，按 LIFO 方式取出
	if len(p.workers) > 0 {
		n := len(p.workers)
		w = p.workers[n-1]
		p.workers[n-1] = nil
		p.workers = p.workers[:n-1]
	} else {
		waiting = p.Running() >= p.Cap()
	}
	// 有可用 worker，按 LIFO 方式取出
	p.lock.Unlock()
	if w != nil {
		return w
	}

	if waiting {
		for {
			p.lock.Lock()
			if len(p.workers) == 0 {
				p.lock.Unlock()
				continue
			}
			n := len(p.workers)
			w = p.workers[n-1]
			p.workers[n-1] = nil
			p.workers = p.workers[:n-1]
			p.lock.Unlock()
			break
		}
	} else {
		w = NewWorker(p)
		w.run()
		p.incrRunning()
	}
	return w
}

// putWorker puts a worker back into free pool, recycling the goroutine
func (p *Pool) putWorker(w *Worker) {
	// 记录 worker 的回收时间
	w.recycleTime = time.Now()
	p.lock.Lock()
	p.workers = append(p.workers, w)
	p.lock.Unlock()
}

// resize change the capacity of the pool
func (p *Pool) resize(size int) {
	if size == p.Cap() {
		return
	}
	atomic.StoreInt32(&p.capacity, int32(size))
	diff := p.Running() - size
	if diff > 0 {
		for i := 0; i < diff; i++ {
			p.getWorker().task <- nil
		}
	}
}

// clear expired workers periodically
func (p *Pool) periodicallyPurge() {
	ticker := time.NewTicker(p.expiryDuration)
	currentTime := time.Now()
	for range ticker.C {
		p.lock.Lock()
		if len(p.workers) == 0 && p.Running() == 0 && len(p.release) == 0 {
			p.lock.Unlock()
			return
		}
	}
	n := 0
	for i, w := range p.workers {
		if currentTime.Sub(w.recycleTime) <= p.expiryDuration {
			break
		}
		n = i
		w.task <- nil
		p.workers[i] = nil
	}
	n++
	// 因为采用了 LIFO 后进先出队列存放空闲 worker，所以该队列默认已经是按照 worker 的最后运行时间由远及近排序
	if n >= len(p.workers) {
		p.workers = p.workers[:0]
	} else {
		p.workers = p.workers[n:]
	}
	p.lock.Unlock()
}
