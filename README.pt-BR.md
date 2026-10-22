# Mercadinho API

[English](README.md) | **Português (BR)**

[![CI](https://github.com/Alphka/Mercadinho-API/actions/workflows/ci.yml/badge.svg)](https://github.com/Alphka/Mercadinho-API/actions/workflows/ci.yml) [![Licença: ISC](https://img.shields.io/badge/licen%C3%A7a-ISC-blue.svg)](LICENSE.md)

API REST desenvolvida como projeto de aprendizado da linguagem Go e para aprofundar o conhecimento em PostgreSQL. Ela é inspirada em [um aplicativo Android](https://github.com/Alphka/MercadoCRUDJava), adaptando suas funcionalidades para uma arquitetura REST.

## Funcionalidades

- CRUD de clientes, produtos e compras
- Compras do cliente (cadastro e listagem paginada)
- Listagem de pagamentos pendentes e atrasados por cliente
- Exclusão lógica (soft delete) de produtos
- Validação de requisições com regras customizadas (`date`, `price`, `phone`)
- Documentação Swagger interativa
- Migrações de banco de dados aplicadas automaticamente na inicialização
- CORS permissivo (todas as origens permitidas)

## Tecnologias

- **Go** com [Gin](https://github.com/gin-gonic/gin)
- **Swagger** via [swaggo](https://github.com/swaggo/swag)
- **golang-migrate** para migrações do banco de dados
- **PostgreSQL** com [pgx](https://github.com/jackc/pgx)

## Pré-requisitos

- [Go](https://go.dev/dl/) (o `go.mod` declara `go 1.26.2`)
- [PostgreSQL](https://www.postgresql.org/) em execução e acessível
- CLI [`swag`](https://github.com/swaggo/swag) (usada pelos scripts de inicialização para gerar a documentação Swagger)

## Instalação

```bash
git clone https://github.com/Alphka/Mercadinho-API.git
cd Mercadinho-API

go install github.com/swaggo/swag/cmd/swag@latest
go mod download
```

### Configuração

Copie o arquivo de exemplo e preencha as variáveis de ambiente:

```bash
cp .env.example .env
```

| Variável | Descrição | Padrão |
|---|---|---|
| `DATABASE_URL` | String de conexão do PostgreSQL (obrigatória) | nenhum |
| `HOST` | Host usado pelo servidor e pelo Swagger | `localhost` |
| `PORT` | Porta do servidor | `8000` |

> [!NOTE]
> Se o banco indicado em `DATABASE_URL` não existir, a API tenta criá-lo na inicialização (para isso ela se conecta ao banco `postgres` do mesmo servidor), então o usuário configurado precisa ter essa permissão.

## Uso

Os scripts geram a documentação Swagger, compilam o binário em `output/` e iniciam o servidor:

```bash
# Bash
./bin/start.sh

# Windows PowerShell
./bin/start.ps1

# CMD
./bin/start.bat
```

Com o servidor em execução, a interface do Swagger fica disponível em:

```
http://localhost:8000/swagger/index.html
```

### Banco de dados

As migrations ficam em [`migrations/`](migrations) e são embutidas no binário. Elas são aplicadas automaticamente na inicialização, exceto quando executado na Vercel (detectado pela variável de ambiente `VERCEL`). Para apenas aplicá-las, sem iniciar o servidor:

```bash
go run . migrate
```

### Endpoints

Todas as rotas usam o prefixo `/api/v1`. O contrato completo está na interface do Swagger.

| Recurso | Rotas |
|---|---|
| Clientes | `GET /customers`, `GET /customers/:id`, `POST /customers`, `PUT /customers/:id`, `DELETE /customers/:id` |
| Pagamentos | `GET /customers/:id/pending-payments`, `GET /customers/:id/overdue-payments` |
| Compras do cliente | `GET /customers/:id/purchases`, `POST /customers/:id/purchases` |
| Produtos | `GET /products`, `GET /products/:id`, `POST /products`, `PUT /products/:id`, `DELETE /products/:id` |
| Compras | `GET /purchases`, `GET /purchases/:id`, `PUT /purchases/:id`, `DELETE /purchases/:id` |

Os endpoints de listagem aceitam os parâmetros de query `page` (padrão `1`) e `pageSize` (padrão `20`).

### Hot reload (opcional)

O [nodemon](https://nodemon.io/) pode reiniciar o servidor automaticamente ao salvar alterações (veja [`nodemon.json`](nodemon.json)). Requer [Node.js](https://nodejs.org/):

```bash
npm install -g nodemon
nodemon
```

## Estrutura do projeto

```
bin/          Scripts de inicialização (Bash, PowerShell, CMD)
controllers/  Handlers das requisições (anotados para o Swagger)
database/     Conexão com o PostgreSQL e execução das migrations
docs/         Arquivos Swagger gerados (swag init)
dtos/         Payloads das requisições e regras de validação
helpers/      Auxiliares de validação, datas e telefone
migrations/   Migrations SQL (golang-migrate)
models/       Modelos de resposta e de banco de dados
routes/       Registro das rotas
```

## Testes

O repositório ainda não possui testes automatizados. Um workflow do GitHub Actions ([`.github/workflows/ci.yml`](.github/workflows/ci.yml)) executa a verificação do `gofmt`, `go build`, `go vet` e `go test` em pushes para `main` e em pull requests.

## Licença

[ISC](LICENSE.md) © 2025 Kayo Souza
