package dialog

import (
	"strings"
	"util"
	"github.com/michlabs/fbbot"
	"os"
)

type SelectBAGDVBot struct {
	fbbot.BaseStep
}

func (s SelectBAGDVBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	//if bot.STMemory.For(msg.Sender.ID).Get("HiddenMess") != "true" {
	//	bot.SendText(msg.Sender, util.Personalize(T("Welcome_bot"), &msg.Sender))
	//}

	CreateQuickReplyTypeOfMenuBAGDV(bot, msg, util.Personalize(T("luachon_bot"), &msg.Sender))
	//bot.SendText(msg.Sender, util.Personalize(T("luachon_bot"), &msg.Sender))
	return StayEvent


}


func (s SelectBAGDVBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	strMsg := msg.Text

	var isValid bool = false
	var optionValue string = ""
	stTypeOfBAGDVOption := os.Getenv("TYPE_OF_BAGDV")
	arrTypeOfBAGDVOption := strings.Split(stTypeOfBAGDVOption, "|")
	for _, val := range arrTypeOfBAGDVOption {
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

	case "bagdv_hdbh":
		bot.STMemory.For(msg.Sender.ID).Set("Action","bagdv_hdbh")
		return InputContractNo_BAGDV_Event
	case "bagdv_kh":
		bot.STMemory.For(msg.Sender.ID).Set("Action","bagdv_kh")
		return InputPolicyHolderName_BAGDV_Event

	}
	////Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
	//return StayEvent
}
