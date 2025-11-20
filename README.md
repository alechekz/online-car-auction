# 🚗 Этап 1. Online Car Auction (базовая версия)

### 1. База данных (минимальный набор таблиц)
- **Vehicles**
  - VIN (PK)
  - Year
  - Odometer
  - ExteriorColor
  - InteriorColor
  - Engine, Transmission (можно подтягивать из `GetBuildData`)
  - MSRP

- **Inspections**
  - VIN (FK → Vehicles)
  - Grade
  - Defects (scratches, electrics, suspension)

- **Prices**
  - VIN (FK → Vehicles)
  - RecommendedPrice
  - Timestamp

---

### 2. API (микросервисы)
- **VehicleService**
  - `POST /vehicles` — добавить машину
  - `GET /vehicles/{vin}` — получить данные
  - `PUT /vehicles/{vin}` — обновить
  - `DELETE /vehicles/{vin}` — удалить

- **InspectionService**
  - `POST /inspect` — расчёт grade по VIN, Odometer, Year, дефектам
  - `GET /build-data/{vin}` — подтянуть данные по VIN

- **PricingService**
  - `POST /price` — расчёт рекомендованной цены (учёт grade, пробега, цветов и MSRP)

---

### 3. Дополнительно
- Bulk-загрузка машин (например, `POST /vehicles/bulk`).
- Батчинг при массовой загрузке.
- Кэширование результатов цен (Redis).

👉 На этом этапе у тебя уже есть система, которая умеет хранить машины, оценивать их состояние и выдавать цену.

---

# 🏁 Этап 2. Расширение до Better Car Auction

### 1. Новые таблицы
- **Users**
  - Id, Name, Role (user/admin)

- **Auctions**
  - Id
  - Name
  - StartDateTime
  - EndDateTime
  - Status (open/closed)

- **AuctionVehicles**
  - AuctionId (FK)
  - VIN (FK)

- **Bids**
  - Id
  - AuctionId (FK)
  - VIN (FK)
  - UserId (FK)
  - Amount
  - Timestamp
  - IsWinning (bool)

---

### 2. Новые сервисы
- **UserService**
  - Регистрация, авторизация, роли.

- **AuctionService**
  - CRUD для аукционов (только админ).
  - Назначение машин на аукцион.

- **BiddingService**
  - `POST /bids` — сделать ставку.
  - `GET /bids/{auctionId}` — список ставок.
  - Автоматический выбор победителя после закрытия аукциона.

---

### 3. Логика
- Пользователь добавляет машину → получает grade и цену.
- Админ создаёт аукцион и назначает туда машины.
- Пользователи делают ставки, пока аукцион открыт.
- После закрытия система автоматически определяет победителя по каждой машине.

---

# 🔧 Этап 3. Доработки
- Логирование и мониторинг.
- Масштабирование (например, вынести ставки в отдельный сервис).
- UI (веб или мобильное приложение).
- История цен для машин (чтобы PricingService мог учитывать аналоги).





Отлично 🙌 Давай тогда разберём **пошаговый сценарий обработки запроса** внутри твоего первого микросервиса — **Vehicle Service**, чтобы стало понятно, как работает Clean Architecture на практике.
---
# 🔄 Пример: `POST /vehicles`

### 1. **Delivery (HTTP Handler)**
- Приходит HTTP‑запрос с JSON:
  ```json
  {
    "vin": "1HGCM82633A004352",
    "year": 2018,
    "odometer": 45000,
    "exteriorColor": "Black",
    "interiorColor": "Beige",
    "msrp": 35000
  }
  ```
- Handler (`vehicle_handler.go`) принимает запрос, валидирует JSON и вызывает **Usecase**.

---

### 2. **Usecase**
- Метод `CreateVehicle(v *Vehicle)`:
  - Проверяет бизнес‑правила (например, VIN = 17 символов).
  - Может вызвать **Inspection Service** (`GetBuildData`) для дополнения недостающих данных.
  - Передаёт сущность в **Repository** для сохранения.

---

### 3. **Domain**
- Сущность `Vehicle` описана в `domain/vehicle.go`.
- Здесь нет логики работы с БД или HTTP — только структура и базовые правила.

---

### 4. **Repository (интерфейс)**
- Usecase вызывает метод `Save(vehicle *Vehicle) error`.
- Repository — это интерфейс, он не знает, какая БД используется.

---

### 5. **Infrastructure (реализация)**
- Конкретная реализация `PostgresVehicleRepository` сохраняет данные в PostgreSQL.
- Если завтра решишь перейти на Mongo — меняется только этот слой.

---

### 6. **Ответ**
- Usecase возвращает результат в Handler.
- Handler формирует HTTP‑ответ:
  ```json
  {
    "status": "created",
    "vin": "1HGCM82633A004352"
  }
  ```

---

# 📌 Итог
- **Handler** принимает запрос.
- **Usecase** решает, что делать.
- **Domain** хранит сущности и правила.
- **Repository** — интерфейс.
- **Infrastructure** — конкретная реализация.

👉 В итоге бизнес‑логика не зависит от того, используешь ли ты Postgres, Mongo или даже in‑memory storage.

---

Хочешь, я распишу такой же **сценарий для `GET /vehicles/{vin}`**, чтобы у тебя был полный CRUD‑пример на Clean Architecture?




# vehicle service API examples
curl -i -X POST http://localhost:7071/vehicles \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","year":2022,"odometer":12000}'

curl -i http://localhost:8081/vehicles

curl -i http://localhost:8081/vehicles/1HGBH41JXMN109186

curl -i -X PUT http://localhost:8081/vehicles/1HGCM82633A004352 \
  -H "Content-Type: application/json" \
  -d '{
    "vin":"1HGCM82633A004352",
    "year":1999,
    "msrp":25999.99,
    "odometer":12500
  }'

curl -i -X DELETE http://localhost:8081/vehicles/5YJSA1E26MF168123

# Inspection Service API examples
curl -i http://localhost:7072/inspections/get-build-data/5YJSA1E26MF168123

curl -i -X POST http://localhost:8082/inspections/inspect \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","year":2022}'

curl -i -X POST http://localhost:7074/pricing/get-recommended-price \
  -H "Content-Type: application/json" \
  -d '{"vin":"5YJSA1E26MF168123","grade":47,"odometer":30000}'

curl -i -X GET http://localhost:7074/health

### Описание задачи - Inspection Service

InspectionService — сервис для получения дополнительных данных о машине (GetBuildData) и расчёта её оценочного состояния (InspectVehicle). Требования кратко:
- GetBuildData принимает только VIN и возвращает данные сборки (color, engine, transmission и т.д.).
- InspectVehicle принимает VIN, odometer, year и вычисляет grade ∈ [0,50] по набору правил и модификаторов.
- При интеграции с VehicleService данные из GetBuildData записываются в хранилище; если поле предоставлено пользователем, оно имеет приоритет над системным.

---

### API (HTTP/JSON)

1) POST /inspection/builddata
- Purpose: GetBuildData
- Request:
  - { "vin": "STRING" }
- Response 200:
  - {
      "vin":"STRING",
      "color":"STRING|null",
      "engine":"STRING|null",
      "transmission":"STRING|null",
      "trim":"STRING|null",
      "other":{...}
    }
- Errors:
  - 400 — invalid vin
  - 404 — not found (optional)

2) POST /inspection/inspect
- Purpose: InspectVehicle
- Request:
  - {
      "vin":"STRING",
      "year": INT,
      "odometer": INT,
      "small_scratches": BOOLEAN (optional, default false),
      "strong_scratches": BOOLEAN (optional, default false),
      "electrics_fail": BOOLEAN (optional, default false),
      "suspension_fail": BOOLEAN (optional, default false)
    }
- Response 200:
  - { "vin":"STRING", "grade": INT, "raw_score": FLOAT, "applied_modifiers":[ "small_scratches", ... ] }
- Errors:
  - 400 — validation failed (missing/invalid fields)

Notes:
- Use JSON content-type; prefer POST to keep body.

---

### Алгоритм расчёта grade

1. Базовая идеальная оценка = 50.
2. Возрастное уменьшение: age = currentYear - year. score_after_age = base - clamp(age, 0, 50)
   (каждый год минус 1 балл).
3. Ограничения по максимуму:
   - Если age > 15 → cap_max = 35.
   - Если odometer > 300000 → cap_max = 30.
   Постфактум ограничиваем итоговый результат сверху: final = min(calculated_value, cap_max_if_applicable, 50).
4. Модификаторы (мультипликативные): применяются последовательно к скору после age:
   - strong_scratches → multiply by 1/1.08 (уменьшение в 1.08 раза).
   - small_scratches → multiply by 1/1.04.
   - electrics_fail → multiply by 1/1.08.
   - suspension_fail → multiply by 1/1.06.
   Примечание: формула "reduces by 1.08 times" я трактую как разделение на 1.08 (score = score / 1.08).
5. Округление и финал:
   - После всех модификаторов округляем вниз до целого (floor).
   - Наложить минимум 0 и максимум 50, затем применить caps из (3).
6. Приоритет пользовательских данных:
   - Если VehicleService передаёт данные для того же VIN и пользователь также передал поля в запись машины, храните оба, при чтении отдавайте пользовательские поля первыми; GetBuildData при записи в storage не перезаписывает непустые пользовательские поля.

Пример выполнения:
- base=50, year=2018, current=2025 → age=7 → score=43.
- small_scratches=true → score = 43 / 1.04 ≈ 41.346 → final grade = 41.

---

### Валидация входных данных

- vin: обязательный, непустой, длина 11..20 (или точно 17 если стандартизируем). Валидировать regex: /^[A-HJ-NPR-Z0-9]{17}$/ (если требуем VIN17).
- year: целое, 1886 <= year <= currentYear.
- odometer: целое, 0 <= odometer <= 1_000_000 (ограничение для sanity).
- булевы поля опциональны.
- Если валидация не прошла — 400 с JSON { "error": "validation failed", "details": { field: msg } }.

---

### Схема хранения (простая, SQL)

Таблица inspections (для сохранённых данных сборки / результатов):

- inspections
  - vin VARCHAR PRIMARY KEY
  - color VARCHAR NULL
  - engine VARCHAR NULL
  - transmission VARCHAR NULL
  - trim VARCHAR NULL
  - source JSONB NULL (доп. данные)
  - user_overrides JSONB NULL  — хранит поля, заданные пользователем
  - last_inspected_at TIMESTAMP
  - last_grade INT
  - created_at TIMESTAMP
  - updated_at TIMESTAMP

Поведение записи GetBuildData:
- При получении системных данных:
  - INSERT ... ON CONFLICT (vin) DO UPDATE SET ... только для полей, которых нет в user_overrides (не перезаписывать user_provided).
  - user_overrides храним отдельно и не трогаем.

Пример SQL-upsert (psuedocode):
- Если record not exists → INSERT with build data in color/engine...
- Else → UPDATE for each field F: SET F = COALESCE(user_overrides->>F, incoming.F).

---

### Интеграция с VehicleService

- Когда VehicleService загружает список машин, оно вызывает InspectionService.GetBuildData(vin) для каждого VIN.
- InspectionService возвращает системные данные и сохраняет их в storage (см. логику записи выше).
- При возврате объединённого объекта VehicleService должен:
  - взять поля из user-provided (если есть),
  - иначе — из inspection storage,
  - вернуть merged result клиенту.

Рекомендация: вызвать InspectVehicle (расчёт grade) когда:
- новая запись добавлена,
- или при запросе детального view (on‑demand).
Сохранять last_grade и last_inspected_at для быстрых ответов.

---

### Пример реализации: Go — псевдокод handler

```go
// InspectRequest {Vin string; Year int; Odometer int; SmallScratches bool; StrongScratches bool; ElectricsFail bool; SuspensionFail bool}
func InspectHandler(w http.ResponseWriter, r *http.Request) {
  var req InspectRequest
  decodeJSON(r.Body, &req) // валидировать

  age := currentYear() - req.Year
  if age < 0 { age = 0 }

  score := 50.0 - float64(age)

  // modifiers (делим на фактор)
  if req.StrongScratches { score /= 1.08 }
  if req.SmallScratches  { score /= 1.04 }
  if req.ElectricsFail  { score /= 1.08 }
  if req.SuspensionFail { score /= 1.06 }

  // caps
  if age > 15 && score > 35 { score = 35 }
  if req.Odometer > 300000 && score > 30 { score = 30 }

  grade := int(math.Floor(score))
  if grade < 0 { grade = 0 }
  if grade > 50 { grade = 50 }

  // persist last grade, last_inspected_at
  upsertInspectionGrade(req.Vin, grade)

  writeJSON(w, InspectResponse{Vin:req.Vin, Grade:grade, RawScore:score, AppliedModifiers:...})
}
```

---

### Тесты (unit cases)

1. Age only:
- year = currentYear → grade = 50.
- year = currentYear - 3 → grade = 47.

2. Caps:
- year = currentYear - 16, odometer small → grade <= 35.
- odometer = 350000, year small → grade <= 30.

3. Multiplicative modifiers:
- small_scratches true on 50: 50/1.04 ≈ 48 -> floor 48.
- strong + electrics: apply both multiplicatively.

4. Combined:
- year = current-7 (score=43), odometer 400k (cap 30) → final = min(applied modifiers result, 30).

5. Edge:
- negative year, huge odometer → 0.

Добавить property tests: генерация случайных комбинаций, убедиться, что grade ∈ [0,50] и caps правильно применяются.

---

### Observability и производственные замечания

- Логировать входные данные и итоговую grade (без персональных данных) с correlation id.
- Экспорт метрик: count requests, histogram latency, histogram grade distribution.
- Ограничение запросов по VIN (rate limiting) и входная валидация, чтобы защититься от шумных вызовов.
- Версионировать API (например /v1/inspection/inspect).

---

Если хочешь, могу:
- сгенерировать готовые Go‑handlers + DTO + SQL миграции,
- подготовить OpenAPI спецификацию для быстрых интеграций,
- или написать unit тесты для алгоритма grade. Что делаем дальше?

protoc -I=services/inspection/delivery/grpc/proto \
  --go_out=paths=source_relative:services/inspection/delivery/grpc/proto \
  --go-grpc_out=paths=source_relative:services/inspection/delivery/grpc/proto \
  services/inspection/delivery/grpc/proto/inspection.proto

protoc -I=services/pricing/delivery/grpc/proto \
  --go_out=paths=source_relative:services/pricing/delivery/grpc/proto \
  --go-grpc_out=paths=source_relative:services/pricing/delivery/grpc/proto \
  services/pricing/delivery/grpc/proto/pricing.proto