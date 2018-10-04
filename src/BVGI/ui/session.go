package ui

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func setSession(w http.ResponseWriter, username string) error {
	value := map[string]string{
		"username": username,
	}

	encoded, err := cookieHandler.Encode("session", value)
	if err != nil {
		log.Errorf("failed to encode session: ", err)
		return err
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: encoded,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	return nil
}

func getUsername(r *http.Request) string {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Debug("failed to read cookie: ", err)
		return ""
	}

	value := make(map[string]string)
	if err := cookieHandler.Decode("session", cookie.Value, &value); err != nil {
		//log.Debugf("failed to read username from cookie: %s", err.Error())
		return ""
	}

	return value["username"]
}

func clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}