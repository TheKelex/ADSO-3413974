# Registro de ajustes

| # | Hallazgo | Antes | Después | Segunda verificación |
|---|---|---|---|---|
| 1 | No quedaba claro el rol de cocina en RF-05. | Cocina aparecía como actor sin indicar qué hacía. | Se aclara que recibe la comanda enviada por el sistema. | Sí: coincide con el flujo de RF-04 y RF-05. |
| 2 | La plantilla de inventario no trazaba la alerta. | Indicaba solo RF-10. | Incluye RF-11 para la alerta por existencia mínima. | Sí: RF-11 también aparece en la matriz. |
| 3 | Faltaba reflejar el descuento automático de stock. | No se mencionaba el cambio de inventario al procesar pedidos. | Se agrega esa regla y el vínculo con movimientos de inventario. | Sí: aparece en la plantilla y el modelo de dominio. |
| 4 | El diagrama de inventario no mostraba el error al guardar. | El flujo terminaba sin una salida para ese error. | Se agrega la decisión y el mensaje de error. | Sí: coincide con la excepción de la plantilla. |
| 5 | La plantilla de pago no listaba los métodos del SRS. | Se nombraba el método sin detallar opciones. | Se incluyen efectivo, tarjeta, transferencia, Nequi y Daviplata. | Sí: coincide con RF-13 del SRS. |
| 6 | El SRS no define cuántos roles puede tener un usuario. | El primer modelo asumía un solo rol. | Se representa la relación como muchos a muchos y se deja pendiente su confirmación. | Sí: la relación coincide en el dominio y el modelo ER. |
