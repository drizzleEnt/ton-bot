package main

import (
	"context"
	"log"

	"github.com/drizzleent/ton-bot/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.New(ctx)
	if err != nil {
		log.Fatalf("failed to init app %v", err)
	}

	err = app.Run(ctx)
	if err != nil {
		log.Fatalf("failed to run app %v", err)
	}

}
