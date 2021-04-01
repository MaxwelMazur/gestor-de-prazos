package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

// [CorrectAnswers retorna uma resposta em JSON para a requisição]
func CorrectAnswers(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

// [WrongAnswers retorna um erro em formato JSON]
func WrongAnswers(w http.ResponseWriter, statusCode int, erro error) {
	CorrectAnswers(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
