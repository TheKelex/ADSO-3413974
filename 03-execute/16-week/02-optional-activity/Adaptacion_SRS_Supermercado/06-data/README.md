# Datos

> Estado: 🟡 Pendiente | Última actualización: 2026-07-03
> Autor: Kevin Santiago Saavedra Chantris | Equipo: Kevin Santiago Saavedra Chantris

## Contenido

Documenta modelos de datos, diccionario y estrategia de migración.

> **Diferencia con `02-domain`:** esta sección describe la implementación de persistencia (tablas, columnas, tipos, relaciones, índices, migraciones). Las entidades de negocio y sus reglas se documentan en [`02-domain/`](../02-domain/). Un mismo concepto (ej: "Horario") tendrá una definición de dominio allá y un esquema de base de datos aquí.

> **Diferencia con `09-microservices`:** aquí vive el modelo de datos **global** del sistema — diccionario conceptual compartido, estrategia de migración y analítica. El modelo **transaccional propio de cada servicio** (esquema local, tablas internas) se documenta en `09-microservices/services/<servicio>/data-model.md`. No duplicar: si un concepto es solo del servicio, va allá; si es compartido entre varios servicios, va aquí.

## Archivos

| Archivo | Descripción | Estado |
|---------|-------------|--------|
| [models.md](./models.md) | Modelos conceptuales, lógicos o físicos de datos | 🟡 |
| [data-dictionary.md](./data-dictionary.md) | Diccionario de entidades, campos y reglas | 🟡 |
| [migration-strategy.md](./migration-strategy.md) | Estrategia para cambios y migraciones de datos | 🟡 |
