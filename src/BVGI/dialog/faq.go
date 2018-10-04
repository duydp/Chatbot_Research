package dialog

import (
	"github.com/michlabs/fbbot"
)

type FAQ struct {
	fbbot.BaseStep
}

func (s FAQ) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.STMemory.For(msg.Sender.ID).Get("previous_step") == "Emoji" {
		bot.STMemory.For(msg.Sender.ID).Delete("previous_step")
	}
	question := bot.STMemory.For(msg.Sender.ID).Get("question")
	if question != "" {
		return s.Process(bot, msg)
	}

	return StayEvent
}


func (s FAQ) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.STMemory.For(msg.Sender.ID).Get("previous_step") == "Emoji" {
		bot.STMemory.For(msg.Sender.ID).Delete("previous_step")
	}else {
		bot.TypingOn(msg.Sender)
		return HandleQuestion(bot, msg)
	}
	return StayEvent
}
