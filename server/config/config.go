package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	// [StringConexaoBanco é a string de conexão com o MySQL]
	StringConexaoBanco = ""

	// [Porta onde a Server vai estar rodando]
	Porta = 0

	// [SecretKey é a chave que vai ser usada para assinar o token]
	SecretKey []byte
)

// [LoadConfig vai inicializar as variáveis de ambiente]
func LoadConfig() {
	var erro error

	if erro = Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9090
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
