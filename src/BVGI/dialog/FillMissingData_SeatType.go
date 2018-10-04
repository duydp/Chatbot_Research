package dialog

import "github.com/michlabs/fbbot"

type fillSeatType struct {
	fbbot.BaseStep
}

var seats []OptionMess

//fill address
func (s fillSeatType) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	_seat := bot.STMemory.For(msg.Sender.ID).Get(seattype)
	if checkAddress(_seat){
		bot.SendText(msg.Sender, "[debug] seattype: "+_seat)
		return FillDatetimeEvent
	}else {
		//show address option
		createSelectOption(bot,msg,T("ask_book_seat"),seats)

		return StayEvent
	}
}
// fill sport process
func (s fillSeatType) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.STMemory.For(msg.Sender.ID).Set(address, msg.Text)
	return FillAddressEvent
}
