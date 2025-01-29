package services

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/vantutran2k1-commits-collector/producer/app/payloads"
	"github.com/vantutran2k1-commits-collector/producer/config"
)

type CommitService interface {
	Send(payload payloads.CommitPayload) error
}

func NewCommitService(kafkaProducer sarama.SyncProducer) CommitService {
	return &commitService{kafkaProducer: kafkaProducer}
}

type commitService struct {
	kafkaProducer sarama.SyncProducer
}

func (s *commitService) Send(payload payloads.CommitPayload) error {
	messageBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, _, err = s.kafkaProducer.SendMessage(&sarama.ProducerMessage{
		Topic: config.AppEnv.KafkaGithubCommitsTopic,
		Value: sarama.ByteEncoder(messageBytes),
	})
	if err != nil {
		return err
	}

	return nil
}
