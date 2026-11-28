# B4 · Diccionario de datos

## Convención

Tablas y columnas en singular `snake_case`; PK autoincremental `id_<tabla>`; las tablas asociativas usan PK compuesta. Motor: MySQL 8.0.16+, InnoDB, `utf8mb4`. `DECIMAL(12,2)` representa importes en COP y `DECIMAL(12,3)` cantidades de insumos.

Abreviaturas: PK = llave primaria; FK = llave foránea; UQ = única; NN = `NOT NULL`; AI = `AUTO_INCREMENT`.

## Tablas y columnas

| Tabla | Columna | Tipo y tamaño | Restricción | Dominio / regla |
|---|---|---|---|---|
| `rol` | `id_rol` | INT UNSIGNED | PK, NN, AI | Identificador positivo generado por MySQL. |
| `rol` | `nombre` | VARCHAR(40) | NN, UQ | Nombre de rol; no vacío. |
| `rol` | `descripcion` | VARCHAR(200) | NN | Propósito del rol. |
| `permiso` | `id_permiso` | INT UNSIGNED | PK, NN, AI | Identificador positivo. |
| `permiso` | `codigo` | VARCHAR(60) | NN, UQ | Código estable de autorización; mayúsculas y guion bajo. |
| `permiso` | `descripcion` | VARCHAR(200) | NN | Acción que concede el permiso. |
| `usuario` | `id_usuario` | INT UNSIGNED | PK, NN, AI | Identificador del empleado. |
| `usuario` | `id_rol` | INT UNSIGNED | FK → `rol.id_rol`, NN | Un rol asignado a cada usuario, según RF-01. |
| `usuario` | `nombre_cifrado` | VARBINARY(512) | NN | Nombre cifrado en la aplicación con AES-256-GCM; almacena nonce, etiqueta y texto cifrado. |
| `usuario` | `apellido_cifrado` | VARBINARY(512) | NN | Apellido cifrado en la aplicación con AES-256-GCM; almacena nonce, etiqueta y texto cifrado. |
| `usuario` | `identificacion_cifrada` | VARBINARY(512) | NN | Identificación cifrada con AES-256-GCM; almacena nonce, etiqueta y texto cifrado. |
| `usuario` | `identificacion_hmac` | BINARY(32) | NN, UQ | HMAC-SHA-256 con clave separada para validar unicidad sin guardar identificación en claro. |
| `usuario` | `nombre_usuario` | VARCHAR(50) | NN, UQ | Cuenta de inicio de sesión; sin espacios al inicio o final. |
| `usuario` | `password_hash` | VARCHAR(255) | NN | Hash de contraseña generado con bcrypt; la sal única queda incorporada por `password_hash()` de PHP. No almacenar texto claro. |
| `usuario` | `telefono_cifrado` | VARBINARY(512) | NULL | Teléfono opcional cifrado con AES-256-GCM; clave fuera de la base de datos. |
| `usuario` | `activo` | TINYINT(1) | NN, DEFAULT 1, CHECK 0/1 | `1` activo; `0` desactivado. |
| `rol_permiso` | `id_rol` | INT UNSIGNED | PK compuesta, FK → `rol`, NN | Rol que integra la relación N:M. |
| `rol_permiso` | `id_permiso` | INT UNSIGNED | PK compuesta, FK → `permiso`, NN | Permiso otorgado al rol. |
| `mesa` | `id_mesa` | INT UNSIGNED | PK, NN, AI | Identificador de mesa. |
| `mesa` | `numero` | SMALLINT UNSIGNED | NN, UQ | Número positivo único dentro del restaurante. |
| `mesa` | `capacidad` | TINYINT UNSIGNED | NN, CHECK > 0 | Cantidad máxima de comensales. |
| `mesa` | `estado` | VARCHAR(20) | NN, CHECK | `Disponible`, `Ocupada` o `Reservada`. |
| `categoria_menu` | `id_categoria` | INT UNSIGNED | PK, NN, AI | Identificador de categoría. |
| `categoria_menu` | `nombre` | VARCHAR(80) | NN, UQ | Nombre único; por ejemplo, entradas o bebidas. |
| `categoria_menu` | `descripcion` | VARCHAR(250) | NULL | Descripción de la categoría. |
| `categoria_menu` | `activa` | TINYINT(1) | NN, DEFAULT 1, CHECK 0/1 | Si la categoría se muestra en el menú. |
| `producto_menu` | `id_producto` | INT UNSIGNED | PK, NN, AI | Identificador del producto. |
| `producto_menu` | `id_categoria` | INT UNSIGNED | FK → `categoria_menu`, NN | Categoría del producto. |
| `producto_menu` | `nombre` | VARCHAR(100) | NN, UQ con `id_categoria` | Nombre dentro de la categoría. |
| `producto_menu` | `descripcion` | VARCHAR(500) | NN | Texto del menú QR y de la aplicación. |
| `producto_menu` | `precio` | DECIMAL(12,2) | NN, CHECK ≥ 0 | Precio vigente de venta en COP. |
| `producto_menu` | `url_imagen` | VARCHAR(500) | NN | Ruta HTTPS de la imagen del menú. |
| `producto_menu` | `disponible` | TINYINT(1) | NN, DEFAULT 1, CHECK 0/1 | `1` disponible; `0` oculto/no disponible. |
| `pedido` | `id_pedido` | INT UNSIGNED | PK, NN, AI | Identificador de pedido. |
| `pedido` | `id_usuario` | INT UNSIGNED | FK → `usuario`, NN | Mesero que registra el pedido. |
| `pedido` | `id_mesa` | INT UNSIGNED | FK → `mesa`, NN | Mesa asociada al pedido. |
| `pedido` | `fecha_hora` | DATETIME | NN | Fecha y hora del registro. |
| `pedido` | `numero_comensales` | TINYINT UNSIGNED | NN, CHECK > 0 | Comensales atendidos en la mesa. |
| `pedido` | `estado` | VARCHAR(20) | NN, CHECK | `Pendiente`, `EnPreparacion`, `Listo`, `Entregado` o `Cancelado`. |
| `detalle_pedido` | `id_pedido` | INT UNSIGNED | PK compuesta, FK → `pedido`, NN | Pedido al que pertenece la línea. |
| `detalle_pedido` | `numero_linea` | SMALLINT UNSIGNED | PK compuesta, NN | Posición positiva dentro del pedido. |
| `detalle_pedido` | `id_producto` | INT UNSIGNED | FK → `producto_menu`, NN | Producto vendido. |
| `detalle_pedido` | `cantidad` | SMALLINT UNSIGNED | NN, CHECK > 0 | Unidades solicitadas. |
| `detalle_pedido` | `precio_unitario` | DECIMAL(12,2) | NN, CHECK ≥ 0 | Precio COP capturado al confirmar; instantánea histórica. |
| `detalle_pedido` | `nota` | VARCHAR(500) | NULL | Instrucción especial del cliente para la línea. |
| `factura` | `id_factura` | INT UNSIGNED | PK, NN, AI | Identificador interno. |
| `factura` | `numero` | VARCHAR(30) | NN, UQ | Número visible de factura. |
| `factura` | `id_pedido` | INT UNSIGNED | FK → `pedido`, NN, UQ | Como máximo una factura por pedido. |
| `factura` | `id_cajero` | INT UNSIGNED | FK → `usuario`, NN | Cajero que emite la factura. |
| `factura` | `fecha_hora` | DATETIME | NN | Fecha y hora de emisión. |
| `factura` | `subtotal` | DECIMAL(12,2) | NN, CHECK ≥ 0 | Suma de las líneas antes de impuestos. |
| `factura` | `impuestos` | DECIMAL(12,2) | NN, CHECK ≥ 0 | Impuestos aplicables. |
| `factura` | `total` | DECIMAL(12,2) | NN, CHECK = subtotal + impuestos | Total emitido en COP. |
| `factura` | `estado` | VARCHAR(15) | NN, CHECK | `Pendiente`, `Pagada` o `Anulada`. |
| `factura` | `motivo_anulacion` | TEXT | NULL | Obligatorio y no vacío al anular. |
| `factura` | `anulada_por_usuario_id` | INT UNSIGNED | FK → `usuario` (UPDATE/DELETE RESTRICT), NULL | Administrador que autorizó la anulación. |
| `factura` | `anulada_en` | DATETIME | NULL | Fecha y hora de anulación. |
| `metodo_pago` | `id_metodo_pago` | INT UNSIGNED | PK, NN, AI | Identificador del método. |
| `metodo_pago` | `codigo` | VARCHAR(30) | NN, UQ | `EFECTIVO`, `TARJETA`, `TRANSFERENCIA` o `BILLETERA_DIGITAL`. |
| `metodo_pago` | `nombre` | VARCHAR(60) | NN, UQ | Nombre visible del método. |
| `metodo_pago` | `activo` | TINYINT(1) | NN, DEFAULT 1, CHECK 0/1 | Si se ofrece actualmente al cajero. |
| `pago` | `id_pago` | INT UNSIGNED | PK, NN, AI | Identificador del pago. |
| `pago` | `id_factura` | INT UNSIGNED | FK → `factura`, NN, UQ | Un pago registrado por factura; supuesto de P1. |
| `pago` | `id_metodo_pago` | INT UNSIGNED | FK → `metodo_pago`, NN | Método elegido. |
| `pago` | `fecha_hora` | DATETIME | NN | Fecha y hora del pago. |
| `pago` | `monto` | DECIMAL(12,2) | NN, CHECK > 0 | Monto pagado; debe igualar el total pendiente de la factura. |
| `pago` | `propina` | DECIMAL(12,2) | NN, DEFAULT 0, CHECK ≥ 0 | Propina registrada aparte del total facturado. |
| `insumo` | `id_insumo` | INT UNSIGNED | PK, NN, AI | Identificador de insumo. |
| `insumo` | `nombre` | VARCHAR(100) | NN, UQ | Nombre del insumo. |
| `insumo` | `categoria` | VARCHAR(80) | NN | Clasificación del insumo para el módulo de inventario. |
| `insumo` | `existencia` | DECIMAL(12,3) | NN, CHECK ≥ 0 | Saldo actual; unidad indicada en `unidad`. |
| `insumo` | `unidad` | VARCHAR(20) | NN | Unidad consistente con la receta: g, kg, ml, L, unidad, etc. |
| `insumo` | `minimo` | DECIMAL(12,3) | NN, CHECK ≥ 0 | Umbral que activa la alerta cuando existencia ≤ mínimo. |
| `insumo` | `costo_unitario` | DECIMAL(12,2) | NN, CHECK ≥ 0 | Costo por unidad de compra, en COP. |
| `insumo` | `proveedor` | VARCHAR(120) | NULL | Proveedor del insumo, si se registra. |
| `producto_insumo` | `id_producto` | INT UNSIGNED | PK compuesta, FK → `producto_menu`, NN | Producto cuya receta se define. |
| `producto_insumo` | `id_insumo` | INT UNSIGNED | PK compuesta, FK → `insumo`, NN | Insumo utilizado. |
| `producto_insumo` | `cantidad_por_producto` | DECIMAL(12,3) | NN, CHECK > 0 | Cantidad necesaria para una unidad del producto. |
| `movimiento_inventario` | `id_movimiento` | INT UNSIGNED | PK, NN, AI | Identificador del movimiento. |
| `movimiento_inventario` | `id_insumo` | INT UNSIGNED | FK → `insumo`, NN | Insumo afectado. |
| `movimiento_inventario` | `id_usuario` | INT UNSIGNED | FK → `usuario`, NN | Empleado que registra o confirma el movimiento. |
| `movimiento_inventario` | `id_pedido` | INT UNSIGNED | FK compuesta opcional → `detalle_pedido` (UPDATE/DELETE RESTRICT) | Pedido que originó el consumo automático. |
| `movimiento_inventario` | `numero_linea` | SMALLINT UNSIGNED | FK compuesta opcional → `detalle_pedido` (UPDATE/DELETE RESTRICT) | Línea de pedido que originó el consumo; nula junto con `id_pedido` para entradas manuales. |
| `movimiento_inventario` | `fecha_hora` | DATETIME | NN | Fecha y hora del cambio. |
| `movimiento_inventario` | `tipo` | VARCHAR(20) | NN, CHECK | `Entrada`, `Salida`, `AjusteEntrada` o `AjusteSalida`. |
| `movimiento_inventario` | `cantidad` | DECIMAL(12,3) | NN, CHECK > 0 | Magnitud positiva del movimiento; el tipo indica el sentido. |
| `movimiento_inventario` | `existencia_resultante` | DECIMAL(12,3) | NN, CHECK ≥ 0 | Saldo del insumo después de aplicar el movimiento. |
| `movimiento_inventario` | `observacion` | VARCHAR(250) | NULL | Justificación breve de un ajuste manual. |
| `alerta_inventario` | `id_alerta` | INT UNSIGNED | PK, NN, AI | Identificador de alerta. |
| `alerta_inventario` | `id_insumo` | INT UNSIGNED | FK → `insumo`, NN | Insumo bajo el mínimo. |
| `alerta_inventario` | `atendida_por_usuario_id` | INT UNSIGNED | FK → `usuario` (UPDATE/DELETE RESTRICT), NULL | Administrador que marcó la alerta atendida. |
| `alerta_inventario` | `fecha_hora` | DATETIME | NN | Fecha y hora de generación. |
| `alerta_inventario` | `existencia_actual` | DECIMAL(12,3) | NN, CHECK ≥ 0 | Saldo observado al generar la alerta. |
| `alerta_inventario` | `minimo_alerta` | DECIMAL(12,3) | NN, CHECK ≥ 0 | Umbral observado al generar la alerta. |
| `alerta_inventario` | `estado` | VARCHAR(10) | NN, CHECK | `Activa` o `Atendida`. |
| `alerta_inventario` | `atendida_en` | DATETIME | NULL | Fecha y hora de atención. |

## Vista derivada

`vw_reporte_ventas` no es tabla persistida. Expone `fecha_venta` (DATE), `id_mesero` (INT UNSIGNED), `id_producto` (INT UNSIGNED), `producto` (VARCHAR(100)), `cantidad_vendida` (DECIMAL(32,0)) y `total_vendido` (DECIMAL(34,2)); se calcula desde facturas pagadas y sus detalles para RF-15. Sus filtros son por fecha, producto y mesero.

## Reglas referenciales

- Tablas asociativas y líneas: borrado en cascada solo dentro de la asociación (`rol_permiso`, `producto_insumo`, `detalle_pedido`).
- Usuarios, mesas, productos, insumos, facturas, pagos y movimientos con historia se protegen con `ON DELETE RESTRICT`.
- Las FK usan `ON UPDATE CASCADE`, excepto las llaves que también participan en restricciones `CHECK`; allí se usa `ON UPDATE RESTRICT`, ya que los identificadores referenciados no se modifican. Las acciones exactas aparecen en `05-B5-Crear-Base-de-Datos.sql`.
