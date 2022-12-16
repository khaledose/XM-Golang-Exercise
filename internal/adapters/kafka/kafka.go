package kafka

import (
	"encoding/json"
	"time"
	"xm/companies/config"
	"xm/companies/internal/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type KafkaBroker struct {
	Producer *kafka.Producer
	log      *zap.SugaredLogger
}

func NewKafkaBroker(logger *zap.SugaredLogger, conf config.KafkaConfigurations) *KafkaBroker {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": conf.Servers,
		"security.protocol": conf.SecurityProtocol,
		"sasl.mechanisms":   conf.SaslMechanism,
		"sasl.username":     conf.User,
		"sasl.password":     conf.Pass,
		"client.id":         conf.ClientName,
		"acks":              "all"})

	if err != nil {
		panic(err)
	}

	return &KafkaBroker{
		Producer: p,
		log:      logger,
	}
}

func (b *KafkaBroker) Produce(topic string, data interface{}, messageStatus string) error {
	delivery_chan := make(chan kafka.Event, 10000)

	event := models.KafkaEvent{
		Id:        uuid.New().String(),
		Status:    messageStatus,
		Source:    models.AppName,
		TimeStamp: time.Now().UTC(),
		Data:      data,
	}

	b.log.Infof("Producing event: %s to topic: %s", event, topic)

	out, err := json.MarshalIndent(event, "", "  ")
	if err != nil {
		b.log.Errorf("Error while marshalling event: %s into json", event)
		return err
	}

	err = b.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: 0},
		Value:          []byte(out)},
		delivery_chan,
	)
	if err != nil {
		b.log.Errorf("Error while producing event: %s with topic: %s", event, topic)
		return err
	}

	b.log.Info("Event produced")
	return nil
}
