package main

import (
	"github.com/stacoviaki/api-mave/db"
	"github.com/stacoviaki/api-mave/internal/router"
)

func main() {
	//Conexão com o banco de dados
	db.ConnectDB()

	//Função que executa as migrações no banco
	db.ExecMigrations()

	//Inicializa GIN e as rotas presentes na api
	router.Routes()
}
