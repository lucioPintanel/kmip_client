package mytls

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func LoadTLSConfig() (*tls.Config, error) {
	// Caminho para os certificados gerados
	caCertPath := "scripts\\ca.crt"
	clientCertPath := "scripts\\client.crt"
	clientKeyPath := "scripts\\client.key"

	// Carrega a CA
	caCert, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Carrega o certificado e a chave do cliente
	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}, nil
}
