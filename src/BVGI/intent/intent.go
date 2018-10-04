package intent
import (
	"strings"
	"BVGI/db"
	"fptai-sdk-go"
)

type IntentUtterance struct {
	Intent string `json:"intent"`
	Utterance string `json:"utterance"`
}

// Detect returns intent of a text. 
func Detect(text string) string {
	i := engine.Detect(strings.ToLower(text))

	h := db.History{
		Question: text,
		Answer: i,
	}
	db.InsertHistory(h)

	return i
}

func AddUtterance(intent, utterance string) error {
	return engine.AddUtterance(intent, utterance)
}

func AddIntentUtterances(ius []IntentUtterance) error {
	return engine.AddIntentUtterances(ius)
}
func DeleteAllIntents() error {
	return engine.DeleteAllIntents()
}

func Train() error {
	return engine.Train()
}

func GetAll() ([]string, error) {
	return db.GetAllIntents()
}

func ExtractEntity(text string) (fptai.Meaning, error) {
	return engine.ExtractEntity(text)
}