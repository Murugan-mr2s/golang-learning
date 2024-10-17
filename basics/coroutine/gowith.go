package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Stopped %s: %v\n", name, ctx.Err())
			return
		default:
			fmt.Printf("Doing work in %s...\n", name)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go doWork(ctx, "go with cancel")
	time.Sleep(3 * time.Second)
	cancel() // calling cancel manually close the goroutine

	ctx, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)
	go doWork(ctx, "go with timeout")
	time.Sleep(5 * time.Second) // enough  time to see timout
	cancelTimeout()

	deadline := time.Now().Add(2 * time.Second)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	go doWork(ctxDeadline, "go with Deadline")
	time.Sleep(5 * time.Second) // Wait for deadline to expire
	cancelDeadline()

}
