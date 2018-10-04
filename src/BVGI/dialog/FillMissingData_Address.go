package dialog

import (
	"github.com/michlabs/fbbot"
	//"go/types"
)
//fill address
type fillAddress struct {
	fbbot.BaseStep
}

var places []OptionMess

//fill address
func (s fillAddress) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	_address := bot.STMemory.For(msg.Sender.ID).Get(address)
	if checkAddress(_address){
		bot.SendText(msg.Sender, "[debug] address: "+_address)
		return FillDatetimeEvent
	}else {
		//show address option
		createSelectOption(bot,msg,T("ask_book_adress"),places)

		return StayEvent
	}
}
// fill sport process
func (s fillAddress) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.STMemory.For(msg.Sender.ID).Set(address, msg.Text)
	return FillAddressEvent
}

