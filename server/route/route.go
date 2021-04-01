package route

import (
	"Gestor-de-prazos/server/controllers"
	"net/http"
)

// LoadRoutes carrega todas as rotas
func LoadRoutes() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/service-worker.js", controllers.SendServiceWorker)
	http.HandleFunc("/manifest.json", controllers.SendManifest)
	http.HandleFunc("/", controllers.RootHandler)
	http.HandleFunc("/ListTarefas", controllers.ListTarefas)
}
