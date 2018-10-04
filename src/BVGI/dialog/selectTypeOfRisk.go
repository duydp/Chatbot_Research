package dialog

import (
	"github.com/michlabs/fbbot"
	//"soap"
	"os"
	"strings"
)

type SelecTypeOfRiskBot struct {
	fbbot.BaseStep
}

func (s SelecTypeOfRiskBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	CreateQuickReplyTypeOfRisk(bot, msg, T("select_loairuiro_title"))

	return StayEvent
}


func (s SelecTypeOfRiskBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	LoaiRuiRo := msg.Text

	// Check loai rui ro phai chon tu list
	var isValid bool = false
	stTypeOfRisk := os.Getenv("TYPE_OF_RISK")
	arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")
	for _, val := range arrTypeOfRisk {
		arrItems := strings.Split(val, "-")
		if (arrItems[1] == LoaiRuiRo) {
			isValid = true
			break
		}
	}

	if isValid == false {
		bot.SendText(msg.Sender, T("select_loairuiro_error"))
		return s.Enter(bot, msg)
	}


	bot.STMemory.For(msg.Sender.ID).Set("LoaiRuiRo", LoaiRuiRo)

	return ImageDocDialogEvent
}
