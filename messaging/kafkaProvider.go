package messaging

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/swagchat/rtm-api/logging"
	"github.com/swagchat/rtm-api/services"
	"github.com/swagchat/rtm-api/utils"
	"go.uber.org/zap/zapcore"
)

var KafkaConsumer *kafka.Consumer

type KafkaProvider struct{}

func (provider *KafkaProvider) Subscribe() {
	cfg := utils.Config()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	var hostname string
	hostname, err := os.Hostname()
	if err != nil {
		hostname = utils.CreateUuid()
	}
	KafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": fmt.Sprintf("%s:%s", cfg.Messaging.Kafka.Host, cfg.Messaging.Kafka.Port),
		"group.id":          hostname,
		// "session.timeout.ms":   6000,
		// "default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"}
	})

	if err != nil {
		logging.Log(zapcore.ErrorLevel, &logging.AppLog{
			Kind:     "messaging-subscribe",
			Provider: "kafka",
			Message:  err.Error(),
		})
	} else {
		logging.Log(zapcore.InfoLevel, &logging.AppLog{
			Kind:     "messaging-subscribe",
			Provider: "kafka",
			Message:  KafkaConsumer.String(),
		})
	}

	err = KafkaConsumer.SubscribeTopics([]string{cfg.Messaging.Kafka.Topic}, nil)
	if err != nil {
		logging.Log(zapcore.ErrorLevel, &logging.AppLog{
			Kind:     "messaging-subscribe",
			Provider: "kafka",
			Message:  err.Error(),
		})
	}

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			run = false
			logging.Log(zapcore.InfoLevel, &logging.AppLog{
				Kind:     "messaging-subscribe-terminated",
				Provider: "kafka",
				Message:  fmt.Sprintf("terminated by %s", sig.String()),
			})
		default:
			ev := KafkaConsumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				services.Srv.Broadcast <- e.Value
				logging.Log(zapcore.InfoLevel, &logging.AppLog{
					Kind:     "messaging-subscribe-receive",
					Provider: "kafka",
					Message:  fmt.Sprintf("Receive a message [%s]", string(e.Value)),
				})
			case kafka.PartitionEOF:
				logging.Log(zapcore.InfoLevel, &logging.AppLog{
					Kind:     "messaging-subscribe",
					Provider: "kafka",
					Message:  e.String(),
				})
			case kafka.Error:
				run = false
				logging.Log(zapcore.ErrorLevel, &logging.AppLog{
					Kind:     "messaging-subscribe",
					Provider: "kafka",
					Message:  e.String(),
				})
			default:
				logging.Log(zapcore.ErrorLevel, &logging.AppLog{
					Kind:     "messaging-subscribe",
					Provider: "kafka",
					Message:  e.String(),
				})
			}
		}
	}

	KafkaConsumer.Close()
	logging.Log(zapcore.InfoLevel, &logging.AppLog{
		Kind:     "messaging-subscribe",
		Provider: "kafka",
		Message:  "close",
	})
}

func (provider *KafkaProvider) Unsubscribe() {
	logging.Log(zapcore.InfoLevel, &logging.AppLog{
		Kind:     "messaging-unsubscribe",
		Provider: "kafka",
	})
	KafkaConsumer.Unsubscribe()
}