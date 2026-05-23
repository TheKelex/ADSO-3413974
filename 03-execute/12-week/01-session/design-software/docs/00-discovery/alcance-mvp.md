# Alcance MVP - Sistema de Horarios SENA

## Objetivo del MVP

Construir una primera version funcional para registrar, consultar y organizar horarios de formacion SENA asociados a instructores, fichas y ambientes de formacion. El MVP debe validar que el dominio puede salir del Excel hacia una aplicacion web con API en Go y persistencia en PostgreSQL, manteniendo las reglas de negocio en backend.

## Hipotesis de producto

Si el equipo responsable de programacion cuenta con una aplicacion que centraliza instructores, fichas, ambientes, horarios y observaciones, entonces podra consultar la programacion con mayor consistencia, reducir cruces basicos y dejar trazabilidad inicial de novedades sin depender exclusivamente del Excel.

## Capacidades Must Have

- Gestion basica de instructores, incluyendo tipo de instructor: planta o contratista.
- Gestion basica de fichas.
- Gestion basica de ambientes de formacion.
- Registro y consulta de programas de formacion sin inventar codigos institucionales.
- Registro de horarios o asignaciones de formacion relacionando instructor, ficha y ambiente.
- Registro de competencia, RAP y actividad de aprendizaje cuando el dato este disponible.
- Registro de observaciones o novedades asociadas a la programacion.
- Identificacion explicita de dias sin programacion.
- Validacion inicial de posibles cruces de horario por instructor, ficha y ambiente.

## Capacidades Should Have

- Filtros por fecha, ficha, instructor y ambiente.
- Vista resumida por dia o rango de fechas.
- Estados simples de programacion, por ejemplo pendiente de confirmacion o confirmada, sujetos a validacion.
- Exportacion simple para revision operativa, si no retrasa el MVP.

## Fuera de alcance del MVP

- Gestion individual de aprendices.
- Matricula, asistencia, evaluacion, evidencias y criterios de evaluacion.
- Inventario de ambientes o equipos.
- Integracion con sistemas institucionales externos.
- Autenticacion institucional compleja o SSO.
- Flujos formales de aprobacion multinivel.
- Optimizacion automatica de horarios.
- Microservicios definitivos por dominio.
- Reproduccion visual exacta del Excel.
- Codigos oficiales, sedes, centros, regionales o catalogos no suministrados.

## Entidades candidatas del MVP

- Instructor.
- Tipo de instructor.
- Ficha.
- Programa de formacion.
- Ambiente de formacion.
- Horario o asignacion.
- Competencia.
- RAP.
- Actividad de aprendizaje.
- Observacion o novedad.
- Dia sin programacion.

## Reglas minimas para validar en MVP

- No permitir o al menos advertir cruces de horario de un mismo instructor.
- No permitir o al menos advertir cruces de horario de una misma ficha.
- No permitir o al menos advertir cruces de horario de un mismo ambiente de formacion.
- Permitir registrar observaciones sin bloquear la programacion cuando falte informacion secundaria.
- Marcar como Pendiente cualquier dato institucional no confirmado.
- Mantener en backend las validaciones de cruces y reglas principales.

## Indicadores candidatos

- Porcentaje de horarios registrados con instructor, ficha y ambiente completos.
- Numero de cruces detectados por semana o periodo.
- Numero de dias sin programacion identificados por ficha o periodo.
- Tiempo estimado para consultar la programacion de una ficha.
- Porcentaje de asignaciones con observaciones o novedades registradas.

## Supuestos criticos

- El equipo puede proporcionar al menos una muestra del Excel actual para validar lenguaje y casos reales.
- Los usuarios iniciales aceptan una vista web distinta al Excel si conserva la informacion operativa necesaria.
- Los catalogos iniciales pueden cargarse manualmente o por importacion simple en una fase posterior.
- Las reglas de cruce de horarios se pueden definir con fecha, hora de inicio y hora de fin.
- PostgreSQL sera la fuente principal de persistencia.

## Pendientes por confirmar

- Si los cruces deben bloquear el guardado o solo generar advertencias.
- Definicion exacta de franja horaria y granularidad minima.
- Si las actividades se programan por dia individual, rango continuo o recurrencia.
- Si el MVP requiere importacion inicial desde Excel o solo registro manual.
- Roles de permiso: quien crea, edita, consulta o elimina programacion.
- Campos obligatorios para ficha, instructor, ambiente y programa.

## Riesgos de alcance

- Intentar cubrir asistencia, matricula o evaluacion antes de estabilizar horarios.
- Tratar competencias y RAP como catalogos completos sin fuente oficial.
- Diseñar pantallas con demasiada fidelidad al Excel y poco criterio de producto.
- Incluir importacion masiva antes de tener reglas de validacion claras.
- Definir microservicios antes de validar limites reales del dominio.

## Criterio de salida de discovery minimo

La fase queda lista para avanzar cuando los stakeholders confirmen que el MVP se concentra en programacion basica, cruces de horario, observaciones y dias sin programacion, dejando datos institucionales no confirmados como Pendiente.

## Recomendacion de siguiente activacion

Siguiente capacidad recomendada: A05 con `cap-product-definition` para convertir este alcance en PRD, backlog inicial y criterios de aceptacion. Si se prioriza arquitectura antes del PRD, activar A06 para una arquitectura inicial modular web/api/db basada en estos limites.
