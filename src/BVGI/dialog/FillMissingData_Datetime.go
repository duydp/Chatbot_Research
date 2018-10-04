package dialog

import "github.com/michlabs/fbbot"

//fill datetime
type fillDatetime struct {
	fbbot.BaseStep
}

var times []OptionMess

//fill datetime
func (s fillDatetime) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	_datetime := bot.STMemory.For(msg.Sender.ID).Get(datetime)
	if checkAddress(_datetime){
		bot.SendText(msg.Sender, "[debug] datetime: "+_datetime)
		return FillDatetimeEvent
	}else {
		//show datetime option
		createSelectOption(bot,msg, T("ask_book_time"),times)

		return StayEvent
	}
}
