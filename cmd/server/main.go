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
	repo.DB = sqlx.MustConnect("postgres", fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v", os.Getenv("POSTGRESQL_USERNAME"), os.Getenv("DATABASE_NAME"), os.Getenv("PGPASSWORD"), os.Getenv("DBSSLMODE")))

	assets.LoadAssetMap()
	router.Route()
	slog.Info("Listening on :8888")
	http.ListenAndServe(":8888", nil)
}
