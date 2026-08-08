# PRD — Venta y control de tiquetes aéreos

## Contexto de negocio

Una aerolínea necesita vender y controlar tiquetes aéreos: un pasajero hace una
reserva, de la reserva se emite un tiquete para un vuelo, se le asigna un asiento
de la aeronave, puede registrar equipaje y pagar, y al viajar se registra su
embarque. El sistema debe permitir consultar, entre otras cosas, los pasajeros
que compraron tiquete pero **no viajaron** (no-show).

## Usuarios

- **Agente de la aerolínea** (rol único del MVP): inicia sesión y gestiona
  reservas, tiquetes, vuelos, asignación de asientos, equipaje, pagos y embarques,
  y consulta reportes. El sistema arranca con un agente administrador sembrado.

## Entidades (modelo de dominio)

1. **passenger** — id: document · nombre, birth_date. Compra y viaja.
2. **reservation** — id: code · date, status. Solicitud previa al tiquete.
3. **ticket** — id: number · issue_date, service_class. Se genera desde la reserva.
4. **flight** — **identificación compuesta**: flight_number + departure_date ·
   scheduled_time. El número se repite por fecha.
5. **airport** — id: code · nombre, ciudad. **Una sola entidad con doble rol**:
   origen y destino de un vuelo.
6. **aircraft** — id: registration (matrícula) · model, capacity. Asignada al vuelo.
7. **seat** — **entidad débil** respecto a aircraft: seat_number + aircraft_registration ·
   row, location.
8. **payment** — id: reference · date, amount. **Opcional** (un tiquete puede o no
   estar pagado).
9. **baggage** — id: tag · weight, status. **Opcional; 0..N** por pasajero/tiquete.
10. **boarding** — derivada del tiquete · entry_time, gate, presentation_condition.
    **Opcional** → su ausencia identifica al pasajero que no viajó (no-show).

## Puntos de modelado a resolver (obligatorios)

- **airport** genera dos relaciones distintas con **flight**: `departs_from` y
  `arrives_at` (doble rol de la misma entidad).
- La asignación **asiento ↔ pasajero depende del vuelo** → relación ternaria o
  entidad asociativa **seat_assignment** (seat, passenger, flight).
- **boarding** con participación **opcional** es la clave de la consulta de
  pasajeros que no viajaron.
- **seat** y su identificación dependen de **aircraft** (entidad débil / id compuesto).

## Requisitos funcionales

1. El agente inicia sesión con credenciales válidas y recibe un token; las
   inválidas se rechazan con un error claro.
2. El agente crea y lista pasajeros (documento, nombre, fecha de nacimiento).
3. El agente crea reservas para un pasajero y las lista con su estado.
4. El agente emite un tiquete a partir de una reserva, con clase de servicio.
5. El agente crea vuelos con número + fecha de salida (id compuesto), origen y
   destino (dos aeropuertos distintos) y aeronave asignada.
6. El agente asigna un asiento de la aeronave del vuelo a un pasajero con tiquete
   (seat_assignment: asiento + pasajero + vuelo); un asiento no puede asignarse dos
   veces en el mismo vuelo.
7. El agente registra equipaje (0..N) y pagos (opcional) asociados al tiquete.
8. El agente registra el embarque de un pasajero (opcional).
9. El agente consulta el reporte de **pasajeros con tiquete que no registraron
   embarque** (no-show) para un vuelo.

## Requisitos no funcionales

- **security**: rutas de escritura con token válido; contraseñas hasheadas.
- **reliability**: un asiento asignado en un vuelo no puede reasignarse
  (unicidad); operaciones atómicas.
- **performance**: consultas de listado bajo 500 ms para hasta 10.000 tiquetes.
- **maintainability**: backend, frontend y base de datos como proyectos separados;
  identificadores técnicos en inglés y singular.
- **usability**: mensajes al usuario en español; fechas y valores con formato local.

## Criterios de aceptación

- Un usuario inicia sesión, crea un pasajero, una reserva, emite un tiquete, crea
  un vuelo con origen/destino distintos, asigna un asiento, y consulta el reporte
  de no-show correctamente.
- El despliegue arranca con `docker compose` (backend, frontend, base de datos con
  agente sembrado) y **la aplicación queda arriba y navegable**.
- Existe evidencia de pruebas automatizadas de backend y frontend.

## Stack

- Backend: Go · Frontend: Angular · Base de datos: PostgreSQL · Operación: Docker Compose
