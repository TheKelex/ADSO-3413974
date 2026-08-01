## MoSCoW

Para la realizacion de la matriz MoSCoW no se plantearan soluciones y RF aparte de los que se mencionan en el archivo [case-1md](/case-1md), siguiendo este entendimiento la matriz se enfoca en una primera version que es la MVP, por lo que desarrolla las pantallas señaladas principales para un MVP en el archivo [Desing Thinkingmd](/Desing%20Thinkingmd)

---------------------------------------------------------------------

### Must Have (Debe Tener)

| Prioridad     | Funcionalidad                   | Justificacion                             |
| ------------- | ------------------------------- | ----------------------------------------- |
| **Must Have** | Login con autenticacion         | El sistema no funciona sin acceso seguro |
| **Must Have** | CRUD de pasajeros               | Punto de partida del proceso             |
| **Must Have** | Crear reservas                  | Base para emitir tiquetes                |
| **Must Have** | Emitir tiquetes                 | Es el objetivo principal del negocio     |
| **Must Have** | Crear vuelos                    | Sin vuelos no existe operacion           |
| **Must Have** | Asignar aeronave                | Necesario para la operacion del vuelo    |
| **Must Have** | Asignar asiento                 | Regla critica del negocio                |
| **Must Have** | Validar asiento unico por vuelo | Regla de negocio obligatoria             |
| **Must Have** | Registrar embarque              | Permite conocer quien viajo              |
| **Must Have** | Reporte No-Show                 | Requisito funcional principal del PRD    |
| **Must Have** | Docker Compose funcionando      | Criterio de aceptacion                   |
| **Must Have** | Pruebas automatizadas           | Criterio de aceptacion                   |

---------------------------------------------------------------------

### Should Have (Deberia Tener)

| Prioridad       | Funcionalidad        |Justificación                                                                         |
| --------------- | -------------------- | ------------------------------------------------------------------------------------- |
| **Should Have** | Registrar pagos      | El PRD lo marca como opcional para el tiquete, pero si debe existir la funcionalidad |
| **Should Have** | Registrar equipaje   | No todos los pasajeros registran equipaje, pero la funcionalidad es necesaria        |
| **Should Have** | Listados de reservas | Apoya la operacion diaria                                                            |
| **Should Have** | Listado de pasajeros | Facilita la gestion                                                                  |
---------------------------------------------------------------------

### Could Have (Podria Tener)

| Prioridad      | Funcionalidad                 | Justificacion                      |
| -------------- | ----------------------------- | ---------------------------------- |
| **Could Have** | Filtros avanzados en reportes | No son requeridos para el MVP     |
| **Could Have** | Historial de cambios          | Aporta auditoria futura           |
| **Could Have** | Busqueda rapida por documento | Mejora la experiencia del usuario |
| **Could Have** | Dashboard operativo           | Valor agregado, no indispensable  |

---------------------------------------------------------------------


### Won´t Have (MVP)

| Prioridad            | Funcionalidad                     | Justificacion                  |
| -------------------- | --------------------------------- | ------------------------------ |
| **Won't Have (MVP)** | Compra de tiquetes por pasajeros  | El unico usuario es el agente |
| **Won't Have (MVP)** | Check-in en linea                 | No aparece en el PRD          |
| **Won't Have (MVP)** | Seleccion de asiento por pasajero | La realiza el agente          |
| **Won't Have (MVP)** | Integracion con pasarelas de pago | Solo se registra el pago      |
| **Won't Have (MVP)** | Notificaciones por correo o SMS   | No forman parte del alcance   |
| **Won't Have (MVP)** | Multiples roles de usuario        | Solo existe el agente         |
