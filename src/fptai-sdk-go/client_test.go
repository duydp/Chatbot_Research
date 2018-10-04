package fptai

import (
	"log"
	"os"
	"testing"
	"time"
)

type IntentTest struct {
	Name string
	Description string
	Utterances []string
}

var client *Client

var intents []IntentTest = []IntentTest{
	{
		Name:        "go_swimming",
		Description: "description for go_swimming intent",
		Utterances:  []string{"i want to go to swimming", "swimming is great"},
	},
	{
		Name:        "play_football",
		Description: "description for play_footbal intent",
		Utterances:  []string{"i want to play football", "football please"},
	},
}

func init() {
	token := os.Getenv("token")
	if token == "" {
		log.Fatal("getting token failed")
	}

	client = NewClient(token)
}

func TestCreateIntent(t *testing.T) {
	for _, i := range intents {
		_, err := client.CreateIntent(i.Name, i.Description)
		if err != nil {
			t.Error(err)
			return
		}
	}
	t.Run("CreateUtterances", testCreateUtterances)
	t.Run("GetIntents", testGetIntents)
	t.Run("TrainIntent", testTrainIntent)
	time.Sleep(3) // wait for training
	t.Run("RecognizeIntent", testRecognizeIntents)
	t.Run("DeleteIntent", testDeleteIntent)
}

func testGetIntents(t *testing.T) {
	intents, err := client.GetIntents()
	if err != nil {
		t.Error(err)
		return
	}
	if len(intents) != 2 {
		t.Error("missing some intents")
		return
	}
}

func testDeleteIntent(t *testing.T) {
	for _, i := range intents {
		err := client.DeleteIntent(i.Name)
		if err != nil {
			t.Error(err)
		}
	}
}

func testCreateUtterances(t *testing.T) {
	for _, i := range intents {
		err := client.CreateUtterances(i.Name, i.Utterances)
		if err != nil {
			t.Error(err)
		}
	}
}

func testTrainIntent(t *testing.T) {
	err := client.TrainIntent()
	if err != nil {
		t.Error(err)
	}
}

func testRecognizeIntents(t *testing.T) {
	for _, i := range intents {
		for _, u := range i.Utterances {
			m, err := client.RecognizeIntents(u)
			if err != nil {
				t.Error(err)
				return
			}
			if len(m.Intents) == 0 {
				t.Error("recognize failed")
				return
			}
			if m.Intents[0].Name != i.Name {
				t.Error("recognize failed")
				return
			}
		}
	}
}