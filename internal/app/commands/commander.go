package commands

import (
	"github.com/Antchel/GoTgBot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(bot *tgbotapi.BotAPI, s *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: s,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("Panic value %v", panicValue)
		}
	}()
	if update.Message != nil { // If we got a message
		command, ok := registeredCommands[update.Message.Command()]

		if ok {
			command(c, update.Message)
		} else {
			c.Default(update.Message)
		}

	}
}

func callMe() {

}
