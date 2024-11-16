package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a flag para a porta
	port := flag.String("port", "8080", "Porta para o proxy ouvir")
	flag.Parse()

	// Exibe a porta selecionada
	fmt.Printf("Iniciando o proxy na porta %s...\n", *port)

	// Configura o servidor HTTP
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Proxy rodando!")
		fmt.Fprintln(w, "codigo:333444")
	})

	// Inicia o servidor na porta especificada
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o proxy: %v", err)
	}
}
