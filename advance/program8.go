package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type PipelineGroup struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
	errOnce sync.Once
	err    error
}

func NewPipelineGroup(ctx context.Context) *PipelineGroup {
	cCtx, cancel := context.WithCancel(ctx)
	return &PipelineGroup{ctx: cCtx, cancel: cancel}
}

func (g *PipelineGroup) Go(fn func(ctx context.Context) error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		if err := fn(g.ctx); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				g.cancel()
			})
		}
	}()
}

func (g *PipelineGroup) Wait() error {
	g.wg.Wait()
	g.cancel()
	return g.err
}

func main() {
	g := NewPipelineGroup(context.Background())

	g.Go(func(ctx context.Context) error {
		fmt.Println("Task 1 completed")
		return nil
	})

	g.Go(func(ctx context.Context) error {
		return errors.New("Task 2 failed with critical error")
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Pipeline halted due to failure:", err)
	}
}