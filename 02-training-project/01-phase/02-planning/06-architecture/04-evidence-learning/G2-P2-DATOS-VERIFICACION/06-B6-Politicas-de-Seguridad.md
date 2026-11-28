# B6 · Políticas de seguridad de datos

## Alcance

Se protege la información del SRS de SIGRES: cuentas y datos de empleados, pedidos, pagos e inventario. El cliente QR no crea cuenta y no se almacenan sus datos personales. Las siguientes medidas se implementan en la aplicación y en el motor MySQL 8; los permisos de la matriz se aplican mediante cuentas de servicio y autorización del backend, no con una cuenta raíz compartida.

## Políticas por principio

### Integridad

| ID | Política concreta | Control / evidencia |
|---|---|---|
| I-01 | Las 16 tablas de negocio usan InnoDB. Las relaciones usan PK/FK, `UNIQUE`, `NOT NULL` y `CHECK` para impedir referencias huérfanas, duplicados y cantidades o importes inválidos. | B2, B4, B5; 32 restricciones `CHECK` verificadas en MySQL 8.0.46. |
| I-02 | Confirmar un pedido y descontar la receta es una sola transacción. El servicio bloquea los saldos de `insumo` con `SELECT ... FOR UPDATE`, calcula con `producto_insumo`, crea `movimiento_inventario` y actualiza `existencia`; si un saldo quedaría negativo, revierte todo. | RF-10; `detalle_pedido`, `producto_insumo`, `insumo`, `movimiento_inventario`. |
| I-03 | Una factura `Pagada` no permite cambiar sus datos; el único cambio permitido es pasar a `Anulada` conservando importes e identificadores, con motivo, Administrador responsable y fecha. Una factura anulada queda inmutable. | Trigger `trg_factura_paga_inmutable` en B5; campos de anulación en `factura`. La autorización del Administrador se verifica también en el backend. |
| I-04 | Cada ajuste manual de existencias debe indicar usuario, hora, tipo, cantidad y observación. El saldo resultante y el movimiento se guardan en la misma transacción; no se borran movimientos ni facturas históricas. | `movimiento_inventario`; FK con `ON DELETE RESTRICT`; permisos sin borrado para roles operativos. |

### Confidencialidad

| ID | Política concreta | Control / evidencia |
|---|---|---|
| C-01 | Aplicar mínimo privilegio por rol y restringir filas del mesero a sus pedidos. Cocina ve únicamente pedidos necesarios para preparar; el cliente QR ve menú y consumo activo, nunca cuentas, pagos ni datos de empleados. | Matriz de acceso de este artefacto y autorización en backend. |
| C-02 | Cifrar `nombre_cifrado`, `apellido_cifrado`, `identificacion_cifrada` y `telefono_cifrado` en la aplicación con AES-256-GCM. Guardar nonce y etiqueta junto al texto cifrado; mantener la clave fuera de MySQL en un gestor de secretos. `identificacion_hmac` es HMAC-SHA-256 con clave separada para validar unicidad sin almacenar el documento en claro. | B2/B4; accesos a la clave auditados y limitados al servicio autorizado. |
| C-03 | Las contraseñas se almacenan exclusivamente como bcrypt con sal aleatoria incluida por `password_hash()` de PHP; validar con `password_verify()`. Nunca registrar, transmitir ni almacenar la contraseña en claro. | `usuario.password_hash`; parámetros de coste actualizados según capacidad y política del despliegue. |
| C-04 | TLS 1.2 o superior protege navegador/app y app/MySQL; MySQL solo acepta conexiones desde la red privada de aplicación. Usar cuentas MySQL distintas por ambiente y rotar secretos; no exponer credenciales en el repositorio. | RNF de seguridad; TLS en el proxy y en el canal de base de datos. |
| C-05 | Tratar identificación, nombres y teléfono de empleados solo para autenticación, operación y auditoría, con aviso y autorización conforme a la normativa aplicable. No guardar datos personales de clientes QR ni números completos/CVV de tarjetas; conservar solo el método de pago y los valores requeridos por el SRS. | Campos cifrados de `usuario`; la vista pública no consulta datos personales. |

### Disponibilidad

| ID | Política concreta | Control / evidencia |
|---|---|---|
| D-01 | Ejecutar respaldo completo cifrado cada 24 horas y conservar 30 días en una cuenta o ubicación separada. Mantener registros binarios cifrados para recuperación entre respaldos. Objetivo inicial: RPO ≤ 24 h y RTO ≤ 4 h. | Monitoreo de trabajos y alertas ante respaldo fallido. |
| D-02 | Probar restauración completa en ambiente aislado al menos una vez al mes; registrar fecha, duración, versión y resultado. No declarar un respaldo válido sin haber comprobado que restaura y que el esquema pasa sus verificaciones. | Acta mensual de restauración y verificación de tablas/relaciones. |
| D-03 | Monitorear conexión, espacio y errores del motor cada minuto; alertar al Administrador y mantener procedimiento de recuperación documentado. Aplicación con reintentos limitados e idempotentes, sin confirmar un pedido si falla la transacción de inventario. | Registros de salud y alertas operativas; disponibilidad de pedidos e inventario. |

## Matriz de acceso por rol

Abreviaturas: `S` = SELECT; `I` = INSERT; `U` = UPDATE; `—` = sin acceso. Para tablas sensibles, el backend limita columnas y filas; no se otorga `DELETE` a roles operativos. El Administrador no usa su cuenta de aplicación como `root`.

| Rol | `usuario` | `rol` / `permiso` | `mesa` / menú | `pedido` / `detalle_pedido` | `factura` / `pago` | Inventario y alertas | `vw_reporte_ventas` |
|---|---|---|---|---|---|---|---|
| Administrador | S/I/U (datos descifrados solo bajo necesidad) | S/I/U | S/I/U | S/I/U | S/I/U | S/I/U | S |
| Cajero | S (campos mínimos) | — | S | S (consulta) | S/I; anulación solo tras autorización de Administrador | — | S (reportes autorizados) |
| Mesero | S (solo datos mínimos) | — | S/U (estado operativo) | S/I/U (solo pedidos propios; sin editar facturados) | — | — | — |
| Cocina | — | — | — | S/U (solo estado de preparación) | — | S solo de recetas necesarias | — |
| Cliente QR | — | — | S solo de menú público | S solo del consumo activo asociado al token QR | — | — | — |

**Aclaración de alcance:** la tabla `rol_permiso` permite configurar permisos, pero las filas y columnas autorizadas se deben aplicar en el backend; no se habilita acceso SQL directo a clientes ni a usuarios de interfaz.

## Verificación operativa

1. Rechazar cualquier inicio de sesión que no valide `password_hash` con bcrypt.
2. Verificar que todas las conexiones se negocien con TLS y que el usuario MySQL no tenga acceso desde redes públicas.
3. Probar cada permiso con una cuenta de prueba de cada rol y confirmar que las operaciones no autorizadas fallan.
4. Probar una restauración mensual y una anulación de factura en ambiente de prueba; conservar los registros sin datos personales en claro.
