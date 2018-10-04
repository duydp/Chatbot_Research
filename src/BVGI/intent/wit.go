package intent

import (
	"fmt"
	"strings"
	
	log "github.com/Sirupsen/logrus"
	"github.com/michlabs/gowit"
	"fptai-sdk-go"
)

type Wit struct {
	client *gowit.Client
}

func InitWit(config map[string]string) (Engine, error) {
	var w Wit
	token := config["token"]
	if token == "" {
		return nil, fmt.Errorf("Wit token must not be empty")
	}
	w.client = gowit.NewClient(token)
	return &w, nil
}

func (wit *Wit) Detect(text string) string {
	meaning, err := wit.client.Detect(strings.ToLower(text))
	if err != nil {
		log.Errorf("Wit failed to detect intent. Error: %s. Text: %s\n", err.Error(), text)
		return ""
	}
	return meaning.Intent()
}

func (wit *Wit) AddUtterance(intent, utterance string) error {
	return wit.client.AddExpression(intent, strings.ToLower(utterance))
}

func (wit *Wit) AddIntentUtterances(ius []IntentUtterance) error {
	i, err := wit.client.GetEntity("intent")
	if err != nil {
		log.Error("Wit failed to get intent")
		return err
	}

	for _, iu := range ius {
		var v gowit.Value
		v.Name = iu.Intent
		v.Expressions = append(v.Expressions, strings.ToLower(iu.Utterance))
		i.Values = append(i.Values, v)
	}

	return wit.client.UpdateEntity(&i)
}

func (wit *Wit) DeleteAllIntents() error {
	return nil
}

func (wit *Wit) Train() error {
	return nil
}
func (wit *Wit) ExtractEntity(text string) (fptai.Meaning, error) {
	return fptai.Meaning{},nil
}