# G1 · Producto 1 · Caracterización y selección de metodología

## 1. Caracterización de la solucion

| Aspecto | Descripción | Estado |
|---|---|---|
| Tipo de proyecto | Sistema para restaurantes. Incluye una aplicación web, una aplicación móvil y una base de datos. Tiene módulos de pedidos, cocina, menú, inventario, caja, mesas y reportes. | Descrito en el SRS. |
| Tamaño del equipo | El SRS registra cuatro integrantes. | Dato del SRS. |
| Estabilidad de  requisitos | El SRS tiene 17 requisitos funcionales (RF) y 31 no funcionales (RNF). Algunos datos y referencias deben confirmarse. | Estabilidad media; falta validar con el instructor o cliente. |
| Disponibilidad del cliente | El SRS describe a los usuarios, pero no indica fechas para revisar ni aprobación del cliente. | Pendiente de acordar. |

### Alcance y restricciones

- Los RF-01 a RF-17 cubren ocho módulos.
- Los RNF incluyen seguridad, facilidad de uso, rendimiento, disponibilidad, mantenimiento y uso en móviles.
- El modelo usa lo que está descrito en el SRS. Las funciones futuras, como una pasarela de pago, no se presentan como parte confirmada del sistema.
- La «Normativa del Proyecto» sirve como contexto. No confirma una aprobación formal ni reemplaza las normas legales.

## 2. Comparación de metodologías

La tabla compara qué tan bien se adapta cada metodología al caso. No significa que el proyecto ya haya usado alguna de ellas.

| Necesidad de SIGRES | RUP | Scrum | Programación extrema (XP) | RAD |
|---|---|---|---|---|
| Organizar el SRS y los módulos | Alta: organiza el trabajo y sus documentos. | Alta: permite ordenar los módulos en una lista de trabajo. | Media: permite trabajar por historias, pero se enfoca más en programar y probar. | Alta: ayuda a crear prototipos rápidos. |
| Mantener la relación entre 17 RF y 31 RNF | Alta: da importancia a los requisitos y documentos. | Media-alta: requiere mantener la matriz y revisar los requisitos. | Media: se debe agregar trabajo para mantener la documentación. | Media: se debe cuidar que los prototipos no dejen requisitos sin registrar. |
| Revisar riesgos de seguridad, caja e inventario | Alta: permite revisar los riesgos desde el inicio. | Media-alta: permite priorizarlos y revisarlos en cada ciclo. | Media: requiere complementar las pruebas con revisión de riesgos. | Media: los prototipos ayudan con las pantallas, pero no cubren todos los RNF. |
| Equipo y alcance de esta entrega | Media: puede ser mucho proceso para una entrega individual corta. | Alta: se puede adaptar a una lista priorizada y entregas cortas. | Media-alta: permite recibir comentarios, pero algunas prácticas requieren un equipo. | Media-alta: permite probar pantallas pronto, pero también se debe cuidar la trazabilidad. |

## 3. Metodología seleccionada: Scrum adaptado

Se trabajara con Scrum adaptado para organizar el análisis. SIGRES tiene varios módulos y requisitos que se pueden priorizar y revisar por partes.

Para esta actividad se usarán una lista de trabajo, prioridades y revisiones.

RUP es otra opción si se necesita más documentación. RAD ayuda a probar pantallas, pero no reemplaza la matriz de requisitos. XP se enfoca más en programar y probar, mientras que esta entrega trata del análisis.

## 4. Decisiones UML para el diagrama

- `Registrar pedido` incluye `Enviar pedido a cocina`. Según el SRS, la comanda se envía al confirmar el pedido.
- `Consultar consumo actual` extiende `Ver menú digital por QR`. Es opcional y aplica si hay un pedido activo para la mesa.
- Se separan registrar, consultar y actualizar el estado del pedido para relacionar mejor los RF-05, RF-06 y RF-07. En el SRS aparecen agrupados en un caso de uso de cocina.
- No se incluye un proveedor de pagos. El SRS describe la pasarela como una integración futura.
