package src

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Conn() *sql.DB {
	errEnv := godotenv.Load()

	if errEnv != nil {
		log.Fatal(errEnv.Error())
	}

	user := os.Getenv("USER")
	pass := os.Getenv("PASS")
	db := os.Getenv("DB")
	dsn := fmt.Sprintf("host=postgres port=5432 user=%s password=%s dbname=%s sslmode=disable", user, pass, db)

	conn, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Erro: %s", err.Error())
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	return conn
}
