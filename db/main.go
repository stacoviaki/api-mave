package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error

	connStr := "host=localhost port=5432 user=postgres password=123456 dbname=api-mave sslmode=disable"

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão: %v", err)
	}

	// testa se conecta de verdade
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
	}

	fmt.Println("✅ Conectado ao PostgreSQL!")
}
