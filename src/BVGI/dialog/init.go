package dialog

import (
	"github.com/michlabs/bottext"
)

var T bottext.BotTextFunc

func Init(languageFP string) {
	bottext.MustLoad(languageFP)
	T = bottext.New("vi")
}
