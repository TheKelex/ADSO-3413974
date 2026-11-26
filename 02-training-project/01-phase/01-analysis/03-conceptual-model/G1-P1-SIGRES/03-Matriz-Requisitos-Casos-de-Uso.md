# Matriz requisito → caso de uso

| ID | Requisito funcional | Caso(s) del diagrama G1 | Caso de uso del SRS | Cobertura |
|---|---|---|---|---|
| RF-01 | Registrar usuarios | G1-UC-01 Registrar usuarios | CU-01 Registrar Usuario | Cubierto. |
| RF-02 | Iniciar sesión | G1-UC-02 Iniciar sesión | CU-02 Inicio de Sesión | Cubierto. |
| RF-03 | Gestionar roles y permisos | G1-UC-03 Gestionar roles y permisos | CU-03 Gestión de Roles y Permisos | Cubierto. |
| RF-04 | Registrar pedido / comanda | G1-UC-04 Registrar pedido | CU-04 Registrar Pedido | Cubierto; el envío a cocina se muestra como comportamiento incluido. |
| RF-05 | Enviar pedido a cocina | G1-UC-05 Enviar pedido a cocina | CU-05 Enviar y Visualizar Pedido en Cocina | Cubierto; acción automática tras confirmar el pedido. |
| RF-06 | Consultar estado del pedido | G1-UC-06 Consultar estado del pedido | CU-05 (parte del flujo de cocina/pedidos) | Cubierto como objetivo separado para facilitar trazabilidad; confirmar la separación. |
| RF-07 | Actualizar estado del pedido | G1-UC-07 Actualizar estado del pedido | CU-05 (parte del flujo de cocina/pedidos) | Cubierto como objetivo separado para facilitar trazabilidad; confirmar la separación. |
| RF-08 | Administrar productos del menú | G1-UC-08 Administrar productos del menú | Sin caso exclusivo inequívoco | Cubierto según la ficha detallada de RF-08; se propone el caso separado. |
| RF-09 | Gestionar categorías del menú | G1-UC-09 Gestionar categorías del menú | Sin caso inequívoco en el índice de casos | Cubierto; crear o confirmar caso específico al validar. |
| RF-10 | Registrar y actualizar inventario | G1-UC-10 Gestionar inventario | CU-06 Registrar y Actualizar Inventario | Cubierto. |
| RF-11 | Generar alertas de existencias mínimas | G1-UC-11 Generar alertas de inventario | CU-06 (flujo automático) | Cubierto; se activa cuando el stock cae bajo el mínimo configurado. |
| RF-12 | Generar factura | G1-UC-12 Generar factura | CU-07 Generar Factura y Procesar Pago | Cubierto; caso separado del pago para mayor trazabilidad. |
| RF-13 | Procesar pago | G1-UC-13 Procesar pago | CU-07 Generar Factura y Procesar Pago | Cubierto; proveedor externo de pagos queda fuera del alcance confirmado. |
| RF-14 | Gestionar mesas | G1-UC-14 Gestionar mesas | Sin caso inequívoco entre los nueve casos listados | Cubierto según ficha detallada; validar el nombre y límites del caso. |
| RF-15 | Generar reportes de ventas | G1-UC-15 Generar reportes | CU-08 Generar Reportes de Ventas | Cubierto. |
| RF-16 | Ver menú digital por QR | G1-UC-16 Ver menú digital por QR | CU-09 Ver Menú Digital y Consultar Consumo | Cubierto; comportamiento base del caso. |
| RF-17 | Consultar consumo actual | G1-UC-17 Consultar consumo actual | CU-09 Ver Menú Digital y Consultar Consumo | Cubierto como extensión condicional del menú digital. |

## Resultado de cobertura

- Requisitos funcionales del SRS: **17**.
- Requisitos funcionales trazados: **17**.
- Cobertura preliminar: **17/17 = 100 %**.
- Requisitos funcionales sin caso: **0** en esta propuesta.