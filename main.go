package main

import (
	"fmt"
	"net/http"
	"os"
)

var portaServer = obterPortaServidor()

func main() {

	http.HandleFunc("/pessoa", obterTodasPessoas)
	http.HandleFunc("/pessoa/{:id}", obterPessoaPorId)

	fmt.Println("Iniciando servidor de aplicação na porta " + portaServer)
	if err := http.ListenAndServe(portaServer, nil); err != nil {
		panic(err)
	}
}

func obterPortaServidor() string {
	porta := os.Getenv("PORT")
	if porta == "" {
		return ":3233"
	}
	return ":" + porta
}

func obterTodasPessoas(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//listaTodos
	case http.MethodPost:
		// Create a new record.
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Invalid request method.", 405)
	}
}

func obterPessoaPorId(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		//listaTodos
	case http.MethodPost:
		// Create a new record.
	case http.MethodPut:
		// Update an existing record.
	case http.MethodDelete:
		// Remove the record.
	default:
		http.Error(w, "Invalid request method.", 405)
	}
}
