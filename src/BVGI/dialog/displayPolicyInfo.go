package dialog

import (
	"github.com/michlabs/fbbot"
	//"soap"
	"os"
	"strings"
	"util"
)

type DisplayPolicyInfoBot struct {
	fbbot.BaseStep
}

func (s DisplayPolicyInfoBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	var msgText = T("display_policy_title")

	if bot.STMemory.For(msg.Sender.ID).Get("Action")=="add document" {
		msgText = T("display_confirm_title")

	} else {
		msgText +=	"\r\n- Số HĐ1: " + bot.STMemory.For(msg.Sender.ID).Get("SoHD1")
		if (bot.STMemory.For(msg.Sender.ID).Get("SoHD2") != "") {
			msgText += "\r\n- Số HĐ2: " + bot.STMemory.For(msg.Sender.ID).Get("SoHD2")
		}
		if (bot.STMemory.For(msg.Sender.ID).Get("SoHD3") != "") {
			msgText += "\r\n- Số HĐ3: " + bot.STMemory.For(msg.Sender.ID).Get("SoHD3")
		}

		msgText += "\r\n- Người yêu cầu: " + bot.STMemory.For(msg.Sender.ID).Get("NguoiYC")
		msgText += "\r\n- Số ĐT: " + bot.STMemory.For(msg.Sender.ID).Get("SoDienThoai")
		msgText += "\r\n- Người rủi ro: " + bot.STMemory.For(msg.Sender.ID).Get("NguoiRuiRo")
		msgText += "\r\n- Loại rủi ro: " + bot.STMemory.For(msg.Sender.ID).Get("LoaiRuiRo")

		//bot.Logger.Infof("stringTypeOfRisk: ", bot.STMemory.For(msg.Sender.ID).Get("LoaiRuiRo"))

		stTypeOfRisk := os.Getenv("TYPE_OF_RISK")
		arrTypeOfRisk := strings.Split(stTypeOfRisk, "|")
		for _, val := range arrTypeOfRisk {
			arrItems := strings.Split(val, "-")
			if (arrItems[1] == bot.STMemory.For(msg.Sender.ID).Get("LoaiRuiRo")) {
				bot.STMemory.For(msg.Sender.ID).Set("MaLoaiRuiRo", arrItems[0])
				break
			}
		}

		msgText += "\r\n- Ghi chú: " + bot.STMemory.For(msg.Sender.ID).Get("GhiChu")
	}

	// Hien thi thong tin ho so
	bot.SendText(msg.Sender, msgText)

	// Confirm xac nhan hoan thanh
	var items []fbbot.QuickRepliesItem
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirm_complete_ok"),
		Payload:     "confirm_complete_ok",
		// ImageURL:    T("select_bot_icon"),
	})
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirm_complete_not"),
		Payload:     "confirm_complete_not",
		// ImageURL:    T("select_bot_icon"),
	})


	botSelect := new(fbbot.QuickRepliesMessage)
	botSelect.Text = util.Personalize(T("confirmComplete_title"), &msg.Sender)
	botSelect.Items = items
	bot.Send(msg.Sender, botSelect)

	return StayEvent
}


func (s DisplayPolicyInfoBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case strings.ToLower(util.RemoveAccent(T("confirm_complete_ok"))):
		return CreateRequestEvent

		// Chuyen sang buoc nhap thong tin: Chon nguoi yeu cau
	case strings.ToLower(util.RemoveAccent(T("confirm_complete_not"))):
		bot.STMemory.For(msg.Sender.ID).Set("SoHDCount", "0")
		bot.STMemory.For(msg.Sender.ID).Set("SoHD", "")
		bot.STMemory.For(msg.Sender.ID).Set("SoHD1", "")
		bot.STMemory.For(msg.Sender.ID).Set("SoHD2", "")
		bot.STMemory.For(msg.Sender.ID).Set("SoHD3", "")
		return NhapSoHDEvent
	}
	//Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)

}
