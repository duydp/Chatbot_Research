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

type SelectTauthuyBot struct {
	fbbot.BaseStep
}

func (s SelectTauthuyBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	CreateQuickReplyTypeOfMenuTauthuy(bot, msg, util.Personalize(T("luachon_bot"), &msg.Sender))
	//bot.SendText(msg.Sender, util.Personalize(T("luachon_bot"), &msg.Sender))
	return StayEvent

}


func (s SelectTauthuyBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	strMsg := msg.Text

	var isValid bool = false
	var optionValue string = ""
	stTypeOfTauOption := os.Getenv("TYPE_OF_TAUTHUY")
	arrTypeOfTauOption := strings.Split(stTypeOfTauOption, "|")
	for _, val := range arrTypeOfTauOption {
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
	case "tau_sodon":
		bot.STMemory.For(msg.Sender.ID).Set("Action","tau_sodon")
		return InputPolicyUrnVesselEvent
	case "tau_ten":
		bot.STMemory.For(msg.Sender.ID).Set("Action","tau_ten")
		return InputVesselNameEvent
	case "tau_IMO":
		bot.STMemory.For(msg.Sender.ID).Set("Action","tau_IMO")
		return InputregNumberVesselEvent
	}
	////Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
	//return StayEvent
}
