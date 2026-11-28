# B2 · Modelo lógico relacional SIGRES

**Motor elegido:** MySQL 8.0.16 o superior, InnoDB, `utf8mb4`.  
**Convención:** tablas y columnas en singular `snake_case`; PK `id_<tabla>`; las relaciones N:M usan tablas asociativas.

## Diagrama lógico

```mermaid
erDiagram
    ROL {
        int id_rol PK
        varchar nombre UK
        varchar descripcion
    }
    PERMISO {
        int id_permiso PK
        varchar codigo UK
        varchar descripcion
    }
    USUARIO {
        int id_usuario PK
        int id_rol FK
        varbinary nombre_cifrado
        varbinary apellido_cifrado
        varbinary identificacion_cifrada
        binary identificacion_hmac UK
        varchar nombre_usuario UK
        varchar password_hash
        varbinary telefono_cifrado
        boolean activo
    }
    ROL_PERMISO {
        int id_rol PK, FK
        int id_permiso PK, FK
    }
    MESA {
        int id_mesa PK
        int numero UK
        int capacidad
        varchar estado
    }
    CATEGORIA_MENU {
        int id_categoria PK
        varchar nombre UK
        varchar descripcion
        boolean activa
    }
    PRODUCTO_MENU {
        int id_producto PK
        int id_categoria FK
        varchar nombre
        varchar descripcion
        decimal precio
        varchar url_imagen
        boolean disponible
    }
    PEDIDO {
        int id_pedido PK
        int id_usuario FK
        int id_mesa FK
        datetime fecha_hora
        int numero_comensales
        varchar estado
    }
    DETALLE_PEDIDO {
        int id_pedido PK, FK
        int numero_linea PK
        int id_producto FK
        int cantidad
        decimal precio_unitario
        varchar nota
    }
    FACTURA {
        int id_factura PK
        int id_pedido FK, UK
        int id_cajero FK
        varchar numero UK
        datetime fecha_hora
        decimal subtotal
        decimal impuestos
        decimal total
        varchar estado
        text motivo_anulacion
        int anulada_por_usuario_id FK
        datetime anulada_en
    }
    METODO_PAGO {
        int id_metodo_pago PK
        varchar codigo UK
        varchar nombre UK
        boolean activo
    }
    PAGO {
        int id_pago PK
        int id_factura FK, UK
        int id_metodo_pago FK
        datetime fecha_hora
        decimal monto
        decimal propina
    }
    INSUMO {
        int id_insumo PK
        varchar nombre UK
        varchar categoria
        decimal existencia
        varchar unidad
        decimal minimo
        decimal costo_unitario
        varchar proveedor
    }
    PRODUCTO_INSUMO {
        int id_producto PK, FK
        int id_insumo PK, FK
        decimal cantidad_por_producto
    }
    MOVIMIENTO_INVENTARIO {
        int id_movimiento PK
        int id_insumo FK
        int id_usuario FK
        int id_pedido
        int numero_linea
        datetime fecha_hora
        varchar tipo
        decimal cantidad
        decimal existencia_resultante
        varchar observacion
    }
    ALERTA_INVENTARIO {
        int id_alerta PK
        int id_insumo FK
        int atendida_por_usuario_id FK
        datetime fecha_hora
        decimal existencia_actual
        decimal minimo_alerta
        varchar estado
        datetime atendida_en
    }

    ROL ||--o{ USUARIO : asigna
    ROL ||--o{ ROL_PERMISO : agrupa
    PERMISO ||--o{ ROL_PERMISO : concede
    USUARIO ||--o{ PEDIDO : registra
    MESA ||--o{ PEDIDO : recibe
    PEDIDO ||--|{ DETALLE_PEDIDO : contiene
    PRODUCTO_MENU ||--o{ DETALLE_PEDIDO : selecciona
    CATEGORIA_MENU ||--o{ PRODUCTO_MENU : clasifica
    PEDIDO ||--o| FACTURA : genera
    USUARIO ||--o{ FACTURA : cobra
    USUARIO |o--o{ FACTURA : anula
    FACTURA ||--o| PAGO : recibe
    METODO_PAGO ||--o{ PAGO : registra
    PRODUCTO_MENU ||--o{ PRODUCTO_INSUMO : requiere
    INSUMO ||--o{ PRODUCTO_INSUMO : compone
    INSUMO ||--o{ MOVIMIENTO_INVENTARIO : cambia
    USUARIO ||--o{ MOVIMIENTO_INVENTARIO : registra
    DETALLE_PEDIDO |o--o{ MOVIMIENTO_INVENTARIO : origina
    INSUMO ||--o{ ALERTA_INVENTARIO : genera
    USUARIO |o--o{ ALERTA_INVENTARIO : atiende
```

## Esquema de tablas

| Tabla | Columnas, PK y FK |
|---|---|
| `rol` | `id_rol` PK, `nombre` UNIQUE, `descripcion` |
| `permiso` | `id_permiso` PK, `codigo` UNIQUE, `descripcion` |
| `usuario` | `id_usuario` PK, `id_rol` FK → `rol`, nombre, apellido e identificación cifrados, HMAC único de identificación, nombre de usuario, hash de contraseña, teléfono cifrado opcional, estado |
| `rol_permiso` | `(id_rol, id_permiso)` PK compuesta; ambas columnas FK |
| `mesa` | `id_mesa` PK, `numero` UNIQUE, capacidad, estado |
| `categoria_menu` | `id_categoria` PK, nombre UNIQUE, descripción, activa |
| `producto_menu` | `id_producto` PK, `id_categoria` FK → `categoria_menu`, nombre, descripción, precio, URL de imagen, disponible; UNIQUE (`id_categoria`, `nombre`) |
| `pedido` | `id_pedido` PK, `id_usuario` FK → `usuario`, `id_mesa` FK → `mesa`, fecha, comensales, estado |
| `detalle_pedido` | `(id_pedido, numero_linea)` PK compuesta; `id_pedido` FK → `pedido`; `id_producto` FK → `producto_menu`; cantidad, precio vendido, nota |
| `factura` | `id_factura` PK, `numero` UNIQUE, `id_pedido` FK UNIQUE → `pedido`, `id_cajero` FK → `usuario`, fecha, importes, estado y datos de anulación |
| `metodo_pago` | `id_metodo_pago` PK, código y nombre UNIQUE, activo |
| `pago` | `id_pago` PK, `id_factura` FK UNIQUE → `factura`, `id_metodo_pago` FK → `metodo_pago`, fecha, monto, propina |
| `insumo` | `id_insumo` PK, nombre UNIQUE, categoría, existencia, unidad, mínimo, costo unitario, proveedor opcional |
| `producto_insumo` | `(id_producto, id_insumo)` PK compuesta; ambas FK; cantidad de insumo usada por producto |
| `movimiento_inventario` | `id_movimiento` PK, FK a `insumo` y `usuario`; FK compuesta opcional (`id_pedido`, `numero_linea`) → `detalle_pedido`; fecha, tipo, cantidad, saldo resultante |
| `alerta_inventario` | `id_alerta` PK, `id_insumo` FK → `insumo`, usuario que atiende opcional FK → `usuario`, fecha, saldos observados, estado y fecha de atención |

## Reglas de transformación aplicadas

| Regla | Aplicación en B1 | Resultado |
|---|---|---|
| Entidad fuerte | Rol, permiso, usuario, mesa, categoría, producto, pedido, factura, pago, método, insumo, movimiento y alerta | Una tabla con PK propia por entidad. |
| Relación 1:N | Rol-Usuario; Usuario-Pedido; Mesa-Pedido; Categoría-Producto; Insumo-Movimiento | La FK se ubica en el lado N: por ejemplo, `pedido.id_mesa`. |
| Relación 1:1 opcional | Pedido-Factura | `factura.id_pedido` es FK `UNIQUE`; un pedido puede no tener factura o tener una. |
| Relación 1:1 opcional | Factura-Pago | `pago.id_factura` es FK `UNIQUE`; refleja el supuesto de P1 de un solo pago por factura. |
| Relación N:M | Rol-Permiso | Tabla intermedia `rol_permiso` con PK compuesta. |
| Relación N:M con atributos | Pedido-ProductoMenu | Tabla `detalle_pedido`; clave compuesta por pedido y número de línea. Conserva cantidad, nota y precio histórico. |
| Relación N:M con atributos | ProductoMenu-Insumo | Tabla `producto_insumo`; cantidad requerida por unidad del producto. Permite calcular el descuento exigido por RF-10. |
| FK compuesta opcional | DetallePedido-MovimientoInventario | Un movimiento automático puede apuntar a la línea que lo originó; las entradas manuales no llevan pedido ni línea. |
| Atributo multivaluado | No hay un atributo con varios valores indivisibles en el SRS | El teléfono de empleado es opcional y singular; los productos de un pedido se normalizan mediante `detalle_pedido`. |
| Recursividad / generalización | No identificadas en el alcance SIGRES | No se agrega una relación o jerarquía sin sustento del SRS. |
| Reporte calculado | RF-15 | Vista `vw_reporte_ventas`, no tabla de almacenamiento. |

## Reglas de negocio que no caben en una FK

- Al confirmar un pedido, en una transacción se guarda el pedido, sus detalles, los movimientos de salida calculados con `producto_insumo` y el nuevo saldo de cada insumo. Si una salida deja saldo negativo, se revierte la operación.
- El precio de `detalle_pedido` es una copia histórica del precio de menú al confirmar el pedido; no se actualiza al cambiar el precio del producto.
- Una factura tiene el estado `Pendiente`, `Pagada` o `Anulada`. Una factura pagada no se modifica; solo el Administrador puede anularla con motivo.
- No se almacena información personal del cliente QR. La consulta del consumo se limita al pedido activo asociado a la mesa identificada por el QR.
- La vista de reportes considera facturas pagadas y permite filtrar por fechas, producto y mesero.
