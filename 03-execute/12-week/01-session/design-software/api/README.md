# API Go

Servicio backend del MVP. Expone endpoints HTTP para catalogos, programacion y dias sin programacion.

## Variables de entorno

- `DATABASE_URL`: cadena PostgreSQL. Valor local por defecto: `postgres://sena:sena@localhost:5432/sena_horarios?sslmode=disable`
- `CORS_ORIGIN`: origen permitido para el frontend. Valor por defecto: `http://localhost:5173`
- `PORT`: puerto HTTP. Valor por defecto: `8080`

## Endpoints iniciales

- `GET /health`
- `GET|POST /api/instructors`
- `GET|POST /api/programs`
- `GET|POST /api/fichas`
- `GET|POST /api/environments`
- `GET|POST /api/schedules`
- `GET|POST /api/no-programming-days`

La validacion de cruces de horario se ejecuta en `POST /api/schedules`.

