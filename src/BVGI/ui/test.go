package ui

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"BVGI/db"
	"BVGI/intent"
)

func UnderstadingViewRender(w http.ResponseWriter, r *http.Request) {
	var data = make(map[string]interface{})

	intents, err := intent.GetAll()
	if err != nil {
		log.Println("Get all intents error " + err.Error())
	}

	data["intents"] = intents
	tplHelper.Render(w, "test", data)
}

func CorrectQuestionTest(intentName, question string) {
	log.Infof("Recorrect intent: %s, question: %s \n", intentName, question)

	if err := db.DeleteOldQuestion(question); err != nil {
		log.Error("Error in Delete old question CorrectQuestionHistory ", err)
		return
	}

	if err := db.InsertNewQuestion(intentName, question); err != nil {
		log.Error("Error in CorrectQuestion ", err)
		return
	}

	err := intent.AddUtterance(intentName, question)
	if err != nil {
		log.Error("Error in CorrectQuestion ", err)
		return
	}

	log.Info("Done.")
}

func CorrectQuestionHistory(intentName string, question string) {
	log.Infof("Recorrect intent: %s, question: %s \n", intentName, question)

	if err := db.DeleteOldQuestion(question); err != nil {
		log.Error("Error in Delete old question CorrectQuestionHistory ", err)
		return
	}


	if err := db.InsertNewQuestion(intentName, question); err != nil {
		log.Error("Error in Insert new question CorrectQuestionHistory ", err)
		return
	}

	if err := db.UpdateHistoryAnswer(question, intentName); err != nil {
		log.Error("Error in update history CorrectQuestionHistory ", err)
		return
	}

	err := intent.AddUtterance(intentName, question)
	if err != nil {
		log.Error("Error in add utterance CorrectQuestionHistory ", err)
		return
	}

	log.Info("Done.")
}

func DetectIntent(w http.ResponseWriter, r * http.Request) {
	question := r.FormValue("question")
	if question == "" {
		fmt.Fprintln(w, "question must not null")
		return
	}

	intent := intent.Detect(question)
	//log.Errorf("Intent %s\n", intent)
	fmt.Fprintln(w, intent)
}

func CorrectQuestionTestHandler(w http.ResponseWriter, r * http.Request) {
	question := r.FormValue("question")
	intent   := r.FormValue("intent")

	log.Errorf("question %s, intent %s", question, intent)
	if question == "" || intent == "_" {
		fmt.Fprintln(w, "question and intent must not null")
		return
	}

	CorrectQuestionTest(intent, question)

	fmt.Fprintln(w, "Updated success")
}

func CorrectQuestionHistoryHandler(w http.ResponseWriter, r * http.Request) {
	question := r.FormValue("question")
	intent 	 := r.FormValue("intent")

	question = strings.TrimSpace(question)
	intent   = strings.TrimSpace(intent)

	log.Errorf("question %s, intent %s \n", question, intent)
	if question == "" || intent == "_"{
		fmt.Fprintln(w, "question and intent must not null")
		return
	}

	go CorrectQuestionHistory(intent, question)

	fmt.Fprintln(w, "Updated success")
}