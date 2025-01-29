package config

import (
	"github.com/IBM/sarama"
	"log"
)

var KafkaProducerClient sarama.SyncProducer

func InitKafkaProducerClient() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	brokers := []string{AppEnv.KafkaBroker}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}

	KafkaProducerClient = producer
}
