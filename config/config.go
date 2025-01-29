package config

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/vantutran2k1-commits-collector/producer/app/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB
var KafkaProducerClient sarama.SyncProducer

func InitDb() {
	l := logger.Default.LogMode(logger.Silent)
	if AppEnv.GinMode == constants.GinDebugMode {
		l = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		AppEnv.DbHost,
		AppEnv.DbUser,
		AppEnv.DbPass,
		AppEnv.DbName,
		AppEnv.DbPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	Db = db
}

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
