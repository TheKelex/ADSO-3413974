# models

> Estado: 🟡 | Última actualización: 2026-07-03
> Autor: Kevin Santiago Saavedra Chantris | Equipo: Kevin Santiago Saavedra Chantris / SENA ADSO 3413974

## Modelo conceptual

- Usuario
- Producto
- Venta
- DetalleVenta
- Inventario
- Proveedor
- Pedido
- Cliente
- Fiado
- Abono

## Relaciones principales

- Una venta tiene varios detalles.
- Un detalle de venta referencia un producto.
- Un fiado pertenece a un cliente.
- Un fiado puede tener varios abonos.
- Un proveedor puede tener muchos pedidos.
