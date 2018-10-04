package dialog

import (
	"github.com/michlabs/fbbot"
	"BVGI/intent"
	log "github.com/Sirupsen/logrus"
	"BVGI/db"
)

type newBook struct {
	fbbot.BaseStep
}

func (s newBook) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	log.Debug("Entering newBookEvent")

	//nhap noi dung con thieu theo format
	bot.SendText(msg.Sender, T("ask_book_all"))
	return FillSportEvent
}
//jub

func (s newBook) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	log.Debug("Entering newBookEvent Process")
	_intent := intent.Detect(msg.Text)

	switch _intent {
	case BookIntent:
		return ProcessBook(bot,msg)
	case OfftopicIntent:
		bot.SendText(msg.Sender, "noi dung khong lien quan den book ve")
	}
	return ErrorEvent
}

func ProcessBook(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	meaning,err := intent.ExtractEntity(msg.Text)
	log.Debug("ExtractEntity Done")
	if err != nil {
		log.Error("Error in ExtractEntity ", err.Error())
		return ErrorEvent
	}

	log.Debug("anaylyzing Entity ")
	for _, v :=  range meaning.Entities{
		switch v.Name {
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

	address := bot.STMemory.For(msg.Sender.ID).Get(address)
	datetime := bot.STMemory.For(msg.Sender.ID).Get(datetime)
	seattype := bot.STMemory.For(msg.Sender.ID).Get(seattype)
	if isValidBookInfo(address,datetime,seattype) {
		log.Debug("Entering newBookEvent Process")
		// TO-DO confirm logic
		_answer := db.GetAnswerFor(BookIntent)
		bot.SendText(msg.Sender, _answer)

		//Thông báo book ok->#faq
		bot.SendText(msg.Sender, T("book_ok"))
		return GoFAQEvent
	} else {
		log.Debug("calling book event")
		return FillSportEvent
	}
}
