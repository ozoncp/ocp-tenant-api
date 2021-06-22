package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ozoncp/ocp-tenant-api/internal/metrics"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	Send(msgType MessageType, id uint64, timestamp time.Time) error
}

type producer struct {
	dataProducer sarama.SyncProducer
	topic        string
	messageChan  chan *sarama.ProducerMessage
}

type MessageType uint

const (
	Create MessageType = iota
	Update
	Remove
)

type message struct {
	msgType MessageType
	Body    map[string]interface{}
}

func New(ctx context.Context, addresses []string, topic string, capacity uint64) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	syncProducer, err := sarama.NewSyncProducer(addresses, config)

	if err != nil {
		log.Error().Err(err).Msg("failed to create a producer")
		return nil, err
	}

	newProducer := &producer{
		dataProducer: syncProducer,
		topic:        topic,
		messageChan:  make(chan *sarama.ProducerMessage, capacity),
	}

	go newProducer.handleMessage(ctx)

	return newProducer, nil
}

func (dProducer *producer) handleMessage(ctx context.Context) {
	for {
		select {
		case msg := <-dProducer.messageChan:
			_, _, err := dProducer.dataProducer.SendMessage(msg)

			if err != nil {
				log.Error().Msgf("failed to send message to kafka: %v", err)
				log.Error().Msgf("retry ...")

				dProducer.messageChan <- msg
			}
		case <-ctx.Done():
			close(dProducer.messageChan)
			dProducer.dataProducer.Close()
			return
		}
	}
}

func (dProducer *producer) Send(msgType MessageType, id uint64, timestamp time.Time) error {
	message := createMessage(msgType, id, timestamp)
	dataBytes, err := json.Marshal(message)

	if err != nil {
		log.Error().Err(err).Msg("failed to marshal message to json")
		return err
	}

	dProducer.messageChan <- &sarama.ProducerMessage{
		Topic:     dProducer.topic,
		Key:       sarama.StringEncoder(dProducer.topic),
		Value:     sarama.StringEncoder(dataBytes),
		Partition: -1,
		Timestamp: time.Time{},
	}
	addMetrics(msgType)
	return nil
}

func addMetrics(msgType MessageType) {
	strType := convertMessageTypeToString(msgType)
	switch msgType {
	case Create:
		metrics.CreateCounterInc(strType)
	case Update:
		metrics.UpdateCounterInc(strType)
	case Remove:
		metrics.RemoveCounterInc(strType)
	}
}

func createMessage(msgType MessageType, id uint64, timestamp time.Time) message {
	return message{
		msgType: msgType,
		Body: map[string]interface{}{
			"Id":        id,
			"Operation": fmt.Sprintf("%s tenant", convertMessageTypeToString(msgType)),
			"Timestamp": timestamp,
		},
	}
}

func convertMessageTypeToString(msgType MessageType) string {
	switch msgType {
	case Create:
		return "Created"
	case Update:
		return "Updated"
	case Remove:
		return "Removed"
	}

	return "unknown message type"
}
