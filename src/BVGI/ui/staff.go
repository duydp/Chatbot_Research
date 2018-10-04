package ui

import (
	"fmt"
	"net/http"
	"strconv"

	"BVGI/db"
	
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func ListStaffsHandler(w http.ResponseWriter, r *http.Request) {
	staffs, err := db.GetAllStaffs()
	if err != nil {
		log.Error("failed to Get All Staffs: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var result = make(map[string]interface{})

	result["staffs"] = staffs
	tplHelper.Render(w, "staff", result)
}

func DeleteStaffHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tmp_id := vars["id"]
	id, err := strconv.Atoi(tmp_id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if err := db.DeleteStaff(id); err != nil {
		log.Error("failed to delete staff: ", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "<tr></tr>")
}