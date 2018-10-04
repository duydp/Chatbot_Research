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

type SelectHangBot struct {
	fbbot.BaseStep
}

func (s SelectHangBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	//if bot.STMemory.For(msg.Sender.ID).Get("HiddenMess") != "true" {
	//	bot.SendText(msg.Sender, util.Personalize(T("Welcome_bot"), &msg.Sender))
	//}

	CreateQuickReplyTypeOfMenuCargo(bot, msg, util.Personalize(T("luachon_bot"), &msg.Sender))
	//bot.SendText(msg.Sender, util.Personalize(T("luachon_bot"), &msg.Sender))
	return StayEvent
}


func (s SelectHangBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	strMsg := msg.Text

	var isValid bool = false
	var optionValue string = ""
	stTypeOfCargoOption := os.Getenv("TYPE_OF_CARGO")
	arrTypeOfCargoOption := strings.Split(stTypeOfCargoOption, "|")
	for _, val := range arrTypeOfCargoOption {
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
	case "cargo_sodon":
		bot.STMemory.For(msg.Sender.ID).Set("Action","cargo_sodon")
		return InputPolicyUrnCargoEvent
	case "cargo_ten":
		bot.STMemory.For(msg.Sender.ID).Set("Action","cargo_ten")
		return InputCargoNameEvent
	case "cargo_kh":
		bot.STMemory.For(msg.Sender.ID).Set("Action","cargo_kh")
		return InputCargoCustomerNameEvent
	/*case "cargo_sodk":
		bot.STMemory.For(msg.Sender.ID).Set("Action","cargo_sodk")
		return InputregNumberCargoEvent*/
	}
	////Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
	//return StayEvent
}
