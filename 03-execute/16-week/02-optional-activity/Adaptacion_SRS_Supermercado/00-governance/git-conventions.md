# Convenciones de Git

> Estado: 🟡 En progreso | Última actualización: 2026-07-03
> Autor: Kevin Santiago Saavedra Chantris | Equipo: Kevin Santiago Saavedra Chantris

Este repositorio usa ramas protegidas, Pull Requests y Conventional Commits para mantener trazabilidad documental.

## Ramas protegidas

`dev`, `qa` y `main` representan ambientes padre y no se trabajan directamente.

| Rama | Propósito | Regla |
|------|-----------|-------|
| `dev` | Integración de trabajo en desarrollo | Recibe PRs desde ramas hijas |
| `qa` | Validación funcional y técnica | Recibe PRs o cherry-picks aprobados desde `dev` |
| `main` | Producción / documentación estable | Recibe solo PRs desde `release/*` |

## Ramas documentales

| Tipo de rama | Cuándo usarla | Ejemplo | Tipo de commit |
|--------------|---------------|---------|----------------|
| `feat` | Documento nuevo | `feat/doc-api-guidelines` | `docs` |
| `fix` | Corrección de contenido | `fix/doc-scope` | `fix` |
| `chore` | Reorganización o renombrado | `chore/doc-move-adr-003` | `chore` |
| `docs` | Actualización de documento existente | `docs/doc-service-catalog` | `docs` |

El tipo de rama describe intención. El tipo del commit sigue Conventional Commits.

## Ramas por historia de usuario

| Caso | Rama base | Formato | Ejemplo |
|------|-----------|---------|---------|
| Desarrollo de HU | `dev` | `hu-<numero>-dev` | `hu-01-dev` |
| Ajuste o validación QA | `qa` | `hu-<numero>-qa` | `hu-01-qa` |
| Release de iteración | `main` | `release/<iteracion>` | `release/iteration-01` |

Las ramas `hu-*` son un caso especial para trazabilidad por historia. No siguen el formato `<tipo>/doc-*`.

## Flujo hacia dev

```bash
git checkout dev
git pull origin dev
git checkout -b hu-01-dev

git add <archivos>
git commit -m "docs(04-requirements): add scheduling availability user story"
git push origin hu-01-dev
```

Abrir PR de `hu-01-dev` hacia `dev`.

## Flujo hacia qa

Crear rama hija desde `qa`:

```bash
git checkout qa
git pull origin qa
git checkout -b hu-01-qa
```

Llevar cambios con merge cuando la HU completa pasa igual:

```bash
git merge origin/hu-01-dev
git push origin hu-01-qa
```

O con cherry-pick cuando solo pasan commits específicos:

```bash
git cherry-pick <commit-sha>
git push origin hu-01-qa
```

Abrir PR de `hu-01-qa` hacia `qa`.

## Release hacia main

`main` representa documentación estable. Para producción, crear una rama release desde `main`:

```bash
git checkout main
git pull origin main
git checkout -b release/iteration-01
```

La rama release puede acumular varias HUs de una iteración:

```bash
git cherry-pick <commit-hu-01>
git cherry-pick <commit-hu-02>
git cherry-pick <commit-hu-03>
git push origin release/iteration-01
```

Abrir PR de `release/iteration-01` hacia `main`.

## Conventional Commits

Formato obligatorio:

```text
<type>(NN-section): short description in English
```

Tipos permitidos:

| Tipo | Uso |
|------|-----|
| `docs` | Crear o actualizar documentación |
| `fix` | Corregir contenido incorrecto |
| `chore` | Mover, renombrar, reordenar o actualizar metadatos |
| `refactor` | Reestructurar documentación sin cambiar significado |

No usar `feat`, `style`, `test`, `perf`, `build` ni `ci` para commits de este repositorio documental.

Ejemplos:

```bash
docs(04-requirements): add scheduling user stories
docs(09-microservices): register auth service
fix(01-context): clarify project scope
chore(08-uml): export sequence diagrams to SVG
refactor(00-governance): split contribution rules by topic
```

## Reglas de commits

- La descripción del commit va en inglés.
- El contenido de los documentos puede estar en español.
- Los commits deben ser pequeños y trazables.
- Si se documentan varios microservicios, usar un commit por microservicio cuando sea posible.
- No mezclar cambios funcionales de varias secciones sin razón clara.

## Hotfix en main

Cuando se detecta un error crítico en `main` que no puede esperar el flujo normal de release:

| Caso | Rama base | Formato | Ejemplo |
|------|-----------|---------|---------|
| Corrección urgente en documentación estable | `main` | `fix/doc-<descripcion>` | `fix/doc-broken-api-contract` |

Flujo:

```bash
git checkout main
git pull origin main
git checkout -b fix/doc-broken-api-contract

git add <archivos>
git commit -m "fix(07-api): correct broken endpoint reference in contract"
git push origin fix/doc-broken-api-contract
```

Abrir PR directo de `fix/doc-*` hacia `main`. Una vez mergeado, aplicar el mismo fix a `qa` y `dev` con cherry-pick:

```bash
git checkout qa
git pull origin qa
git cherry-pick <commit-sha>
git push origin qa
```
