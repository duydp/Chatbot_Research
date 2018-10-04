package dialog

import (
	"github.com/michlabs/fbbot"
	"BVGI/intent"
)
//fill sport
type fillSport struct {
	fbbot.BaseStep
}
var sports []OptionMess
// fill sport enter
func (s fillSport) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	_sport := bot.STMemory.For(msg.Sender.ID).Get(Sport)
	if checkSport(_sport){
		bot.SendText(msg.Sender, "[debug] sport: "+_sport)
		return FillAddressEvent
	}else {
		//send sport option
		bot.SendText(msg.Sender,T("ask_book_sport"))
		return StayEvent
	}
}

// fill sport process
func (s fillSport) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	//intent sport
	i := intent.Detect(msg.Text)
	if i == Sport {
		bot.SendText(msg.Sender, "[debug] sport: " + msg.Text)
		bot.STMemory.For(msg.Sender.ID).Set(Sport, msg.Text)
	}

	//entities sport
	entities,err := intent.ExtractEntity(msg.Text)
	if err != nil {
		return ErrorEvent
	}

	for _, v :=  range entities.Entities {

		bot.SendText(msg.Sender, v.Name+": "+v.Value)
		switch v.Name {
		case Sport:
			bot.SendText(msg.Sender, "[debug] sport: " + msg.Text)
			bot.STMemory.For(msg.Sender.ID).Set(Sport, v.Value)
			break
		}
	}

	if i==OfftopicIntent {
		bot.SendText(msg.Sender, "[debug] off topic ")
		return GoFAQEvent
	}

	return FillSportEvent
}

type OptionMess struct {
	name string
	imagelink string
}
