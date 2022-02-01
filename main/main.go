package main

import (
	"RabbitMQ/simple"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go simple.Produce(&wg)
	go simple.Consume(&wg)
	wg.Wait()
}
