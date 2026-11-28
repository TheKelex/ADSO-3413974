# B3 · Evidencia de normalización

Se muestran dos relaciones del propio modelo SIGRES paso a paso. La primera usa llave compuesta y dependencia parcial; la segunda muestra dependencia transitiva. El estado final coincide con B2 y el diccionario B4.

## Caso 1: líneas de pedido (`detalle_pedido`)

### 0FN: grupo repetido

`PEDIDO_0FN(id_pedido, fecha_hora, id_mesa, numero_mesa, id_mesero, nombre_mesero, items[(numero_linea, id_producto, nombre_producto, id_categoria, nombre_categoria, cantidad, precio_unitario, nota)])`

`items` contiene una lista de líneas dentro de la fila. La llave candidata del detalle es `(id_pedido, numero_linea)`.

### 1FN: valores atómicos

Se separa una fila por línea de producto:

`DETALLE_PEDIDO_1FN(id_pedido, numero_linea, fecha_hora, id_mesa, numero_mesa, id_mesero, nombre_mesero, id_producto, nombre_producto, id_categoria, nombre_categoria, cantidad, precio_unitario, nota)`

PK: `(id_pedido, numero_linea)`. Todos los valores son atómicos, pero datos del pedido y del producto se repiten en cada línea.

**Dependencias funcionales (DF):**

- `(id_pedido, numero_linea) → id_producto, cantidad, precio_unitario, nota`
- `id_pedido → fecha_hora, id_mesa, id_mesero`
- `id_mesa → numero_mesa`
- `id_mesero → nombre_mesero`
- `id_producto → nombre_producto, id_categoria`
- `id_categoria → nombre_categoria`

### 2FN: eliminar dependencias parciales de la llave compuesta

`fecha_hora`, `id_mesa` e `id_mesero` dependen solo de `id_pedido`, componente de la llave. Por ello se separa `pedido`. Las líneas conservan su llave compuesta:

- `PEDIDO_2FN(id_pedido PK, fecha_hora, id_mesa, numero_mesa, id_mesero, nombre_mesero)`
- `PRODUCTO_MENU_2FN(id_producto PK, nombre_producto, id_categoria, nombre_categoria)`
- `DETALLE_PEDIDO(id_pedido PK/FK, numero_linea PK, id_producto FK, cantidad, precio_unitario, nota)`

En `DETALLE_PEDIDO`, los atributos no llave dependen de la llave completa `(id_pedido, numero_linea)`. El precio es el precio vendido en ese pedido, no el precio actual del menú.

### 3FN: eliminar dependencias transitivas

Persisten `id_mesa → numero_mesa`, `id_mesero → nombre_mesero` e `id_categoria → nombre_categoria`. Se mueven esos atributos a sus entidades:

- `MESA(id_mesa PK, numero, capacidad, estado)`
- `USUARIO(id_usuario PK, nombre, apellido, ...)`
- `CATEGORIA_MENU(id_categoria PK, nombre, descripcion, activa)`
- `PEDIDO(id_pedido PK, fecha_hora, id_mesa FK, id_usuario FK, numero_comensales, estado)`
- `PRODUCTO_MENU(id_producto PK, id_categoria FK, nombre, descripcion, precio, url_imagen, disponible)`
- `DETALLE_PEDIDO(id_pedido PK/FK, numero_linea PK, id_producto FK, cantidad, precio_unitario, nota)`

Cada atributo no llave de estas relaciones depende de la llave, de toda la llave y de nada más que la llave. La asociación N:M entre pedido y producto queda resuelta por `detalle_pedido`.

## Caso 2: facturación (`factura`)

### 0FN: datos de cajero anidados

`FACTURA_0FN(id_factura, numero, fecha_hora, id_pedido, subtotal, impuestos, total, cajero[(id_cajero, nombre_cajero, apellido_cajero, id_rol, nombre_rol)])`

El grupo `cajero` es compuesto; `id_factura` es la llave candidata.

### 1FN: atributos atómicos

`FACTURA_1FN(id_factura, numero, fecha_hora, id_pedido, subtotal, impuestos, total, estado, id_cajero, nombre_cajero, apellido_cajero, id_rol, nombre_rol)`

**DF:**

- `id_factura → numero, fecha_hora, id_pedido, subtotal, impuestos, total, estado, id_cajero, nombre_cajero, apellido_cajero, id_rol, nombre_rol`
- `id_cajero → nombre_cajero, apellido_cajero, id_rol`
- `id_rol → nombre_rol`

### 2FN

La llave es simple (`id_factura`), así que no hay dependencia parcial. La relación permanece igual que en 1FN.

### 3FN: eliminar la dependencia transitiva

`id_factura → id_cajero → nombre_cajero, apellido_cajero, id_rol → nombre_rol` es transitiva. Los datos de empleado y rol ya pertenecen a sus tablas:

- `FACTURA(id_factura PK, numero UNIQUE, id_pedido FK UNIQUE, id_cajero FK, fecha_hora, subtotal, impuestos, total, estado, motivo_anulacion, anulada_por_usuario_id FK, anulada_en)`
- `USUARIO(id_usuario PK, id_rol FK, nombre, apellido, ...)`
- `ROL(id_rol PK, nombre, descripcion)`

La factura guarda el identificador del cajero, no duplica su nombre ni el nombre del rol. El estado y los datos de anulación permiten aplicar RF-12 y la regla del SRS que impide modificar una factura pagada, salvo anulación justificada por el Administrador.

## Datos derivados conservados con justificación

- `detalle_pedido.precio_unitario`: instantánea del precio de venta para conservar el valor histórico aunque cambie `producto_menu.precio`.
- `factura.subtotal`, `impuestos` y `total`: instantánea legal/operativa de los importes al emitir la factura; se verifican y quedan protegidos cuando se paga.
- `insumo.existencia`: saldo actual para responder a RF-10/RF-11; cada variación se registra en `movimiento_inventario` y ambas escrituras se hacen en una misma transacción.
- `movimiento_inventario.existencia_resultante` y los valores de `alerta_inventario`: instantáneas de auditoría, no sustituyen el historial de movimientos.
