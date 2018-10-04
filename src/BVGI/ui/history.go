package ui

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"BVGI/db"
	"BVGI/intent"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func ListHistoriesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmp_id := vars["id"]
	if tmp_id == "" {
		tmp_id = "1"
	}
	page, err := strconv.Atoi(tmp_id)
	if err != nil {
		log.Error("failed to strconv: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	offset := (page - 1) * HistoriesPerPage
	
	if offset < 0 {
		offset = 0
	}

	histories, err := db.GetHistoryForPagination(HistoriesPerPage, offset)
	if err != nil {
		log.Error("failed to list histories: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	result := make(map[string]interface{})
	result["history"] = histories

	historyRows, err := db.GetHistoryRows()
	if err != nil {
		log.Errorf("Count history error: ", err)
	}

	pages := math.Ceil(float64(historyRows)/float64(HistoriesPerPage))

	result["pages"] = pages
	result["page"] = page

	intents, err := intent.GetAll()
	if err != nil {
		log.Println("Get all intents error " + err.Error())
	}

	result["intents"] = intents

	tplHelper.Render(w, "history", result)

	return
}


func DeleteHistoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmp_id := vars["id"]
	id, err := strconv.Atoi(tmp_id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := db.DeleteHistory(id); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<tr></tr>")
}