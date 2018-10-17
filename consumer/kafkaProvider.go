package consumer

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/betchi/tracer"
	logger "github.com/betchi/zapper"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/swagchat/rtm-api/config"
	"github.com/swagchat/rtm-api/rtm"
)

var client *kafka.Consumer

type kafkaProvider struct {
	ctx context.Context
}

func (kp *kafkaProvider) SubscribeMessage() error {
	span := tracer.StartSpan(kp.ctx, "SubscribeMessage", "consumer")
	defer tracer.Finish(span)

	cfg := config.Config()

	host := cfg.Consumer.Kafka.Host
	if host == "" {
		err := errors.New("consumer.kafka.host is empty")
		logger.Error(err.Error())
		return err
	}

	port := cfg.Consumer.Kafka.Port
	if port == "" {
		err := errors.New("consumer.kafka.port is empty")
		logger.Error(err.Error())
		return err
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = uuid.NewV4().String()
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	client, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", host, port),
		"group.id":          hostname,
		// "session.timeout.ms":   6000,
		// "default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"}
	})
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	topic := cfg.Consumer.Kafka.Topic
	err = client.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info(fmt.Sprintf("%s group.id[%s] topic[%s]", client.String(), hostname, topic))

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			run = false
			logger.Info(fmt.Sprintf("terminated by %s", sig.String()))
		default:
			ev := client.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				logger.Info("Receive a message")
				rtm.Server().Broadcast <- e.Value
			case kafka.PartitionEOF:
				logger.Info(e.String())
			case kafka.Error:
				run = false
				logger.Error(e.String())
			default:
				logger.Info(e.String())
			}
		}
	}

	client.Close()
	logger.Info("kafka close")

	return nil
}

func (kp *kafkaProvider) UnsubscribeMessage() error {
	span := tracer.StartSpan(kp.ctx, "UnsubscribeMessage", "consumer")
	defer tracer.Finish(span)

	if client == nil {
		return nil
	}

	logger.Info("kafka unsubscribe")
	return client.Unsubscribe()
}
