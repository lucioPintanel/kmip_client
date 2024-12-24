package kmip

import (
	"crypto/tls"
	"log"

	"kmip_cli/internal/mytls"
)

func KMIPClient() {
	tlsConfig, err := mytls.LoadTLSConfig()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração TLS: %v", err)
	}

	conn, err := tls.Dial("tcp", "test.kmip-server.com:5696", tlsConfig)
	if err != nil {
		log.Fatalf("Erro ao conectar ao servidor KMIP: %v", err)
	}
	defer conn.Close()

	log.Println("Conexão estabelecida com sucesso com o servidor KMIP.")
	// Aqui você pode implementar operações KMIP, como criação de chaves, obtenção de chaves, etc.
}
