# Base de datos

PostgreSQL es la persistencia principal del MVP.

El archivo `init/001_schema.sql` crea las tablas iniciales:

- `instructors`
- `training_programs`
- `fichas`
- `training_environments`
- `schedule_entries`
- `no_programming_days`

No incluye datos oficiales SENA ni codigos institucionales inventados.

El archivo `init/002_seed_demo.sql` carga datos de demostracion marcados como `Pendiente Demo`:

- Catalogos parametricos: instructores, programas, fichas y ambientes.
- Datos volumetricos: bloques de horario semanales y dias sin programacion.

Para recargar el seed en Docker se debe eliminar el volumen de PostgreSQL:

```powershell
docker compose down -v
docker compose up --build
```
