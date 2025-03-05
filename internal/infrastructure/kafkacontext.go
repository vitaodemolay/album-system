package infrastructure

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

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
