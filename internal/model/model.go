package model

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

// Make sure to call defer model.Init()() in main
func Init() func() {
	dbConnString := fmt.Sprintf(
		"user=%v dbname=%v password=%v sslmode=%v",
		os.Getenv("POSTGRESQL_USERNAME"), os.Getenv("DATABASE_NAME"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("DBSSLMODE"),
	)
	DB = sqlx.MustConnect("postgres", dbConnString)
	return Cleanup
}

func Cleanup() {
	DB.Close()
}
