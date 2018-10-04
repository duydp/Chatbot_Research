package dialog

import (
	"util"
	"github.com/michlabs/fbbot"
	"strings"
	"soap"
	"encoding/json"
)

type GetInfoBySoGQQLBot struct {
	fbbot.BaseStep
}

func (s GetInfoBySoGQQLBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("getInfoSoBVGI_title"), &msg.Sender))

	return StayEvent
}


func (s GetInfoBySoGQQLBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	text := strings.TrimSpace(msg.Text)

	bot.STMemory.For(msg.Sender.ID).Set("SoGQQL", text)
	result, _ := soap.GetInfoBySoGQQL(text, msg.Sender.ID)

	if result==nil {
		bot.SendText(msg.Sender, util.Personalize(T("error"),&msg.Sender))
		return s.Enter(bot, msg)
	}
	if result.SearchCommonInfoResult == "OK" {
		var raw map[string]string
		json.Unmarshal([]byte(result.SearchCommonInfoResult), &raw)
		bot.STMemory.For(msg.Sender.ID).Set("TaskID",raw["TaskID"])
		//ac := accounting.Accounting{Symbol: "", Precision: 0}
		//phidk, _ := strconv.Atoi(raw["PhiDK"])
		var msgText = `Thông tin hồ sơ:
				- BMBH: ` + raw["BMBH"] + `
				- NBRR: ` + raw["NBRR"] + `
				- Số HĐ1: ` + raw["SOHD"] + `
				- Số HĐ2: ` + raw["SOHD2"] + `
				- Số HĐ3: ` + raw["SOHD3"]

		bot.SendText(msg.Sender, msgText)

		return ConfirmSoGQQLEvent
	} else {
		bot.SendText(msg.Sender,result.SearchCommonInfoResult)
		return s.Enter(bot, msg)
	}

	return GetInfoSoGQQLEvent
}