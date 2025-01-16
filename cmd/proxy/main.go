package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	// Define os parâmetros do programa
	port := flag.String("port", "8080", "Porta para o proxy ouvir")
	destination := flag.String("dest", "http://example.com", "URL de destino para onde o proxy redirecionará as requisições")
	help := flag.Bool("help", false, "Exibe a ajuda do programa")
	logRequests := flag.Bool("log", false, "Habilita o log das requisições interceptadas")
	flag.Parse()

	// Exibe a ajuda e sai se o parâmetro --help for passado
	if *help {
		fmt.Println("Uso:")
		fmt.Println("  proxy [opções]")
		fmt.Println("Opções:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Valida a URL de destino
	targetURL, err := url.Parse(*destination)
	if err != nil {
		log.Fatalf("Erro ao analisar a URL de destino: %v", err)
	}

	// Exibe informações sobre a configuração atual
	fmt.Printf("Iniciando o proxy em http://localhost:%s\n", *port)
	fmt.Printf("Redirecionando requisições para %s\n", *destination)

	// Configura o proxy handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if *logRequests {
			fmt.Printf("Interceptado: %s %s\n", r.Method, r.URL.Path)
		}

		// Cria o proxy reverso
		proxy := httputil.NewSingleHostReverseProxy(targetURL)

		// Intercepta a resposta (se necessário)
		proxy.ModifyResponse = func(resp *http.Response) error {
			if *logRequests {
				fmt.Printf("Resposta com status: %d\n", resp.StatusCode)
			}
			return nil
		}

		// Envia a requisição para o destino
		proxy.ServeHTTP(w, r)
	})

	// Inicia o servidor HTTP
	err = http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o proxy: %v", err)
	}
}
