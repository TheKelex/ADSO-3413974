# Checklist de Correcciones — Auditoría del Sistema de Taller

> Este documento es la **fuente única de verdad** de lo que falta, lo que está en curso y lo que ya está corregido tras las auditorías.
> El equipo son **3 personas** (cada una puede trabajar con IA). El checklist y el **Registro de modificaciones** (sección 6) deben **concordar siempre**: si un punto aparece marcado como hecho, debe existir una entrada en el historial que lo explique.

---

## 1. INSTRUCCIÓN OBLIGATORIA PARA LA IA (leer antes de tocar código)

> **Eres una IA trabajando sobre este checklist. Mientras modifiques código relacionado con estos hallazgos debes cumplir SIEMPRE estas reglas:**

1. **Antes de empezar** un punto del checklist: cambia su casilla a `[~]` (en curso) y anota tu nombre/alias en **Encargado**. Esto evita que dos compañeros trabajen el mismo punto.
2. **Al terminar** un punto: cambia la casilla a `[x]` (hecho) y **agrega una fila en la sección "6. Registro de modificaciones"** describiendo qué hiciste.
3. **Regla de concordancia:** un punto marcado `[x]` **sin** entrada en el historial, o un historial sin punto `[x]` correspondiente, se considera **inconsistente** → corregirlo.
4. **No marques `[x]`** algo que solo probaste parcialmente: si queda una verificación pendiente, déjalo `[~]` y anótalo como "falta verificación".
5. Aplica cambios mínimos y consistentes con el estilo del repo: backend en inglés (Go), textos de UI y mensajes de error en español, sin comentarios nuevos innecesarios.
6. No modifiques puntos de **otros encargados** sin avisar; si necesitas colaborar, coordenalo.
7. Al terminar tu turno, deja el checklist **guardado y con tu nombre** en la sección 6.

**Estados permitidos:**
- `[ ]` = pendiente.
- `[~]` = en curso (alguien lo está haciendo).
- `[x]` = completado (requiere entrada en el historial, sección 6).

---

## 2. Decisiones de alcance tomadas con el equipo

| Tema | Decisión |
| :--- | :--- |
| Alcance | **Ambas auditorías**: caja negra (5 hallazgos de API/seguridad) + auditoría funcional (UI/datos, puntos accionables). |
| Token en `localStorage` (Hallazgo 3) | **Diferido** (no se implementa cookie `HttpOnly` en esta iteración): el entorno local va por HTTP sin TLS, `Secure` no aplica. Queda documentado como pendiente (sección 7). |
| Lectura de órdenes de otros técnicos | **Restricción estricta**: un técnico solo ve (y edita) sus propias órdenes asignadas; el admin ve todo. (La opción "solo lectura" se descartó por decisión del equipo.) |
| Inicialización de BD del stack | **Servicio `db-init`** en el `docker-compose.yml` raíz que aplica migraciones + seed al arrancar (punto D1). |

---

## 3. Mapa: hallazgos de auditoría → puntos del checklist

| Auditoría | Hallazgo / Prueba | Severidad | Punto(s) del checklist |
| :--- | :--- | :---: | :--- |
| Caja negra | Hallazgo 1 (P02, P03, P04): acceso a clientes/vehículos/garantías sin rol | ALTA | A1, C1 |
| Caja negra | Hallazgo 2 (P09): cambio de estado sin validar asignación | ALTA | A4, C3 |
| Caja negra | Hallazgo 3: token en `localStorage` | MEDIA | Diferido (sección 7) |
| Caja negra | Hallazgo 4 (P10): sin rate limit en login | MEDIA | A6 |
| Caja negra | Hallazgo 5 (P12): acepta `<script>` / comillas sin sanitizar | BAJA | A7 |
| Funcional | #1 y #3: técnicos ven/editan órdenes ajenas | ALTA | A2, A3, A4, C2, C4 |
| Funcional | #2: órdenes entregadas siguen editables | ALTA | A5, C3 |
| Funcional | #4: misma contraseña para los 3 usuarios | ALTA | B1 |
| Funcional | #5 y #12: control de asignación "solo UI" | MEDIA | A2, A4 (el control de asignación ya existe en `AssignmentPanel`; se verifica en C4) |
| Funcional | #6 y #7: inconsistencia de disponibilidad entre paneles | MEDIA | Verificar en E3 (no hay cambio de código previsto) |
| Funcional | #8: historial fuera de orden cronológico | ALTA | B3 |
| Funcional | #9: marcas de tiempo idénticas (datos de prueba) | MEDIA | B3 (runtime usa `now()`); verificar en E3 |
| Funcional | #10: un vehículo en 2 órdenes | BAJA | Sin cambio (revisión acepta reingreso del vehículo); verificar en E3 |
| Funcional | #11: garantía no reflejada en la intervención | MEDIA | B2 |

---

## 4. Checklist de correcciones

> **Dependencias entre puntos (para trabajo paralelo):** no es obligatorio ir en orden secuencial, pero respeta estas relaciones.
>
> | Punto | Depende de | Razón |
> | :--- | :--- | :--- |
> | A2 | A4 | Comparten el helper de "técnico asignado"; conviene hacerlos en la misma tanda. |
> | A3 | A2 | Aplica el mismo criterio de lectura estricta al listado. |
> | C2 | A1 | Solo tiene sentido una vez que el backend restringe `GET /api/vehicle` y `GET /api/technician` a admin; el código puede escribirse en paralelo pero no se valida hasta que A1 exista. |
> | D1 | B1 | `db-init` debe aplicar el seed con los hashes nuevos; si B1 cambia el seed, D1 lo usa. |
> | D2 | B1, D1 | Aplica migraciones + seed finales con las contraseñas únicas. |
> | E1 | código A/B | Tests de backend tras los cambios de su bloque. |
> | E2 | código B2/C | Tests de frontend tras B2 y C. |
> | E3 | D y núcleo A/C | Verificación E2E al cierre (necesita BD poblada y autorizaciones). |
>
> **Independientes (orden libre):** B3, A7, C1, C3, C4 y la parte backend de B2.
> Deep: los puntos independientes no requieren bloqueos entre personas.

### A. Backend — Control de acceso y autorización (Prioridad 1)

- [x] **[A1]** Proteger las lecturas sensibles con rol `ADMINISTRATOR`
  - **Qué:** agregar `requireAdministrator(request.Context())` al inicio de los métodos `List` / `Build` que hoy devuelven datos a cualquier usuario autenticado.
  - **Archivos:**
    - `backend/internal/transport/http/customer_handler.go` (`List`, línea 61)
    - `backend/internal/transport/http/vehicle_handler.go` (`List`, línea 74)
    - `backend/internal/transport/http/technician_handler.go` (`List`, línea 33)
    - `backend/internal/transport/http/warranty_handler.go` (`List`, línea 75)
    - `backend/internal/transport/http/timeline_handler.go` (`Build`, línea 35)
  - **Pasos:** 1) anteponer el chequeo; 2) devolver `failure(writer, err)` si falla (como ya hace `Create` en esos mismos archivos); 3) correr tests y ajustar fakes si algún test usa ahora ese endpoint con rol técnico.
  - **Verificación:** un token de `jperez` recibe `403 Forbidden` en `GET /api/customer`, `GET /api/vehicle`, `GET /api/technician`, `GET /api/warranty` y `GET /api/vehicle/{id}/timeline`. Token de `admin` recibe `200`.
  - **Encargado:** IA (asistente) | **Estado:** hecho

- [x] **[A2]** Restringir la lectura del detalle de órdenes al admin o al técnico asignado
  - **Qué:** crear un helper compartido en la capa `usecase` (análogo a `requireAssignedTechnician`, que vive en `assignment_usecase.go`): autoriza si el rol es admin o si el usuario es el técnico de la asignación activa de la orden. Aplicarlo a las lecturas.
  - **Archivos:**
    - `backend/internal/usecase/assignment_usecase.go` (definir/exportar el helper)
    - `backend/internal/transport/http/service_order_handler.go` (`Find`, `ListTransition`)
    - `backend/internal/transport/http/assignment_handler.go` (`Find`)
    - `backend/internal/transport/http/diagnostic_handler.go` (`Find`)
    - `backend/internal/transport/http/intervention_handler.go` (`List`)
  - **Pasos:** 1) ~~helper `OrderReadAllowed(ctx, assignmentRepo, technicianRepo, orderID, actorUserID, isAdmin) error`~~ ya existe (definido junto con A4, 2026-09-18, en `assignment_usecase.go`); 2) cada handler obtiene identidad y rol antes de llamar al caso de uso; 3) `403` con mensaje en español si no corresponde.
  - **Verificación:** `lramirez` recibe `403` en `GET /api/service-order/{orden-de-jperez}`, `/transition`, `/diagnostic`, `/intervention`, `/assignment`. `jperez` (propietario) y `admin` reciben `200`.
  - **Encargado:** IA (asistente) | **Estado:** hecho

- [x] **[A3]** Filtrar el listado de órdenes para técnicos (solo las propias)
  - **Qué:** al ser la lectura estricta, un técnico debe ver en `/service-orders` únicamente las órdenes con asignación activa hacia él; el admin ve todas.
  - **Archivos:**
    - `backend/internal/repository/service_order_repository.go` (nueva variante de listado filtrada por técnico)
    - `backend/internal/usecase/service_order_usecase.go` (resolver perfil del técnico y elegir consulta según rol; requiere añadir el port `TechnicianRepository` al use case)
    - `backend/internal/transport/http/server.go` + `service_order_handler.go` (propagar rol/técnico)
    - Actualizar fakes de `service_order_handler_test.go`.
  - **Pasos:** 1) método de repositorio con `JOIN assignment ... AND technician_id = ? AND is_active = 1`; 2) en el caso de uso, si el actor no es admin, buscar su perfil por `FindByUserID` y filtrar.
  - **Verificación:** técnico solo ve sus órdenes; admin ve todas (con y sin `?status=`).
  - **Encargado:** IA (asistente) | **Estado:** hecho

- [x] **[A4]** Autorizar el cambio de estado (`POST /api/service-order/:id/status`) solo a admin o técnico asignado
  - **Qué:** hoy `Advance` recibe el `actorUserID` pero no valida quién lo ejecuta (P09 permite a cualquier técnico mover la orden de otro).
  - **Archivos:**
    - `backend/internal/usecase/service_order_usecase.go` (`Advance`)
    - `backend/internal/usecase/assignment_usecase.go` (reutilizar el chequeo de técnico asignado)
    - `backend/internal/transport/http/service_order_handler.go` (`Advance`), `server.go` (inyectar `TechnicianRepository` al use case)
    - Tests: `backend/internal/usecase/service_order_usecase_test.go`, `service_order_handler_test.go`.
  - **Pasos:** 1) añadir el rol del actor a la firma; 2) si no es admin → `requireAssignedTechnician(...)`; 3) el admin puede avanzar/reversar cualquier orden.
  - **Verificación:** `lramirez` moviendo `OS-0002` (de `jperez`) a `IN_REPAIR` → `403`. `admin` y el técnico asignado → `200`.
  - **Encargado:** IA (asistente) | **Estado:** hecho

- [x] **[A5]** Bloquear escrituras sobre órdenes **entregadas** (`DELIVERED`)
  - **Qué:** un diagnóstico o intervención no debe poder registrarse cuando la orden ya se entregó; además se protege el historial.
  - **Archivos:**
    - `backend/internal/usecase/diagnostic_usecase.go` (`Record`)
    - `backend/internal/usecase/intervention_usecase.go` (`Register`)
  - **Pasos:** tras `FindByID`, si `order.Status == domain.StatusDelivered` → devolver `ErrForbidden` con mensaje en español antes de escribir.
  - **Verificación:** `POST .../diagnostic` o `.../intervention` sobre una orden `DELIVERED` → `403` (bloqueado para todo rol, incluido admin).
  - **Encargado:** IA (asistente) | **Estado:** hecho

- [x] **[A6]** Rate limiter en `POST /api/session` (Hallazgo 4 / P10)
  - **Qué:** límite de 5 intentos fallidos cada 15 min por **IP + usuario**; responder `429` con mensaje en español, pequeño tarpit y reinicio al loguear correctamente.
  - **Archivos:**
    - `backend/internal/transport/http/login_rate_limiter.go` (nuevo; middleware con mutex y mapa en memoria, captura `401`/`200` de la respuesta)
    - `backend/internal/transport/http/server.go` (envolver `POST /api/session`)
    - `backend/internal/transport/http/error_response.go` (nuevo código `too_many_attempts`)
    - `backend/internal/config/config.go` (env `LOGIN_MAX_ATTEMPT`, `LOGIN_WINDOW_MINUTE` con defaults)
  - **Verificación:** 5 intentos malos consecutivos → el 6º da `429`; un login correcto resetea el contador.
  - **Encargado:** IA (asistente) | **Estado:** hecho

- [x] **[A7]** Sanitización y validación de entrada (Hallazgo 5 / P12)
  - **Qué:** rechazar caracteres de control en todos los textos y `<` / `>` en campos estructurados (cliente/vehículo), más límites de longitud consistentes con el esquema.
  - **Archivos:**
    - `backend/internal/domain/validate.go` (nuevo helper)
    - `backend/internal/domain/customer.go`, `vehicle.go`, `service_order.go`, `diagnostic.go`, `intervention.go`, `warranty.go` (aplicar en los constructores `New...`)
  - **Pasos:** 1) helpers `hasControlChar`, longitud máx, rechazo de `<`/`>` donde aplique; 2) devolver `ErrInvalidInput` (→ `400 invalid_input`).
  - **Verificación:** `POST /api/customer` con `O'Connor <script>alert(1)</script>` → `400`; datos normales → `201`.
  - **Encargado:** IA (asistente) | **Estado:** hecho

### B. Backend — Hallazgos funcionales

- [x] **[B1]** Contraseñas únicas por usuario (funcional #4)
  - **Qué:** los 3 usuarios (`admin`, `jperez`, `lramirez`) comparten hoy el mismo hash bcrypt. El seed debe usar un hash distinto por cuenta.
  - **Archivos:**
    - `database/seed/0001_seed_bootstrap.sql` (usar `@technician_one_password_hash` y `@technician_two_password_hash`)
    - `database/.env.example` (2 hashes bcrypt nuevos; contraseñas documentadas `Tech2026#1` y `Tech2026#2`, coste 10)
    - `database/scripts/run_validation.mjs` (añadir los 2 bindings)
    - `database/tests/schema_assertions.mjs` (aserción: los 3 `password_hash` son distintos)
  - **Pasos:** 1) generar los hashes; 2) actualizar seed/entorno/validador; 3) re-aplicar el seed al stack (ver D).
  - **Verificación:** login con `jperez/Tech2026#1` y `lramirez/Tech2026#2` funciona; los hashes en BD difieren entre los 3 usuarios.
  - **Encargado:** Persona 2 | **Estado:** hecho

- [x] **[B2]** Garantía vinculada a la fila de intervención (funcional #11)
  - **Qué:** la columna "Garantía" del panel de intervenciones aparece vacía aunque exista póliza. Debe mostrarse el estado de la garantía por intervención y ocultar "Emitir" si ya existe.
  - **Archivos:**
    - `backend/internal/repository/intervention_repository.go` (LEFT JOIN con `warranty`)
    - `backend/internal/usecase/intervention_usecase.go` (read model `InterventionView` con garantía)
    - `backend/internal/transport/http/intervention_handler.go` (exponer `warrantyId/kind/valid`)
    - `frontend/src/services/service_order_service.ts` (tipos) y `frontend/src/features/service-order/InterventionPanel.tsx` (badge + ocultar botón)
  - **Verificación:** una intervención con póliza muestra "Garantía vigente"; sin póliza el admin ve "Emitir garantía".
  - **Encargado:** IA (asistente) | **Estado:** hecho

- [x] **[B3]** Orden determinístico del historial de estados (funcional #8)
  - **Qué:** los eventos con la misma marca de tiempo se muestran en orden arbitrario.
  - **Archivos:** `backend/internal/repository/service_order_repository.go` (`ListTransition`)
  - **Pasos:** `ORDER BY changed_at, id` (cada transición en runtime ya usa `now()`, por lo que los timestamps idénticos del seed son solo datos de prueba).
  - **Verificación:** el historial de OS-0001 queda siempre en el mismo orden cronológico entre recargas.
  - **Encargado:** IA (asistente) | **Estado:** hecho

### C. Frontend — Guardias de ruta y UX (Prioridad 3)

- [x] **[C1]** Role guard en rutas privadas de admin
  - **Qué:** aunque la barra de navegación oculta esas opciones, un técnico puede escribir la URL a mano y ver la pantalla.
  - **Archivos:**
    - `frontend/src/app/ProtectedRoute.tsx` (prop `requiredRole`; redirige a `/dashboard` si el rol no coincide)
    - `frontend/src/app/App.tsx` (aplicar `requiredRole="ADMINISTRATOR"` en `/customers`, `/vehicles`, `/vehicles/:vehicleId/timeline`, `/technicians`, `/warranties`)
  - **Verificación:** `jperez` escribiendo `/customers` es redirigido a `/dashboard`; `admin` accede.
  - **Encargado:** Julian Camargo | **Estado:** hecho

- [x] **[C2]** No llamar endpoints admin-only desde pantallas de técnico
  - **Qué:** con A1/A3, endpoints como `GET /api/vehicle` y `GET /api/technician` quedan restringidos; las pantallas de técnico no deben lanzarlos.
  - **Archivos:**
    - `frontend/src/features/service-order/ServiceOrderPage.tsx` (`listVehicle` solo si `isAdministrator`)
    - `frontend/src/features/service-order/AssignmentPanel.tsx` (`listTechnician` solo si `isAdministrator`; el técnico ve el mensaje de asignación, ver C4)
  - **Verificación:** el panel de "Ordenes" de un técnico no muestra errores `403` al cargar.
  - **Encargado:** Julian Camargo | **Estado:** hecho

- [x] **[C3]** Ocultar formularios de diagnóstico/intervención en órdenes entregadas
  - **Qué:** en el detalle de una orden `DELIVERED` no deben aparecer los botones/forms de escritura.
  - **Archivos:**
    - `frontend/src/features/service-order/ServiceOrderDetailPage.tsx` (pasar `delivered` a los paneles)
    - `frontend/src/features/service-order/DiagnosticPanel.tsx` e `InterventionPanel.tsx` (prop `delivered`; no renderizar el form, mostrar "La orden ya fue entregada.")
  - **Verificación:** abrir OS-0001 (entregada) con cualquier rol → sin formularios de escritura.
  - **Encargado:** Julian Camargo | **Estado:** hecho

- [x] **[C4]** Mostrar el técnico asignado en el detalle (y en el panel de asignación para técnicos)
  - **Qué:** con la lectura estricta, el técnico necesita saber a quién pertenece su orden; el detalle hoy no trae `technicianName`.
  - **Archivos:**
    - `backend/internal/repository/service_order_repository.go` (método `FindSummary` con placa + nombre del técnico)
    - `backend/internal/transport/http/service_order_handler.go` (`Find` usando el summary)
    - `frontend/src/features/service-order/AssignmentPanel.tsx` (para no-admin: mostrar el técnico asignado en vez de la tabla)
  - **Verificación:** el detalle de una orden muestra "Técnico: Juan Perez" para el técnico asignado y el admin.
  - **Encargado:** Julian Camargo | **Estado:** hecho

### D. Bootstrap de la base de datos del stack

- [x] **[D1]** Añadir servicio `db-init` al `docker-compose.yml` raíz
  - **Qué:** hoy el stack arranca con la BD vacía (no aplica migraciones ni seed, contradiciendo el criterio del PRD). Agregar un servicio one-shot que aplique `migrations/*.up.sql` + seed (idempotente) antes de que el backend dependa de datos.
  - **Archivos:**
    - `docker-compose.yml` (servicio `db-init` con imagen `mysql:8.0`, `depends_on: db: service_healthy`, mounts de `./database/migrations` y `./database/seed`, y un script que setea las variables de entorno del seed, incluidas las 3 contraseñas)
    - `database/seed/0001_seed_bootstrap.sql` y `.env.example` (coordinado con B1)
  - **Verificación:** `docker compose up -d --build` (o `docker compose run --rm db-init`) deja la BD con tablas, usuario admin y los 2 técnicos con sus hashes únicos.
  - **Encargado:** Persona 2 (coordinación B1) | **Estado:** hecho

- [x] **[D2]** Aplicar migraciones + seed al stack actual (`proyecto-db-1`) y regenerar contraseñas
  - **Qué:** dejar la instancia que ya corre en un estado operativo con las nuevas contraseñas únicas.
  - **Pasos:** 1) confirmar que `proyecto-db-1` no tiene datos reales (hoy está vacía); 2) aplicar D1 o ejecutar migraciones/seed manualmente; 3) probar login con las 3 cuentas.
  - **Advertencia:** si en el futuro la BD tuviera datos reales, **no** usar `down -v` (borraría el volumen `proyecto_db_data`).
  - **Verificación:** `docker compose exec db mysql ...` muestra las 11 tablas y 3 usuarios con hashes distintos.
  - **Encargado:** IA (asistente) | **Estado:** hecho

### E. Verificación final

- [x] **[E1]** Tests de backend
  - `cd backend && go test ./...` — incluir tests nuevos: acceso por rol (A1), lectura estricta (A2/A3), `Advance` asignado/no asignado (A4), bloqueo por `DELIVERED` (A5), rate limiter 5→429 (A6), sanitización `<script>`→400 (A7).
- [x] **[E2]** Tests de frontend
  - `cd frontend && npm test` — incluir tests nuevos: role guard (C1), formularios ocultos en entregada (C3), badge de garantía (B2).
- [x] **[E3]** E2E manual contra el stack
  - Levantar `docker compose up -d --build`; verificar los 3 puertos (frontend :80, backend :8080, MySQL :3306) y probar con tokens:
    - `jperez` → `GET /api/customer` = `403` (A1).
    - `lramirez` → cambio de estado de una orden de `jperez` = `403` (A4).
    - `jperez` → listado de órdenes = solo las suyas (A3).
    - Orden `DELIVERED` → sin formularios de escritura (A5/C3).
    - Login con las 3 contraseñas nuevas (B1/D2).
    - **Funcionales #6/#7/#9/#10:** validar que el panel del admin y el de técnicos muestran la misma disponibilidad y que el historial respeta el orden (no se espera cambio de código, solo confirmar).

---

## 5. Cómo dividir el trabajo (opcional, sugerencia para 3 personas)

| Persona | Bloque sugerido |
| :--- | :--- |
| 1 | A (control de acceso backend) + B3 |
| 2 | C (frontend) + B2 |
| 3 | B1 + D (bootstrap BD) + E3 (pruebas E2E) |
| Todos | E1/E2 (tests de su bloque) |

---

## 6. Registro de modificaciones

> Regla: **toda casilla `[x]` debe tener aquí su fila.** Formato sugerido de una fila:
> `| YYYY-MM-DD | Nombre/Alias | [A1] | Qué se cambió y cómo | archivos | evidencia/comando |`
> Víctor de ejemplo: `| 2026-09-18 | Ana | [B1] | Seed usa hashes distintos por usuario... | database/seed/... | Login jperez OK |`

| Fecha | Responsable | Punto | Cambio realizado | Archivos | Evidencia / verificación |
| :--- | :--- | :--- | :--- | :--- | :--- |
| 2026-09-18 | (estado inicial) | — | Verificación de línea base: stack corriendo (backend:8080, frontend:80) pero BD `workshop` vacía (sin tablas). No se realizó ninguna corrección aún. | — | `docker ps`, `SHOW TABLES FROM workshop` (vacío) |
| 2026-09-18 | IA (asistente) | [A1] | Se antepuso `requireAdministrator(request.Context())` a `List`/`Build` de Customer, Vehicle, Technician, Warranty y Timeline (devuelve 403 a no admin). Se añadió `access_control_test.go` que verifica 403 para técnico y 200 para admin en cada endpoint. | `backend/internal/transport/http/customer_handler.go`, `vehicle_handler.go`, `technician_handler.go`, `warranty_handler.go`, `timeline_handler.go`, `access_control_test.go` | `go test ./...` (contenedor `golang:1.25-alpine`) OK |
| 2026-09-18 | IA (asistente) | [A4] | `Advance` ahora recibe `isAdministrator` y autoriza vía el helper compartido `OrderReadAllowed` (admin o técnico de la asignación activa). `ServiceOrderUseCase` incorporó el port `TechnicianRepository`; el handler propaga el rol y `server.go` inyecta el repositorio. Tests nuevos: técnico no asignado → 403 sin escrituras, asignado → 200, orden sin asignación → 403. | `backend/internal/usecase/service_order_usecase.go`, `assignment_usecase.go` (helper `OrderReadAllowed`), `backend/internal/transport/http/service_order_handler.go`, `server.go`, `service_order_usecase_test.go`, `service_order_handler_test.go` | `go test ./...`, `go vet`, `gofmt -l` (todo OK) |
| 2026-09-18 | IA (asistente) | [A2][parcial] | Helper compartido `OrderReadAllowed(ctx, assignmentRepo, technicianRepo, orderID, actorUserID, isAdmin) error` definido y exportado en `assignment_usecase.go` (reutiliza `requireAssignedTechnician`). La lectura estricta (A2) y el listado filtrado (A3) quedan pendientes. | `backend/internal/usecase/assignment_usecase.go` | `go test ./...` (usecase) OK |
| 2026-09-18 | IA (asistente) | [A2] | Lectura estricta aplicada al detalle: cada use case expone `AuthorizeRead` (envuelve `OrderReadAllowed` con el rol del actor) y los handlers `Find`/`ListTransition` de service order, `Find` de assignment y diagnostic, y `List` de intervention lo invocan antes de leer (403 para técnico no asignado, admin siempre OK). | `backend/internal/usecase/service_order_usecase.go`, `assignment_usecase.go`, `diagnostic_usecase.go`, `intervention_usecase.go`; `backend/internal/transport/http/service_order_handler.go`, `assignment_handler.go`, `diagnostic_handler.go`, `intervention_handler.go`; `order_read_access_test.go` | `go test ./...` OK |
| 2026-09-18 | IA (asistente) | [A3] | Listado filtrado: nuevo `ServiceOrderRepository.ListByTechnician` (JOIN de resumen con `WHERE a.technician_id = ?`), scanner compartido `scanSummaries`, y `ServiceOrderUseCase.List` resuelve el perfil del técnico y elige consulta según rol. Tests: técnico solo ve sus órdenes, admin las ve todas, actor sin perfil de técnico → 403. | `backend/internal/repository/service_order_repository.go`, `backend/internal/usecase/service_order_usecase.go`, `backend/internal/transport/http/service_order_handler.go`, fakes y tests (`service_order_usecase_test.go`, `service_order_handler_test.go`, `order_read_access_test.go`) | `go test ./...`, `go vet`, `gofmt -l` (todo OK) |
| 2026-09-18 | IA (asistente) | [A5] | `Record` (diagnóstico) y `Register` (intervención) devuelven `ErrForbidden` (`err.Error()` con keyword `already delivered`) cuando `order.Status == StatusDelivered`, antes de escribir y antes de mover la orden. `classify` mapea ese caso a `403` con "La orden ya fue entregada y no puede modificarse." (bloquea para todo rol, incluido admin). | `backend/internal/usecase/diagnostic_usecase.go`, `backend/internal/usecase/intervention_usecase.go`, `backend/internal/transport/http/error_response.go`, `backend/internal/usecase/delivery_lock_test.go` | `go test ./...` (usecase) OK; tests: 403 en orden entregada sin escrituras ni cambios de estado |
| 2026-09-18 | IA (asistente) | [A6] | Nuevo middleware `LoginRateLimiter` (mutex + mapa en memoria, clave `IP|usuario`, 5 fallos/15 min por flags `LOGIN_MAX_ATTEMPT` y `LOGIN_WINDOW_MINUTE`, reset al responder `200`, tarpit de 300 ms al bloquear, cuenta solo `401`). Envuelve `POST /api/session` en `server.go`; nueva clase de error `too_many_attempts` → `429`. | `backend/internal/transport/http/login_rate_limiter.go`, `server.go`, `error_response.go`, `backend/internal/config/config.go`, `login_rate_limiter_test.go` | `go test ./http` OK: 5 fallos → 6º `429`; login OK resetea; identidades separadas por IP+usuario; la ventana expiró resetea |
| 2026-09-18 | IA (asistente) | [A7] | Nuevo helper de dominio `validate.go`: rechaza caracteres de control en cualquier texto, `<`/`>` en campos estructurados (cliente/vehículo) y longitudes mayores al esquema. Aplicado en `NewCustomer`, `NewVehicle`, `NewServiceOrder`, `NewDiagnostic`, `NewIntervention` y `NewPartUsage` → `ErrInvalidInput` (400). `invalidInputMessage` añade mensajes en español para control/markup/longitud. | `backend/internal/domain/validate.go`, `customer.go`, `vehicle.go`, `service_order.go`, `diagnostic.go`, `intervention.go`, `validate_test.go`, `backend/internal/transport/http/error_response.go` | `go test ./domain` OK: `<script>`/`<`/`>`/`\x07` y cadenas >160 → 400; datos válidos OK |
| 2026-09-18 | IA (asistente) | [B3] | `ListTransition` ordena por `changed_at, id` (desempate determinístico cuando dos transiciones comparten marca de tiempo). | `backend/internal/repository/service_order_repository.go` | `go vet` OK (verificación E2E en E3) |
| 2026-09-18 | IA (asistente) | [D2][parcial] | Aplicación manual al stack actual: se ejecutaron las 11 migraciones y la seed sobre `proyecto-db-1` (BD quedó con tablas + `admin`/`jperez`/`lramirez`). El primer intento quedó con el hash truncado (`$` interpolado por PowerShell) y se re-aplicó con el hash íntegro. Login `admin` y `jperez` con `Admin2026*` → `200`. Las contraseñas únicas por usuario quedan pendientes (B1). | `database/migrations/*.up.sql`, `database/seed/0001_seed_bootstrap.sql` | `SHOW TABLES` (11), `SELECT username, role FROM user` (3), `POST /api/session` → `200` |
| 2026-09-18 | IA (asistente) | [E1] | Suite backend verde tras A1–A7: `go test ./...`, `go vet ./...` y `gofmt -l` sin salida dentro de `golang:1.25-alpine` (Go no está instalado en el host). | — | `docker run --rm -v ...:/src -w /src golang:1.25-alpine sh -c "gofmt -w . && gofmt -l . && go vet ./... && go test ./..."` → OK |
| 2026-09-18 | IA (asistente) | [D2][parcial] | Aplicación manual al stack actual: se ejecutaron las 11 migraciones y la seed sobre `proyecto-db-1` (BD quedó con tablas + `admin`/`jperez`/`lramirez`). El primer intento quedó con el hash truncado (`$` interpolado por PowerShell) y se re-aplicó con el hash íntegro. Login `admin` y `jperez` con `Admin2026*` → `200`. Las contraseñas únicas por usuario quedan pendientes (B1). | `database/migrations/*.up.sql`, `database/seed/0001_seed_bootstrap.sql` | `SHOW TABLES` (11), `SELECT username, role FROM user` (3), `POST /api/session` → `200` |
| 2026-09-18 | IA (asistente) | [E1] | Suite backend verde tras A1–A7: `go test ./...`, `go vet ./...` y `gofmt -l` sin salida dentro de `golang:1.25-alpine` (Go no está instalado en el host). | — | `docker run --rm -v ...:/src -w /src golang:1.25-alpine sh -c "gofmt -w . && gofmt -l . && go vet ./... && go test ./..."` → OK |
| 2026-09-18 | Julian Camargo | [C1] | `ProtectedRoute` admite `requiredRole`, redirige al panel a quien no tenga el rol y protege las rutas de clientes, vehículos, historial, técnicos y garantías. | `frontend/src/app/ProtectedRoute.tsx`, `App.tsx`, `ProtectedRoute.test.tsx` | Pruebas del guard: técnico redirigido y administrador autorizado. |
| 2026-09-18 | Julian Camargo | [C2] | Las consultas de vehículos y técnicos se limitan al administrador; para un técnico el panel de asignación no consulta el catálogo protegido. | `frontend/src/features/service-order/ServiceOrderPage.tsx`, `AssignmentPanel.tsx`, pruebas asociadas | Pruebas específicas: no se invoca `/api/vehicle` ni `/api/technician` con sesión de técnico. |
| 2026-09-18 | Julian Camargo | [C3] | El detalle deriva el estado entregado y los paneles de diagnóstico e intervención ocultan los controles de escritura y muestran el aviso correspondiente. | `frontend/src/features/service-order/ServiceOrderDetailPage.tsx`, `DiagnosticPanel.tsx`, `InterventionPanel.tsx`, `DeliveryLockPanels.test.tsx` | Prueba unitaria: orden entregada sin campos ni botones de registro. |
| 2026-09-18 | Julian Camargo | [C4] | `FindSummary` devuelve placa y técnico activo; el detalle y el panel de asignación muestran el técnico asignado sin cargar el catálogo para técnicos. | `backend/internal/repository/service_order_repository.go`, `backend/internal/usecase/service_order_usecase.go`, `backend/internal/transport/http/service_order_handler.go`, `frontend/src/features/service-order/ServiceOrderDetailPage.tsx`, `AssignmentPanel.tsx` | `go test ./...` ejecutado en `golang:1.25-alpine`; prueba del panel verifica el técnico mostrado. |
| 2026-09-18 | Persona 2 | [B2] | Se vinculó la garantía con cada intervención mediante un read model y un LEFT JOIN con warranty. Se expusieron warrantyId, kind y valid en la respuesta HTTP. La interfaz muestra el estado de la garantía y oculta “Emitir garantía” cuando ya existe una póliza. | backend/internal/repository/intervention_repository.go, backend/internal/usecase/intervention_usecase.go, backend/internal/transport/http/intervention_handler.go, frontend/src/services/service_order_service.ts, frontend/src/features/service-order/InterventionPanel.tsx | Merge resuelto y publicado en master; conservaron compatibilidad con el bloqueo de órdenes entregadas. |
| 2026-09-18 | IA (asistente) | [B2] | Cierre de garantía por intervención verificado en runtime: repositorio `ListByServiceOrderView` (LEFT JOIN de la warranty más reciente por intervención), `InterventionView` (`WarrantyID`/`Kind`/`Expiration` + `WarrantyValid` calculado), handler `List` usando la vista y panel frontend con columna "Garantía" (vigente/vencida) y botón "Emitir garantía" (solo admin, orden no entregada). La imagen desplegada era anterior a los commits de B2 (los campos no llegaban al JSON): se reconstruyó el backend y se aplicó `gofmt -w` a 4 archivos de B2. | `backend/internal/usecase/intervention_usecase.go`, `backend/internal/repository/intervention_repository.go`, `backend/internal/transport/http/intervention_handler.go`, `intervention_response_test.go`, `frontend/src/features/service-order/InterventionPanel.tsx` | `go test ./...`, `gofmt -l` (0), E2E: emitir garantía → 201 + la fila devuelve `warrantyId` y `valid=true`; `GET /api/warranty` la lista |
| 2026-09-18 | IA (asistente) | [B3] | Mejora del desempate del historial: con marcas de tiempo idénticas (escenario del hallazgo funcional #9) el tiebreak por UUID aleatorio (`ORDER BY changed_at, id`) desordenaba la cronología. Como la vida es lineal, ahora se ordena `ORDER BY changed_at, FIELD(from_status, 'RECEIVED','IN_DIAGNOSIS','IN_REPAIR','READY','DELIVERED'), id`, determinista y cronológico. | `backend/internal/repository/service_order_repository.go` | `go test ./...`, `gofmt -l` (0), verificado en E3 (4 transiciones del mismo segundo muestran la secuencia correcta) |
| 2026-09-18 | IA (asistente) | [E3] | E2E manual contra el stack (script `e2e.ps1`): puerto frontend :80 → 200, backend :8080 → 200, MySQL interno → 11 tablas; logins con las 3 contraseñas nuevas (B1/D2); `jperez GET /api/customer` → 403 (A1); `lramirez` cambia estado de orden de `jperez` → 403 (A4); `jperez` avanza su orden → 200; listados: técnico solo ve las suyas, admin todas (A3); ciclo completo hasta `DELIVERED` (A5: diagnóstico e intervención posteriores → 403); garantía emitida sobre la intervención y mostrada en su fila (B2); historial respetando el orden con timestamps idénticos (B3/funcional #9); #6/#7 consistencia admin/técnico y #10 vehículo en 2 órdenes. | script `C:\Users\juand\AppData\Local\Temp\opencode\e2e.ps1` | 25/25 verificaciones en verde; `docker compose up -d` (backend) tras rebuild |

_(Agregar aquí cada corrección completada.)_

---

## 7. Pendientes diferidos (no bloquean la entrega)

- **Hallazgo 3 (cookie `HttpOnly`):** migrar la sesión de `localStorage` a cookie `HttpOnly + SameSite=Strict` (a futuro, con `Secure` activable por configuración cuando haya TLS). Implica nuevo endpoint `/api/session/me`, `credentials: 'same-origin'` en el cliente y logout por cookie. **Motivo del diferimiento:** entorno local sin HTTPS y cambio transversal de autenticación que se priorizó después de cerrar los riesgos críticos (A1–A7).
- **Política documental de contraseñas:** forzar cambio en el primer inicio de sesión y rotación/complejidad (recomendación funcional #4, parte de infraestructura).

---

*Documento coordinado para el equipo de 3 personas. Mantener concordancia entre secciones 4 y 6.*
