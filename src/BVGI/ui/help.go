package ui

import (
	"net/http"
)

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	var result = make(map[string]interface{})

	result["name"] = "help"
	tplHelper.Render(w, "help", result)
}