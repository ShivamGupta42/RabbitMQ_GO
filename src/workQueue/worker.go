package workQueue

import (
	"bytes"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

func Consume(wg *sync.WaitGroup) {
	defer wg.Done()
	/*server heartbeat interval of 10
	seconds and sets the handshake deadline to 30 seconds
	*/
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//Multiple channels can be opened per connection
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count - channel
		0,     // prefetch size - for consumer no limit
		false, // global (set prefetch on a connection level)
	)
	failOnError(err, "Failed to set QoS")

	//Consumer decl on one channel of one connection
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	failOnError(err, "Failed to register a consumer")

	var localWg sync.WaitGroup
	localWg.Add(1)

	go func() {
		defer localWg.Done()
		for i := 0; i < 5; i++ {
			d := <-msgs
			log.Printf("Received a message: %s", d.Body)
			dotCount := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dotCount)
			time.Sleep(t * time.Second)
			d.Ack(false) //For batch processing, call this method with multiple true
		}

	}()

	localWg.Wait()
	log.Printf("Exiting consumer")
}
