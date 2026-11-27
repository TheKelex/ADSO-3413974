# Matriz de trazabilidad

| RF | Caso de uso G1 | Plantilla / actividad | Clase de dominio | Entidad | Pantalla crítica |
|---|---|---|---|---|---|
| RF-01 | G1-UC-01 Registrar usuarios | Diagrama de Administración | Usuario | USUARIO | — |
| RF-02 | G1-UC-02 Iniciar sesión | Diagrama de Administración | Usuario | USUARIO | — |
| RF-03 | G1-UC-03 Gestionar roles y permisos | Diagrama de Administración | Usuario, Rol, Permiso | USUARIO, USUARIO_ROL, ROL, PERMISO, ROL_PERMISO | — |
| RF-04 | G1-UC-04 Registrar pedido | Plantilla G1-UC-04; actividad 01 | Pedido, Mesa, DetallePedido, ProductoMenu | PEDIDO, MESA, DETALLE_PEDIDO, PRODUCTO_MENU | [G1-UC-04 · Registrar pedido](./05-MockUp/G1-UC-04-Registrar-pedido.png) |
| RF-05 | G1-UC-05 Enviar pedido a cocina | Incluido en G1-UC-04; actividad 01 | Pedido, DetallePedido | PEDIDO, DETALLE_PEDIDO | [G1-UC-04 · Registrar pedido](./05-MockUp/G1-UC-04-Registrar-pedido.png) |
| RF-06 | G1-UC-06 Consultar estado del pedido | Diagrama de Pedidos y atención | Pedido | PEDIDO | — |
| RF-07 | G1-UC-07 Actualizar estado del pedido | Diagrama de Pedidos y atención | Pedido | PEDIDO | — |
| RF-08 | G1-UC-08 Administrar productos del menú | Diagrama de Administración | ProductoMenu, CategoriaMenu | PRODUCTO_MENU, CATEGORIA_MENU | — |
| RF-09 | G1-UC-09 Gestionar categorías del menú | Diagrama de Administración | CategoriaMenu | CATEGORIA_MENU | — |
| RF-10 | G1-UC-10 Gestionar inventario | Plantilla G1-UC-10; actividad 02 | Insumo, MovimientoInventario | INSUMO, MOVIMIENTO_INVENTARIO | [G1-UC-10 · Gestionar inventario](./05-MockUp/G1-UC-10-Gestionar-Inventario.png) |
| RF-11 | G1-UC-11 Generar alertas de inventario | Plantilla G1-UC-10; actividad 02 | Insumo, AlertaInventario | INSUMO, ALERTA_INVENTARIO | — |
| RF-12 | G1-UC-12 Generar factura | Diagrama de Pedidos y atención | Factura, Pedido, DetallePedido | FACTURA, PEDIDO, DETALLE_PEDIDO | — |
| RF-13 | G1-UC-13 Procesar pago | Plantilla G1-UC-13 | Factura, Pago | FACTURA, PAGO | [G1-UC-13 · Procesar pago](./05-MockUp/G1-UC-13-Procesar-Pago.png) |
| RF-14 | G1-UC-14 Gestionar mesas | Diagrama de Pedidos y atención | Mesa, Pedido | MESA, PEDIDO | — |
| RF-15 | G1-UC-15 Generar reportes | Diagrama de Administración | Factura, Pago | FACTURA, PAGO | — |
| RF-16 | G1-UC-16 Ver menú digital por QR | Diagrama de Pedidos y atención | ProductoMenu, CategoriaMenu | PRODUCTO_MENU, CATEGORIA_MENU | — |
| RF-17 | G1-UC-17 Consultar consumo actual | Diagrama de Pedidos y atención | Pedido, Mesa, DetallePedido | PEDIDO, MESA, DETALLE_PEDIDO | — |

**Cobertura preliminar:** 17 de 17 RF (100 %).  
Las tres pantallas críticas están vinculadas al mockup.
