package ui

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-IC-CancelPolling", "true")
	http.Redirect(w, r, "/faq", http.StatusFound)
}