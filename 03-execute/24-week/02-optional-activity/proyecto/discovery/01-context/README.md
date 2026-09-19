# 01 — Project Context

> **What is this?** The "why" of the system. Anyone new must be able to read this
> folder and understand what problem the project solves, what it includes, and what it does NOT include.

## Why this section exists

Before designing anything, the team needs to agree on:
- What problem are we solving? (Gestión ineficiente y pérdida de trazabilidad en soporte automotriz)
- For whom? (Talleres, Jefes de Taller, Técnicos Mecánicos y Clientes)
- What is in scope and what is out of scope? (MVP enfocado en ciclo técnico, analítica preventiva y garantías; excluye facturación fiscal directa)
- What does each term we use mean? (Definiciones compartidas de Orden de Servicio, Diagnóstico, Mantenimiento Predictivo, etc.)

Without this, each team member works with different assumptions and the project fragments.

---

## What is here and how to fill it in

### `overview.md` ⭐
Executive description of the system in maximum 1 page.
**Filled in:** Management System for Automotive Technical Support, user pain points (paperwork, lack of history), key roles (Client, Shop Boss, Technician), Node.js/PostgreSQL stack, and current status (Under Construction).

### `scope.md` ⭐
System boundaries: what it does and what it does NOT do.
**Filled in:** Comprehensive features including service tracking, warranty management, chronological customer timeline by vehicle, and intelligent predictive maintenance analysis. Explicitly marks legal billing and online payments as out of scope.

### `glossary.md` ⭐
Dictionary of the project domain.
**Filled in:** Core automotive and analytical definitions (Check-in, Predictive Maintenance, Technical Intervention, VIN) alongside core technical microservices architecture architecture terms and standard acronyms.

### `_template-project-profile.md`
Project technical sheet for internal records.
**To fill in:** When the project is formally named and stakeholders are assigned.

### `_template-scope-declaration.md`
Formal scope declaration template for presentations or deliverables.

---

## Correlations with other sections

| If you change this... | Also review... |
|-----------------------|----------------|
| The problem described in `overview.md` | Product vision in `03-product/vision.md` |
| The scope in `scope.md` | Requirements in `04-requirements/`, PRD in `03-product/` |
| A term in `glossary.md` | Every document where that term appears |

---

## Recommended fill order

1. `overview.md` — Completed (Baseline set for the automotive support system)
2. `scope.md` — Completed (Includes advanced core logic for history tracking and analytics)
3. `glossary.md` — Completed (Initialized with 10 business terms + architectural standards)

---

## Questions this section must answer

- What does this system exist for? (To optimize, track, and provide predictive automation to automotive technical support services)
- Who are the users? (Clients, Workshop Managers, and Field Technicians)
- What does the system NOT do? (Handle fiscal invoice generation or live payment gateways)
- What does [term X] mean in this project? (Refer strictly to the unified definitions in `glossary.md`)
