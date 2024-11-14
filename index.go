package main

import (
	"app/queue"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		queue.InitQueue()
	}()

	wg.Wait()
}
