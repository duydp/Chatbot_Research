package ui

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"BVGI/db"
	"BVGI/intent"
	"github.com/gorilla/mux"
)

func ListFAQsHandler(w http.ResponseWriter, r *http.Request) {
	faqs, err := db.GetAllFAQs()
	if err != nil {
		log.Error("failed to get FAQs: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["faqs"] = faqs

	tplHelper.Render(w, "faq", result)

	return
}

func NewFAQHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result = make(map[string]interface{})
		result["name"] = "new faq"
		tplHelper.Render(w, "faq_new", result)	
		return
	}

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Error("failed to read request NewFAQ ", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
    }
   	
    var faq db.FAQ
    err = json.Unmarshal(body, &faq)
    if err != nil {
        log.Errorf("failed to Unmarshal New Faq request. Error = %s. Body = %s \n", err, string(body))
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
    }

	if faq.IntentName == "" || faq.Answer == "" || len(faq.Questions) == 0 {
		log.Error("IntentName or Answer of Questions not found")
		fmt.Fprintln(w, "IntentName or Answer of Questions not found")
		
		return
	}

	exist, err := db.IsIntentExist(faq.IntentName)
	if err != nil {
		log.Error("Error in check exist intent")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if exist {
		log.Error("Intent has exist")
		fmt.Fprintln(w, "Intent has Exist")
		return
	}

	if err := db.InsertFAQ(faq); err != nil {
		log.Error("Error in insert faq. Error: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var ius []intent.IntentUtterance
	for i, v := range faq.Questions {
		ius = append(ius, intent.IntentUtterance{faq.IntentName, v.Question})
		log.Infof("Add utterance #%d (%s, %s) \n", i+1, faq.IntentName, v.Question)
	}

	go func(){
		err = intent.AddIntentUtterances(ius)
		if err != nil {
			log.Errorf("Failed to insert %d utterances to FPT.AI\n", len(ius))
		}

		log.Infof("Inserted %d utterances to FPT.AI\n", len(ius))
	}()

	time.Sleep(1 * time.Second)

	fmt.Fprintln(w, "success")
}

//Render FAQ edit view
func FAQEditViewRender(w http.ResponseWriter, r *http.Request) {
	sID := r.FormValue("id")
	if sID == "" {
		log.Error("Form value not have id")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(sID)
	if err != nil {
		log.Error("Error convert Id from string to int " + err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	faq, err := db.GetFAQByID(id)
	if err == NotFound {
		log.Error("Get FAQ error " + err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//log.Error(faq)

	faqs := []db.FAQ{faq}

	var result = make(map[string]interface{})

	result["faq"] = faqs

	tplHelper.Render(w, "edit", result)
}

func UpdateNewUtterances(faq, faqOld db.FAQ, w http.ResponseWriter)  {
	var oldFaqSet map[string]bool
	oldFaqSet = make(map[string]bool)

	for _, v := range faqOld.Questions {
		v.Question = strings.TrimSpace(v.Question)
		oldFaqSet[strings.ToLower(v.Question)] = true
	}

	var ius []intent.IntentUtterance

	counter := 0
	for _, v := range faq.Questions {
		v.Question = strings.TrimSpace(v.Question)
		if !oldFaqSet[strings.ToLower(v.Question)] {
			ius = append(ius, intent.IntentUtterance{faq.IntentName, v.Question})
			log.Infof("Add utterance #%d (%s, %s) \n", counter, faq.IntentName, v.Question)
			counter = counter + 1
		}
	}

	err := intent.AddIntentUtterances(ius)
	if err != nil {
		log.Errorf("Failed to insert %d utterance\n", len(ius))
	}
}

//Handler FAQ edit
func FAQEditHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        log.Error("failed to read request edit faq ", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
        return
    }
   	
    var faq db.FAQ
    err = json.Unmarshal(body, &faq)
    if err != nil {
        log.Error("failed to Unmarshal edit Faq request ", err)
        http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
    }

	faqOld, err := db.GetFAQByIntent(faq.IntentName)
	if err != nil {
		log.Error("Error when get faq by intent in edit handler")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := db.UpdateFAQ(faq); err != nil {
		log.Println("Error when update faq ", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	go UpdateNewUtterances(faq, faqOld, w)

	time.Sleep(1 * time.Second)
	fmt.Fprintln(w, "success")
}

//Delete a question of a FAQ
func DeleteQuestionHandler(w http.ResponseWriter, r *http.Request) {
	questionId := r.FormValue("id")
	if questionId == "" {
		fmt.Fprintln(w, "not found id")
		return
	}

	qId, err := strconv.Atoi(questionId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := db.DeleteQuestion(qId); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "success")
}

//Delete a FAQ
func DeleteFAQHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmp_id := vars["id"]
	id, err := strconv.Atoi(tmp_id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := db.DeleteFAQ(id); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<tr></tr>")
	
}

//Search FAQ by intent service
func SearchFAQByIntentHandler(w http.ResponseWriter, r *http.Request) {
	intent := r.FormValue("intent")
	if intent == "" {
		log.Error("Form value not have intent")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var faqs []db.FAQ = []db.FAQ{}
	//faqs = make([]FAQ, 1, 1)

	faq, err := db.GetFAQByIntent(intent)
	if err == NotFound {
		log.Error("Not found FAQ " + err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	} else if err != nil {
		fmt.Fprintln(w, "Internal Error")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	} else {
		faqs = append(faqs, faq)
	}

	result := make(map[string]interface{})
	result["faq"] = faqs

	json, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, string(json))
}

//Function to retrain all faqs
func FAQTrain(w http.ResponseWriter, r *http.Request) {
	log.Info("Retraining model........")
	isTraining = true
	err := intent.Train()
	if err != nil {
		log.Error("Error in FAQ Retrain: ", err)
		isTraining = false
		return
	}

	time.Sleep(30 * time.Second) //TODO: Remove later hungnq: work around because lacking of checking training status api
	isTraining = false
	log.Info("Done")
	fmt.Fprintln(w, "Done")
}

//Function to retrain all faqs
func FAQTrainAll(w http.ResponseWriter) {
	log.Info("Retraining all........")
	isTraining = true

	faqs, err := db.GetAllFAQs()
	if err != nil {
		log.Error("failed to get FAQs: ", err)
		return
	}
	var ius []intent.IntentUtterance

	counter := 0
	for _, faq := range faqs {
		for _, v := range faq.Questions {
			v.Question = strings.TrimSpace(v.Question)
			ius = append(ius, intent.IntentUtterance{faq.IntentName, v.Question})
			log.Infof("Add utterance #%d (%s, %s) \n", counter, faq.IntentName, v.Question)
			counter = counter + 1
		}
	}

	err = intent.AddIntentUtterances(ius)
	if err != nil {
		log.Errorf("Failed to insert %d utterance. Error: %+v\n", len(ius), err)
		return
	}

	log.Infof("Start training..")
	err = intent.Train()
	if err != nil {
		log.Error("Error in FAQ Retrain: ", err)
		isTraining = false
		return
	}
	//time.Sleep(30 * time.Second) //TODO: Remove later hungnq: work around because lacking of checking training status api
	isTraining = false
	log.Info("Done")

	fmt.Fprintln(w, "Done")
}