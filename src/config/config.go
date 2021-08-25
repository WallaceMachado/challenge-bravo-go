package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConnectionDB é a string de conexão com o banco de dados
	StringConnectionDB = ""

	DB_Name = ""

	// Port onde a API vai estar rodando
	Port = 0
)

// Loader vai inicializar as variáveis de ambiente
func Loader() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	DB_Name = os.Getenv("DB_NAME")

	//passando de string para inteiro
	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	StringConnectionDB = os.Getenv("STRING_CONNECTION_DB")

}
