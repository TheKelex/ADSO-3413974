# Convenciones Ágiles del Equipo (Agile Team Conventions)

> Define cómo trabaja el equipo a lo largo de sus ciclos de desarrollo. Debe ser acordado y firmado por todo el equipo antes del primer sprint. Se actualizará cuando el equipo decida cambiar algo.

---

## Estructura del Sprint

| Campo | Valor |
|-------|-------|
| Duración | 2 semanas |
| Inicio del Sprint | Lunes |
| Fin del Sprint | Viernes de la semana 2 |
| Cronograma del Proyecto | 8 meses (1 de septiembre de 2026 — 30 de abril de 2027) |
| Sprint Actual | Sprint 1 — 1 de septiembre de 2026 al 11 de septiembre de 2026 |
| Capacidad Estimada | 45–55 puntos de historia por sprint (Basado en 10 desarrolladores activos) |

---

## Ceremonias

### Planificación del Sprint (Sprint Planning)
- **Cuándo:** Primer lunes del sprint — 09:00 AM
- **Duración:** Máximo 2 horas
- **Quiénes:** Todo el equipo (1 Tech Lead, 1 Arquitecto, 1 Delivery Manager, 2 QAs, 5 Desarrolladores). *Nota: Como parte de la estructura de aprendizaje, los 10 miembros programan activamente y participan en todas las fases técnicas.*
- **Objetivo:** Seleccionar y comprometerse con las historias de usuario del sprint, desglosándolas en tareas técnicas.
- **Artefacto de salida:** Sprint Backlog actualizado en GitHub Projects.

### Sincronización Diaria (Daily Stand-up)
- **Cuándo:** Todos los días — 09:15 AM
- **Duración:** Máximo 15 minutos (Seguimiento estricto del tiempo ya que es un equipo grande de 10 personas)
- **Formato:**
  1. ¿Qué hice ayer?
  2. ¿Qué voy a hacer hoy?
  3. ¿Tengo algún impedimento o bloqueo?
- **Regla:** Las discusiones técnicas y arquitectónicas profundas ocurren después de la daily (estacionando el tema), nunca durante ella.

### Revisión del Sprint (Sprint Review)
- **Cuándo:** Último viernes del sprint — 03:00 PM
- **Duración:** Máximo 45 minutos
- **Quiénes:** Todo el equipo + Product Owner + Interesados clave (Jefes de Taller)
- **Objetivo:** Mostrar software totalmente funcional (ej: lógica central del dominio, validación de garantías o el motor de analítica predictiva) y recolectar retroalimentación.

### Retrospectiva del Sprint (Sprint Retrospective)
- **Cuándo:** Último viernes del sprint — 04:00 PM (después de la revisión)
- **Duración:** Máximo 45 minutes
- **Formato:** ¿Qué salió bien? / ¿Qué podemos mejorar? / Compromisos de acción
- **Regla:** Cada retrospectiva debe generar al menos 1 acción de mejora con un responsable asignado y una fecha límite.

### Refinamiento del Backlog (Backlog Refinement)
- **Cuándo:** Miércoles de la segunda semana (mitad del sprint) — 02:00 PM
- **Duración:** Máximo 1.5 horas (Extendido para permitir debates técnicos profundos entre el Tech Lead, Arquitecto y Desarrolladores).
- **Objetivo:** Detallar, dividir y estimar historias de usuario para el próximo sprint.
- **Criterio de salida:** La historia de usuario cumple con la Definición de Listo (DoR).

---

## Estimación

### Escala

| Puntos | Significado |
|--------|-------------|
| 1 | Trivial — se hace en horas (ej: añadir un campo simple de marca de vehículo) |
| 2 | Pequeño — se hace en un día (ej: crear un validador de reglas de garantía) |
| 3 | Mediano — toma 2–3 días (ej: construir la lógica central de seguimiento de servicios) |
| 5 | Grande — toma casi un sprint completo (ej: implementar el algoritmo de analítica predictiva) |
| 8 | Muy grande — debe ser dividido |
| 13 | Épica — DEBE ser dividida obligatoriamente antes de entrar al sprint |

**Técnica:** Planning Poker  
**Herramienta:** Hatjitsu / PlanitPoker

### Regla de estimación
- Si hay un desacuerdo de 2 o más niveles (ej: alguien vota 2 y otro 5), el Arquitecto Líder y el Tech Lead guiarán una discusión técnica sobre la implementación antes de votar nuevamente.
- Si una historia se estima en 5 puntos o más, debe ser revisada minuciosamente y dividida en subtareas más pequeñas o características independientes si es posible.

---

## Herramienta de Backlog

**Herramienta:** GitHub Projects (integrado directamente con el código del repositorio)  
**URL del Tablero:** [Inserta la URL de tu proyecto en GitHub aquí]

### Columnas del tablero

| Columna | Significado |
|---------|-------------|
| Backlog | Pendiente de refinamiento y priorización de negocio |
| Ready (Listo) | Listo para entrar al sprint (cumple con el DoR) |
| In Progress | Un desarrollador está trabajando activamente en ello |
| In Review | En Pull Request / Revisión de código por un par y aprobación técnica del Tech Lead/Arquitecto |
| QA Testing | Columna de validación y pruebas funcionales para los 2 QAs dedicados |
| Done (Hecho) | Cumple con el DoD, pasó todas las tuberías automatizadas de CI/CD y está cerrado |

---

## Velocidad del Equipo

| Sprint | Puntos de historia completados | Notas |
|--------|--------------------------------|-------|
| Sprint 1 | — | Configuración base, alineación de 01-context y estructura arquitectónica inicial |
| Sprint 2 | — | — |
| Sprint 3 | — | — |
| Promedio | — | — |

---

## Documentos relacionados

- Definición de Listo (DoR) → `00-governance/definition-of-ready.md`
- Definición de Hecho (DoD) → `00-governance/definition-of-done.md`
- Gestión de riesgos → `15-project-control/risks.md`
- Backlog de deuda técnica → `15-project-control/tech-backlog.md`
