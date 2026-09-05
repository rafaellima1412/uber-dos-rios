Uber dos Rios

API REST em Go para gestão de **transporte fluvial de passageiros** (o "Uber dos Rios"): cadastro de embarcações, configuração de assentos e camarotes, terminais, rotas, horários, viagens, reservas de passagens, upload de fotos e rastreamento por GPS. Construída em Clean Architecture, com Postgres, Gin e documentação Swagger.

## Visão geral

O domínio modela a operação de uma empresa de navegação:

- **Ships (embarcações)** — cadastro de barcos, com fotos enviadas para S3.
- **Ships-config (assentos e camarotes)** — configuração de cadeiras e camarotes de cada embarcação, com busca de unidades disponíveis.
- **Terminals (terminais)** — portos/terminais de embarque e desembarque, agrupados por cidade.
- **Routes (rotas)** — trajetos entre terminais.
- **Schedules (horários)** — grade de horários associada a rotas e terminais.
- **Trips-config / Trips (viagens)** — instâncias de viagem geradas a partir de uma configuração, com busca de conexões entre terminais usando **Dijkstra** (menor custo) e **DFS** (todas as combinações possíveis) sobre um grafo de viagens.
- **Reservations (reservas)** — reserva de passageiros em uma ou mais viagens (inclusive com conexões), com validação de disponibilidade de assentos/camarotes.
- **GPS** — ingestão de posições de GPS das embarcações.
- **Photos** — upload direto de fotos (S3) para embarcações.

## Arquitetura

Clean Architecture / Ports & Adapters:

```
cmd/api/main.go                        # bootstrap: config, DB pool, injeção de dependências, router

internal/
├── application/
│   ├── domain/        # entidades de negócio (Ship, Trip, Reservation, Route, Terminal, Seat, Cabin, GPS...)
│   ├── ports/
│   │   ├── input/      # interfaces de casos de uso (o que a API oferece)
│   │   └── output/     # interfaces de repositórios/serviços externos (o que a API precisa)
│   └── usecase/        # regras de negócio (create/update/delete/list/search por entidade,
│                        #   + Dijkstra/DFS para busca de viagens conectadas)
├── adapters/
│   ├── input/
│   │   ├── http/        # router (Gin) + DTOs de request/response
│   │   └── handlers/    # handlers HTTP por entidade
│   └── output/
│       ├── repository/postgres/  # implementação dos repositórios em Postgres (pgx)
│       └── service/               # integrações externas (organization service, GPS, S3)
├── config/             # carregamento de configuração (YAML)
└── logger/             # logger estruturado (zap)

migrations/             # migrations SQL (golang-migrate), organizações → usuários → embarcações →
                         #   rotas → terminais → horários → viagens → reservas → bagagens
docs/                   # Swagger (swaggo) gerado a partir das anotações do código
configs/                # config.yaml / config.local.yaml / config.docker.yaml
```

## Stack

- **Go 1.24**, [Gin](https://github.com/gin-gonic/gin) (HTTP), [pgx v5](https://github.com/jackc/pgx) (Postgres)
- **AWS SDK v2 (S3)** para upload de fotos das embarcações
- **swaggo/swag** + **gin-swagger** para documentação OpenAPI (`/docs/*any`)
- **zap** para logging estruturado
- **golang-migrate** para migrations SQL versionadas (44 migrations até o momento)
- **CORS** configurado via `gin-contrib/cors`
- Deploy em container: `Dockerfile` (produção, multi-stage) e `Dockerfile.local`

## Requisitos

- Go 1.24+
- PostgreSQL
- (Opcional) Docker, para build/execução em container
- Credenciais AWS, se for usar upload de fotos para S3

## Configuração

A aplicação lê o caminho do arquivo de configuração pela variável de ambiente `CONFIG_PATH` (padrão: `configs/config.local.yaml`). Cada ambiente tem seu próprio arquivo:

- `configs/config.local.yaml` — desenvolvimento local
- `configs/config.docker.yaml` — execução via Docker
- `configs/config.yaml` — produção

Cada um define `server.port`, `server.host` e `database.url`.

> ⚠️ Os arquivos de config incluídos no repositório têm URLs de conexão com credenciais do banco em texto plano. Vale mover isso para variáveis de ambiente/secrets antes de expor o repositório publicamente ou reutilizar essas credenciais.

Variáveis de ambiente adicionais usadas pela integração com o serviço de organizações:
```
ORG_SERVICE_CLIENT_ID
ORG_SERVICE_CLIENT_SECRET
ORG_SERVICE_AUTH_URL
SWAGGER_HOST   # opcional, sobrescreve o host mostrado no Swagger
```

## Como rodar

### Localmente

```bash
git clone https://github.com/rafaellima1412/uber-dos-rios.git
cd uber-dos-rios
go mod download

# rode as migrations no Postgres configurado em configs/config.local.yaml
# (usando golang-migrate ou a ferramenta de sua preferência)

go run ./cmd/api/main.go
```

A API sobe em `http://localhost:8001` (conforme `config.local.yaml`) e a documentação Swagger fica disponível em `/docs/index.html`.

### Com Docker

```bash
docker build -t nautical-logistics-api .
docker run -p 8002:8002 -e CONFIG_PATH=configs/config.docker.yaml nautical-logistics-api
```

O `Dockerfile` faz build multi-stage (builder + alpine) e expõe um healthcheck em `/health`.

## Principais rotas (`/api/v1`)

| Recurso | Rotas |
|---|---|
| Health | `GET /` |
| GPS | `POST /gps` |
| Embarcações | `GET,POST /ships`, `GET,PUT,DELETE /ships/:id` |
| Config. de assentos/camarotes | `POST,GET /ships-config/seats`, `PUT,DELETE /ships-config/seats/:id`, `GET /ships-config/units/search`, `POST,GET /ships-config/cabins`, `PUT,DELETE /ships-config/cabins/:id` |
| Terminais | `GET,POST /terminals`, `GET,PUT,DELETE /terminals/:id`, `GET /terminals/search`, `GET /terminals/cities` |
| Rotas | `GET,POST /routes`, `GET,PUT,DELETE /routes/:id`, `GET /routes/search` |
| Horários | `GET,POST /schedules`, `GET,PUT,DELETE /schedules/:id` |
| Config. de viagem | `GET,POST /trips-config`, `GET,PUT,DELETE /trips-config/:id` |
| Viagens | `GET /trips/search`, `POST /trips`, `POST /trips/catalogo`, `POST /trips/catalogo/menor-preco` |
| Reservas | `GET,POST /reservations`, `GET,PUT,DELETE /reservations/:id` |
| Fotos | `POST /photos/upload-urls` |

A lista completa e os schemas de request/response estão no Swagger (`/docs/index.html`), gerado a partir de `docs/swagger.json` / `docs/swagger.yaml`.

## Banco de dados

As migrations em `migrations/` (formato `NNNNNN_descricao.up.sql` / `.down.sql`, compatível com `golang-migrate`) descrevem a evolução do schema: tipos de organização → organizações/usuários/memberships → embarcações e suas configurações de assento/camarote → rotas e terminais → horários e viagens → reservas, passageiros/tickets e bagagens → posições de GPS → índices de performance.


## Possíveis próximos passos

- Externalizar credenciais de banco/S3 para variáveis de ambiente ou um gerenciador de secrets.
- Adicionar testes automatizados (não há arquivos `_test.go` no repositório atual).
- Documentar o fluxo de autenticação/autorização entre este serviço e o serviço de organizações (`ORG_SERVICE_*`).
