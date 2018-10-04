package dialog

import (
	"util"
	"BVGI/db"
	
	"github.com/michlabs/fbbot"
	log "github.com/Sirupsen/logrus"
)


type StaffRegister struct {
	fbbot.BaseStep
}

func (s StaffRegister) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	bot.SendText(msg.Sender, util.Personalize(T("register_staff_success"), &msg.Sender))

	staff := db.Staff{
		Fullname : msg.Sender.FirstName() + " " +  msg.Sender.LastName(),
		FbID : msg.Sender.ID,
	}

	err := db.InsertStaff(staff)
	if err != nil {
		log.Error(err.Error())
	}

	return GoodbyeEvent
}
