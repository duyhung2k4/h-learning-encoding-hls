package queue

func InitQueue() {
	queueMp4Quantity := NewQueueMp4Quantity()
	go func() {
		queueMp4Quantity.Worker()
	}()
}
