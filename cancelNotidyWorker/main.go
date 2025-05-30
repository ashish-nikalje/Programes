// Problem Statement: Notify the other goroutines that his task is completed , the other goroutine should stop execution as got the notification.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, _ := context.WithCancel(context.Background())
	cancelChan := make(chan struct{}) // Initialize the channel

	wg.Add(2)

	go startWorker(ctx, 0, &wg, func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("---> ", i)

			if i == 9 {
				fmt.Println("cancelling channel ---> ", 0)
				cancelChan <- struct{}{}
			}
		}
	})

	go startWorker(ctx, 1, &wg, func() {
		for i := 0; i < 100; i++ {
			select {
			case <-cancelChan:
				fmt.Println("Received cancel signal from 0. Exiting... ", 1)
				return
			default:
				fmt.Println("printing id --> ", i)
				time.Sleep(100 * time.Millisecond)
			}
		}
	})

	wg.Wait()
}

func startWorker(ctx context.Context, id int, wg *sync.WaitGroup, fn func()) {
	defer wg.Done()
	fn()
}
