# entities-and-rules

> Estado: 🟢 | Última actualización: 2026-07-03
> Autor: Kevin Santiago Saavedra Chantris | Equipo: Kevin Santiago Saavedra Chantris / SENA ADSO 3413974

## Entidades principales

- **Usuario:** accede al sistema según su rol.
- **Producto:** ítem comercial con precio, stock y estado.
- **Venta:** transacción registrada con fecha, total y cambio.
- **Inventario:** control del stock disponible y movimientos.
- **Proveedor:** tercero que suministra productos.
- **Fiado:** registro de deuda asociada a un cliente.
- **Cliente:** persona asociada a fiados y abonos.
- **Pedido:** compra o solicitud registrada a un proveedor.

## Reglas del negocio

- Toda venta debe registrar productos, cantidades, total y cambio.
- El inventario debe actualizarse al confirmar una venta.
- Un producto agotado o con bajo stock debe generar alerta.
- Cada fiado debe guardar cliente, fecha y saldo pendiente.
- Los pedidos a proveedores deben conservar historial.
- El acceso a funciones críticas debe depender del rol del usuario.
