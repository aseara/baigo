package workerpool

import (
	"errors"
	"fmt"
	"sync"
)

var ErrWorkerPoolFreed = errors.New("workpool freed")

const _defaultCapacity = 12
const _maxCapacity = 1000

// Pool worker pool 实例
type Pool struct {
	capacity int

	active chan struct{}
	tasks  chan Task

	wg   sync.WaitGroup
	quit chan struct{}
}

// New 创建 Pool
func New(capacity int) *Pool {
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
	fmt.Printf("workerpool start\n")

	go p.run()

	return p
}

// Free 销毁 Pool
func Free(pool *Pool) {

}

// Schedule 调度
func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil

	}
}

// run workerpool 处理task
func (p *Pool) run() {
	idx := 0

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
