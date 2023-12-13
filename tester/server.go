package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá, você acessou: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler) // Define o handler para a rota raiz

	fmt.Println("Servidor rodando em http://localhost:8080")
	err := http.ListenAndServe(":8080", nil) // Inicia o servidor na porta 8080
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
