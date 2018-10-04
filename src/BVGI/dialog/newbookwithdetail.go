package dialog

import (
	"github.com/michlabs/fbbot"
	log "github.com/Sirupsen/logrus"
	"BVGI/intent"
)

type newBookWithDetail struct {
	fbbot.BaseStep
}

func (s newBookWithDetail) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	log.Debug("Entering newBookWithDetailEvent")

	bot.SendText(msg.Sender, "[debug] Entering newBookWithDetailEvent")
	//luu cac data trich xuat dc
	meaning,err := intent.ExtractEntity(msg.Text)
	if err != nil {
		log.Error("Error in ExtractEntity ", err.Error())
		return ErrorEvent
	}
	for _, v :=  range meaning.Entities{
		bot.SendText(msg.Sender, v.Name+": "+v.Value)
		switch v.Name {
		case Sport:
			bot.STMemory.For(msg.Sender.ID).Set(Sport, v.Value)
			break
		case Address_intent:
			bot.STMemory.For(msg.Sender.ID).Set(Address_intent, v.Value)
			break
		case Datetime_intent:
			bot.STMemory.For(msg.Sender.ID).Set(Datetime_intent, v.Value)
			break
		case seattype:
			bot.STMemory.For(msg.Sender.ID).Set(seattype, v.Value)
			break
		}
	}

	// hoi cac data con thieu
	return 	FillSportEvent
}
