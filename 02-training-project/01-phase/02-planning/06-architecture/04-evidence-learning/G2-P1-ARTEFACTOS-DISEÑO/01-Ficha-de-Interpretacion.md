# A1 · Ficha de interpretación y plataforma

## Caso y alcance

SIGRES es una solución cliente-servidor para restaurantes. Incluye plataforma web, aplicación móvil para meseros, módulo de cocina, caja, inventario, mesas, reportes y menú QR de consulta. Los clientes no registran pedidos desde el QR. El SRS contiene 17 requisitos funcionales y 31 no funcionales.

## Requisitos funcionales para el diseño

| ID | Requisito resumido |
|---|---|
| RF-01 | Registrar usuarios. |
| RF-02 | Iniciar sesión. |
| RF-03 | Gestionar roles y permisos. |
| RF-04 | Registrar pedidos y comandas. |
| RF-05 | Enviar pedidos a cocina. |
| RF-06 | Consultar el estado de pedidos. |
| RF-07 | Actualizar el estado de pedidos desde cocina. |
| RF-08 | Administrar productos del menú. |
| RF-09 | Gestionar categorías del menú. |
| RF-10 | Registrar y actualizar inventario. |
| RF-11 | Alertar por existencias mínimas. |
| RF-12 | Generar facturas. |
| RF-13 | Procesar pagos. |
| RF-14 | Gestionar mesas. |
| RF-15 | Generar reportes de ventas. |
| RF-16 | Mostrar el menú digital por QR. |
| RF-17 | Consultar el consumo actual por QR. |

## Requisitos no funcionales y restricciones

| ID | Requisito resumido | Implicación de diseño |
|---|---|---|
| RNF-01 | Autenticación segura y acceso según rol. Contraseñas con bcrypt. | Centralizar autenticación y autorización. |
| RNF-03 | Comunicación cliente-servidor por HTTPS. | Cifrar todas las conexiones públicas. |
| RNF-04 | ORM o consultas parametrizadas. | Acceso a datos con Laravel ORM o consultas preparadas. |
| RNF-07 | Registrar un pedido en menos de 3 minutos desde la app de meseros. | Flujo móvil corto, claro y con pocos pasos. |
| RNF-14 | Carga web inicial máxima de 4 segundos a 10 Mbps. | Carga diferida de módulos y recursos optimizados. |
| RNF-16 | Hasta 20 usuarios activos sin degradación perceptible. | Empezar con una aplicación modular y dimensionamiento acorde al caso. |
| RNF-17 | Avisar a cocina en máximo 5 segundos; WebSocket o polling optimizado. | Usar eventos en tiempo real por WebSocket. |
| RNF-19 | Pedidos y transacciones confirmados no se pierden; transacciones ACID. | Usar transacciones del SGBDR. |
| RNF-20 | Respaldo diario en una ubicación distinta al servidor principal. | Programar copia y restauración probada fuera del nodo de aplicación. |
| RNF-22 | La app muestra el menú almacenado sin conexión; no permite confirmar pedidos offline. | Mantener copia de lectura y bloquear registro sin conexión. |
| RNF-28 | Backend gratuito: Node/Express, Python/Django o PHP/Laravel. | Se selecciona PHP con Laravel. |
| RNF-29 | SGBDR gratuito: MySQL 8+, PostgreSQL 13+ o MariaDB 10.5+. | Se selecciona MySQL 8 o superior. |
| RNF-30 | Despliegue en plataforma de bajo costo o gratuita. | Mantener despliegue compatible con servidor pequeño o VPS accesible. |
| RNF-31 | App móvil compatible con Android 8+; permite React Native. | Se selecciona React Native para Android API 26 o superior. |

## Actores y condiciones

- **Administrador:** usuarios, roles, menú e inventario.
- **Mesero:** pedidos, mesas y consulta de estados.
- **Personal de cocina:** consulta y actualización de comandas.
- **Cajero:** facturación y pagos.
- **Gerente:** reportes y consulta administrativa.
- **Cliente:** consulta menú y consumo desde QR.

El restaurante requiere conexión a internet para sincronización y operación en tiempo real. La app móvil puede mostrar el menú previamente cargado sin conexión, pero no confirmar pedidos. El despliegue debe ser de bajo costo y utilizar herramientas permitidas por el SRS. Las integraciones con domicilios y pasarelas externas quedan fuera del alcance inicial.

## Plataforma seleccionada

| Parte | Selección | Justificación |
|---|---|---|
| Backend | PHP con Laravel | Está permitido por RNF-28. |
| Datos | MySQL 8 o superior | Está permitido por RNF-29; el modelo tiene relaciones entre pedidos, mesas, productos, facturas e inventario; sus transacciones apoyan la integridad ACID de RNF-19. |
| App de meseros | React Native para Android 8 o superior | Está permitido por RNF-31 y cubre el uso móvil de meseros requerido por el SRS. |
| Web y menú QR | Navegadores modernos con HTML5, CSS3 y JavaScript | Coincide con la restricción de tecnologías web estándar del SRS y permite consultar el QR sin instalar una app. |
| Comunicación | API HTTPS con JSON; WebSocket seguro para eventos de cocina | HTTPS responde a RNF-03; WebSocket responde al límite de 5 segundos de RNF-17. |

**Condiciones usadas para decidir:** el stack fue fuertemente influido por mis conocimientos y los del equipo en tecnologias, debido a esto se decidio usar un stack tecnologico conocido por el equipo.

## Entregables de diseño

1. Diagrama de clases y sus patrones.
2. Arquitectura y decisiones frente a RNF.
3. Vista de componentes.
4. Vista de despliegue.
5. Matriz de trazabilidad entre los 17 RF y el diseño.

## Fuente

SRS SIGRES: `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx`.
