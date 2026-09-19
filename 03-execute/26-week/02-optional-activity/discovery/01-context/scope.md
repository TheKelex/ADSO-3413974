# System Scope

> **Why this document exists:** Scope prevents scope creep and aligns expectations.
> It is equally important to define what the system does NOT do as what it does.
> Review this document at the start of each planning cycle.

---

## In Scope

What the system **DOES build and maintain**:

### MVP Features

| # | Feature | Description | Responsible service |
|---|---------|-------------|---------------------|
| 1 | Gestión de Solicitudes y Check-in | Permite registrar datos del cliente, del vehículo (Placa/VIN) y la falla inicial reportada de forma digital. | customer-service |
| 2 | Diagnósticos Técnicos | Módulo para que el mecánico registre la evaluación técnica detallada de la falla y los componentes a reparar. | workshop-service |
| 3 | Asignación de Personal | Panel que permite al Jefe de Taller asignar mecánicos específicos a las órdenes de servicio según disponibilidad. | allocation-service |
| 4 | Seguimiento de Estados | Monitoreo en tiempo real del ciclo de vida del servicio (*Recibido*, *En Diagnóstico*, *En Reparación*, *Listo*). | tracking-service |
| 5 | Intervenciones y Mantenimiento | Registro minucioso de cada acción física ejecutada sobre el vehículo y repuestos internos utilizados. | workshop-service |
| 6 | Gestión de Garantías | Registro, control y validación de las garantías vigentes otorgadas al cliente tanto en repuestos instalados como en mano de obra. | warranty-service |
| 7 | Cronología e Historial Clínico | Línea de tiempo unificada con todos los servicios prestados a un cliente desde el primer día, con filtros dinámicos por cada automotor asociado. | customer-service |
| 8 | Planificación Inteligente de Mantenimientos | Motor analítico que procesa el historial de datos para programar y sugerir alertas automáticas de mantenimiento predictivo, preventivo y correctivo. | analytics-service |

### Included integrations

| External system | Integration type | Purpose |
|----------------|-----------------|---------|
| ERP Corporativo | REST API | Sincronización de repuestos e inventario disponible en taller. |
| Proveedor de SMS / WhatsApp | Webhook / SDK | Envío de notificaciones automáticas al cliente cuando cambie el estado de su vehículo o se sugiera un mantenimiento predictivo. |

### Environments being built

| Environment | Purpose |
|-------------|---------|
| Local | Desarrollo en las máquinas locales de los desarrolladores. |
| Development (dev) | Integración continua (CI) y pruebas automáticas de desarrollo. |
| Staging | Entorno pre-producción para pruebas de aceptación del Product Owner (PO). |
| Production | Entorno productivo real para uso de los talleres automotrices. |

---

## Out of Scope

What the system **does NOT build** in this version and why:

| # | What is out of scope | Reason | Future version? |
|---|---------------------|--------|----------------|
| 1 | Procesamiento de Pagos en Línea | El cobro final se realiza en las cajas físicas del taller; no es crítico para el MVP. | Sí — H1 2027 |
| 2 | Facturación Fiscal y Legal | La emisión de facturas legales con impuestos se delega completamente al ERP externo. | No — N/A |
| 3 | Gestión de Compra a Proveedores | El sistema solo consume stock existente, no maneja la cadena de suministro externa. | Sí — H2 2027 |

### What another system / team handles (and why not us)

| Feature | Who builds it | Why not us |
|---------|--------------|-----------|
| Autenticación Única (SSO) | Equipo Central de Seguridad | Reutilización de la infraestructura de identidad existente en la empresa. |
| Reportes Financieros Avanzados | Sistema de BI / Equipo de Analytics | Fuera del dominio principal del negocio técnico (Core Domain). |

---

## Scope assumptions

> These assumptions are taken to be true. If they change, the scope must be renegotiated.

| # | Assumption | Consequence if false |
|---|-----------|---------------------|
| 1 | El ERP Corporativo cuenta con una API REST estable de inventario. | Tendríamos que construir un módulo de inventario interno completo, retrasando el MVP. |
| 2 | Los técnicos operarán el sistema mediante tablets/dispositivos móviles en el taller. | El diseño UX/UI tendría que cambiar drásticamente a interfaces de escritorio de pantalla ancha. |
| 3 | El volumen inicial de datos es menor a 50,000 registros de vehículos al año. | La estrategia de indexación y base de datos podría requerir un esquema de particionamiento prematuro para procesar las líneas de tiempo. |

---

## Constraints

| Type | Description |
|------|-------------|
| **Time** | El MVP del sistema de soporte técnico debe estar listo en un plazo máximo de 12 semanas. |
| **Budget** | Limitado a un total de 480 horas de desarrollo estimadas para el equipo actual. |
| **Technology** | Obligatorio el uso del stack corporativo: Node.js + TypeScript para backend y PostgreSQL para base de datos. |
| **Regulatory** | Debe cumplir con la ley local de protección de datos personales de los clientes (ej: RGPD o equivalente local). |
| **Team** | El equipo de ingeniería disponible está compuesto por 3 desarrolladores full-stack y 1 QA. |

---

## External dependencies

| Dependency | Team / Provider | Required date | Status |
|-----------|----------------|--------------|--------|
| API de Inventario ERP | Equipo de Sistemas Internos | Semana 4 | 🟢 Available |
| Credenciales del Proveedor de SMS | Proveedor Externo de Telecomunicaciones | Semana 6 | 🟡 In progress |
| Configuración de Infraestructura AWS | Equipo de DevOps Central | Semana 2 | 🔴 Pending |

---

## How to update the scope

The scope can change, but the change has a process:

1. Document the proposed change in this file
2. Evaluate the impact on schedule and effort
3. Obtain approval from the Product Owner and Tech Lead
4. Update the roadmap in `03-product/vision.md`
5. Create or update HUs in `04-requirements/user-stories.md`

---

## Correlations

- Vision and roadmap → `03-product/vision.md`
- Term glossary → `01-context/glossary.md`
- System overview → `01-context/overview.md`
- Scope-related risks → `15-project-control/risks.md`
