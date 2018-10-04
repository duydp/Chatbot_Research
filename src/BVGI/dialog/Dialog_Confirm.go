package dialog

import (
	"github.com/michlabs/fbbot"
	"strings"
)

type dialog_confirm struct {
	fbbot.BaseStep
}

func (s dialog_confirm) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	createSelectOption(bot,msg,ConfirmMess(bot,msg),yesno)
	return StayEvent
}

func (s dialog_confirm) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	//msg.Text
	switch msg.Text {
	case T("no_answer_confirm_yes"):
		break
	default:
		return GoFAQEvent
		break
	}
	return GoFAQEvent
}

func ConfirmMess(bot *fbbot.Bot, msg *fbbot.Message) string{
	var Mess = T("ask_book_all")

	_sport := bot.STMemory.For(msg.Sender.ID).Get(Sport)
	strings.Replace(Mess,"@"+Sport,_sport,1)

	_address := bot.STMemory.For(msg.Sender.ID).Get(address)
	strings.Replace(Mess,"@"+address,_address,1)

	_datetime := bot.STMemory.For(msg.Sender.ID).Get(datetime)
	strings.Replace(Mess,"@"+datetime,_datetime,1)

	_seat := bot.STMemory.For(msg.Sender.ID).Get(seattype)
	strings.Replace(Mess,"@"+seattype,_seat,1)

	return Mess
}