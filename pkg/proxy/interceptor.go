package proxy

import (
	"fmt"
	"io"
	"net/http"
)

// HandleRequestAndResponse intercepta requisições e respostas
func HandleRequestAndResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Intercepting request to %s\n", r.URL)

	// Faz a requisição para o servidor alvo
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copia os headers da resposta
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	// Copia o corpo da resposta
	io.Copy(w, resp.Body)
}
