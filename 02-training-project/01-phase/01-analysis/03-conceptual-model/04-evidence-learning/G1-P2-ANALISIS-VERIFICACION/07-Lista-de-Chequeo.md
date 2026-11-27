# Lista de chequeo · G1-P2 SIGRES

| # | Criterio verificable | Evidencia | ¿Cumple? | Hallazgo |
|---|---|---|---|---|
| 1 | Hay dos o más diagramas de actividades. | `02-Diagramas-de-Actividad/` | Sí | — |
| 2 | Cada diagrama tiene inicio, fin y carriles por actor. | Diagramas 01 y 02 | Sí | — |
| 3 | Las decisiones tienen guardas para todas sus salidas. | Diagramas 01 y 02 | Sí | — |
| 4 | Las ramas alternas y excepciones se describen en su plantilla. | Diagramas 01 y 02; casos extendidos | Sí | — |
| 5 | Hay tres o más plantillas extendidas. | `01-Casos-de-Uso-Extendidos.md` | Sí | — |
| 6 | Cada flujo principal tiene seis o más pasos. | `01-Casos-de-Uso-Extendidos.md` | Sí | — |
| 7 | Cada plantilla tiene un flujo alterno. | `01-Casos-de-Uso-Extendidos.md` | Sí | — |
| 8 | Cada plantilla tiene al menos una excepción. | `01-Casos-de-Uso-Extendidos.md` | Sí | — |
| 9 | Cada plantilla define precondiciones y postcondiciones. | `01-Casos-de-Uso-Extendidos.md` | Sí | — |
| 10 | El dominio tiene ocho o más clases conceptuales. | `03-Modelo-de-Dominio.puml` | Sí | — |
| 11 | Cinco o más clases tienen al menos dos atributos de negocio. | `03-Modelo-de-Dominio.puml` | Sí | — |
| 12 | El dominio no incluye métodos ni claves técnicas. | `03-Modelo-de-Dominio.puml` | Sí | — |
| 13 | El dominio incluye una clase asociativa. | `DetallePedido` en el modelo de dominio | Sí | — |
| 14 | El modelo ER tiene seis o más entidades con identificadores. | `04-Modelo-Entidad-Relacion.puml` | Sí | — |
| 15 | La relación pedido-producto se resuelve con una entidad asociativa con atributos. | `DETALLE_PEDIDO` en el modelo ER | Sí | — |
| 16 | El modelo ER identifica una entidad débil o justifica una relación recursiva. | Nota de `DETALLE_PEDIDO` en el modelo ER | Sí | — |
| 17 | La notación indica cardinalidad y participación. | Leyenda y relaciones del modelo ER | Sí | — |
| 18 | Al menos 90 % de los RF se traza a caso, clase y entidad. | `06-Matriz-de-Trazabilidad.md` | Sí | 17 de 17 RF; cobertura preliminar. |
| 19 | Cada pantalla crítica está rotulada con su caso de uso. | `05-MockUp/` | Sí | Las capturas llevan G1-UC-04, G1-UC-10 y G1-UC-13 en el nombre. |
| 20 | Hay cinco o más ajustes con segunda verificación. | `09-Registro-de-Ajustes.md` | Sí | — |
| 21 | El informe incluye requisitos, modelos, verificación y ajustes. | `08-Informe-de-Analisis.md` | Sí | — |
| 22 | El README tiene los enlaces y datos de entrega completos. | `README.md` y `05-MockUp/README.md` | Pendiente | Completar commits después de publicar. |
