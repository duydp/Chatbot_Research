package dialog

import (
	"strings"
	//"github.com/Sirupsen/logrus"
	"util"
	"github.com/michlabs/fbbot"
	//"soap"
	//"os"
	//"soap"
	//"fmt"

	"fmt"

	"os"
)

type WelcomeBot struct {
	fbbot.BaseStep
}

func (s WelcomeBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	//fmt.Print(strings.ToUpper(msg.Sender.Name()))

	stFbIDs:= os.Getenv("TYPE_OF_FBID")
	var check_prefix =strings.Contains(strings.ToUpper(msg.Sender.Name()),`BVGI`)

	if (check_prefix == true || strings.Contains(stFbIDs, msg.Sender.ID)) {
		if bot.STMemory.For(msg.Sender.ID).Get("HiddenMess") != "true" {
			bot.SendText(msg.Sender, util.Personalize(T("Welcome_bot"), &msg.Sender))
		}
		botSelect := new(fbbot.ButtonMessage)
		botSelect.AddPostbackButton(T("select_xecogioi"),"xecogioi")
		botSelect.AddPostbackButton(T("select_taisan"),"taisan")
		botSelect.AddPostbackButton(T("select_hanghoa"),"hanghoa")

		botSelect.Text = util.Personalize(T("say_hello_bot"), &msg.Sender)
		bot.Send(msg.Sender, botSelect)

		// THUCDH Bổ sung
		botExtraSelect := new(fbbot.ButtonMessage)
		botExtraSelect.AddPostbackButton(T("select_tauthuy"), "tauthuy")
		//botExtraSelect.AddPostbackButton(T("select_bagdv"), "bagdv")
		//select_bagdv
		botExtraSelect.AddPostbackButton(T("select_tracuunhanh"), "syntax")

		//url :=  os.Getenv("WEBVIEW_URL")
		//var strURL string
		//strURL=string(url+"?id="+msg.Sender.ID+"&strkey=")

		//botExtraSelect.AddWebURLButton("Tìm kiếm địa điểm", strURL)

		botExtraSelect.Text = "..."
		bot.Send(msg.Sender, botExtraSelect)
	} else {
		bot.SendText(msg.Sender, util.Personalize(T("Welcome_bot_fail"), &msg.Sender))
	}

	//CreateQuickReplyTypeOfMenu(bot, msg, util.Personalize(T("luachon_bot"), &msg.Sender))

	return StayEvent
}


func (s WelcomeBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
		bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	}

	///fmt.Print(msg.Text)

	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case "xecogioi":
		bot.STMemory.For(msg.Sender.ID).Set("Action","xecogioi")
		return SelectXCGBotEvent
	case "taisan":
		bot.STMemory.For(msg.Sender.ID).Set("Action","taisan")
		return NhapTaiSanEvent
	case "hanghoa":
		bot.STMemory.For(msg.Sender.ID).Set("Action","hanghoa")
		return NhapHangEvent
	case "tauthuy":
		bot.STMemory.For(msg.Sender.ID).Set("Action","tauthuy")
		return NhapTauThuyEvent
	/*case "bagdv":
		bot.STMemory.For(msg.Sender.ID).Set("Action","bagdv")
		return InputBAGDVEvent*/
	case "syntax":
		bot.STMemory.For(msg.Sender.ID).Set("Action","syntax")
		return TracuuSyntaxEvent
	 default://when user used to syntax at any areas on bot
	 	fmt.Println(msg.Text)
		bot.STMemory.For(msg.Sender.ID).Set("SyntaxInvalid", "true")
		return ConfirmSyntaxEvent

	}
	//Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)
}