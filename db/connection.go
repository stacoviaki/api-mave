package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=api-mave sslmode=disable"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão: %v", err)
	}

	// testa se conecta de verdade
	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	fmt.Printf("\n✅ Conectado ao PostgreSQL!\n")
	return db, nil
}
