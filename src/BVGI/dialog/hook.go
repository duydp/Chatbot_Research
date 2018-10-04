package dialog

import (
	"strings"
	log "github.com/Sirupsen/logrus"
	"BVGI/config"
	"util"
	"github.com/michlabs/fbbot"

)

func checkTiming(bot *fbbot.Bot, msg *fbbot.Message) bool {
	lastEcho := bot.STMemory.For(msg.Sender.ID).Get("lastEcho")

	if lastEcho != "" && !TimeExpired(lastEcho, config.Bot.SilenceDuration) {
		log.Debugf("Bot do nothing since staff is chatting, lastEcho = %s", util.ToTime(lastEcho).Format("15:04:05"))
		return true
	}

	lastActiveTime := bot.STMemory.For(msg.Sender.ID).Get("lastMessage")

	if TimeExpired(lastActiveTime, config.Bot.ConversationTimeout) {
		log.Debug("Dialog is expired, move to Welcome state")
		insujs.Move(msg, SelectBot{})
		return true
	}

	return false
}

func AdHocPostBack(bot *fbbot.Bot, msg *fbbot.Message) bool {


	log.Infof("\n msg = %s \n", msg.Text)
	log.Infof("\n len = %d %d %d %d \n", len(msg.Audios), len(msg.Files), len(msg.Images), len(msg.Videos))
	//bot.Logger.Infof("text data:", msg.Text)

	// Kiem tra neu la cu phap tra cuu
	strMsg := strings.ToUpper(msg.Text)
	//strings.HasPrefix("foobar", "foo")

	if (strings.HasPrefix(strMsg, "CAR ") || strings.HasPrefix(strMsg, "GCN ")|| strings.HasPrefix(strMsg, "SK ") || strings.HasPrefix(strMsg, "CMT ") || strings.HasPrefix(strMsg, "MST ") || strings.HasPrefix(strMsg, "SD ") || strings.HasPrefix(strMsg, "TAU ") || strings.HasPrefix(strMsg, "HH ") || strings.HasPrefix(strMsg, "IMO ")|| strings.HasPrefix(strMsg, "DK ")|| strings.HasPrefix(strMsg, "DD ")|| strings.HasPrefix(strMsg, "KH ")|| strings.HasPrefix(strMsg, "HD ")|| strings.HasPrefix(strMsg, "TEN ")) {

		var strSyntax = strings.Replace(strMsg, "CAR ", "", 1)
		strSyntax = strings.Replace(strSyntax, "GCN ", "", 1)
		strSyntax = strings.Replace(strSyntax, "SK ", "", 1)
		strSyntax = strings.Replace(strSyntax, "CMT ", "", 1)
		strSyntax = strings.Replace(strSyntax, "MST ", "", 1)
		strSyntax = strings.Replace(strSyntax, "SD ", "", 1)
		strSyntax = strings.Replace(strSyntax, "TAU ", "", 1)
		strSyntax = strings.Replace(strSyntax, "HH ", "", 1)
		strSyntax = strings.Replace(strSyntax, "IMO ", "", 1)
		strSyntax = strings.Replace(strSyntax, "TEN ", "", 1)
		strSyntax = strings.Replace(strSyntax, "DK ", "", 1)
		strSyntax = strings.Replace(strSyntax, "DD ", "", 1)
		strSyntax = strings.Replace(strSyntax, "KH ", "", 1)
		strSyntax = strings.Replace(strSyntax, "HD ", "", 1)


		bot.STMemory.For(msg.Sender.ID).Set("KEY_BSX", strings.Trim(strSyntax, " "))
		bot.STMemory.For(msg.Sender.ID).Set("Redirect_From_Hook","true")
		insujs.Move(msg, GetInfoBot{})
	}

	//fmt.Print(msg.Text)

	switch strings.ToLower(msg.Text) {
	case "/restart":
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "hook")
		insujs.Move(msg, Goodbye{})
		bot.STMemory.For(msg.Sender.ID).Delete("lastEcho")
		bot.STMemory.For(msg.Sender.ID).Delete("lastMessage")
		return true
	case "get_started_payload":
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "hook")
		insujs.Move(msg, Goodbye{})
		bot.STMemory.For(msg.Sender.ID).Delete("lastEcho")
		bot.STMemory.For(msg.Sender.ID).Delete("lastMessage")
		return true
	case "/bot":
		bot.STMemory.For(msg.Sender.ID).Set("previous_step", "hook")
		insujs.Move(msg, Welcome{})
		bot.STMemory.For(msg.Sender.ID).Delete("lastEcho")
		bot.STMemory.For(msg.Sender.ID).Delete("lastMessage")
		return true
	case "/enter_bsx":
		insujs.Move(msg, WelcomeBot{})
		return true
	case strings.ToLower(config.Bot.Hex):
		insujs.Move(msg, StaffRegister{})
		return true
	}

	return false
}

func PreHandlerMessageHook(bot *fbbot.Bot, msg *fbbot.Message) bool{
	return AdHocPostBack(bot, msg) || checkTiming(bot, msg)
}