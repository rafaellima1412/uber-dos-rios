# Issues conhecidos — para revisão futura

Levantamento feito em análise do código em 2026-09-04. Nenhum item aqui foi corrigido ainda — é só o registro do que foi encontrado, pra retomar quando houver tempo.

---

## 1. [CRÍTICO] Reserva/overbooking: código Go dessincronizado do schema atual

**O que é:** o fluxo de criação de reserva e a checagem de disponibilidade de assentos usam colunas que já não existem no banco.

**Onde:**
- `internal/adapters/output/repository/postgres/reservation_repository_impl.go`, função `CreateTripReservation` — grava em `reservations_trips.occupied_seats` / `occupied_cabins`.
- `internal/adapters/output/repository/postgres/ship_config_repository_impl.go`, função `SearchSeatByShipAndDate` — lê `occupied_seats` / `occupied_cabins` via `jsonb_array_elements_text`.

**Por que está quebrado:**
- A migration `000041_remove_jsonb_legacy.up.sql` **dropou** as colunas `occupied_seats` e `occupied_cabins` de `reservations_trips`.
- A migration `000043_refactor_occupancy.up.sql`, posterior, introduziu o modelo correto: uma linha por assento em `reservations_trips`, com colunas `seat_code`, `passenger_id`, `ticket_id`, `ship_cabin_id`, `origin_stop_order`, `destination_stop_order`, `status`, **e uma constraint física** (`EXCLUDE USING gist`, com `btree_gist`) que impede duas reservas com o mesmo `seat_code` no mesmo `trip_instances_id` cujos trechos se sobreponham.
- **Nenhum arquivo `.go` do projeto referencia `seat_code`, `passenger_id`, `ticket_id` ou `ship_cabin_id`.** O código nunca foi atualizado para o modelo novo.

**Impacto esperado:** se as 44 migrations estiverem aplicadas no banco em uso, tanto `POST /reservations` (fluxo de adicionar viagem à reserva) quanto a busca de assentos disponíveis (`GET /ships-config/units/search`) devem falhar em runtime com erro do tipo `column "occupied_seats" does not exist`.

**Consequência de negócio:** a trava de overbooking existe e está bem desenhada no banco, mas está inalcançável — nenhuma reserva popula `seat_code`/`origin_stop_order`/`destination_stop_order`, então a constraint nunca é exercitada.

**O que precisa ser feito:**
- Reescrever `CreateTripReservation` para inserir uma linha por assento/passageiro em `reservations_trips`, preenchendo `seat_code`, `origin_stop_order` e `destination_stop_order` (a origem desses dois últimos precisa ser definida — provavelmente vem da posição dos terminais de embarque/desembarque na rota).
- Reescrever `SearchSeatByShipAndDate` para consultar disponibilidade por `seat_code` em vez do array JSONB antigo.
- Validar se `ValidatedTripsUseCase.Execute` (que hoje faz a checagem "olha e depois grava", sem transação nem lock) ainda é necessário depois disso, já que a constraint do banco passaria a ser a fonte de verdade contra corrida — o ideal é confiar na constraint (capturar a violação como erro de negócio) em vez de tentar prevenir a colisão só na aplicação.

---

## 2. [ALTO] Nenhuma autenticação nas rotas desta API

`internal/adapters/input/http/router.go` não tem nenhum middleware de auth. O único mecanismo de token existente (`TokenManager` em `organization_service_impl.go`) autentica esta API *contra outro serviço*, não protege as próprias rotas.

Combinado com o item 1: mesmo depois de corrigido, qualquer pessoa com a URL pode criar/editar/apagar embarcações, rotas e reservas — a autenticação foi projetada para acontecer em outro serviço/gateway; só falta confirmar que ele realmente está na frente disso em todos os ambientes.

---

## 3. [MÉDIO] Métodos do repositório de viagens sem implementação

`internal/adapters/output/repository/postgres/trip_repository_impl.go`: `DeleteTrip`, `GetTrip`, `ListTrips` e `UpdateTrip` só têm `panic("unimplemented")`.

Hoje nenhuma rota chama esses métodos (verificado), então não quebram nada em produção — mas quebram (com panic, não erro tratado) assim que alguém conectar um handler a eles.

---

## 4. [BAIXO] Dependências nulas passadas de propósito

Em `cmd/api/main.go`, `searchSeatUseCase` e `searchSchedulesUseCase` são passados como `nil` na montagem dos handlers (`ShipConfigHandler`, `ScheduleHandler`). A rota `/ships-config/units/search` na prática usa outro usecase (`listSeatAvailableUseCase`), então não há crash hoje — mas indica casos de uso planejados e nunca implementados.

---

## 5. [BAIXO] Credenciais de banco em texto plano nos arquivos de config

`configs/config.local.yaml` e `configs/config.yaml` têm usuário e senha do Postgres direto na URL de conexão, commitados no repositório.

---

## 6. [BAIXO] Regra de negócio dentro do handler HTTP

`internal/adapters/input/handlers/trips_handler.go` tem um comentário do próprio autor original: `// TODO remover as regras do handler` — lógica de negócio que deveria estar em usecase está no adapter HTTP, dificultando testar isoladamente.
