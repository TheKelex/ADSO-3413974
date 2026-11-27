# A3 · Patrones de diseño

Los patrones se aplican al diseño de SIGRES y aparecen en [el diagrama de clases](./02-Diagrama-de-Clases.puml).

| Patrón | Problema concreto | Aplicación en SIGRES | Clases | Beneficio |
|---|---|---|---|---|
| Estrategia | El sistema admite efectivo, tarjeta, transferencia y billeteras digitales. Cada método procesa el pago de forma distinta. | `Pago` usa la interfaz `MetodoPago`; la estrategia concreta se selecciona según el método elegido. No se integra una pasarela externa, que está fuera del alcance inicial. | `MetodoPago`, `PagoEfectivo`, `PagoTarjeta`, `PagoTransferencia`, `PagoBilleteraDigital` | Permite agregar o cambiar un método sin cargar todas sus reglas en `Pago`. |
| Fachada | Confirmar un pedido coordina su registro, el cambio de estado de la mesa, el inventario y el aviso a cocina. | `GestionPedidoFacade` ofrece `registrarPedido` y `confirmarPedido`, y coordina las clases del flujo. | `GestionPedidoFacade`, `ServicioInventario`, `NotificadorCocina`, `Pedido`, `Mesa` | Mantiene el flujo de atención en un punto y evita que la interfaz coordine directamente cada subsistema. |

Ambos patrones atienden necesidades descritas en RF-04, RF-05, RF-10, RF-11 y RF-13. El patrón Estrategia no es Singleton.
