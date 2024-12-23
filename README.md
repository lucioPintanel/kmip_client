# KMIP Client with TLS/SSL

[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](#)
[![Git](https://img.shields.io/badge/Git-F05032?logo=git&logoColor=fff)](#)

![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)

Este projeto implementa um cliente KMIP (Key Management Interoperability Protocol) que se comunica com um servidor KMIP utilizando TLS/SSL para garantir a segurança na transmissão de dados. Através deste cliente, é possível realizar operações de gerenciamento de chaves criptográficas de forma segura e eficiente.

## Estrutura do Projeto
# kmip_client

* [cmd/](.\kmip_client\cmd)
  * [kmip-client/](.\kmip_client\cmd\kmip-client)
    * [main.go](.\kmip_client\cmd\kmip-client\main.go)
* [internal/](.\kmip_client\internal)
  * [kmip/](.\kmip_client\internal\kmip)
  * [tls/](.\kmip_client\internal\tls)
    * [certs/](.\kmip_client\internal\tls\certs)
    * [tls.go](.\kmip_client\internal\tls\tls.go)
* [scripts/](.\kmip_client\scripts)
  * [ca.crt](.\kmip_client\scripts\ca.crt)
  * [ca.key](.\kmip_client\scripts\ca.key)
  * [ca.srl](.\kmip_client\scripts\ca.srl)
  * [client.crt](.\kmip_client\scripts\client.crt)
  * [client.csr](.\kmip_client\scripts\client.csr)
  * [client.key](.\kmip_client\scripts\client.key)
  * [generate_certs.ps1](.\kmip_client\scripts\generate_certs.ps1)
* [.gitignore](.\kmip_client\.gitignore)
* [LICENSE](.\kmip_client\LICENSE)
* [README.md](.\kmip_client\README.md)

## Instalação

Para clonar e executar este projeto, siga estas etapas:

```bash
git clone https://github.com/lucioPintanel/kmip_client.git
cd kmip_client
go mod tidy
go run internal/cmd/main.go
```

## Funcionalidades
* Conexão Segura: Estabelecimento de comunicação segura usando TLS/SSL.
* Geração de Certificados: Script para geração de certificados TLS/SSL utilizando OpenSSL.
* Cliente KMIP: Implementação de um cliente para se comunicar com servidores KMIP e realizar operações de gerenciamento de chaves.
* Testes: Scripts de teste para verificar a funcionalidade e segurança da comunicação.

## Instalação do OpenSSL no Windows
1. Baixe o OpenSSL:
Acesse Win32 OpenSSL e baixe o instalador adequado para o seu sistema (geralmente Win64 OpenSSL v1.1.1L para sistemas Windows 64-bit).

2. Instale o OpenSSL:
Execute o instalador baixado e siga as instruções. Certifique-se de escolher a opção que copia o arquivo openssl.cnf para o diretório de instalação do OpenSSL.

3. Verifique a Instalação:
Abra o PowerShell e execute:
```
openssl version
```
Isso deve exibir a versão do OpenSSL instalada.

## Habilitando a Execução de Scripts no PowerShell
1. Abra o PowerShell como Administrador:
Clique com o botão direito no ícone do PowerShell e selecione "Executar como administrador".

2. Verifique a Política de Execução Atual:
```
Get-ExecutionPolicy
```

3. Defina a Política de Execução para RemoteSigned:
Isso permite a execução de scripts locais que não são assinados e scripts baixados da internet que são assinados por um editor confiável.
```
Set-ExecutionPolicy RemoteSigned
```

4. Confirme a Alteração:
Se solicitado, digite "A" e pressione Enter para confirmar a alteração.

## Como Usar

1. **Gerar Certificados TLS:**
   Execute o script no PowerShell para gerar os certificados TLS necessários.
```powershell
   cd C:\caminho\para\seu\projeto\scripts
   .\generate_certs.ps1
```

Compilar o Cliente: Compile o cliente KMIP.
```
go build -o kmip-client ./cmd/kmip-client
```

Executar o Cliente: Execute o cliente para se conectar ao servidor KMIP.

```
./kmip-client
```

## Requisitos
* Go 1.23 ou superior
* OpenSSL

## Contribuições
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests para melhorar este projeto.

## Licença
Este projeto está licenciado sob os termos da licença MIT.
