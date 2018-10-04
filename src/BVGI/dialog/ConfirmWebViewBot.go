package dialog

import (
	"github.com/michlabs/fbbot"
	"os"

	"util"
	"strings"

	"net/url"
	"fmt"
)

type ConfirmWebViewBot struct{
	fbbot.BaseStep
}

func (s ConfirmWebViewBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {

	////////////////////////////////////////////
	bot.STMemory.For(msg.Sender.ID).Get("skey")
	bot.STMemory.For(msg.Sender.ID).Get("svtype")

	var strkey = bot.STMemory.For(msg.Sender.ID).Get("skey")

	//strkey= PathEscape(strkey);

	var svtype = bot.STMemory.For(msg.Sender.ID).Get("svtype")

	surl :=  os.Getenv("WEBVIEW_URL")
	//fmt.Print(url)

	var strURL string

	switch strings.ToLower(util.RemoveAccent(svtype)) {
	  case "ts_sodon","ts_location","ts_tkh":
		  //strURL = url.PathEscape(string(surl+"/TaiSan.aspx?svtype="+svtype+"&id="+msg.Sender.ID+"&strkey="+strkey))
		  strURL = string(surl+"/TaiSan.aspx?svtype="+svtype+"&id="+msg.Sender.ID+"&strkey="+ url.QueryEscape(strkey))
	  case "biensoxe","sogcnbh","sok":
		  strURL = string(surl+"/XCG.aspx?svtype="+svtype+"&id="+msg.Sender.ID+"&strkey="+ url.QueryEscape(strkey))
	  case "cargo_sodon","cargo_kh","cargo_ten":
		  strURL = string(surl+"/Cargo.aspx?svtype="+svtype+"&id="+msg.Sender.ID+"&strkey="+ url.QueryEscape(strkey))
	  case "tau_sodon","tau_ten","tau_imo":
		  strURL = string(surl+"/Vessel.aspx?svtype="+svtype+"&id="+msg.Sender.ID+"&strkey="+url.QueryEscape(strkey))
	  case "cmt","mst":
		  strURL =string(surl+"/Common.aspx?svtype="+svtype+"&id="+msg.Sender.ID+"&strkey="+url.QueryEscape(strkey))

	}
	fmt.Print(strURL)
		///Common.aspx?svtype=CMT&strkey=1111%
	//var strURL string=string(url+"/Default.aspx?id="+msg.Sender.ID+"&strkey="+strkey)

	botExtraSelect := new(fbbot.ButtonMessage)
	botExtraSelect.AddWebURLButton("Tìm kiếm mở rộng", strURL)

	if (strings.ToLower(util.RemoveAccent(svtype))=="cmt"||strings.ToLower(util.RemoveAccent(svtype))=="mst"){
		botExtraSelect.AddPostbackButton(T("confirmBackHome_btn"),"backhome")
		botExtraSelect.Text = util.Personalize(T("getwelcome_webview"), &msg.Sender)
		bot.Send(msg.Sender, botExtraSelect)
	}else{

		botExtraSelect.AddPostbackButton(T("confirmReInput_btn"),"nhaplai")
		botExtraSelect.AddPostbackButton(T("confirmBackHome_btn"),"backhome")
		botExtraSelect.Text = util.Personalize(T("getwelcome_webview"), &msg.Sender)
		bot.Send(msg.Sender, botExtraSelect)
	}

	return StayEvent

}


func (s ConfirmWebViewBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	//if bot.LTMemory.For(msg.Sender.ID).Get("isFirstConversation") == "" {
	//	bot.LTMemory.For(msg.Sender.ID).Set("isFirstConversation", "true")
	//}
	bot.STMemory.For(msg.Sender.ID).Get("svtype")
	var svtype = bot.STMemory.For(msg.Sender.ID).Get("svtype")

	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case "nhaplai":
		bot.STMemory.For(msg.Sender.ID).Set("Action","nhaplai")

		switch strings.ToLower(util.RemoveAccent(svtype)) {
		case "biensoxe":
			return NhapBienSoXeEvent
		case "sogcnbh":
			return NhapGCNBHEvent
		case "sok":
			return InputChassisNoEvent
		case "ts_sodon":
			return InputPolicyUrnFireEvent
		case "ts_location":
			return InputLocationFireEvent
		case "ts_tkh":
			return InputPolicyHolderNameFireEvent
		case "tau_sodon":
			return InputPolicyUrnVesselEvent
		case "tau_ten":
			return InputVesselNameEvent
		case "tau_imo":
			return InputregNumberVesselEvent
		case "cargo_sodon":
			return InputPolicyUrnCargoEvent
		case "cargo_kh":
			return InputCargoCustomerNameEvent
		case "cargo_ten":
			return InputCargoNameEvent
		}

	case "backhome":
		bot.STMemory.For(msg.Sender.ID).Set("Action","backhome")
		return SelectBotEvent
	}
	return s.Enter(bot, msg)

}
