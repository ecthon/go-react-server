package main

import (
	"log"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Definir o comando tern migrate
	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// Executar o comando e capturar a saída
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Erro ao executar o comando tern migrate: %v\nSaída: %s", err, output)
	}

	// Exibir a saída do comando
	log.Printf("Saída do comando: %s", output)
}
