package main

import "RabbitMQ/pubsub"

func main() {
	/*
		var wg sync.WaitGroup

		wg.Add(2)
		go simple.Produce(&wg)
		go simple.Consume(&wg)
		wg.Wait()
	*/

	go pubsub.Publish()

}
