package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type LockFreeSpinLock struct {
	state int32
}

func (l *LockFreeSpinLock) Lock() {
	for !atomic.CompareAndSwapInt32(&l.state, 0, 1) {
		// Spin until state flips from 0 to 1
	}
}

func (l *LockFreeSpinLock) Unlock() {
	atomic.StoreInt32(&l.state, 0)
}

func main() {
	var spinlock LockFreeSpinLock
	var wg sync.WaitGroup
	counter := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			spinlock.Lock()
			counter++
			spinlock.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Lock-free updated counter:", counter)
}