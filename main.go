package main

import (
	"Gestor-de-prazos/server/config"
	"Gestor-de-prazos/server/route"
	"Gestor-de-prazos/tool/dice"
	"log"
	"net/http"
	"os"
)

func main() {

	log.SetOutput(os.Stdout)
	config.LoadConfig()
	route.LoadRoutes()

	porta := ":" + dice.InterfaceToString(config.Porta)
	log.Println("Servindo na porta" + porta)
	err := http.ListenAndServe(porta, nil)
	if err != nil {
		log.Fatal(err)
	}
}
