# 00-governance — Documentación Completa del Equipo

Este documento consolida los archivos requeridos para el bloque de gobernanza, personalizados para un equipo de **10 integrantes aprendices** (1 Tech Lead, 1 Arquitecto, 1 Delivery Manager, 2 QAs y 5 Desarrolladores) donde **todos programan** y participan de las fases del **Sistema de Gestión de Soporte Técnico Automotriz** durante un ciclo de **8 meses** (1 de septiembre de 2026 al 30 de abril de 2027).

---

## 1. agile-conventions.md

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

## Ceremonies

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
- **Duración:** Máximo 45 minutos
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

---

## 2. definition-of-done.md

# Definición de Hecho (Definition of Done - DoD)

> Una Historia de Usuario está **HECHA (DONE)** cuando cumple con TODOS los criterios de esta lista de verificación. Si falta un solo criterio, la historia NO está hecha y regresa a la columna "En Progreso".

## Lista de verificación obligatoria (Mandatory checklist)

### Código y Arquitectura
- [ ] El código implementa todos los criterios de aceptación de la historia de usuario.
- [ ] El diseño del código respeta los límites del dominio definidos por el Arquitecto (Arquitectura Hexagonal).
- [ ] El código fue revisado y aprobado formalmente por al menos 1 desarrollador par y cuenta con el visto bueno técnico del Tech Lead (aprobación del PR).
- [ ] El código sigue los estándares de estilo del proyecto (las herramientas de linter y formateo pasan sin errores en la integración continua - CI).
- [ ] No se introduce deuda técnica. En caso de ser inevitable, el Tech Lead debe aprobarlo y registrarlo inmediatamente en `15-project-control/tech-backlog.md`.

### Pruebas de Calidad (QA) y TDD
- [ ] Se escribieron pruebas unitarias utilizando la metodología TDD para toda la nueva lógica de negocio (ej: validadores de garantías o reglas de analítica predictiva).
- [ ] La cobertura (coverage) de pruebas no disminuye respecto a la línea base del proyecto (Mínimo 80%).
- [ ] Los 2 QAs del equipo ejecutaron y validaron manualmente los criterios de aceptación en el entorno correspondiente.
- [ ] Todas las pruebas automatizadas pasan con éxito de forma local y en el servidor de CI.

### Integración y Contratos de Datos
- [ ] Los cambios no rompen otros microservicios (todas las pruebas de integración pasan).
- [ ] Si la funcionalidad interactúa con el inventario del taller, se verificó la integración o el mock con el ERP Corporativo.
- [ ] Si la API del microservicio cambió: el contrato OpenAPI fue actualizado rigurosamente en la carpeta `07-api/`.
- [ ] Si el modelo de datos cambió (ej: tablas de vehículos o clientes): los archivos de la carpeta `06-data/` y sus respectivas migraciones fueron actualizados.
- [ ] Si se crearon o modificaron eventos de dominio (ej: `GarantiaReclamada`): se actualizó el catálogo de eventos en `02-domain/`.

### Despliegue (Deployment)
- [ ] El código se puede fusionar (merge) a la rama `develop` sin conflictos.
- [ ] La tubería de CI/CD está completamente en verde para la rama de la característica.
- [ ] La funcionalidad está desplegada con éxito en el entorno de **Staging**.
- [ ] Se realizó una prueba de humo básica (smoke test) en Staging y el sistema responde correctamente.

### Documentación
- [ ] El archivo `README.md` del microservicio afectado fue actualizado si su interfaz pública o flujos sufrieron cambios.
- [ ] Si se tomó una decisión técnica o de diseño de alto impacto: se redactó y aprobó el correspondiente registro de decisión arquitectónica (ADR) en `05-architecture/`.

---

## Excepciones permitidas

Las siguientes excepciones deben ser aprobadas explícitamente y por escrito por el Tech Lead o el Arquitecto del proyecto:
- Omisión de pruebas de integración complejas por limitaciones extremas del entorno externo (se documenta el riesgo en la carpeta 15).
- Postergación de documentación técnica no crítica por entregas de urgencia máxima (se genera obligatoriamente un ticket en el backlog técnico).

---

## Lo que NO es un criterio de "Hecho"

- *"El código está en mi máquina"* — debe estar integrado en el repositorio central de GitHub.
- *"Funciona en mi entorno local"* — debe funcionar correctamente en el entorno de Staging.
- *"El Product Owner ya lo vio y le gustó"* — esa es la aceptación del producto, no la definición de calidad de ingeniería del código.

---

## 3. definition-of-ready.md

# Definición de Listo (Definition of Ready - DoR)

> Una Historia de Usuario está **LISTA (READY)** cuando todo el equipo comprende su alcance y puede iniciar su desarrollo en el próximo sprint sin necesidad de resolver preguntas fundamentales a mitad de camino. Si una historia no cumple con este DoR, regresa a refinamiento.

---

## Lista de verificación del DoR (DoR checklist)

