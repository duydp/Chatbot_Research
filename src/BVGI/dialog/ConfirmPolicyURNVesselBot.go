//author: duydp
//wrote at: 28/3/2018
package dialog
import (
	"util"
	"github.com/michlabs/fbbot"
	"fmt"
	"strings"
)

type ConfirmPolicyURNVesselBot struct {
	fbbot.BaseStep
}

func (s ConfirmPolicyURNVesselBot) Enter(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {
	var items  []fbbot.QuickRepliesItem
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirmContinue_btn"),//nhap lai
		Payload:     "tieptuc",
	})

	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirmReInput_btn"),//nhap lai
		Payload:     "nhaplai",
	})
	items = append(items, fbbot.QuickRepliesItem{
		ContentType: "Text",
		Title:       T("confirmBackHome_btn"),//boqua
		Payload:     "backhome",
	})


	botSelect := new(fbbot.QuickRepliesMessage)

	botSelect.Text = util.Personalize(T("getAnswer_title"), &msg.Sender)


	botSelect.Items = items
	bot.Send(msg.Sender, botSelect)

	return StayEvent

}

func (s ConfirmPolicyURNVesselBot) Process(bot *fbbot.Bot, msg *fbbot.Message) fbbot.Event {


	fmt.Print(strings.ToLower(util.RemoveAccent(msg.Text)))

	switch strings.ToLower(util.RemoveAccent(msg.Text)) {
	case strings.ToLower(util.RemoveAccent(T("confirmContinue_btn"))): // Neu xac nhan dung soHD
		return InputPolicyUrnVesselEvent
	case strings.ToLower(util.RemoveAccent(T("confirmReInput_btn"))):
		return NhapTauThuyEvent
	case strings.ToLower(util.RemoveAccent(T("confirmBackHome_btn"))):
		return SelectBotEvent

	}
	//Khong nam trong cac lua chon, tu dong hoi lai NSD
	return s.Enter(bot, msg)



}