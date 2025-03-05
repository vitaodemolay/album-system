package infrastructure

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

const (
	kafkaconnection = "localhost:9092"
	topic           = "test-topic"
)

func TestSendMessage(t *testing.T) {
	// Arrange
	publisher, err := NewPublisherKafka(kafkaconnection)
	assert.NoError(t, err)

	request := SendMessageRequest{
		Message: fmt.Sprintf("{'teste': 'value_test', 'price': %d}", rand.Int()),
		Topic:   "test-topic",
	}

	// Act
	err = publisher.SendMessage(request)

	// Assert
	assert.NoError(t, err)
}

func TestMapToKafkaMessage(t *testing.T) {
	// Arrange
	request := SendMessageRequest{
		Message:   "Test message",
		Topic:     "test-topic",
		Partition: 1,
		Key:       "test-key",
		Headers:   map[string]string{"header1": "value1"},
	}

	// Act
	result := request.mapToKafkaMessage()

	// Assert
	assert.Equal(t, request.Topic, *result.TopicPartition.Topic)
	assert.Equal(t, request.Partition, result.TopicPartition.Partition)
	assert.Equal(t, []byte(request.Key), result.Key)
	assert.Equal(t, []byte(request.Message), result.Value)
	assert.Len(t, result.Headers, 1)
	assert.Equal(t, "header1", result.Headers[0].Key)
	assert.Equal(t, []byte("value1"), result.Headers[0].Value)
}

func TestMapToTopicPartition(t *testing.T) {
	// Arrange
	request := SendMessageRequest{
		Topic:     "test-topic",
		Partition: 1,
	}

	// Act
	result := request.mapToTopicPartition()

	// Assert
	assert.Equal(t, request.Topic, *result.Topic)
	assert.Equal(t, request.Partition, result.Partition)
}

func TestMapToHeaders(t *testing.T) {
	// Arrange
	request := SendMessageRequest{
		Headers: map[string]string{
			"header1": "value1",
			"header2": "value2",
		},
	}

	// Act
	result := request.mapToHeaders()

	// Assert
	assert.Len(t, result, 2)
	assert.Contains(t, result, kafka.Header{Key: "header1", Value: []byte("value1")})
	assert.Contains(t, result, kafka.Header{Key: "header2", Value: []byte("value2")})
}
