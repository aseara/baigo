package workerpool

import (
	"errors"
	"fmt"
	"sync"
)

const (
	_defaultCapacity = 100
	_maxCapacity     = 1000
)

var (
	// ErrWorkerPoolFreed workerpool已终止运行
	ErrWorkerPoolFreed = errors.New("workpool freed")
	// ErrNoIdleWorkerInPool workerpool中任务已满，没有空闲goroutine用于处理新任务
	ErrNoIdleWorkerInPool = errors.New("no idle worker in pool")
)

// Task task 实例
type Task func()

// Pool worker pool 实例
type Pool struct {
	preAlloc bool
	block    bool

	capacity int

	active chan struct{}
	tasks  chan Task

	wg   sync.WaitGroup
	quit chan struct{}
}

// New 创建 Pool
func New(capacity int, opts ...Option) *Pool {
	if capacity <= 0 {
		capacity = _defaultCapacity
	}
	if capacity > _maxCapacity {
		capacity = _maxCapacity
	}
	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}

	for _, opt := range opts {
		opt(p)
	}
	fmt.Printf("workerpool start(preAlloc=%t)\n", p.preAlloc)

	if p.preAlloc {
		for i := 0; i < p.capacity; i++ {
			p.newWorker(i + 1)
			p.active <- struct{}{}
		}
	}

	go p.run()
	return p
}

// Free 销毁 Pool
func (p *Pool) Free() {
	close(p.quit)
	p.wg.Wait()
	fmt.Printf("workerpool freed(preAlloc=%t)\n", p.preAlloc)
}

// Schedule 调度
func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	default:
		if p.block {
			p.tasks <- t
			return nil
		}
		return ErrNoIdleWorkerInPool
	}
}

// run workerpool 处理task
func (p *Pool) run() {
	idx := 0

	if !p.preAlloc {
	loop:

		for t := range p.tasks {
			p.returnTAsk(t)
			select {
			case <-p.quit:
				return
			case p.active <- struct{}{}:
				idx++
				p.newWorker(idx)
			default:
				break loop
			}
		}
	}

	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			// create a new worker
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) returnTAsk(t Task) {
	go func() {
		p.tasks <- t
	}()
}

func (p *Pool) newWorker(idx int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", idx, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", idx)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: exit\n", idx)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", idx)
				t()
			}
		}

	}()
}
