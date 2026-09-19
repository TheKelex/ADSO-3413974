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
