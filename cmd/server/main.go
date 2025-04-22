package main

import (
	"log/slog"
	"net/http"

	"github.com/bdtomlin/gostak/internal/assets"
	"github.com/bdtomlin/gostak/internal/repo"
	"github.com/bdtomlin/gostak/internal/router"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	repo.Init()
	assets.LoadAssetMap()
	router := router.NewRouter()
	slog.Info("Listening on :8888")
	http.ListenAndServe(":8888", router)
}
