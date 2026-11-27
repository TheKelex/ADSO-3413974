# A4 · Arquitectura y decisiones

## Estilo

Arquitectura cliente-servidor por capas, implementada como aplicación modular en PHP/Laravel:

1. **Presentación:** app React Native, web administrativa y menú QR.
2. **Aplicación y negocio:** API, autenticación, pedidos, menú, caja, inventario, reportes y eventos.
3. **Datos:** MySQL 8 o superior, accedido solo desde los servicios del servidor.

La app móvil, la plataforma web y el menú QR no acceden directamente a MySQL. La vista está en [04-Arquitectura.puml](./04-Arquitectura.puml).

## Decisiones frente a requisitos no funcionales

| RNF | Decisión de diseño | Cómo responde |
|---|---|---|
| RNF-01 | Autenticación en servidor, contraseñas con bcrypt y autorización por rol en cada operación. | Impide depender solo de controles visuales del cliente. |
| RNF-03 | HTTPS para clientes y WSS para notificaciones. | Protege credenciales y datos durante el tránsito. |
| RNF-04 | Laravel ORM o consultas parametrizadas en el acceso a datos. | Evita formar SQL concatenando entradas del usuario. |
| RNF-07 | Flujo móvil de pedido en pasos breves y con selección de productos y cantidades. | Apoya la meta de completar un pedido en menos de 3 minutos; debe validarse con usuarios. |
| RNF-14 | Cargar primero funciones esenciales y diferir módulos secundarios; optimizar recursos estáticos. | Apoya la carga web máxima de 4 segundos. |
| RNF-16 | Aplicación modular en lugar de servicios distribuidos; dimensionar instancia para 20 usuarios y medir con prueba de carga. | Mantiene bajo el costo y corresponde al volumen indicado. |
| RNF-17 | Publicar cambios de pedidos a cocina mediante WebSocket sobre TLS. | Apoya la notificación en máximo 5 segundos. |
| RNF-19 | Usar transacciones ACID al confirmar pedidos, facturas y pagos. | Evita guardar parcialmente una operación confirmada. |
| RNF-20 | Programar respaldo diario de MySQL en almacenamiento distinto al servidor y probar restauración. | Reduce el riesgo de perder los datos ante falla del servidor. |
| RNF-22 | Guardar una copia de lectura del menú en la app; bloquear el envío y confirmación de pedidos sin conexión. | Cumple el alcance offline del SRS sin simular sincronización de pedidos. |
| RNF-28 / RNF-29 | PHP/Laravel y MySQL 8+. | Usa tecnologías gratuitas permitidas expresamente por el SRS. |
| RNF-30 / RNF-31 | Desplegar en VPS o plataforma accesible; distribuir app para Android 8+. | Se ajusta al despliegue de bajo costo y a los dispositivos mínimos. |

## Comunicación

- Cliente a API: HTTPS, puerto 443, datos JSON.
- Cliente a eventos en tiempo real: WSS, puerto 443.
- Laravel a MySQL: TCP 3306 en red privada.
- Agente de respaldos a MySQL: MySQL/TCP 3306 en red privada local.
- Agente de respaldos a almacenamiento separado: HTTPS/TLS, puerto 443. El proveedor se define en el despliegue.

## Alternativa descartada

Se descarta iniciar con microservicios. Para 20 usuarios activos y un equipo con presupuesto limitado, aumenta despliegue y operación sin una necesidad indicada por el SRS. La aplicación modular mantiene separadas las responsabilidades y permite dividir servicios si el volumen futuro lo exige.
