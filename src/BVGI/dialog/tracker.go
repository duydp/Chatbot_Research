package dialog

import (
	"strings"
	"time"

	"util"
	"github.com/michlabs/fbbot"
)

type ActivityTracker struct{}

func (h *ActivityTracker) HandleEcho(bot * fbbot.Bot, msg *fbbot.Message)  {
	if msg.AppID > 0 {
		return
	}

	if msg.Text == "/bot" {
		bot.SendText(msg.Sender, "Move to bot")
		insujs.Move(msg, Welcome{})
	}

	if strings.HasPrefix(msg.Text, "/") {
		return
	}
	bot.STMemory.For(msg.Page.ID).Set("lastEcho", util.FromTime(time.Now()))
}

func (h *ActivityTracker) HandleMessage(bot *fbbot.Bot, msg *fbbot.Message) {
	bot.STMemory.For(msg.Sender.ID).Set("lastMessage", util.FromTime(time.Now()))
}

func (h *ActivityTracker) HandlePostback(bot *fbbot.Bot, msg *fbbot.Postback) {
	bot.STMemory.For(msg.Sender.ID).Set("lastPostback", util.FromTime(time.Now()))
}