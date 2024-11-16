package proxy

import (
	"crypto/tls"
)

// Configura certificados SSL/TLS para interceptação HTTPS
func SetupTLS() *tls.Config {
	// Placeholder para setup de certificados
	return &tls.Config{
		InsecureSkipVerify: true, // Atenção: apenas para testes
	}
}
