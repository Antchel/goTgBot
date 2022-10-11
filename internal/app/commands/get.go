package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)

	if err != nil {
		log.Println("wrong arg", args)
		return
	}

	product, err := c.productService.Get(idx)
	if err != nil {
		log.Println(fmt.Sprintf("fail to get product with index %d: %v", idx, product))
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, product.Title)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
