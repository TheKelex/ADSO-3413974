# Estructura tecnica inicial - Sistema de Horarios SENA

## Objetivo tecnico

Definir una base tecnica minima para un MVP local ejecutable con contenedores separados: frontend React, API Go y base de datos PostgreSQL. Esta estructura no implementa todavia reglas completas de negocio ni microservicios finales; prepara el producto para evolucionar con separacion clara de responsabilidades.

## Principios iniciales

- El frontend no debe contener reglas de negocio criticas.
- La API en Go concentra validaciones, reglas de cruce, persistencia y contratos.
- PostgreSQL es la fuente principal de persistencia.
- Cada componente debe poder levantarse en su propio contenedor.
- El MVP inicia como arquitectura modular simple, no como microservicios definitivos.
- La estructura debe permitir crecimiento futuro hacia asistencia, inventario, matricula u otros modulos sin definirlos funcionalmente ahora.

## Separacion esperada

### web

Aplicacion React para uso operativo. Responsabilidades iniciales:

- Consultar y registrar programacion.
- Mostrar fichas, instructores, ambientes y horarios.
- Capturar observaciones o novedades.
- Presentar advertencias devueltas por la API.
- Consumir contratos HTTP de la API.

No debe decidir por si sola si existe un cruce de horario; esa validacion debe venir de la API.

### api

Servicio backend en Go. Responsabilidades iniciales:

- Exponer endpoints para catalogos y programacion.
- Validar reglas de negocio minimas.
- Detectar posibles cruces por instructor, ficha y ambiente.
- Persistir y consultar datos en PostgreSQL.
- Mantener contratos de entrada y salida para el frontend.
- Prepararse para auditoria y trazabilidad futura.

### db

Base de datos PostgreSQL. Responsabilidades iniciales:

- Persistir entidades candidatas del MVP.
- Soportar consultas por fecha, ficha, instructor y ambiente.
- Proteger integridad referencial basica.
- Preparar migraciones controladas desde la API o herramienta definida en arquitectura.

## Propuesta minima de carpetas

```text
design-software/
  docs/
    00-discovery/
      entendimiento-inicial.md
      alcance-mvp.md
      estructura-tecnica-inicial.md
  web/
    README.md
  api/
    README.md
  db/
    README.md
  deploy/
    local/
      README.md
```

Esta propuesta es de estructura, no de implementacion funcional completa. Los archivos README pueden crearse en una fase posterior para documentar comandos y contratos iniciales.

## Propuesta minima de contenedores

- `web`: contenedor para la aplicacion React.
- `api`: contenedor para el servicio Go.
- `db`: contenedor para PostgreSQL.

Para ejecucion local se espera un orquestador tipo Docker Compose con redes internas entre `web`, `api` y `db`. La definicion concreta del compose queda para fase tecnica posterior.

## Contratos tecnicos candidatos

- `web` consume HTTP desde `api`.
- `api` accede a `db` mediante variables de entorno.
- `db` expone solo el puerto necesario para desarrollo local.
- Las reglas de cruce viven en `api`, no en `web`.
- Las migraciones de esquema deben versionarse y ejecutarse de forma reproducible.

## Datos iniciales no definidos

Marcar como Pendiente hasta recibir fuente institucional:

- Codigos de programa.
- Codigos de competencia.
- Codigos o nombres oficiales de RAP.
- Numeros o identificadores reales de ficha.
- Nombres de ambientes de formacion.
- Centros de formacion, sedes y regionales.
- Calendarios institucionales o reglas oficiales de jornada.

## Decisiones diferidas

- Framework especifico de React y herramienta de build.
- Router frontend, libreria de formularios y estrategia de estado.
- Framework HTTP en Go.
- Herramienta de migraciones.
- Convencion final de errores de API.
- Estrategia de autenticacion.
- Importacion desde Excel.
- Modelo definitivo de recurrencia de horarios.

## Riesgos tecnicos

- Crear un modelo rigido antes de validar ejemplos reales del Excel.
- Duplicar reglas de cruce en frontend y backend.
- Dejar PostgreSQL como detalle accidental sin migraciones controladas.
- Sobre-disenar microservicios antes de tener limites funcionales claros.
- No separar observaciones operativas de datos academicos.

## Recomendacion de siguiente activacion

Activar A06 para `cap-software-patterns` y luego arquitectura inicial cuando el alcance MVP sea confirmado. Si aun faltan reglas de negocio o prioridades, activar A05 con `cap-product-definition` antes de arquitectura.
