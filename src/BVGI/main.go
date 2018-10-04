package main
//D1D606D57489A58D3E26DDF5FD2D3E00
import (
	"os"
	log "github.com/Sirupsen/logrus"
	"BVGI/config"
	"BVGI/db"
	"BVGI/dialog"
	"BVGI/intent"
	"BVGI/ui"
	"github.com/michlabs/fbbot"
)

var insujs *fbbot.Dialog

func init() {
	if err := config.LoadFromEnv(); err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	if err := db.Init(&config.DB); err != nil {
		log.Fatal("failed to connect to db: ", err)
	}

	if err := intent.Init("fptai", config.Bot.FPTAI); err != nil {
	//if err := intent.Init("wit", config.Bot.Wit); err != nil {
		log.Fatal("failed to init intent package: ", err)
	}
	dialog.Init(config.Bot.LanguageFile)
}

func main() {
	ui.Init(&config.UI)
	go ui.Run()

	insujs = dialog.New()

	bot := fbbot.New(config.Bot.Port, config.Bot.VerifyToken, config.Bot.PageAccessToken)

	tracker := new(dialog.ActivityTracker)

	bot.AddMessageHandler(tracker)
	bot.AddEchoHandler(tracker)
	bot.AddPostbackHandler(tracker)

	bot.AddMessageHandler(insujs)
	bot.AddPostbackHandler(insujs)

	// Minhnt: Add menu
	bot.AddGetStartedButton()
	bot.AddMenuPersistent()

	switch config.Bot.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	logWriter, err := os.OpenFile(config.Bot.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logWriter)

	// debugLogWriter, err := os.OpenFile(config.Bot.DebugFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer debugLogWriter.Close()
	// debug.Init(bot, debugLogWriter)

	bot.Run()
}
