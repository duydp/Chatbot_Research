package dialog

import (
	"strings"

	"util"
	"github.com/michlabs/fbbot"
)

type NoAnswer struct {
	fbbot.BaseStep
}

func (s NoAnswer) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	var items  []fbbot.QuickRepliesItem
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("no_answer_confirm_yes"),
		Payload:     "yes",
		ImageURL:    T("no_answer_confirm_yes_image"),
	})
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("no_answer_confirm_no"),
		Payload:     "no",
		ImageURL:    T("no_answer_confirm_no_image"),
	})

	staffSupportOptions := new(fbbot.QuickRepliesMessage)
	staffSupportOptions.Text = util.Personalize(T("confirm_go_for_staff"), &msg.Sender)
	staffSupportOptions.Items = items

	bot.Send(msg.Sender, staffSupportOptions)

	return StayEvent
}

func (s NoAnswer) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case strings.ToLower(util.RemoveAccent(T("no_answer_confirm_yes"))):
		return GoSlienceEvent
	case strings.ToLower(util.RemoveAccent(T("no_answer_confirm_no"))):
		bot.SendText(msg.Sender, util.Personalize(T("back_faq"), &msg.Sender))
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "noAnswer")
		return GoFAQEvent
	default:
		bot.STMemory.For(msg.Sender.ID).Set("question", msg.Text)
		return GoFAQEvent
	}
}