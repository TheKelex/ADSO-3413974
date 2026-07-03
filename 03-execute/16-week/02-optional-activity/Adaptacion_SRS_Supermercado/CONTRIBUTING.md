# Guía de contribución

Este repositorio es la fuente de verdad documental del proyecto **El Apunte**. Todo equipo que trabaje en el proyecto documenta aquí, siguiendo las reglas de gobierno documental.

## Regla de oro

No trabajar directamente sobre `dev`, `qa` ni `main`. Todo cambio debe hacerse en una rama hija, pasar por Pull Request, revisión y aprobación.

## Antes de empezar

- Tener acceso al repositorio con permisos de escritura.
- Configurar Git localmente con `user.name` y `user.email`.
- Leer las reglas aplicables antes de crear, mover o editar documentos.

```bash
git clone <url-del-repo>
cd design-software-docs
```

## Estructura general

| Sección | Propósito |
|---------|-----------|
| `00-governance/` | Reglas del repositorio documental |
| `01-context/` | Contexto, alcance y glosario |
| `02-domain/` | Dominio, entidades y reglas de negocio |
| `03-product/` | Visión, roadmap y backlog de producto |
| `04-requirements/` | Requisitos, HUs y trazabilidad |
| `05-architecture/` | Arquitectura y ADRs |
| `06-data/` | Modelos y diccionario global de datos |
| `07-api/` | Guidelines, autenticación y contratos API |
| `08-uml/` | Diagramas con fuente editable y exportaciones |
| `09-microservices/` | Catálogo y documentación de servicios |
| `10-devops/` | Setup, CI/CD y ambientes |
| `11-quality/` | Testing y code review |
| `12-ux-ui/` | Design system, navegación y wireframes |
| `13-operations/` | Observabilidad, incidentes y recuperación |
| `14-training/` | Manuales y onboarding |
| `15-project-control/` | Riesgos, dependencias y preguntas abiertas |
| `99-archive/` | Documentos deprecados o históricos |

El `README.md` de cada carpeta es el índice de esa sección.

## Reglas por tema

| Tema | Documento |
|------|-----------|
| Naming, estructura mínima, estados, índices y diagramas | [00-governance/documentation-rules.md](./00-governance/documentation-rules.md) |
| Ramas, ambientes, releases y commits | [00-governance/git-conventions.md](./00-governance/git-conventions.md) |
| Registro de microservicios | [00-governance/microservices-documentation.md](./00-governance/microservices-documentation.md) |
| Seguridad documental | [00-governance/security-rules.md](./00-governance/security-rules.md) |
| Definition of Ready | [00-governance/definition-of-ready.md](./00-governance/definition-of-ready.md) |
| Definition of Done | [00-governance/definition-of-done.md](./00-governance/definition-of-done.md) |
| ADRs | [05-architecture/decisions/README.md](./05-architecture/decisions/README.md) |

## Flujo rápido

1. Identifica la sección correcta según la tabla de estructura.
2. Revisa [documentation-rules.md](./00-governance/documentation-rules.md).
3. Crea una rama según [git-conventions.md](./00-governance/git-conventions.md).
4. Edita o crea el documento con estado, fecha, autor/equipo, contexto, contenido y referencias.
5. Registra el archivo en el `README.md` de su sección.
6. Valida seguridad con [security-rules.md](./00-governance/security-rules.md).
7. Haz commits con Conventional Commits en inglés.
8. Abre Pull Request y completa el checklist.

## Agregar un documento

1. Confirma que el documento no exista ya en otra sección.
2. Crea el archivo en `kebab-case.md`.
3. Usa la estructura mínima definida en [documentation-rules.md](./00-governance/documentation-rules.md#estructura-mínima).
4. Agrega el enlace al `README.md` de la sección.
5. Haz commit y PR.

Ejemplo:

```bash
git checkout -b feat/doc-auth-flows
git add 07-api/auth-flows.md 07-api/README.md
git commit -m "docs(07-api): add OAuth2 authentication flows"
git push origin feat/doc-auth-flows
```

## Agregar una carpeta

Crear carpetas de sección nueva en la raíz requiere aprobación de arquitectura. Para subcarpetas:

- Crear carpeta solo si agrupa varios documentos o tiene ciclo de vida propio.
- Incluir siempre `README.md`.
- Registrar la subcarpeta en el `README.md` padre.
- Evitar nombres genéricos como `misc/`, `otros/` o `temp/`.

## Microservicios

No crear carpetas de microservicios hasta que el servicio exista en el repositorio de código o su creación esté aprobada por arquitectura.

El flujo completo está en [microservices-documentation.md](./00-governance/microservices-documentation.md).

## ADRs

Las decisiones arquitectónicas se registran en `05-architecture/decisions/records/` y nunca se eliminan. Usar:

- [Proceso de ADR](./05-architecture/decisions/README.md)
- [Plantilla de ADR](./05-architecture/decisions/_template-adr.md)

## Antes de abrir PR

- [ ] El documento cumple [Definition of Ready](./00-governance/definition-of-ready.md).
- [ ] El archivo está enlazado desde el `README.md` de su sección.
- [ ] Los commits siguen [git-conventions.md](./00-governance/git-conventions.md).
- [ ] No se publican secretos, datos personales ni información sensible.
- [ ] Si hay diagramas, tienen fuente editable y exportación.
- [ ] Si es ADR, está registrado en el índice de decisiones.
- [ ] Si es microservicio, está registrado en `09-microservices/service-catalog.md`.
- [ ] `CHANGELOG.md` fue actualizado si el cambio afecta a varios equipos.

## Dudas

Abrir un PR sobre este mismo archivo o registrar una pregunta en [15-project-control/open-questions.md](./15-project-control/open-questions.md).
