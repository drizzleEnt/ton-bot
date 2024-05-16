package telegram

import (
	"fmt"
	"log"
	"strings"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
	ShowCmd  = "/show"
)

func (p *processor) doCmd(text string, chatID int, username string) error {

	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID, username)
	case ShowCmd:
		return p.sendGpusStats(chatID)
	default:
		p.sendUnknownCommand(chatID)
	}

	return nil
}

func (p *processor) sendGpusStats(chatID int) error {
	resp, err := p.srv.CheckGpu("")
	if err != nil {
		return err
	}
	return p.tg.SendMesage(chatID, resp)
}

func (p *processor) sendHello(chatID int, username string) error {
	msg := fmt.Sprintf(msgHello, username)
	return p.tg.SendMesage(chatID, msg)
}

func (p *processor) sendHelp(chatID int) error {
	return p.tg.SendMesage(chatID, msgHelp)
}

func (p *processor) sendUnknownCommand(chatID int) error {
	return p.tg.SendMesage(chatID, msgUnknownCommand)
}
