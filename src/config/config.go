package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConnectionDB string de conexão com o banco de dados
	StringConnectionDB = ""

	//Nome do banco de dados criado no mongo
	DB_Name = ""

	// Port porta em que o aplicativo será executado.
	ApiPort = 0

	// AddrRedis local  em que o Redis está sendo executado (host:port)
	AddrRedis = ""

	//PassRedis Senha do Redis
	PassRedis = ""

	//KeyApiHGBRASIL chave de acesso a api HGBRASIL
	KeyApiHGBRASIL = ""
)

// Loader vai inicializar as variáveis de ambiente
func Loader() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	//passando de string para inteiro
	ApiPort, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		ApiPort = 3000
	}

	StringConnectionDB = os.Getenv("STRING_CONNECTION_DB")
	DB_Name = os.Getenv("DB_NAME")

	AddrRedis = os.Getenv("ADRRESS_REDIS")
	PassRedis = os.Getenv("PASSWORD_REDIS")

	KeyApiHGBRASIL = os.Getenv("KEY_API_HGBRASIL")

}
