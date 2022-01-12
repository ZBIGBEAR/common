package util

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

const (
	DefaultParallelCount int64 = 10
)

// define job
type Job func() error

// define parallel interface
type Parallel interface {
	Add(Job)
	Start()
	Result() []error
	Reset()
}

// implement interface
type parallel struct {
	jobs []Job   // 任务
	errs []error // 结果
	ctx  context.Context
	sema *semaphore.Weighted // 信号量，用于控制任务的并发
}

// new Parallel
func NewParallel(parallelCount int64) Parallel {
	return &parallel{
		ctx:  context.Background(),
		sema: semaphore.NewWeighted(parallelCount),
	}
}

// new Parallel
func NewDefaultParallel() Parallel {
	return NewParallel(DefaultParallelCount)
}

// add job
func (p *parallel) Add(job Job) {
	p.jobs = append(p.jobs, job)
}

// start jobs
func (p *parallel) Start() {
	jobCount := len(p.jobs)

	if jobCount == 0 {
		return
	}

	errs := make(chan error, jobCount)

	wg := sync.WaitGroup{}
	wg.Add(jobCount)

	for i := range p.jobs {
		if err := p.sema.Acquire(p.ctx, 1); err != nil {
			p.addErr(err)
			return
		} else {
			go func(job Job) {
				defer func() {
					wg.Done()
					p.sema.Release(1)
				}()
				errs <- job()
			}(p.jobs[i])
		}
	}

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			p.addErr(err)
		}
	}
}

func (p *parallel) addErr(err error) {
	p.errs = append(p.errs, err)
}

// get result
func (p *parallel) Result() []error {
	return p.errs
}

// reset
func (p *parallel) Reset() {
	p.jobs = make([]Job, 0)
	p.errs = make([]error, 0)
}
