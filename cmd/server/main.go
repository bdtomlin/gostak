package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/bdtomlin/gostak/internal/assets"
	"github.com/bdtomlin/gostak/internal/repo"
	"github.com/bdtomlin/gostak/internal/router"
	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	dbConnString := fmt.Sprintf(
		"user=%v dbname=%v password=%v sslmode=%v",
		os.Getenv("POSTGRESQL_USERNAME"), os.Getenv("DATABASE_NAME"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("DBSSLMODE"),
	)
	repo.DB = sqlx.MustConnect("postgres", dbConnString)

	assets.LoadAssetMap()
	router := router.NewRouter()
	slog.Info("Listening on :8888")
	http.ListenAndServe(":8888", router)
}
