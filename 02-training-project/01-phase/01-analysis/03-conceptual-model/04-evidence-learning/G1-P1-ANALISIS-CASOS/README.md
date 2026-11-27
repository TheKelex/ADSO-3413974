# README de entrega


## 1. Datos de la entrega

- **Producto:** P1 plan de análisis y casos de uso
- **Usuario de GitHub del autor:** TheKelex
- **Caso de estudio:** Sistema Integral de Gestión para Restaurantes (SIGRES)
- **Trabajo en equipo del año:** Sí; el trabajo de SIGRES

## 2. Qué reutilicé

| Artefacto | Ruta original en mi fork (semana) | ¿Lo hice solo o en equipo? |
|---|---|---|
| SRS SIGRES | `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx` | En equipo |
| Normativa del Proyecto SIGRES, como contexto del problema | `03-execute/14-week/02-optional-activity/Normativa del Proyecto-1.docx` | En equipo |

## 3. Qué ajusté

| # | Artefacto | Qué cambié y por qué | Commit |
|---|---|---|---|
| 1 | Casos de uso | Organicé los objetivos de usuario en 17 casos G1 para poder relacionarlos explícitamente con los 17 requisitos funcionales del SRS. | [1add8c7](https://github.com/TheKelex/ADSO-3413974/commit/1add8c7573ba6dde62a9337bf9e92528d14736d6) |
| 2 | Trazabilidad | Preparé una fila por RF y señalé las referencias del SRS que parecen no coincidir con los nombres de las fichas detalladas. | [1add8c7](https://github.com/TheKelex/ADSO-3413974/commit/1add8c7573ba6dde62a9337bf9e92528d14736d6) |
| 3 | Relaciones UML | Propuse `include` entre registrar y enviar pedido a cocina, y `extend` entre consultar consumo y ver menú QR, con su condición. | [365d306](https://github.com/TheKelex/ADSO-3413974/commit/365d30682efe8710cff135c2a64b054f28037163) |

## 4. Qué construí nuevo

- Caracterización de SIGRES y comparación de RUP, Scrum, XP y RAD.
- Selección justificada de Scrum adaptado para organizar este análisis.
- Plan de actividades para analizar SIGRES, con responsable, fecha y resultado esperado.
- Matriz requisito funcional → caso de uso.
- Código editable del diagrama UML en PlantUML.

## 5. Artefactos entregados

**Producto 1**

- [x] Caracterización de la solución: [01-Caracterizacion-y-Metodologia.md](./01-Caracterizacion-y-Metodologia.md)
- [x] Comparación de metodologías y selección justificada: [01-Caracterizacion-y-Metodologia.md](./01-Caracterizacion-y-Metodologia.md)
- [x] Plan de actividades, cronograma y entregables: [02-Plan-de-Analisis.md](./02-Plan-de-Analisis.md)
- [x] Diagrama principal de pedidos y atención (fuente editable): [01-Pedidos-y-Atencion.puml](./Diagrama%20de%20casos%20de%20uso%20UML/01-Pedidos-y-Atencion.puml)
- [x] Diagrama de apoyo de administración (fuente editable): [02-Administracion-del-Restaurante.puml](./Diagrama%20de%20casos%20de%20uso%20UML/02-Administracion-del-Restaurante.puml)
- [x] Imágenes de los diagramas: [01-Pedidos-y-Atencion.png](./Diagrama%20de%20casos%20de%20uso%20UML/01-Pedidos-y-Atencion.png) y [02-Administracion-del-Restaurante.png](./Diagrama%20de%20casos%20de%20uso%20UML/02-Administracion-del-Restaurante.png)
- [x] Matriz requisito → caso de uso: [03-Matriz-Requisitos-Casos-de-Uso.md](./03-Matriz-Requisitos-Casos-de-Uso.md)

## 6. Declaración de autoría

- [x] Declaro que el ajuste y lo construido nuevo en esta carpeta es **mi trabajo individual**.
- [x] Declaro que lo que tomé de mi equipo o de otra fuente está identificado en la sección 2.
- [x] Entiendo que el sábado 28 de noviembre de 2026 el instructor me hará **3 preguntas orales** sobre mis productos (una del producto 1, una del producto 2 y una tercera a su elección) y que, si respondo bien menos de 2, el producto sobre el que fallé queda con tope de 50.

Firma (usuario de GitHub y fecha): TheKelex - 26/11/2026

## 7. Uso de inteligencia artificial

| Herramienta | Para qué la usé | Qué verifiqué o corregí yo |
|---|---|---|
| Copilot | Preparar un borrador de la estructura, la matriz y el código PlantUML a partir de los archivos del proyecto. | Revisé el borrador, adapté el contenido a SIGRES y corregí los diagramas. |

## 8. Autocontrol de la entrega

- [x] Hice commits propios en al menos 2 días distintos entre la publicación de la guía (sábado 21 de noviembre de 2026) y la entrega.
- [x] No subí el marco documental del instructor, plantillas sin diligenciar ni archivos de otras personas.
- [x] Subir este README actualizado y comprobar en GitHub que la carpeta está publicada.