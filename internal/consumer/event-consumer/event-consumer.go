package eventconsumer

import (
	"log"
	"time"

	"github.com/drizzleent/ton-bot/internal/events"
)

type consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchSize int) *consumer {
	return &consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c *consumer) Start() error {
	for {
		gotevents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("{ERR}: Consumer: %s", err.Error())
		}

		if len(gotevents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvent(gotevents); err != nil {
			log.Print(err)
			continue
		}

	}
}

func (c *consumer) handleEvent(events []events.Event) error {
	for _, e := range events {
		log.Printf("got new event: %s", e.Text)

		if err := c.processor.Process(e); err != nil {
			log.Printf("cant process event : %s", err.Error())
			continue
		}
	}

	return nil
}
