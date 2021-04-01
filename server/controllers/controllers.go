package controllers

import (
	"Gestor-de-prazos/server/connection"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

var templates = template.Must(template.ParseFiles("template/index.html"))

func RootHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func SendServiceWorker(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("static/service-worker.js")
	if err != nil {
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	w.Write(data)
}

func SendManifest(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("static/manifest.json")
	if err != nil {
		http.Error(w, "Couldn't read file", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}

type Tarefas struct {
	ID              int    `json:"id,omitempty"`
	NomeCliente     string `json:"nomeCliente,omitempty"`
	DataVencimento  string `json:"dataVencimento,omitempty"`
	DataNotificacao string `json:"dataNotificacao,omitempty"`
	TipoDocumento   string `json:"tipoDocumento,omitempty"`
	Descricao       string `json:"descricao,omitempty"`
	Desativado      string `json:"desativado,omitempty"`
	CriadoEm        string `json:"criadoEm,omitempty"`
}

func ListTarefas(w http.ResponseWriter, r *http.Request) {

	db, err := connection.Connect()
	if err != nil {
		WrongAnswers(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	rows, err := db.Query(
		"select t.* from gestorDePrazos.tarefa t order by t.id asc;",
	)

	if err != nil {
		WrongAnswers(w, -3333, err)
		return
	}
	defer rows.Close()

	var tarefas []Tarefas

	for rows.Next() {
		var tarefa Tarefas

		if err := rows.Scan(
			&tarefa.ID,
			&tarefa.NomeCliente,
			&tarefa.DataVencimento,
			&tarefa.DataNotificacao,
			&tarefa.TipoDocumento,
			&tarefa.Descricao,
			&tarefa.Desativado,
			&tarefa.CriadoEm,
		); err != nil {
			return
		}

		tarefas = append(tarefas, tarefa)
	}

	fmt.Println(tarefas)
	fmt.Println(http.StatusOK)
	CorrectAnswers(w, http.StatusOK, tarefas)
}
