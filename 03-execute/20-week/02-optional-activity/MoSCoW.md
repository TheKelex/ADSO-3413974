## MoSCoW

Para la realizacion de la matriz MoSCoW no se plantearan soluciones y RF aparte de los que se mencionan en el archivo [case-1.md](./case-1.md), siguiendo este entendimiento la matriz se enfoca en una primera version que es la MVP, por lo que desarrolla las pantallas señaladas principales para un MVP en el archivo [Desing Thinking.md](./Desing%20Thinking.md)

---------------------------------------------------------------------

### Must Have

| Prioridad     | Funcionalidad                   | Justificación                             |
| ------------- | ------------------------------- | ----------------------------------------- |
| **Must Have** | Login con autenticación         | El sistema no funciona sin acceso seguro. |
| **Must Have** | CRUD de pasajeros               | Punto de partida del proceso.             |
| **Must Have** | Crear reservas                  | Base para emitir tiquetes.                |
| **Must Have** | Emitir tiquetes                 | Es el objetivo principal del negocio.     |
| **Must Have** | Crear vuelos                    | Sin vuelos no existe operación.           |
| **Must Have** | Asignar aeronave                | Necesario para la operación del vuelo.    |
| **Must Have** | Asignar asiento                 | Regla crítica del negocio.                |
| **Must Have** | Validar asiento único por vuelo | Regla de negocio obligatoria.             |
| **Must Have** | Registrar embarque              | Permite conocer quién viajó.              |
| **Must Have** | Reporte No-Show                 | Requisito funcional principal del PRD.    |
| **Must Have** | Docker Compose funcionando      | Criterio de aceptación.                   |
| **Must Have** | Pruebas automatizadas           | Criterio de aceptación.                   |

---------------------------------------------------------------------

### Should Have

| Prioridad       | Funcionalidad        | Justificación                                                                         |
| --------------- | -------------------- | ------------------------------------------------------------------------------------- |
| **Should Have** | Registrar pagos      | El PRD lo marca como opcional para el tiquete, pero sí debe existir la funcionalidad. |
| **Should Have** | Registrar equipaje   | No todos los pasajeros registran equipaje, pero la funcionalidad es necesaria.        |
| **Should Have** | Listados de reservas | Apoya la operación diaria.                                                            |
| **Should Have** | Listado de pasajeros | Facilita la gestión.                                                                  |
---------------------------------------------------------------------

### Could Have

| Prioridad      | Funcionalidad                 | Justificación                      |
| -------------- | ----------------------------- | ---------------------------------- |
| **Could Have** | Filtros avanzados en reportes | No son requeridos para el MVP.     |
| **Could Have** | Historial de cambios          | Aporta auditoría futura.           |
| **Could Have** | Búsqueda rápida por documento | Mejora la experiencia del usuario. |
| **Could Have** | Dashboard operativo           | Valor agregado, no indispensable.  |

---------------------------------------------------------------------


### Won´t Have (MVP)

| Prioridad            | Funcionalidad                     | Justificación                  |
| -------------------- | --------------------------------- | ------------------------------ |
| **Won't Have (MVP)** | Compra de tiquetes por pasajeros  | El único usuario es el agente. |
| **Won't Have (MVP)** | Check-in en línea                 | No aparece en el PRD.          |
| **Won't Have (MVP)** | Selección de asiento por pasajero | La realiza el agente.          |
| **Won't Have (MVP)** | Integración con pasarelas de pago | Solo se registra el pago.      |
| **Won't Have (MVP)** | Notificaciones por correo o SMS   | No forman parte del alcance.   |
| **Won't Have (MVP)** | Múltiples roles de usuario        | Solo existe el agente.         |
