package infrastructure

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

/************************************************************************************************
**************************** Producer Section implements ****************************************
*************************************************************************************************/

type SendMessageRequest struct {
	Message   string
	Topic     string
	Partition int32
	Key       string
	Headers   map[string]string
}

func (req *SendMessageRequest) mapToKafkaMessage() *kafka.Message {
	msg := &kafka.Message{
		TopicPartition: *req.mapToTopicPartition(),
		Key:            []byte(req.Key),
		Value:          []byte(req.Message),
		Headers:        req.mapToHeaders(),
	}
	return msg
}

func (req *SendMessageRequest) mapToTopicPartition() *kafka.TopicPartition {
	part := kafka.PartitionAny
	if req.Partition > 0 {
		part = req.Partition
	}
	return &kafka.TopicPartition{Topic: &req.Topic, Partition: part}
}

func (req *SendMessageRequest) mapToHeaders() []kafka.Header {
	headers := make([]kafka.Header, 0, len(req.Headers))
	for key, value := range req.Headers {
		headers = append(headers, kafka.Header{Key: key, Value: []byte(value)})
	}
	return headers
}

type IPublisherKafka interface {
	SendMessage(request SendMessageRequest) error
}

type publisherKafka struct {
	producer *kafka.Producer
}

func NewPublisherKafka(connection string) (*publisherKafka, error) {
	producer, err := createProducer(connection)
	if err != nil {
		return nil, err
	}
	return &publisherKafka{producer: producer}, nil
}

func createProducer(connection string) (*kafka.Producer, error) {
	config := &kafka.ConfigMap{"bootstrap.servers": connection}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func (publisher *publisherKafka) SendMessage(request SendMessageRequest) error {
	delivery := make(chan kafka.Event)
	err := publisher.producer.Produce(request.mapToKafkaMessage(), delivery)
	if err != nil {
		return err
	}
	e := <-delivery
	msg := e.(*kafka.Message)
	return msg.TopicPartition.Error
}

/************************************************************************************************
**************************** Consumer Section implements ****************************************
*************************************************************************************************/

type IConsumerKafka interface {
	Consume(sigchan chan os.Signal, handler func(msg *ReceivedMessage)) error
}

type consumerKafka struct {
	consumer    *kafka.Consumer
	poolTimeout int
}

type ConsumerConfigs struct {
	BootstrapServers string
	GroupId          string
	Topic            string
	SectionTimeoutMs int
	PoolTimeoutMs    int
}

type ReceivedMessage struct {
	Key     string
	Value   string
	Headers map[string]string
}

func mapToReceiveMessage(msg *kafka.Message) *ReceivedMessage {
	headers := make(map[string]string, len(msg.Headers))
	for _, header := range msg.Headers {
		headers[header.Key] = string(header.Value)
	}
	return &ReceivedMessage{Key: string(msg.Key), Value: string(msg.Value), Headers: headers}
}

func (c *ConsumerConfigs) mapToKafkaConfig() *kafka.ConfigMap {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  c.BootstrapServers,
		"group.id":           c.GroupId,
		"session.timeout.ms": c.SectionTimeoutMs,
		"auto.offset.reset":  "earliest",
	}
	return config
}

func NewConsumerKafka(configs ConsumerConfigs) (*consumerKafka, error) {
	consumer, err := kafka.NewConsumer(configs.mapToKafkaConfig())
	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe(configs.Topic, nil)
	if err != nil {
		return nil, err
	}
	return &consumerKafka{consumer: consumer, poolTimeout: configs.PoolTimeoutMs}, nil
}

func (c *consumerKafka) Consume(sigchan chan os.Signal, handler func(msg *ReceivedMessage)) error {
	done := make(chan bool)
	var result error = nil

	go func() {
		for {
			select {
			case sig := <-sigchan:
				log.Printf("Caught signal %v: terminating consumer\n", sig)
				c.consumer.Close()
				done <- true
				return
			default:
				ev := c.consumer.Poll(c.poolTimeout)
				if ev == nil {
					continue
				}
				switch e := ev.(type) {
				case *kafka.Message:
					handler(mapToReceiveMessage(e))
				case kafka.Error:
					msg := fmt.Sprintf("Error: %v: %v", e.Code(), e)
					log.Println(msg)
					if e.Code() == kafka.ErrNoError {
						result = errors.New(msg)
						c.consumer.Close()
						done <- true
						return
					}
				default:
					log.Printf("Ignored %v\n", e)
				}
			}
		}
	}()

	<-done
	return result
}
