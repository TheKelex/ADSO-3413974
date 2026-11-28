# B5 · Evidencia de ejecución

## Entorno y resultado

- Motor real: MySQL Community Server 8.0.46 (imagen oficial `mysql:8.0`).
- Script: `05-B5-Crear-Base-de-Datos.sql`.
- Base creada: `sigres_p2`.
- Resultado de ejecución del DDL: sin errores.
- Verificación en `information_schema`: 16 tablas base, 1 vista y 32 restricciones `CHECK`.
- Comparación automatizada B4↔MySQL: 96 columnas en cada uno; coincidencia exacta de nombres, tipos y dimensiones.

## Pruebas de comportamiento

Se usaron únicamente datos sintéticos en un contenedor temporal local; no contienen información real.

| Prueba | Resultado observado |
|---|---|
| Insertar registros relacionados de usuario, menú, pedido, detalle, factura y método de pago | Correcto; PK y FK aceptadas. |
| Cambiar una factura pendiente a pagada y crear el pago | Correcto. |
| Consultar `vw_reporte_ventas` después del pago | Correcto; devuelve fecha, mesero, producto, cantidad y total vendido. |
| Insertar una mesa con capacidad cero | Rechazo esperado: `Check constraint 'chk_mesa_capacidad' is violated.` |
| Modificar el total de una factura pagada | Rechazo esperado por `trg_factura_paga_inmutable`. |
| Anular la factura pagada con motivo, usuario y fecha, sin cambiar importes | Correcto. |
| Modificar una factura ya anulada | Rechazo esperado por `trg_factura_paga_inmutable`. |

## Resumen de salida del motor

```text
mysql_version=8.0.46
base_tables=16
check_constraints=32
views=1
dictionary_columns=96
DDL: exit code 0 (sin errores)
```

La salida completa de estas consultas y pruebas está en [07-Ejecucion-MySQL.log](./07-Ejecucion-MySQL.log).

La prueba confirma la sintaxis y restricciones en MySQL 8.0.46. No sustituye pruebas de concurrencia, permisos de aplicación ni una restauración de respaldo; estas se deben validar en el entorno de despliegue.
