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
)

type DevelopBot struct {
	fbbot.BaseStep
}

func (s DevelopBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {


	var items []fbbot.QuickRepliesItem
	//items = append(items, fbbot.QuickRepliesItem{
	//	ContentType: "Text",
	//	Title:       T("confirm_continue"),
	//	Payload:     "confirm_continue",
	//	// ImageURL:    T("select_bot_icon"),
	//})
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirm_back_home"),
		Payload:     "confirm_back_home",
		// ImageURL:    T("select_bot_icon"),
	})


	botSelect := new(fbbot.QuickRepliesMessage)

	//if bot.STMemory.For(msg.Sender.ID).Get("SyntaxInvalid") == "true" {
	//	botSelect.Text = util.Personalize("Bạn nhập sai cú pháp. Vui lòng đúng cú pháp!", &msg.Sender)
	//
	//}else{
	//	botSelect.Text = util.Personalize("Bạn có muốn tiếp tục tra cứu không?", &msg.Sender)
	//}


	botSelect.Text = util.Personalize(T("getAlert_title"), &msg.Sender)

	botSelect.Items = items
	bot.Send(msg.Sender, botSelect)

	return StayEvent
}


func (s DevelopBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {


	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	//case strings.ToLower(util.RemoveAccent(T("confirm_continue"))): // Neu xac nhan dung soHD
	//	return TracuuSyntaxEvent
	case strings.ToLower(util.RemoveAccent(T("confirm_back_home"))):
		bot.STMemory.For(msg.Sender.ID).Set("HiddenMess", "true")
		return SelectBotEvent
	}
	//Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
}