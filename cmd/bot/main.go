package main

import (
	"github.com/Antchel/GoTgBot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			switch update.Message.Command() {

			case "help":
				helpCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, productService)
			default:
				defaultBehavior(bot, update.Message)
			}

		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "/help - help\n"+
		"/list - list products")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, s *product.Service) {
	outMsg := "Here all products : \n\n"
	list := s.List()
	for _, p := range list {
		outMsg += p.Title
		outMsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMsg)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("Authorized on account %s", bot.Self.UserName)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	bot.Send(msg)
}
