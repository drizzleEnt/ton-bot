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
	"github.com/drizzleent/ton-bot/internal/service"
	"github.com/drizzleent/ton-bot/internal/service/check"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	tgConfig config.BotConfig

	consumer consumer.Consumer

	fetcher   events.Fetcher
	processor events.Processor

	service service.Service

	tgClient       *tgClient.Client
	grpcClientConn *grpc.ClientConn
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

func (s *serviceProvider) GRPCClientConn(_ context.Context) *grpc.ClientConn {
	if nil == s.grpcClientConn {
		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect grpc server %s", err.Error())
		}
		s.grpcClientConn = conn
	}
	return s.grpcClientConn
}

func (s *serviceProvider) TgClient(_ context.Context) *tgClient.Client {
	if nil == s.tgClient {
		s.tgClient = tgClient.New(s.BotConfig().Host(), s.BotConfig().Token())
	}
	return s.tgClient
}

func (s *serviceProvider) Fetcher(ctx context.Context) events.Fetcher {
	if nil == s.fetcher {
		s.fetcher = telegram.New(s.TgClient(ctx), s.Service(ctx))
	}

	return s.fetcher
}

func (s *serviceProvider) Processor(ctx context.Context) events.Processor {
	if nil == s.processor {
		s.processor = telegram.New(s.TgClient(ctx), s.Service(ctx))
	}

	return s.processor
}

func (s *serviceProvider) Consumer(ctx context.Context) consumer.Consumer {
	if nil == s.consumer {
		s.consumer = eventconsumer.New(s.Fetcher(ctx), s.Processor(ctx), 100)
	}

	return s.consumer
}

func (s *serviceProvider) Service(ctx context.Context) service.Service {
	if nil == s.service {
		s.service = check.NewCheckService(s.GRPCClientConn(ctx))
	}

	return s.service
}
