package proxy

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"strings"
)

// CloneRequest cria uma cópia do request HTTP.
func CloneRequest(r *http.Request) (*http.Request, error) {
	// Lê o corpo do request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Restaura o corpo do request original
	r.Body = ioutil.NopCloser(bytes.NewReader(body))

	// Cria um novo request com o mesmo método, URL e cabeçalhos
	req, err := http.NewRequest(r.Method, r.URL.String(), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Copia todos os cabeçalhos do request original
	req.Header = r.Header.Clone()

	return req, nil
}

// AddCustomHeader adiciona um header customizado no request ou resposta.
func AddCustomHeader(r *http.Request, key, value string) {
	// Adiciona um header customizado
	r.Header.Add(key, value)
}

// RemoveHeader remove um header específico do request.
func RemoveHeader(r *http.Request, key string) {
	// Remove um header
	r.Header.Del(key)
}

// SetAuthorizationHeader define o header de autorização com o token fornecido.
func SetAuthorizationHeader(r *http.Request, token string) {
	// Configura o header de autorização
	r.Header.Set("Authorization", "Bearer "+token)
}

// ValidateURL verifica se a URL de destino é válida e segura.
func ValidateURL(u string) (*url.URL, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		log.Println("Erro ao validar URL:", err)
		return nil, err
	}

	// Verifica se a URL está no formato adequado
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		log.Println("URL não é HTTP ou HTTPS:", u)
		return nil, err
	}

	return parsedURL, nil
}
