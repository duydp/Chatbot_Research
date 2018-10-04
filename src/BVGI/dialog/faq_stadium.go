package dialog

import "github.com/michlabs/fbbot"

type FaQ_Stadium struct {
	fbbot.BaseStep
}

func (s FaQ_Stadium) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	return GoodbyeEvent
}
