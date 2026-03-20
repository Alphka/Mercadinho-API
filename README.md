# Mercadinho API

API desenvolvida como projeto de aprendizado da linguagem Go e para aprofundar o conhecimento em PostgreSQL. Ela serve para simular um back-end para [um aplicativo Android](https://github.com/Alphka/MercadoCRUDJava), adaptando suas funcionalidades para uma arquitetura REST.

## Tecnologias

- **Go** com [Gin](https://github.com/gin-gonic/gin)
- **Swagger** via [swaggo](https://github.com/swaggo/swag)
- **golang-migrate** para migrações do banco de dados
- **PostgreSQL** com [pgx](https://github.com/jackc/pgx)

## Endpoints

As rotas seguem o padrão `/api/v1`.

### Clientes
- `GET /api/v1/customers`
- `GET /api/v1/customers/:id`
- `POST /api/v1/customers`
- `PUT /api/v1/customers/:id`
- `DELETE /api/v1/customers/:id`

### Pagamentos
- `GET /api/v1/customers/:id/pending-payments`
- `GET /api/v1/customers/:id/overdue-payments`

### Compras do cliente
- `GET /api/v1/customers/:id/purchases`
- `POST /api/v1/customers/:id/purchases`

### Produtos
- `GET /api/v1/products`
- `GET /api/v1/products/:id`
- `POST /api/v1/products`
- `PUT /api/v1/products/:id`
- `DELETE /api/v1/products/:id`

### Compras
- `GET /api/v1/purchases`
- `GET /api/v1/purchases/:id`
- `PUT /api/v1/purchases/:id`
- `DELETE /api/v1/purchases/:id`

## Como executar

### Pré-requisitos

- [Go](https://go.dev/dl/) instalado
- [PostgreSQL](https://www.postgresql.org/) em execução

### Configuração

Copie o arquivo `.env.example` e preencha as variáveis de ambiente:

```bash
cp .env.example .env
```

### Instalando dependências

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go mod download
```

### Executando o projeto

```bash
# Bash
./bin/start.sh

# Windows PowerShell
./bin/start.ps1

# CMD
./bin/start.bat
```

Os scripts geram a documentação Swagger, compilam e iniciam o servidor automaticamente.

## Banco de dados

A API utiliza PostgreSQL e as migrations são gerenciadas com `golang-migrate`.  
Certifique-se de que o banco esteja em execução e configurado corretamente no arquivo `.env` antes de iniciar o projeto.

## Swagger

Com o servidor em execução, acesse:

```
http://localhost:8000/swagger/index.html
```

## Estrutura do projeto

```
controllers/
dtos/
models/
routes/
migrations/
bin/
```

## Hot reload (opcional)

O [nodemon](https://nodemon.io/) pode ser utilizado para reiniciar o servidor automaticamente ao salvar alterações. Requer [Node.js](https://nodejs.org/) instalado.:

```bash
npm install -g nodemon
nodemon
```
