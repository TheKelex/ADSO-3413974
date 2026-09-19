# Reglas de Documentación (Documentation Rules)

> Estas reglas determinan cómo se escribe, organiza y mantiene la documentación en este proyecto. La documentación que no siga estas reglas puede ser rechazada en la revisión de código (Pull Request).

---

## Principio Central

> **"La documentación es código. Si no está actualizada, está rota."**

Cada Historia de Usuario que modifique el comportamiento del sistema DEBE incluir la actualización de los documentos afectados. La Definición de Hecho (DoD) lo exige de forma estricta.

---

## Idioma (Language)

| Artefacto | Idioma |
|----------|----------|
| Código fuente (variables, funciones, clases) | Inglés (Estándar técnico) |
| Comentarios en el código | Inglés |
| Commits | Inglés (Conventional Commits) |
| Nombres de ramas (Branches) | Inglés |
| Documentación en Markdown (`.md`) | Español (Para facilitar la comprensión del equipo y negocio) |
| Contratos OpenAPI (descripciones) | Inglés (Para compatibilidad con generadores de código) |
| Mensajes de error retornados al frontend | Español |
| Logs internos del sistema | Inglés |

> **Regla:** Una vez elegido el idioma para cada categoría, este es vinculante para todo el proyecto. Mezclar idiomas dentro de una misma categoría es motivo automático de rechazo del Pull Request.

---

## Estructura de Archivos


---

## Qué documentar y qué NO documentar

### SÍ documentar

| Qué | Dónde |
|------|-------|
| Decisiones arquitectónicas no obvias | `05-architecture/decisions/records/ADR-NNN.md` |
| Reglas de negocio e invariantes del dominio (ej: Garantías, Mantenimientos) | `02-domain/entities-and-rules.md` |
| Contratos de API para cada microservicio | `07-api/contracts/openapi/[service].yaml` |
| Cambios en el modelo de datos (PostgreSQL) | `06-data/models.md` |
| Procedimientos operativos y de despliegue | `13-operations/` |
| Riesgos identificados del proyecto | `15-project-control/risks.md` |

### NO documentar

- Lo que el código ya dice claramente (no repetir en comentarios lo que se puede leer de forma evidente en el código).
- Decisiones temporales o experimentos de código que serán revertidos inmediatamente.
- Detalles de implementación de librerías externas (esas ya tienen su propia documentación oficial).
- Historial de cambios manual dentro del texto (para eso se utiliza estrictamente el historial de Git).

---

## Responsables por sección (Owners)

Dado que todo el equipo programa pero cumple roles de buenas prácticas, las responsabilidades se distribuyen así:

| Sección | Responsable (Owner) | Frecuencia de Revisión |
|---------|-------|-----------------|
| `00-governance/` | Tech Lead | Inicio de cada sprint |
| `02-domain/` | Arquitecto + Product Owner | Cuando cambien las reglas del taller |
| `04-requirements/` | Product Owner / Delivery Manager | Cada sprint |
| `05-architecture/` | Arquitecto del proyecto | Con cada decisión de diseño técnico |
| `07-api/contracts/` | Desarrollador asignado a la tarea | Con cada cambio en los endpoints |
| `09-microservices/` | Desarrollador asignado a la tarea | En cada liberación (release) |
| `11-quality/` | Los 2 QAs del equipo | Al cierre de cada sprint |
| `13-operations/` | Equipo DevOps / Desarrollador de guardia | Después de cada incidente |
| `15-project-control/` | Delivery Manager + Tech Lead | Revisión semanal |

---

## Formato del Documento

### Encabezados (Headings)
- `# H1` — Solo se permite uno por archivo; funciona como el título principal.
- `## H2` — Secciones principales del documento.
- `### H3` — Subsecciones técnicas.
- No utilizar H4 o niveles más profundos; si se requiere, el documento tiene demasiada jerarquía y debe dividirse.

### Tablas
Se deben utilizar tablas exclusivamente para comparaciones, registros, matrices e índices. No utilizar tablas para listas simples de texto.

### Código
Utilizar siempre bloques de código especificando el lenguaje de programación correspondiente:

const warrantyDurationDays = 365;
const warrantyDurationDays = 365;

### Instrucciones de Plantillas
Los bloques marcados como `> [!NOTE] INSTRUCTIONS` indican que el documento es una plantilla sin llenar. Deben ser eliminados por el desarrollador una vez el documento esté completo.

---

## Proceso de Actualización

1. El desarrollador identifica qué documentos se ven afectados por su cambio de código antes de iniciar la tarea.
2. Actualiza los documentos correspondientes en el mismo commit/rama junto con el código (en el mismo PR).
3. El revisor (par técnico, Tech Lead o Arquitecto) verifica en GitHub que la documentación esté al día antes de dar el "Approve".
4. Si el PR cierra una Historia de Usuario que impacta la comunicación entre microservicios, el contrato OpenAPI en la carpeta `07-api/` debe ser actualizado de forma obligatoria.

---

## Correlaciones

- Convenciones de Git → `00-governance/git-conventions.md`
- Estándar de documentación por microservicio → `00-governance/microservices-documentation.md`
- Definición de Hecho (DoD) → `00-governance/definition-of-done.md`
```

---

Con esto, las reglas de documentación quedan claras, bilingües donde corresponde y adaptadas a la estructura de tu equipo.

¿Cuál es el siguiente archivo de **`00-governance`** que revisamos? Quedan pendientes documentos clave como `git-conventions.md` o las políticas de seguridad.
