package app

import (
	"context"
	"log"

	tgClient "github.com/drizzleent/ton-bot/internal/clients/telegram"
	"github.com/drizzleent/ton-bot/internal/config"
	"github.com/drizzleent/ton-bot/internal/config/env"
	"github.com/drizzleent/ton-bot/internal/consumer"
	eventconsumer "github.com/drizzleent/ton-bot/internal/consumer/event-consumer"
	"github.com/drizzleent/ton-bot/internal/events"
	"github.com/drizzleent/ton-bot/internal/events/telegram"
)

type serviceProvider struct {
	tgConfig config.BotConfig

	consumer consumer.Consumer

	fetcher   events.Fetcher
	processor events.Processor

	tgClient *tgClient.Client
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) BotConfig() config.BotConfig {
	if nil == s.tgConfig {
		cfg, err := env.NewBotConfig()
		if err != nil {
			log.Fatalf("failed to load tg bot cofig %s", err.Error())
		}
		s.tgConfig = cfg
	}

	return s.tgConfig
}

func (s *serviceProvider) TgClient(_ context.Context) *tgClient.Client {
	if nil == s.tgClient {
		s.tgClient = tgClient.New(s.BotConfig().Host(), s.BotConfig().Token())
	}
	return s.tgClient
}

func (s *serviceProvider) Fetcher(ctx context.Context) events.Fetcher {
	if nil == s.fetcher {
		s.fetcher = telegram.New(s.TgClient(ctx))
	}

	return s.fetcher
}

func (s *serviceProvider) Processor(ctx context.Context) events.Processor {
	if nil == s.processor {
		s.processor = telegram.New(s.TgClient(ctx))
	}

	return s.processor
}

func (s *serviceProvider) Consumer(ctx context.Context) consumer.Consumer {
	if nil == s.consumer {
		s.consumer = eventconsumer.New(s.Fetcher(ctx), s.Processor(ctx), 100)
	}

	return s.consumer
}
