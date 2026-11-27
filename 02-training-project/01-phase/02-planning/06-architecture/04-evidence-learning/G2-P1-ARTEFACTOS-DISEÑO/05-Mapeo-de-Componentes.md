# A5 · Mapeo de componentes

El diagrama editable está en [05-Vista-de-Componentes.puml](./05-Vista-de-Componentes.puml). Los contratos nombrados son interfaces lógicas; su ruta y formato final se confirman durante implementación.

| Componente | Interfaz provista | Interfaces requeridas | Clases que agrupa | Requisitos |
|---|---|---|---|---|
| App móvil React Native | `IMeseroUI` | `IAuthAPI`, `IOrdersAPI`, `IMenuAPI` | Cliente móvil | RF-02, RF-04, RF-06, RF-14; RNF-07, RNF-22, RNF-31 |
| Web administrativa y menú QR | `IWebUI`, `IQRMenuUI` | `IAuthAPI`, `IOrdersAPI`, `IMenuAPI`, `ICashReportsAPI` | Cliente web | RF-01, RF-02, RF-03, RF-08, RF-09, RF-12, RF-13, RF-15, RF-16, RF-17 |
| API Laravel | `IAuthAPI`, `IOrdersAPI`, `IMenuAPI`, `ICashReportsAPI` | `IAuthService`, `IOrderService`, `IMenuStockService`, `ICashReportService` | Controladores API | RF-01 a RF-17; RNF-03, RNF-28 |
| Autenticación y autorización | `IAuthService` | `IRepositories` | `Usuario`, `Rol`, `Permiso` | RF-01, RF-02, RF-03; RNF-01 |
| Pedidos y mesas | `IOrderService` | `IMenuStockService`, `IKitchenEvents`, `IRepositories` | `Pedido`, `DetallePedido`, `Mesa`, `GestionPedidoFacade` | RF-04, RF-05, RF-06, RF-07, RF-14; RNF-17, RNF-19 |
| Menú e inventario | `IMenuStockService` | `IRepositories` | `ProductoMenu`, `CategoriaMenu`, `Insumo`, `MovimientoInventario`, `AlertaInventario`, `ServicioInventario`, `MenuQRService` | RF-08, RF-09, RF-10, RF-11, RF-16, RF-17 |
| Caja y reportes | `ICashReportService` | `IOrderService`, `IRepositories` | `Factura`, `Pago`, `MetodoPago`, estrategias de pago, `ReporteVentas` | RF-12, RF-13, RF-15 |
| Notificaciones de cocina | `IKitchenEvents` | `IRepositories` | `NotificadorCocina` | RF-05, RF-06, RF-07; RNF-17 |
| Acceso a datos | `IRepositories` | Conexión MySQL | Repositorios de entidades | RF-01 a RF-17; RNF-04, RNF-19 |

Las flechas del diagrama indican qué componente requiere un contrato de otro. Los clientes no acceden al componente MySQL.
