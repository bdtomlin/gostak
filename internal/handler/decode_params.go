package handler

import (
	"net/http"

	"github.com/gorilla/schema"
)

func DecodeParams(dst any, r *http.Request) error {
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
