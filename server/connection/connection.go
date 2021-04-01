package connection

import (
	"Gestor-de-prazos/server/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// [Connect abre a conexão com o banco de dados e a retorna]
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
