package dialog

import (
	"github.com/michlabs/fbbot"
	"BVGI/intent"
	"util"
	"strings"
	log "github.com/Sirupsen/logrus"
)

const sport_football	= "football"
const Sport		= "sport"
const address		= "address"
const datetime 		= "datetime"
const seattype 		= "seat_Type"
const seat		= "seat"
const availbleseat	= "availbleseat"

type Book struct {
	fbbot.BaseStep
}
type BookL2 struct {
	fbbot.BaseStep
}
type BookL3 struct {
	fbbot.BaseStep
}
type BookL4 struct {
	fbbot.BaseStep
}
type BookL5 struct {
	fbbot.BaseStep
}


//--------------------------LAYER1--------------------------//
func (s Book) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	log.Debug("Entering BookEvent")
	_address := bot.STMemory.For(msg.Sender.ID).Get(Address_intent)
	if checkAddress(_address){
		log.Debug("address: ",_address)
		return BookL2Event
	}else {
		bot.SendText(msg.Sender, T("ask_book_adress"))
		return StayEvent
	}
}

func (s Book) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	i := intent.Detect(msg.Text)
	if i == Address_intent {
		log.Debug("add ok %s",msg.Text)
		bot.STMemory.For(msg.Sender.ID).Set(Address_intent, msg.Text)
	}
	if i==OfftopicIntent {
		log.Debug("off topic ",msg.Text)
		return GoFAQEvent
	}

	return GoFAQEvent
}


//--------------------------LAYER2--------------------------//
func (s BookL2) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	datetime := bot.STMemory.For(msg.Sender.ID).Get(Datetime_intent)

	if checkDatetime(datetime){
		return BookL3Event
	}else {
		bot.SendText(msg.Sender, T("ask_book_time"))
		return StayEvent
	}
}

func (s BookL2) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	i := intent.Detect(msg.Text)
	if i == Datetime_intent {
		bot.STMemory.For(msg.Sender.ID).Set(Datetime_intent, msg.Text)
	}
	if i==OfftopicIntent {
		log.Debug("off topic ",msg.Text)
		return GoFAQEvent
	}

	return GoFAQEvent
}

//--------------------------LAYER3--------------------------//
func (s BookL3) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	seattype := bot.STMemory.For(msg.Sender.ID).Get(seattype)

	if checkAddress(seattype){
		return BookL4Event
	}else {
		availbleseat := bot.LTMemory.For(seat).Get(availbleseat)
		if availbleseat == "" {
			availbleseat = "0"
		}
		//moi chon loaij ghe
		mess := T("ask_book_seat")+availbleseat
		bot.SendText(msg.Sender, mess)
		//comfirm ok or not
		var items  []fbbot.QuickRepliesItem
		items = append(items, fbbot.QuickRepliesItem{
			ContentType: "Text",
			Title:       T("confirm_book_seat_A"),
			Payload:     "A",
			// ImageURL:    T("select_human_icon"),
		})
		items = append(items, fbbot.QuickRepliesItem{
			ContentType: "Text",
			Title:       T("confirm_book_seat_B"),
			Payload:     "B",
			// ImageURL:    T("select_bot_icon"),
		})
		botSelect := new(fbbot.QuickRepliesMessage)
		botSelect.Text = util.Personalize(T("confirm_book_seat"), &msg.Sender)
		botSelect.Items = items

		bot.Send(msg.Sender, botSelect)

		//s.ProcessBookL3(bot,msg)
		return StayEvent
	}
}

func (s BookL3) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.STMemory.For(msg.Sender.ID).Set(seat, msg.Text)

	return NewBookEvent
}

//--------------------------LAYER4--------------------------//
func (s BookL4) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	//Gọi process book Check xem book có ok hay không

	address := bot.STMemory.For(msg.Sender.ID).Get(Address_intent)
	datetime := bot.STMemory.For(msg.Sender.ID).Get(Datetime_intent)
	seattype := bot.STMemory.For(msg.Sender.ID).Get(seattype)
	if isValidBookInfo(address,datetime,seattype) {
		//Thông báo book ok->#faq
		bot.SendText(msg.Sender, T("book_ok"))
		return GoFAQEvent
	} else {
		//Thông báo book không thành công Hỏi user có muốn gặp cskh không
		bot.SendText(msg.Sender, T("book_false"))
		var items  []fbbot.QuickRepliesItem
		items = append(items, fbbot.QuickRepliesItem{
			ContentType: "Text",
			Title:       T("select_human"),
			Payload:     "yes",
			// ImageURL:    T("select_human_icon"),
		})
		items = append(items, fbbot.QuickRepliesItem{
			ContentType: "Text",
			Title:       T("select_bot"),
			Payload:     "no",
			// ImageURL:    T("select_bot_icon"),
		})

		bot.SendText(msg.Sender, util.Personalize(T("select_bot_title"), &msg.Sender))
		botSelect := new(fbbot.QuickRepliesMessage)
		botSelect.Text = util.Personalize(T("select_bot_options"), &msg.Sender)
		botSelect.Items = items

		bot.Send(msg.Sender, botSelect)
		return StayEvent
	}
}

func (s BookL4) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case strings.ToLower(util.RemoveAccent(T("select_human"))):
		return GoSlienceEvent
	case strings.ToLower(util.RemoveAccent(T("select_bot"))):
		return GoFAQEvent
	}
	return GoFAQEvent
}