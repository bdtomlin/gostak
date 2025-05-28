package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bdtomlin/gostak/internal/model"
)

func setCookie(w http.ResponseWriter, name, value string) {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

type sessionData struct {
	Uid   string `json:"uid"`
	Token string `json:"token"`
}

func setSessionCookie(w http.ResponseWriter, session *model.Session) error {
	sd := sessionData{Uid: session.UserID, Token: session.Token}
	jsonData, err := json.Marshal(sd)
	if err != nil {
		return fmt.Errorf("setSessionCookie Error: %w", err)
	}
	setCookie(w, "session", string(jsonData))
	return nil
}

func readCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("readCookie: %w", err)
	}
	return c.Value, nil
}

func readSessionCookie(r *http.Request) (uid, sid string, err error) {
	value, err := readCookie(r, "session")
	if err != nil {
		return "", "", fmt.Errorf("readSessionCookie: %w", err)
	}
	var sd sessionData
	if err := json.Unmarshal([]byte(value), &sd); err != nil {
		return "", "", fmt.Errorf("readSessionCookie - Error decoding cookie: %w", err)
	}
	return sd.Uid, sd.Token, nil
}
