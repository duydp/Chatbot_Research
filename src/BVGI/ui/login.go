package ui

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"BVGI/config"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var data map[string]interface{}
		data = make(map[string]interface{})
		data["botname"] = config.UI.BotName

		tplHelper.Render(w, "login", data)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != config.UI.Username || password != config.UI.Userpass {
		log.Debugf("unauthorized: username=%s, password=%s\n", username, password)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	if err := setSession(w, username); err != nil {
		log.Error("failed to login, username = ", username)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/faq", http.StatusFound)
}