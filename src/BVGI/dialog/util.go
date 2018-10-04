package dialog

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"BVGI/db"
	"BVGI/intent"
	"util"
	// "util/debug"
	"github.com/michlabs/fbbot"
	"encoding/json"
	"os"
)

func CallStaffs(bot *fbbot.Bot, msg *fbbot.Message) {
	staffs, err := db.GetAllStaffs()
	if err != nil {
		log.Error("Error in call staffs ", err.Error())
	}

	messageToStaff := util.Personalize(T("staff_alert_new_message"), &msg.Sender)

	for _, v :=  range staffs {
		u := fbbot.User{
			ID: v.FbID,
		}
		log.Infof("Send alert message to %s (%s)\n", v.FbID, v.Fullname)
		bot.SendText(u, messageToStaff)
	}
}

func checkAddress(msg string) bool{
	if strings.Compare(msg,"") == 0 {
		return false
	}
	return true
}
func checkSport(msg string) bool{
	if strings.Compare(msg,"") == 0 {
		return false
	}
	return true
}

func checkDatetime(msg string) bool{
	if strings.Compare(msg,"") == 0 {
		return false
	}
	return true
}

func checkSeatType(msg string) bool{
	if strings.Compare(msg,"") == 0 {
		return false
	}
	return true
}

func isValidBookInfo(address string,datetime string,seattype string) bool{
	if strings.Compare(seattype,T("confirm_book_seat_A")) == 0{
		return true
	}else{
		return false
	}
	return true
}

var yesno []OptionMess

func createBookData(){
	sports = append(sports, OptionMess{"sport 1",""})
	sports = append(sports, OptionMess{"sport 2",""})
	sports = append(sports, OptionMess{"sport 3",""})

	places = append(places, OptionMess{"place 1",""})
	places = append(places, OptionMess{"place 2",""})
	places = append(places, OptionMess{"place 3",""})

	times = append(times, OptionMess{"午前",""})
	times = append(times, OptionMess{"午後",""})
	times = append(times, OptionMess{"夜",""})

	yesno = append(yesno, OptionMess{T("no_answer_confirm_yes"),""})
	yesno = append(yesno, OptionMess{T("no_answer_confirm_no"),""})
}

func createSelectOption(bot *fbbot.Bot, msg *fbbot.Message,mess string, options []OptionMess){
	var items  []fbbot.QuickRepliesItem

	bot.SendText(msg.Sender, mess)
	for _, option := range options {
		items = append(items, fbbot.QuickRepliesItem{
			ContentType: "Text",
			Title:       option.name,
			Payload:     option.name,
			// ImageURL:    T("select_human_icon"),
		})
	}
	botSelect := new(fbbot.QuickRepliesMessage)
	botSelect.Text = util.Personalize(T("select_bot_options"), &msg.Sender)
	botSelect.Items = items

	bot.Send(msg.Sender, botSelect)
}

// Minhnt: Hàm tạo danh sách thông tin người yc/người rủi ro
func CreateListPerson(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stBMBHValue := bot.STMemory.For(msg.Sender.ID).Get("BMBHInfo")
	var arrBMBH []string
	if stBMBHValue != "" {
		err := json.Unmarshal([]byte(stBMBHValue), &arrBMBH)
		if err == nil {
			botSelect := new(fbbot.ButtonMessage)

			for i:=0; i < len(arrBMBH); i++ {
				if (i == 0 || i == 3) {
					botSelect = new(fbbot.ButtonMessage)
					botSelect.Text = mess
					if i > 0 { botSelect.Text = "..." }
				}
				botSelect.AddPostbackButton(arrBMBH[i], arrBMBH[i])
				if (i == 2 || i == (len(arrBMBH)-1)) {
					bot.Send(msg.Sender, botSelect)
				}
			}
		} else {

		}
	}

}

// Minhnt: Hàm tạo danh sách thông tin người yc/người rủi ro
func CreateQuickReplyPerson(bot *fbbot.Bot, msg *fbbot.Message, mess string) {
	stBMBHValue := bot.STMemory.For(msg.Sender.ID).Get("BMBHInfo")
	var arrBMBH []string
	if stBMBHValue != "" {
		err := json.Unmarshal([]byte(stBMBHValue), &arrBMBH)
		if err == nil {
			var items  []fbbot.QuickRepliesItem

			for i:=0; i < len(arrBMBH); i++ {
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrBMBH[i],
					Payload:     arrBMBH[i],
					// ImageURL:    T("select_human_icon"),
				})
			}

			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}

}

// Minhnt: Hàm tạo danh sách loại rủi ro
func CreateQuickReplyTypeOfRisk(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stTypeOfRisk := os.Getenv("TYPE_OF_RISK")
	if stTypeOfRisk != "" {
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")

		if len(arrTypeOfRisk) > 0 {
			var items  []fbbot.QuickRepliesItem

			for _, val := range arrTypeOfRisk {
				arrItems := strings.Split(val, "-")
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrItems[1],
					Payload:    arrItems[0],
				})
			}

			// Send message
			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}
}
// THUCDH: Hàm tạo danh sách loại MENU
func CreateQuickReplyTypeOfMenu(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stTypeOfRisk := os.Getenv("TYPE_OF_OPTION")
	if stTypeOfRisk != "" {
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")

		if len(arrTypeOfRisk) > 0 {
			var items  []fbbot.QuickRepliesItem

			for _, val := range arrTypeOfRisk {
				arrItems := strings.Split(val, "-")
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrItems[1],
					Payload:    arrItems[0],
				})
			}

			// Send message
			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}
}

func CreateQuickReplyTypeOfMenuTaisan(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stTypeOfRisk := os.Getenv("TYPE_OF_TAISAN")
	if stTypeOfRisk != "" {
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")

		if len(arrTypeOfRisk) > 0 {
			var items  []fbbot.QuickRepliesItem

			for _, val := range arrTypeOfRisk {
				arrItems := strings.Split(val, "-")
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrItems[1],
					Payload:    arrItems[0],
				})
			}

			// Send message
			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}
}
func CreateQuickReplyTypeOfMenuBAGDV(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stTypeOfRisk := os.Getenv("TYPE_OF_BAGDV")
	if stTypeOfRisk != "" {
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")

		if len(arrTypeOfRisk) > 0 {
			var items  []fbbot.QuickRepliesItem

			for _, val := range arrTypeOfRisk {
				arrItems := strings.Split(val, "-")
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrItems[1],
					Payload:    arrItems[0],
				})
			}

			// Send message
			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}
}
func CreateQuickReplyTypeOfCommon(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stTypeOfRisk := os.Getenv("TYPE_OF_COMMON")
	if stTypeOfRisk != "" {
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")

		if len(arrTypeOfRisk) > 0 {
			var items  []fbbot.QuickRepliesItem

			for _, val := range arrTypeOfRisk {
				arrItems := strings.Split(val, "-")
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrItems[1],
					Payload:    arrItems[0],
				})
			}

			// Send message
			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}
}
func CreateQuickReplyTypeOfMenuTauthuy(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stTypeOfRisk := os.Getenv("TYPE_OF_TAUTHUY")
	if stTypeOfRisk != "" {
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")

		if len(arrTypeOfRisk) > 0 {
			var items  []fbbot.QuickRepliesItem

			for _, val := range arrTypeOfRisk {
				arrItems := strings.Split(val, "-")
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrItems[1],
					Payload:    arrItems[0],
				})
			}

			// Send message
			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}
}
//duydp build on 30/3/2018
func CreateQuickReplyTypeOfMenuCargo(bot *fbbot.Bot, msg *fbbot.Message, mess string) {

	stTypeOfRisk := os.Getenv("TYPE_OF_CARGO")
	if stTypeOfRisk != "" {
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")

		if len(arrTypeOfRisk) > 0 {
			var items  []fbbot.QuickRepliesItem

			for _, val := range arrTypeOfRisk {
				arrItems := strings.Split(val, "-")
				items = append(items, fbbot.QuickRepliesItem{
					ContentType: "Text",
					Title:       arrItems[1],
					Payload:    arrItems[0],
				})
			}

			// Send message
			botSelect := new(fbbot.QuickRepliesMessage)
			botSelect.Text = mess
			botSelect.Items = items

			bot.Send(msg.Sender, botSelect)
		}
	}
}
func HandleQuestion(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	question := bot.STMemory.For(msg.Sender.ID).Get("question")
	if question == "" {
		question = msg.Text
	}
	log.Infof("Question = %s", question)

	bot.STMemory.For(msg.Sender.ID).Delete("question")
	i := intent.Detect(question)
	switch i {
	case "":
		log.Debug("There's no matching sentence for: ", msg.Text)
		bot.SendText(msg.Sender, util.Personalize(T("dontunderstand"), &msg.Sender))
		return NoAnswerEvent
	case "goodbye":
		return GoodbyeEvent
	case CheckBookedIntent:
		return CheckBookedEvent
	case BookIntent:
		log.Debug("going BookIntent")
		createBookData()
		return NewBookEvent
	case BookWithDetailIntent:
		log.Debug("going BookwithdetailIntent")
		createBookData()
		return NewBookWithdetailEvent
	case GuideIntent:
		return GuideEvent
	case FaQStadiumIntent:
		return FaQStadiumEvent
	case FaQAthleteIntent:
		return FaQAthleteEvent
	case FaQSportIntent:
		return FaQSportEvent
	default:
		a := db.GetAnswerFor(i)
		if a == "" {
			bot.SendText(msg.Sender, util.Personalize(T("no_answer"), &msg.Sender))
			return NoAnswerEvent
		}

		if strings.Contains(a, "@call_staff") {
			a = strings.Replace(a, "@call_staff", "", -1)
			CallStaffs(bot, msg)
		}

		var arr []string
		a = strings.TrimSpace(a)
		arr = strings.SplitN(a, "\n", -1)
		for _, v := range arr {
			bot.TypingOn(msg.Sender)
			util.SendTextWithImages(bot, msg.Sender, util.Personalize(v, &msg.Sender))
		}
	}
	return StayEvent
}