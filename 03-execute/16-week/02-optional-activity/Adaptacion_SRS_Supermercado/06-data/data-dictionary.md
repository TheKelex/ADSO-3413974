# data-dictionary

> Estado: 🟡 | Última actualización: 2026-07-03
> Autor: Kevin Santiago Saavedra Chantris | Equipo: Kevin Santiago Saavedra Chantris / SENA ADSO 3413974

| Entidad | Campo | Descripción |
|---------|-------|-------------|
| Producto | id, nombre, precio, stock, estado | Información comercial y de inventario |
| Venta | id, fecha, total, cambio | Registro de transacción |
| DetalleVenta | venta_id, producto_id, cantidad, subtotal | Detalle de productos vendidos |
| Proveedor | id, nombre, contacto, productos, tiempo_entrega | Datos del proveedor |
| Fiado | id, cliente_id, fecha, total_pendiente | Crédito interno |
| Abono | id, fiado_id, fecha, valor | Pago parcial |
