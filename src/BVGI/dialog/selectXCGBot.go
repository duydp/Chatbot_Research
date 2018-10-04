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

type SelectXCGBot struct {
	fbbot.BaseStep
}

func (s SelectXCGBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	//if bot.STMemory.For(msg.Sender.ID).Get("HiddenMess") != "true" {
	//	bot.SendText(msg.Sender, util.Personalize(T("Welcome_bot"), &msg.Sender))
	//}

	CreateQuickReplyTypeOfMenu(bot, msg, util.Personalize(T("luachon_bot"), &msg.Sender))
	//bot.SendText(msg.Sender, util.Personalize(T("luachon_bot"), &msg.Sender))
	return StayEvent


}


func (s SelectXCGBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	strMsg := msg.Text

	var isValid bool = false
	var optionValue string = ""
	stTypeOfOption := os.Getenv("TYPE_OF_OPTION")
	arrTypeOfOption := strings.Split(stTypeOfOption, "|")
	for _, val := range arrTypeOfOption {
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
	case "biensoxe":
		bot.STMemory.For(msg.Sender.ID).Set("Action","biensoxe")
		return NhapBienSoXeEvent
	case "soGCNBH":
		bot.STMemory.For(msg.Sender.ID).Set("Action","soGCNBH")
		return NhapGCNBHEvent
	case "soK":
		bot.STMemory.For(msg.Sender.ID).Set("Action","soK")
		return InputChassisNoEvent
		//NhapCMTNDEvent
	}
	////Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
	//return StayEvent
}