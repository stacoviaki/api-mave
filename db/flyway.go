package db

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecutarMigrations() {
	fmt.Println("Iniciando Migrations")
	cmd := exec.Command("flyway", "migrate", "-configFiles=/home/doss/projects/api-mave/db/flyway.conf")
	resp, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Erro ao executar o Flyway: %s", err)
		fmt.Println(string(resp))
		log.Fatal("Falha nas migrations.")
	}

	fmt.Println(string(resp))
	fmt.Printf("\n✅ Migrations Finalizadas\n")
}
