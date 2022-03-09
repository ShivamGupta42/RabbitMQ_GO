package main

import (
	"RabbitMQ/src/workQueue"
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	//Simple
	//go simple.Produce(&wg)
	//go simple.Consume(&wg)

	//Work Queue
	go workQueue.Produce(&wg)
	go workQueue.Consume(&wg)
	wg.Wait()
	fmt.Println("Execution completed")

}
