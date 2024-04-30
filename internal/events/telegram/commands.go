package telegram

import (
	"log"
	"strings"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *processor) doCmd(text string, chatID int, username string) error {

	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	}

	return nil
}

func (p *processor) sendHello(chatID int) error {
	return p.tg.SendMesage(chatID, msgHello)
}

func (p *processor) sendHelp(chatID int) error {
	return p.tg.SendMesage(chatID, msgHelp)
}
