package sbroker

import (
	"context"

	"github.com/swagchat/rtm-api/config"
)

type provider interface {
	SubscribeMessage() error
	UnsubscribeMessage() error
}

func Provider(ctx context.Context) provider {
	cfg := config.Config()

	var p provider
	switch cfg.SBroker.Provider {
	case "":
		p = &notuseProvider{
			ctx: ctx,
		}
	case "nsq":
		p = &nsqProvider{
			ctx: ctx,
		}
	case "kafka":
		p = &kafkaProvider{
			ctx: ctx,
		}
	}

	return p
}
