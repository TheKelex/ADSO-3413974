# Sistema de Horarios SENA

MVP local para gestionar programacion basica de formacion, relacionando instructores, fichas, ambientes de formacion, horarios, observaciones y dias sin programacion.

## Componentes

- `web`: aplicacion React con Vite.
- `api`: servicio Go que concentra reglas de negocio iniciales.
- `db`: PostgreSQL con esquema base del MVP.

## Ejecutar localmente con contenedores

```powershell
docker compose up --build
```

Luego abrir:

- Web: http://localhost:5173
- API health: http://localhost:8080/health
- PostgreSQL: `localhost:5432`, base `sena_horarios`, usuario `sena`, clave `sena`

Si abres solo la web y la API no esta disponible, la interfaz entra en modo demo local con `localStorage` para permitir ver y probar el flujo MVP. La persistencia real se usa cuando `api` y `db` estan levantados.

## Seed demo

La base local carga seed automatico al crear el volumen de PostgreSQL:

- `db/init/001_schema.sql`: esquema inicial.
- `db/init/002_seed_demo.sql`: datos parametricos y volumetricos de demostracion.

El seed incluye instructores de planta y contratistas, programas/fichas/ambientes marcados como `Pendiente Demo`, 12 bloques de horario y dias sin programacion. No contiene codigos oficiales SENA inventados.

Si ya existe el volumen de PostgreSQL y quieres recargar el seed desde cero:

```powershell
docker compose down -v
docker compose up --build
```

## Alcance materializado

- Creacion de instructores con tipo `planta` o `contratista`.
- Creacion de programas de formacion, fichas y ambientes de formacion.
- Registro de horarios con fecha, hora inicio, hora fin, instructor, ficha y ambiente.
- Campos opcionales para competencia, RAP, actividad de aprendizaje y observaciones.
- Registro de dias sin programacion.
- Validacion en API de cruces por instructor, ficha y ambiente en la misma fecha/franja.
- Vista sintetizada de horario por instructor.

## Pendientes intencionales

- Importacion desde Excel.
- Autenticacion y roles.
- Edicion/eliminacion de registros.
- Catalogos oficiales institucionales.
- Modelo avanzado de recurrencia.
- Auditoria formal de cambios.
