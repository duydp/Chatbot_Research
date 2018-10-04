package debug

import (
	"io"
	"github.com/michlabs/fbbot"
)

const DEBUG_MARK string = "debug_clause_failed"
const DEBUG_MEMORY string = "debug_clause"

var debugger *Debugger 

type Debugger struct {
	bot *fbbot.Bot
	writer io.Writer
}

func Init(bot *fbbot.Bot, writer io.Writer) {
	debugger = &Debugger{
		bot: bot,
		writer: writer,
	}
}

func Verify(message string, clause string, receiver fbbot.User) {
	m := fbbot.NewButtonMessage()
	m.Text = "[Debug] " + message
	m.AddPostbackButton("Sai", DEBUG_MARK)
	debugger.bot.STMemory.For(receiver.ID).Set(DEBUG_MEMORY, clause)
	debugger.bot.Send(receiver, m)
}

func Save(clause string) {
	_, err := debugger.writer.Write([]byte(clause + "\n"))
	if err != nil {
		debugger.bot.Logger.Error("debugger failed: ", err.Error())
	}
}

func CatchResponse(msg *fbbot.Message) bool {
	if msg.Text != DEBUG_MARK {
		return false
	}
	Save(debugger.bot.STMemory.For(msg.Sender.ID).Get(DEBUG_MEMORY))
	debugger.bot.SendText(msg.Sender, "[Debug] Cảm ơn bạn đã phản hồi! Mời bạn tiếp tục hội thoại.")
	return true
}