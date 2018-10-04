package dialog

import "github.com/michlabs/fbbot"

type FaQ_athlete struct {
	fbbot.BaseStep
}

func (s FaQ_athlete) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	return GoodbyeEvent
}
