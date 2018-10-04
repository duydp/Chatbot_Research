package dialog

import (
	"strings"
	"util"
	"github.com/michlabs/fbbot"
	//"soap"
	//"os"
	//"os"
	//"fmt"
	"os"
)

type SelectTaisanBot struct {
	fbbot.BaseStep
}

func (s SelectTaisanBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	//if bot.STMemory.For(msg.Sender.ID).Get("HiddenMess") != "true" {
	//	bot.SendText(msg.Sender, util.Personalize(T("Welcome_bot"), &msg.Sender))
	//}

	CreateQuickReplyTypeOfMenuTaisan(bot, msg, util.Personalize(T("luachon_bot"), &msg.Sender))
	//bot.SendText(msg.Sender, util.Personalize(T("luachon_bot"), &msg.Sender))
	return StayEvent


}


func (s SelectTaisanBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	strMsg := msg.Text

	var isValid bool = false
	var optionValue string = ""
	stTypeOfTaisanOption := os.Getenv("TYPE_OF_TAISAN")
	arrTypeOfTaisanOption := strings.Split(stTypeOfTaisanOption, "|")
	for _, val := range arrTypeOfTaisanOption {
		arrItems := strings.Split(val, "-")
		if (arrItems[1] == strMsg) {
			isValid = true
			optionValue = arrItems[0]
			break
		}
	}

	if isValid == false {
		return s.Enter(bot, msg)
	}

	//fmt.Print(optionValue)

	switch optionValue {
	//case "so_hdbh":
	//	bot.STMemory.For(msg.Sender.ID).Set("Action","sodonInfo")
	//	return StayEvent
	case "ts_sodon":
		bot.STMemory.For(msg.Sender.ID).Set("Action","ts_sodon")
		return InputPolicyUrnFireEvent
	case "ts_mst":
		bot.STMemory.For(msg.Sender.ID).Set("Action","ts_mst")
		return Taisan_MST_Event
	case "ts_cmt":
		bot.STMemory.For(msg.Sender.ID).Set("Action","ts_cmt")
		return InputPolicyHolderIDFireEvent
	case "ts_location":
		bot.STMemory.For(msg.Sender.ID).Set("Action","ts_location")
		return InputLocationFireEvent
		//return InputLocationFireEvent
		//return InputLocationFireEvent_L2
	case "ts_tkh":
		bot.STMemory.For(msg.Sender.ID).Set("Action","ts_tkh")
		return InputPolicyHolderNameFireEvent
	}
	////Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
	//return StayEvent
}
