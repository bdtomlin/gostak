package handler

import (
	"net/http"

	"github.com/gorilla/schema"
)

func decodeParams(dst any, r *http.Request) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = decoder.Decode(dst, r.PostForm)
	if err != nil {
		return err
	}
	return nil
}

func redirect(w http.ResponseWriter, r *http.Request, location string) {
	w.Header().Add("HX-Location", location)
}
