package dialog

import (
	//"strings"
	//"github.com/sirupsen/logrus"
	"util"
	"github.com/michlabs/fbbot"
	//"strconv"
	"strings"
	//"strconv"
	//"util/debug"
	//"fmt"
)

type ConfirmSyntaxBot struct {
	fbbot.BaseStep
}

func (s ConfirmSyntaxBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	//var KetquaAction  = bot.STMemory.For(msg.Sender.ID).Get("Action")
	//fmt.Print(" Kq Action: " + KetquaAction)


	var items []fbbot.QuickRepliesItem
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirm_continue"),
		Payload:     "confirm_continue",
		// ImageURL:    T("select_bot_icon"),
	})
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirm_back_home"),
		Payload:     "confirm_back_home",
		// ImageURL:    T("select_bot_icon"),
	})


	botSelect := new(fbbot.QuickRepliesMessage)
	if bot.STMemory.For(msg.Sender.ID).Get("SyntaxInvalid") == "true" {
		botSelect.Text = util.Personalize("Bạn nhập sai cú pháp. Vui lòng đúng cú pháp!", &msg.Sender)

	}else{
		botSelect.Text = util.Personalize("Bạn có muốn tiếp tục tra cứu không?", &msg.Sender)
	}



	botSelect.Items = items
	bot.Send(msg.Sender, botSelect)

	return StayEvent
}


func (s ConfirmSyntaxBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {



	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case strings.ToLower(util.RemoveAccent(T("confirm_continue"))): // Neu xac nhan dung soHD
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="xecogioi"{
			return SelectXCGBotEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="taisan"{
			return NhapTaiSanEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="hanghoa"{
			return NhapHangEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="tauthuy"{
			return NhapTauThuyEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="syntax"{
			return TracuuSyntaxEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="biensoxe"{
			return NhapBienSoXeEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="soGCNBH"{
			return NhapGCNBHEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="bagdv"{
			return InputBAGDVEvent
		}
		if bot.STMemory.For(msg.Sender.ID).Get("Action")=="soCMT"{
			return NhapCMTNDEvent
		} else{
			return TracuuSyntaxEvent
		}

	case strings.ToLower(util.RemoveAccent(T("confirm_back_home"))):
		bot.STMemory.For(msg.Sender.ID).Set("HiddenMess", "true")
		return SelectBotEvent
	}
	//Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
}