package model

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {
	dbConnString := fmt.Sprintf(
		"user=%v dbname=%v password=%v sslmode=%v",
		os.Getenv("POSTGRESQL_USERNAME"), os.Getenv("DATABASE_NAME"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("DBSSLMODE"),
	)
	DB = sqlx.MustConnect("postgres", dbConnString)
}
