package dialog

import (
	"util"
	"BVGI/db"

	"github.com/michlabs/fbbot"
)

type CheckBooked struct {
	fbbot.BaseStep
}

func (s CheckBooked) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	answer := db.GetAnswerFor(CheckBookedIntent)

	if answer == "" {
		bot.SendText(msg.Sender, util.Personalize(T("no_answer"), &msg.Sender))
		return NoAnswerEvent
	}

	answer = util.Personalize(answer, &msg.Sender)
	bot.SendText(msg.Sender, answer)
	return GoFAQEvent
}