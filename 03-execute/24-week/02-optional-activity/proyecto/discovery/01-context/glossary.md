# Project Glossary

> **Instructions:** Define here all technical and business terms used in the project.
> This is the official dictionary — if there is ambiguity, this document wins.
> Add terms throughout the project, not only at the start.

---

## How to use this glossary

1. Before using a technical or business term in code, docs, or conversations: look it up here.
2. If it's not there: add it with its definition.
3. If there is disagreement about the definition: discuss it as a team and update this document.

---

## Domain terms

| Term | Definition | Notes / Synonyms |
|------|-----------|-----------------|
| **Orden de Servicio** | Documento digital principal que orquesta y registra todo el ciclo de vida del soporte de un vehículo, desde su check-in hasta su entrega final. | Service Order, Ticket. Evitar usar "Ficha". |
| **Check-in** | Proceso de recepción inicial del vehículo en el taller donde se capturan los datos de entrada, kilometraje y síntomas reportados por el cliente. | Recepción, Ingreso. |
| **Diagnóstico** | Evaluación técnica que detalla la causa raíz de la falla, los componentes mecánicos/electrónicos a reparar y las horas estimadas de mano de obra. | Diagnostic, Evaluación. Debe ser aprobado antes de iniciar la intervención. |
| **Intervención Técnica** | Acción física de reparación correctiva o mantenimiento preventivo ejecutada por un técnico asignado sobre el automotor. | Reparación, Tarea. |
| **Garantía** | Cobertura legal y comercial otorgada al cliente que asegura la calidad de los repuestos instalados o de la mano de obra por un periodo de tiempo o kilometraje específico. | Cobertura. Se valida de forma estricta antes de generar costos de reparación. |
| **Historial Clínico** | Línea de tiempo cronológica e inalterable que consolida todas las intervenciones, diagnósticos y servicios prestados a un cliente y sus respectivos automotores. | Timeline, Cronología de Servicios. |
| **Mantenimiento Preventivo** | Tareas programadas regularmente (como cambio de aceite o filtros) para evitar fallas mecánicas basadas en el tiempo o kilometraje del vehículo. | Mantenimiento Programado. |
| **Mantenimiento Correctivo** | Acciones destinadas a reparar averías o fallas ya existentes reportadas por el cliente en el check-in. | Reparación por Avería. |
| **Mantenimiento Predictivo** | Análisis de datos históricos y patrones del vehículo para estimar cuándo ocurrirá una falla mecánica antes de que suceda, permitiendo agendar un soporte proactivo. | Mantenimiento Inteligente. |
| **Técnico** | Personal calificado del taller encargado de ejecutar físicamente los diagnósticos y las intervenciones técnicas en las bahías de trabajo. | Mecánico, Operador. |

---

## Technical terms of the project

| Term | Definition |
|------|-----------|
| Microservice | Independent service with a single responsibility, its own process, and its own database |
| Domain Event | A fact that occurred in the business that other services can observe. Name always in past tense. |
| Bounded Context | Boundary within which a particular domain model has consistent meaning |
| API Gateway | Single entry point to the system that routes requests to the corresponding microservices |
| Circuit Breaker | Pattern that stops calls to a failing service, preventing failure cascades |
| Saga | Sequence of local transactions across different services with compensating transactions on failure |
| Dead Letter Queue | Queue where messages that could not be processed after several retries are sent |
| Idempotence | Property of an operation to produce the same result if executed multiple times |

---

## Acronyms

| Acronym | Meaning |
|---------|---------|
| IAM | Identity and Access Management |
| JWT | JSON Web Token |
| API | Application Programming Interface |
| CRUD | Create, Read, Update, Delete |
| DTO | Data Transfer Object |
| FR | Functional Requirement |
| NFR | Non-Functional Requirement |
| SLO | Service Level Objective |
| SLA | Service Level Agreement |
| ADR | Architecture Decision Record |
| PR | Pull Request |
| DoD | Definition of Done |
| CI/CD | Continuous Integration / Continuous Delivery |
| VIN | Vehicle Identification Number (Número de Chasis único de 17 caracteres) |
| MVP | Minimum Viable Product (Producto Mínimo Viable) |
