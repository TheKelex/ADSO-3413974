# Definición de Listo (Definition of Ready - DoR)

> Una Historia de Usuario está **LISTA (READY)** cuando todo el equipo comprende su alcance y puede iniciar su desarrollo en el próximo sprint sin necesidad de resolver preguntas fundamentales a mitad de camino. Si una historia no cumple con este DoR, regresa a refinamiento.

---

## Lista de verificación del DoR (DoR checklist)

Antes de mover una Historia de Usuario a la columna "Ready" (Listo para el Sprint), verifica que cumpla lo siguiente:

### Claridad y Propósito
- [ ] La historia está escrita bajo el formato estándar: **Como [rol específico], quiero [acción], para [beneficio/valor]**.
- [ ] El rol es sumamente específico dentro del contexto automotriz (ej: *"Como Técnico Mecánico"*, *"Como Jefe de Taller"*, no un genérico *"Como usuario"*).
- [ ] El beneficio esperado para el negocio del taller es claro, cuantificable y completamente verificable.

### Criterios de Aceptación (Acceptance Criteria)
- [ ] Cuenta con al menos 2 criterios de aceptación redactados en formato **Dado (Given) / Cuando (When) / Entonces (Then)**.
- [ ] Los criterios cubren tanto el camino feliz (happy path) como los principales casos de error o flujos alternos.
- [ ] Los criterios son testeables (es técnicamente viable que los 2 QAs diseñen una prueba manual o automatizada para cada uno).
- [ ] Se eliminó cualquier criterio ambiguo o subjetivo (ej: *"la pantalla de analítica predictiva debe cargar rápido"* no es válido; debe decir *"debe responder en menos de 2 segundos"*).

### Dependencias e Integraciones
- [ ] Se identificaron explícitamente todas las dependencias externas (ej: si requiere consultar repuestos en el ERP Corporativo).
- [ ] Las dependencias bloqueantes están resueltas o se definió formalmente una estrategia de simulación (mocking) para no detener el desarrollo.
- [ ] Si la funcionalidad depende de otra historia de usuario, esa historia previa ya se encuentra en estado "Done" o "In Progress".

### Estimación del Equipo
- [ ] La historia fue estimada colectivamente por los miembros del equipo mediante Planning Poker utilizando la escala de Fibonacci.
- [ ] Existe consenso absoluto en el equipo de que la tarea puede completarse holgadamente dentro de las 2 semanas del sprint.
- [ ] Si la historia es demasiado grande, fue dividida en historias de usuario más pequeñas e independientes.

### Preparación Técnica y Arquitectura
- [ ] Los accesos, permisos y entornos necesarios están disponibles para los desarrolladores y QAs.
- [ ] Si la historia implica nuevos endpoints (ej: consultar el historial clínico del vehículo), los contratos de la API (OpenAPI) ya están pre-definidos y acordados.
- [ ] El modelo de datos, tablas o cambios en la base de datos PostgreSQL fueron previamente revisados y validados por el Arquitecto del equipo.
- [ ] Se analizó e identificó el impacto colateral que el cambio puede causar en otros microservicios del ecosistema.

### Requisitos No Funcionales (NFRs)
- [ ] Se especifican las restricciones de rendimiento o capacidad si la historia lo amerita (ej: el procesamiento analítico de mantenimientos predictivos).
- [ ] Se contemplan explícitamente las reglas de seguridad a nivel de código (encriptación de datos de clientes, tokens JWT en el API Gateway y validación de entradas).
- [ ] Se definen los requerimientos de observabilidad mínimos (qué eventos de negocio o errores críticos deben generar logs estructurados).

---

## Razones comunes por las que una historia NO está lista

| Problema | Qué hacer para solucionarlo |
|---------|----------------------------|
| Requerimientos confusos o ambiguos | El Delivery Manager agenda una sesión rápida de refinamiento de 30 minutos con el PO. |
| Faltan criterios de aceptación o formato Gherkin | El Product Owner complementa los criterios antes de que inicie la sesión de Planning. |
| Dependencias técnicas desconocidas | El Tech Lead y el Arquitecto revisan, aclaran y documentan las dependencias en la tarjeta. |
| La historia es demasiado grande (> 5 puntos) | El equipo técnico la desglosa en subtareas independientes o en historias más pequeñas. |
| Falta de contratos API establecidos | El equipo acuerda y escribe el archivo OpenAPI preliminar en `07-api/` antes de programar. |

---

## DoR frente a DoD (DoR vs DoD)

| Característica | Definición de Listo (DoR) | Definición de Hecho (DoD) |
|---|---|---|
| **¿Cuándo se aplica?** | Antes de iniciar el desarrollo de la historia (Fase de Planeación). | Al finalizar todas las actividades de la historia (Fase de Entrega). |
| **¿Quién lo verifica?** | Todo el equipo durante la planeación y el refinamiento. | El equipo técnico, los QAs y el Tech Lead en la revisión del PR y pruebas. |
| **Propósito principal** | Asegurar que el equipo empiece a programar con claridad y sin bloqueos. | Asegurar que la pieza de software construida es de alta calidad y desplegable. |

---

## Correlaciones

- Definición de Hecho completa → `00-governance/definition-of-done.md`
- Plantilla oficial de Historias de Usuario → `04-requirements/_template-hu.md`
- Tablero de Historias de Usuario (Backlog) → `04-requirements/user-stories.md`
