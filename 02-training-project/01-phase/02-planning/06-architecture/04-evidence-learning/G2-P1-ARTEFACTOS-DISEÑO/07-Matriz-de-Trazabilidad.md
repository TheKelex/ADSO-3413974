# A7 · Matriz de trazabilidad análisis → diseño

| Requisito | Clase(s) | Componente(s) | Decisión |
|---|---|---|---|
| RF-01 Registrar usuarios | Usuario, Rol | Autenticación y autorización; API Laravel; Acceso a datos | Autenticación validada en servidor. |
| RF-02 Iniciar sesión | Usuario, Rol | Autenticación y autorización; API Laravel; App móvil y Web | Contraseña con bcrypt y permisos según rol. |
| RF-03 Gestionar roles y permisos | Usuario, Rol, Permiso | Autenticación y autorización; Web | Reglas de acceso centralizadas en servidor. |
| RF-04 Registrar pedido | Pedido, DetallePedido, Mesa, ProductoMenu, GestionPedidoFacade | Pedidos y mesas; App móvil; API Laravel | La fachada coordina el flujo de registro. |
| RF-05 Enviar pedido a cocina | Pedido, NotificadorCocina | Pedidos y mesas; Notificaciones de cocina | Evento en tiempo real por WebSocket. |
| RF-06 Consultar estado del pedido | Pedido, EstadoPedido | Pedidos y mesas; App móvil; Web | Estado consultado desde el servicio de pedidos. |
| RF-07 Actualizar estado del pedido | Pedido, EstadoPedido | Pedidos y mesas; Notificaciones de cocina | La actualización se publica a los clientes autorizados. |
| RF-08 Administrar productos del menú | ProductoMenu | Menú e inventario; Web | Gestión desde el servicio de menú. |
| RF-09 Gestionar categorías del menú | CategoriaMenu | Menú e inventario; Web | Categorías administradas sin cambiar código. |
| RF-10 Registrar y actualizar inventario | Insumo, MovimientoInventario, ServicioInventario | Menú e inventario; Acceso a datos; Web | Movimientos quedan asociados al insumo. |
| RF-11 Alertar por existencias mínimas | Insumo, AlertaInventario, ServicioInventario | Menú e inventario; Web | Se compara existencia con mínimo configurado. |
| RF-12 Generar factura | Factura, Pedido, DetallePedido | Caja y reportes; Pedidos y mesas; Web | La factura se genera para el pedido activo. |
| RF-13 Procesar pago | Pago, MetodoPago, PagoEfectivo, PagoTarjeta, PagoTransferencia, PagoBilleteraDigital | Caja y reportes; Web | Estrategia de pago según método; sin pasarela externa. |
| RF-14 Gestionar mesas | Mesa, Pedido | Pedidos y mesas; App móvil; Web | La mesa refleja su estado operativo. |
| RF-15 Generar reportes de ventas | ReporteVentas, Factura, Pago | Caja y reportes; Web | Reportes consultan operaciones registradas. |
| RF-16 Ver menú digital por QR | MenuQRService, CategoriaMenu, ProductoMenu | Menú e inventario; Web y menú QR | Menú público de solo lectura y responsive. |
| RF-17 Consultar consumo actual por QR | MenuQRService, Pedido, Mesa, Factura | Menú e inventario; Pedidos y mesas; Web y menú QR | El QR consulta consumo asociado a una mesa activa. |

**Cobertura funcional:** 17/17 RF mapeados a clases y componentes.  
Las relaciones producto-insumo y los pagos parciales no se detallan porque el SRS no los define.
