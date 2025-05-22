package main

import (
	"log/slog"
	"net/http"

	"github.com/bdtomlin/gostak/internal/assets"
	"github.com/bdtomlin/gostak/internal/model"
	"github.com/bdtomlin/gostak/internal/router"
)

func main() {
	model.Init()
	assets.LoadAssetMap()
	router := router.NewRouter()
	slog.Info("Listening on :8888")
	http.ListenAndServe(":8888", router)
}
