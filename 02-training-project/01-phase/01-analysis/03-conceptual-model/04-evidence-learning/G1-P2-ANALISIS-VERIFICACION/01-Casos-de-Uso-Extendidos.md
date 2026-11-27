# Casos de uso extendidos

Casos seleccionados por su prioridad alta y marca de críticos en el SRS.

## G1-UC-04 · Registrar pedido

| Campo | Contenido |
|---|---|
| Requisitos que atiende | RF-04. Incluye el envío automático a cocina descrito en RF-05. |
| Actor primario / secundarios | Mesero / Personal de cocina, que recibe la comanda |
| Objetivo del actor | Registrar el pedido de una mesa y enviarlo a cocina. |
| Precondiciones | El mesero inició sesión. |
| Disparador | El mesero selecciona la opción para registrar un pedido. |
| Flujo principal | 1. El mesero selecciona la mesa y el número de comensales.<br>2. Selecciona productos y cantidades.<br>3. Agrega notas especiales, si las necesita.<br>4. Confirma los datos.<br>5. El sistema valida los campos y la disponibilidad.<br>6. El sistema registra el pedido con un identificador único.<br>7. El sistema envía la comanda a cocina.<br>8. El sistema marca la mesa como ocupada.<br>9. El sistema confirma el registro. |
| Flujos alternos | 3a. Si no hay notas especiales, el mesero continúa sin agregarlas. |
| Excepciones | 5a. Si faltan datos o un producto no está disponible, el sistema informa el problema y no registra el pedido.<br>5b. Si falla la conexión, el sistema informa que no pudo registrar el pedido. |
| Postcondiciones | El pedido queda registrado y visible para cocina; la mesa queda ocupada. |
| Reglas de negocio | El pedido debe tener los datos requeridos y productos disponibles. |
| Pantalla del prototipo | Pendiente de vincular al mockup que se incorporará. |

## G1-UC-10 · Gestionar inventario

| Campo | Contenido |
|---|---|
| Requisitos que atiende | RF-10 y la alerta descrita en RF-11. |
| Actor primario / secundarios | Administrador |
| Objetivo del actor | Registrar o actualizar insumos y su existencia. |
| Precondiciones | El usuario tiene permiso para modificar el inventario. |
| Disparador | El administrador selecciona la opción de inventario. |
| Flujo principal | 1. El administrador elige registrar o actualizar un insumo.<br>2. Ingresa nombre, categoría, cantidad, unidad y mínimo de alerta.<br>3. Ingresa el costo unitario y, si aplica, el proveedor.<br>4. El sistema valida los datos.<br>5. El sistema registra o actualiza el insumo.<br>6. El sistema actualiza la existencia.<br>7. Si la cantidad está bajo el mínimo, genera una alerta.<br>8. El sistema confirma el resultado. |
| Flujos alternos | 7a. Si la cantidad está en el mínimo o por encima, el sistema guarda el cambio sin generar alerta. |
| Excepciones | 4a. Si la cantidad no es válida o falta la unidad, el sistema informa el error.<br>4b. Si el insumo ya existe al intentar registrarlo, el sistema lo informa.<br>5a. Si ocurre un error al guardar, el sistema informa que no pudo actualizar el inventario. |
| Postcondiciones | El insumo queda registrado o actualizado; si está bajo el mínimo, queda una alerta generada. |
| Reglas de negocio | Solo usuarios autorizados pueden modificar el inventario. El stock también se actualiza al procesar pedidos. La alerta se genera cuando la existencia es inferior al mínimo. |
| Pantalla del prototipo | Pendiente de vincular al mockup que se incorporará. |

## G1-UC-13 · Procesar pago

| Campo | Contenido |
|---|---|
| Requisitos que atiende | RF-13. |
| Actor primario / secundarios | Cajero |
| Objetivo del actor | Registrar el pago de una factura y entregar su recibo. |
| Precondiciones | La factura está generada y pendiente de pago. |
| Disparador | El cajero selecciona la factura para pagar. |
| Flujo principal | 1. El cajero selecciona el método de pago.<br>2. Si el pago es en efectivo, ingresa el monto recibido.<br>3. Ingresa la propina, si aplica.<br>4. Confirma el pago.<br>5. El sistema valida el método y el monto.<br>6. El sistema registra el pago.<br>7. Si es efectivo, calcula el vuelto.<br>8. El sistema genera el recibo y registra la transacción.<br>9. El sistema confirma el pago. |
| Flujos alternos | 2a. Si el pago no es en efectivo, el cajero continúa sin ingresar un monto recibido. |
| Excepciones | 5a. Si no se selecciona un método o el monto en efectivo es insuficiente, el sistema informa el error.<br>5b. Si la factura ya está pagada, el sistema no procesa otro pago.<br>6a. Si ocurre un error al guardar, el sistema informa que no pudo registrar el pago. |
| Postcondiciones | El pago queda registrado y el recibo generado. |
| Reglas de negocio | Métodos: efectivo, tarjeta, transferencia, Nequi o Daviplata. El vuelto aplica al efectivo. Solo usuarios autorizados pueden procesar pagos. |
| Pantalla del prototipo | Pendiente de vincular al mockup que se incorporará. |
