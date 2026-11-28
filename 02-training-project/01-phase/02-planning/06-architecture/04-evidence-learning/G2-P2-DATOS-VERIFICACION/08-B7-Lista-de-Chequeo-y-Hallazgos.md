# B7 · Lista de chequeo y registro de hallazgos

## Lista de chequeo diligenciada

Se evaluaron los artefactos publicados de G2-P1 y los archivos locales de G2-P2. `Cumple` indica evidencia visible en el archivo citado; `Parcial` indica que el hallazgo se registra abajo y la corrección está en B1/B2/B4/B5. Las severidades son Crítica, Mayor y Menor.

| ID | Producto / artefacto | Criterio de aceptación medible | Severidad | Estado | Evidencia |
|---|---|---|---|---|---|
| LC-01 | P1 A1 | La ficha identifica al menos 8 RF y 3 RNF con ID y selecciona plataforma. | Mayor | Cumple | P1 `01-Ficha-de-Interpretacion.md`: 17 RF, RNF y selección Laravel/MySQL/React Native. |
| LC-02 | P1 A2 | El diagrama contiene al menos 8 clases y atributos tipados y operaciones. | Mayor | Cumple | P1 `02-Diagrama-de-Clases.puml`. |
| LC-03 | P1 A2 | Todas las asociaciones muestran multiplicidad en ambos extremos. | Mayor | Cumple | P1 `02-Diagrama-de-Clases.puml`. |
| LC-04 | P1 A2 | Se observa al menos una composición y una interfaz con dos o más implementaciones. | Mayor | Cumple | `Pedido *-- DetallePedido`; `MetodoPago` implementada por cuatro clases, en P1 `02-Diagrama-de-Clases.puml`. |
| LC-05 | P1 A3 | Se aplican y justifican al menos dos patrones distintos. | Mayor | Cumple | P1 `03-Patrones-de-Diseno.md`: Estrategia y Fachada. |
| LC-06 | P1 A4 | El estilo incluye decisiones para al menos 3 RNF y una alternativa descartada. | Mayor | Cumple | P1 `04-Decisiones-de-Arquitectura.md`. |
| LC-07 | P1 A5 | Se documentan al menos 5 componentes con interfaces provistas/requeridas. | Mayor | Cumple | P1 `05-Mapeo-de-Componentes.md`: 9 componentes. |
| LC-08 | P1 A6 | La vista contiene al menos 3 nodos y sus conexiones indican protocolo/puerto. | Mayor | Cumple | P1 `06-Vista-de-Despliegue.puml` y `06-Vista-de-Despliegue.jpeg`. |
| LC-09 | P1 A7 | La matriz incluye los 17 RF del SRS. | Crítica | Cumple | P1 `07-Matriz-de-Trazabilidad.md`: cobertura declarada 17/17. |
| LC-10 | P1 diagramas | Cada vista principal tiene fuente editable y exportación de imagen. | Menor | Cumple | Fuentes `.puml` y exportaciones `.jpeg` listadas en el README de P1. |
| LC-11 | P2 B1 | El modelo conceptual tiene al menos 6 entidades con identificador y atributos. | Mayor | Cumple | `01-B1-Modelo-Conceptual.md`: 13 entidades. |
| LC-12 | P2 B1 | Cada relación tiene cardinalidad mínima/máxima y verbo; hay N:M y 1:N. | Crítica | Cumple | Tabla “Cardinalidades mínimas y máximas” y diagrama Mermaid de B1. |
| LC-13 | P2 B1 | La selección de MySQL cita condiciones del SRS y justifica integridad transaccional. | Mayor | Cumple | `01-B1-Modelo-Conceptual.md`, alcance y tipo de base de datos. |
| LC-14 | P2 B2 | El modelo lógico contiene al menos 8 tablas, cada una con PK/FK, y resuelve las N:M. | Crítica | Cumple | `02-B2-Modelo-Logico.md`: 16 tablas; `rol_permiso`, `detalle_pedido` y `producto_insumo`. |
| LC-15 | P2 B3 | Se normalizan 2 relaciones propias de 0FN a 3FN con DF escritas. | Mayor | Cumple | `03-B3-Normalizacion.md`: `detalle_pedido` y `factura`. |
| LC-16 | P2 B3 | La evidencia incluye llave compuesta/dependencia parcial y dependencia transitiva. | Mayor | Cumple | B3, casos 1 y 2, pasos 2FN y 3FN. |
| LC-17 | P2 B4 | El diccionario cubre todas las columnas de las 16 tablas con tipo, tamaño, restricciones y dominio. | Crítica | Cumple | `04-B4-Diccionario-de-Datos.md`, tablas y columnas. |
| LC-18 | P2 B5 | El script crea el 100 % de tablas en MySQL real sin errores. | Crítica | Cumple | `05-B5-Crear-Base-de-Datos.sql`; log resumido en `07-Evidencia-Ejecucion-MySQL.md`. |
| LC-19 | P2 B5 | El script incluye al menos 3 `CHECK` y acciones `ON DELETE` en al menos 3 FK. | Crítica | Cumple | Ejecución verificada: 32 `CHECK`; FKs explícitas en B5. |
| LC-20 | P2 B5 | El motor rechaza capacidad cero, bloquea edición de factura pagada y permite anulación justificada. | Crítica | Cumple | Pruebas registradas en `07-Evidencia-Ejecucion-MySQL.md`. |
| LC-21 | P2 B6 | Hay al menos 6 políticas concretas, al menos 2 por integridad, confidencialidad y disponibilidad. | Mayor | Cumple | `06-B6-Politicas-de-Seguridad.md`: 12 políticas. |
| LC-22 | P2 B6 | La matriz cubre al menos 3 roles y limita el acceso a tablas/filas. | Mayor | Cumple | B6, matriz de Administrador, Cajero, Mesero, Cocina y Cliente QR. |
| LC-23 | P2 B6 / B4 / B5 | Se define bcrypt con sal, TLS y protección identificable para datos personales. | Crítica | Cumple | B6 C-02 a C-05; B4/B5 columnas cifradas y HMAC único. |
| LC-24 | P2 B6 | Se indica frecuencia de respaldo y prueba de restauración verificable. | Mayor | Cumple | B6 D-01/D-02: respaldo cada 24 h y restauración mensual. |
| LC-25 | P2 B7 | Hay al menos 20 criterios medibles con severidad, estado y evidencia para P1 y P2. | Mayor | Cumple | Esta tabla contiene 27 criterios, LC-01 a LC-27. |
| LC-26 | README | El README sigue la plantilla, identifica orígenes, declara IA y no inventa publicación o segundo commit. | Mayor | Cumple | `README.md`; historial de dos días señalado como pendiente hasta el commit real. |
| LC-27 | P1 A2 / P2 B1-B5 | El modelo de inventario relaciona cada producto de menú con los insumos y cantidades requeridos para RF-10. | Crítica | Parcial en P1; corregido en P2 | P1 `02-Diagrama-de-Clases.puml` no incluye ProductoMenu-Insumo; P2 añade `producto_insumo` en B1/B2/B4/B5. |

## Registro de hallazgos y ajustes

| ID | Hallazgo / evidencia inicial | Severidad | Ajuste realizado | Estado |
|---|---|---|---|---|
| H-01 | El diagrama de clases P1 no incluye la relación ProductoMenu-Insumo, aunque RF-10 exige descontar existencias por pedido confirmado. | Crítica | Se añadió receta N:M `producto_insumo` a B1/B2, con cantidad por unidad, y se documentó su uso transaccional en B5/B6. | Corregido en P2; confirmar cantidades de receta con el restaurante. |
| H-02 | La factura debe poder anularse con justificación y responsable, pero esos datos no estaban completos en el diseño P1. | Crítica | Se añadieron `motivo_anulacion`, `anulada_por_usuario_id` y `anulada_en`, restricciones de estado y trigger de inmutabilidad. | Corregido y probado en MySQL. |
| H-03 | El MER de G1 presenta Usuario-Rol N:M; RF-01 y P1 asignan un rol por usuario. | Mayor | B1/B2 implementan Usuario N:1 Rol. La divergencia queda declarada como decisión pendiente de confirmar con el equipo. | Ajustado; validar supuesto. |
| H-04 | El primer borrador de esquema guardaba identificación y teléfono como texto ordinario, sin protección de almacenamiento explícita. | Mayor | Se sustituyeron por campos cifrados AES-256-GCM y un HMAC separado para unicidad; las claves quedan fuera de MySQL. | Corregido en B2/B4/B5/B6. |
