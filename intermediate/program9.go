package main

import (
	"context"
	"fmt"
	"time"
)

func longTask(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Task finished completely")
	case <-ctx.Done():
		fmt.Println("Task aborted:", ctx.Err())
	}
}

func main() {
	
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go longTask(ctx)

	time.Sleep(2 * time.Second)
}