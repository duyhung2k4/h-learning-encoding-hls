package queue

import (
	"sync"
)

func InitQueue() {
	queueMp4Quantity := NewQueueMp4Quantity()
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		queueMp4Quantity.Worker()
	}()

	wg.Wait()
}
