package dialog

import (
	"BVGI/intent"
	"util"
	"github.com/michlabs/fbbot"
	"github.com/Sirupsen/logrus"
)

type Welcome struct {
	fbbot.BaseStep
}

func (s Welcome) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	isFirstConversation := bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation")
	question := bot.STMemory.For(msg.Sender.ID).Get("question")
	logrus.Infoln("question = ", question)
	if question == "" {
		question = "hello"
	}

	i := intent.Detect(question)

	if isFirstConversation == "false" {
		bot.TypingOn(msg.Sender)

		if i == "hello" {
			bot.SendText(msg.Sender, util.Personalize(T("welcome_second_time"), &msg.Sender))
			bot.STMemory.For(msg.Sender.ID).Delete("question")
			return GoFAQEvent
		}else {
			bot.SendText(msg.Sender, util.Personalize(T("welcome_second_time_answer"), &msg.Sender))
			return GoFAQEvent
		}

	}else{
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "false")
		bot.TypingOn(msg.Sender)

		if i == "hello" {
			bot.SendText(msg.Sender, util.Personalize(T("welcome_first_time"), &msg.Sender))
			bot.STMemory.For(msg.Sender.ID).Delete("question")
			return GoFAQEvent
		}else{
			bot.SendText(msg.Sender, util.Personalize(T("welcome_first_time_answer"), &msg.Sender))
			return GoFAQEvent
		}
	}
}