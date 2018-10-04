package dialog

import (
	"util"
	"github.com/michlabs/fbbot"
)


type Error struct {
	fbbot.BaseStep
}

func (s Error) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("error"), &msg.Sender))
	return GoSlienceEvent
}
