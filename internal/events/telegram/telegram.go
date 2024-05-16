package telegram

import (
	"fmt"

	"github.com/drizzleent/ton-bot/internal/clients/telegram"
	"github.com/drizzleent/ton-bot/internal/events"
	"github.com/drizzleent/ton-bot/internal/service"
)

type processor struct {
	srv    service.Service
	tg     *telegram.Client
	offset int
}

type Meta struct {
	ChatID   int
	Username string
}

func New(cl *telegram.Client, srv service.Service) *processor {
	return &processor{
		srv: srv,
		tg:  cl,
	}
}

func (p *processor) Process(e events.Event) error {

	switch e.Type {
	case events.Message:
		return p.processMessage(e)
	default:
		return fmt.Errorf("cannot procces message")
	}
}

func (p *processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)

	if err != nil {
		return nil, err
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *processor) processMessage(e events.Event) error {
	meta, err := meta(e)
	if err != nil {
		return err
	}

	if err := p.doCmd(e.Text, meta.ChatID, meta.Username); err != nil {
		return fmt.Errorf("cannot procces message %v", err)
	}
	return nil
}

func meta(e events.Event) (Meta, error) {
	res, ok := e.Meta.(Meta)
	if !ok {
		return Meta{}, fmt.Errorf("cant get meta")
	}
	return res, nil
}

func event(upd telegram.Update) events.Event {
	updType := fetchType(upd)

	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatID:   upd.Message.Chat.ID,
			Username: upd.Message.From.Username,
		}
	}

	return res
}

func fetchText(upd telegram.Update) string {
	if upd.Message == nil {
		return ""
	}

	return upd.Message.Text
}

func fetchType(upd telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}

	return events.Message
}
