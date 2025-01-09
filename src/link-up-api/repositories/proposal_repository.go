package repositories

import (
	"log"

	"github.com/IBM/sarama"
)

type ProposalRepository struct {
	Producer sarama.SyncProducer
	Topic    string
}

func NewProposalRepository(brokers []string, topic string) (*ProposalRepository, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Printf("Failed to create Sarama producer: %v", err)
		return nil, err
	}

	return &ProposalRepository{
		Producer: producer,
		Topic:    topic,
	}, nil
}

func (r *ProposalRepository) PublishProposal(message []byte) error {
	defer r.Producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: r.Topic,
		Value: sarama.ByteEncoder(message),
	}

	partition, offset, err := r.Producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)", r.Topic, partition, offset)
	return nil
}
