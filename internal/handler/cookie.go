package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/bdtomlin/gostak/internal/model"
	"github.com/google/uuid"
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
	Uid   *uuid.UUID `json:"uid"`
	Token string     `json:"token"`
}

func setSessionCookie(w http.ResponseWriter, session *model.Session) error {
	sd := sessionData{Uid: &session.UserID, Token: session.Token}
	jsonData, err := json.Marshal(sd)
	if err != nil {
		return fmt.Errorf("setSessionCookie Error: %w", err)
	}
	setCookie(w, "session", url.QueryEscape(string(jsonData)))
	return nil
}

func readCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("readCookie: %w", err)
	}
	return c.Value, nil
}

func readSessionCookie(r *http.Request) (uid *uuid.UUID, sid string, err error) {
	value, err := readCookie(r, "session")
	if err != nil {
		return nil, "", fmt.Errorf("readSessionCookie: %w", err)
	}
	value, err = url.QueryUnescape(value)
	if err != nil {
		return nil, "", fmt.Errorf("readSessionCookie - Error unescaping cookie: %w", err)
	}
	var sd sessionData
	if err := json.Unmarshal([]byte(value), &sd); err != nil {
		return nil, "", fmt.Errorf("readSessionCookie - Error decoding cookie: %w", err)
	}
	return sd.Uid, sd.Token, nil
}
