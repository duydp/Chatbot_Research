package dialog

import (
	//"strings"
	//"github.com/Sirupsen/logrus"
	"util"
	"github.com/michlabs/fbbot"
	"soap"
	"encoding/json"
	"fmt"
	//"github.com/leekchan/accounting"
	//"strconv"
	"strings"
)

type GetSoHDBot struct {
	fbbot.BaseStep
}

func (s GetSoHDBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	//mess := strings.ToLower(util.RemoveAccent(msg.Text))
	//logrus.Infof("mess %s compare %s %s \n", mess, strings.ToLower(util.RemoveAccent(T("select_human"))), strings.ToLower(util.RemoveAccent(T("select_bot"))))
	/*
	if mess != strings.ToLower(util.RemoveAccent(T("select_newsptt"))) && mess != strings.ToLower(util.RemoveAccent(T("select_newuvl")))&& mess != strings.ToLower(util.RemoveAccent(T("select_bosungtl"))){
		if bot.STMemory.For(msg.Sender.ID).Get("question") == "" && !strings.HasPrefix(msg.Text, "/"){
			bot.STMemory.For(msg.Sender.ID).Set("question", msg.Text)
		}
	}
	*/

	bot.SendText(msg.Sender, util.Personalize(T("getSoHD_title"), &msg.Sender))

	return StayEvent
}


func (s GetSoHDBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}
	var SoHD = strings.TrimSpace(msg.Text)
	result, _ :=  soap.GetInfoBySoHD(SoHD, msg.Sender.ID)

	if result.SearchCommonInfoResult=="OK" {
		bot.STMemory.For(msg.Sender.ID).Set("SoHD", SoHD)
		in := []byte(result.SearchCommonInfoResult)
		var raw map[string]interface{}
		json.Unmarshal(in, &raw)

		var msgText = `Thông tin Hợp đồng:
  - BMBH: ` + fmt.Sprintf("%s", raw["BMBH"]) + ``
  		if (raw["NDBH"].(string) != "") {
  			msgText += `- NDBH: ` + fmt.Sprintf("%s", raw["NDBH"])
  		}

		// Save info BMBH, NDBH to array
		stBMBHValue := bot.STMemory.For(msg.Sender.ID).Get("BMBHInfo")
		//bot.Logger.Infof("stringBMBH: ", stBMBHValue)

		var arrBMBH []string
		if stBMBHValue != "" {
			err := json.Unmarshal([]byte(stBMBHValue), &arrBMBH)
			if err != nil {}
		}

		if raw["BMBH"].(string) != "" {
			arrBMBH = append(arrBMBH, raw["BMBH"].(string))
		}
		if (raw["NDBH"].(string) != "" && raw["NDBH"].(string) != raw["BMBH"].(string)) {
			arrBMBH = append(arrBMBH, raw["NDBH"].(string))
		}
		// Luu lai vao memory
		jsonvar, _ := json.Marshal(arrBMBH)
		bot.STMemory.For(msg.Sender.ID).Set("BMBHInfo", string(jsonvar))
		//bot.Logger.Infof("json data",string(jsonvar))

		bot.STMemory.For(msg.Sender.ID).Set("SoHD", SoHD)

		bot.SendText(msg.Sender, msgText)
		bot.STMemory.For(msg.Sender.ID).Set("GetSoHDError","200")
		return ConfirmSoHDEvent
	} else {
		// Truong hop ko tim thay so HD trong he thong
		bot.SendText(msg.Sender, result.SearchCommonInfoResult)

		bot.STMemory.For(msg.Sender.ID).Set("SoHD", SoHD)
		bot.STMemory.For(msg.Sender.ID).Set("NotFoundSoHD", "true")
		return ConfirmSoHDEvent
	}

	return NhapSoHDEvent
}