package ui

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"BVGI/intent"
)

func UnsafeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var result = make(map[string]interface{})
		result["name"] = "unsafe"
		tplHelper.Render(w, "unsafe", result)
		return
	}

	if r.Method == "DELETE" {
		err := intent.DeleteAllIntents()
		if err != nil {
			log.Error("Failed to delete all intents: ", err.Error())
			fmt.Fprintln(w, "Failed to delete")
			return
		}
		fmt.Fprintln(w, "Success")
	}

	if r.Method == "POST" {
		go FAQTrainAll(w)
	}

}
