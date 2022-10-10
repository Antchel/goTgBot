package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outMsg := "Here all products : \n\n"
	list := c.productService.List()
	for _, p := range list {
		outMsg += p.Title
		outMsg += "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outMsg)
	c.bot.Send(msg)
}
