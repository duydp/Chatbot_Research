package dialog

import "github.com/michlabs/fbbot"

type FaQ_Sport struct {
	fbbot.BaseStep
}

func (s FaQ_Sport) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	return GoodbyeEvent
}
