package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("Authorized on account %s", c.bot.Self.UserName)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["default"] = (*Commander).Default
}
