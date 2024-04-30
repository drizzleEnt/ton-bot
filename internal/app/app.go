package app

import (
	"context"

	"github.com/drizzleent/ton-bot/pkg/closer"
	"github.com/subosito/gotenv"
)

type App struct {
	serviceProvide *serviceProvider
}

func New(ctx context.Context) (*App, error) {
	a := &App{}
	err := a.initdebs(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	err := a.serviceProvide.Consumer(ctx).Start()

	if err != nil {
		return err
	}

	return nil
}

func (a *App) initdebs(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := gotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvide = newServiceProvider()
	return nil
}
