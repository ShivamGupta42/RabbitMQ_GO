package main

import (
	"RabbitMQ/consumers"
	"RabbitMQ/producer"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go producer.Produce(&wg)
	go consumers.Consume(&wg)
	wg.Wait()
}
