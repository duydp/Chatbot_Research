package dialog

import (

	"github.com/michlabs/fbbot"
)

type Guide struct {
	fbbot.BaseStep
}

func (s Guide) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, T("guide"))
	return GoFAQEvent
}
