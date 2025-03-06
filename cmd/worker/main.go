package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/vitaodemolay/album-system/internal/infrastructure"
)

const (
	kafkaConnection  = "localhost:9092"
	topicEvents      = "albums-public-events"
	groupId          = "test-group"
	sectionTimeoutMs = 6000
	poolTimeout      = 100
)

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	var c infrastructure.IConsumerKafka
	var err error

	log.Print("Worker is starting...")

	configs := infrastructure.ConsumerConfigs{
		BootstrapServers: kafkaConnection,
		GroupId:          groupId,
		Topic:            topicEvents,
		SectionTimeoutMs: sectionTimeoutMs,
		PoolTimeoutMs:    poolTimeout,
	}

	log.Print("Creating kafka connection")
	c, err = infrastructure.NewConsumerKafka(configs)
	if err != nil {
		log.Fatal(err)
	}

	handler := func(msg *infrastructure.ReceivedMessage) {
		// Process message
		log.Printf("Received message: %+v\n", msg)
	}

	log.Print("Starting kafka consumer")
	err = c.Consume(sigchan, handler)

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Worker finished")
}
