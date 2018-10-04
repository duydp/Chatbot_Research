package dialog

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"BVGI/config"
	"util"
	"github.com/michlabs/fbbot"
)

type Silence struct {
	fbbot.BaseStep
}

func (s Silence) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("wait_for_staff"), &msg.Sender))

	bot.STMemory.For(msg.Sender.ID).Set("lastEcho", util.FromTime(time.Now()))
	lastEcho := util.ToTime(bot.STMemory.For(msg.Sender.ID).Get("lastEcho"))
	log.Debugf("Bot entered silence, lastEcho = %s", lastEcho.Format("15:04:05"))

	CallStaffs(bot, msg)

	return StayEvent
}

func (s Silence) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	lastEcho := bot.STMemory.For(msg.Sender.ID).Get("lastEcho")
	if TimeExpired(lastEcho, config.Bot.SilenceDuration) {
		log.Debugf("Bot was waken up, lastEcho = %s", util.ToTime(lastEcho).Format("15:04:05"))
		return GoWelcomeEvent
	}else{
		log.Debugf("Bot in silence mode, lastEcho = %s", util.ToTime(lastEcho).Format("15:04:05"))
		return StayEvent
	}
}

func TimeExpired(lastTimeStr string, duration float64) bool {
	currentTime := time.Now()
	if lastTimeStr == "" {
		return false
	}
	lastTime := util.ToTime(lastTimeStr)

	if float64(currentTime.Sub(lastTime).Minutes()) < duration {
		return false
	}

	return true
}