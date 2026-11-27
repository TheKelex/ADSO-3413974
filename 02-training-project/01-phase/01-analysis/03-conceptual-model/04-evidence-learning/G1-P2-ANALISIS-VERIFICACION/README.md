# README de entrega


## 1. Datos de la entrega

- **Producto:** P2 modelo de análisis y verificación
- **Usuario de GitHub del autor:** TheKelex
- **Caso de estudio:** Sistema Integral de Gestión para Restaurantes (SIGRES)
- **Trabajo en equipo del año:** Sí; el trabajo de SIGRES

## 2. Qué reutilicé

| Artefacto | Ruta original en mi fork (semana) | ¿Lo hice solo o en equipo? |
|---|---|---|
| SRS SIGRES | `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx` | En equipo |
| Casos de uso y matriz del Producto 1 | `02-training-project/01-phase/01-analysis/03-conceptual-model/04-evidence-learning/G1-P1-ANALISIS-CASOS/` | En equipo |
| Mockup SIGRES | `03-execute/14-week/02-optional-activity/MockUp SIGRES.pptx` y proyecto Stitch | En equipo |

## 3. Qué ajusté

| # | Artefacto | Qué cambié y por qué | Commit |
|---|---|---|---|
| 1 | Casos de uso extendidos | Organicé tres casos críticos del SRS con precondiciones, flujo principal, flujo alterno, excepciones y postcondiciones. Mejora el criterio C2. | [05d52ba](https://github.com/TheKelex/ADSO-3413974/commit/05d52ba23e634e7d714e571ffb7d35af0cdc1c15) |
| 2 | Diagramas de actividades | Modelé los flujos de registrar pedido y gestionar inventario con carriles, decisiones y guardas. Mejora el criterio C1. | [3ac0337](https://github.com/TheKelex/ADSO-3413974/commit/3ac0337e239fb1a1d0cce59154e989ce2ac01d14) |
| 3 | Matriz de trazabilidad | Relacioné los requisitos funcionales con casos de uso, clases de dominio y entidades para comprobar su cobertura. Mejora el criterio C8. | [f180045](https://github.com/TheKelex/ADSO-3413974/commit/f18004509a752d28febfbeec32eb7eb652e56b2d) |

## 4. Qué construí nuevo

- Tres plantillas extendidas de casos de uso: registrar pedido, gestionar inventario y procesar pago.
- Dos diagramas de actividades en PlantUML.
- Modelo de dominio con conceptos del negocio y sus relaciones.
- Modelo entidad-relación con entidades, identificadores, cardinalidades y una relación N:M resuelta.
- Tres pantallas críticas del mockup con enlace al prototipo Stitch.
- Lista de chequeo, registro de ajustes, informe de análisis y matriz de trazabilidad.

## 5. Artefactos entregados

**Producto 2**

- [x] Diagramas de actividades: [01-Registrar-Pedido.puml](./02-Diagramas-de-Actividad/01-Registrar-Pedido.puml) y [02-Gestionar-Inventario.puml](./02-Diagramas-de-Actividad/02-Gestionar-Inventario.puml)
- [x] Plantillas extendidas de casos de uso: [01-Casos-de-Uso-Extendidos.md](./01-Casos-de-Uso-Extendidos.md)
- [x] Modelo de dominio: [03-Modelo-de-Dominio.puml](./03-Modelo-de-Dominio.puml) y [03-Modelo-de-Dominio.png](./03-Modelo-de-Dominio.png)
- [x] Modelo entidad-relación: [04-Modelo-Entidad-Relacion.puml](./04-Modelo-Entidad-Relacion.puml) y [04-Modelo-Entidad-Relacion.png](./04-Modelo-Entidad-Relacion.png)
- [x] Prototipo inicial: [05-MockUp/README.md](./05-MockUp/README.md)
- [x] Lista de chequeo diligenciada: [07-Lista-de-Chequeo.md](./07-Lista-de-Chequeo.md)
- [x] Registro de ajustes y segunda verificación: [09-Registro-de-Ajustes.md](./09-Registro-de-Ajustes.md)
- [x] Informe de análisis: [08-Informe-de-Analisis.md](./08-Informe-de-Analisis.md)
- [x] Matriz de trazabilidad: [06-Matriz-de-Trazabilidad.md](./06-Matriz-de-Trazabilidad.md)

## 6. Declaración de autoría

- [x] Declaro que el ajuste y lo construido nuevo en esta carpeta es **mi trabajo individual**.
- [x] Declaro que lo que tomé de mi equipo o de otra fuente está identificado en la sección 2.
- [x] Entiendo que el sábado 28 de noviembre de 2026 el instructor me hará **3 preguntas orales** sobre mis productos (una del producto 1, una del producto 2 y una tercera a su elección) y que, si respondo bien menos de 2, el producto sobre el que fallé queda con tope de 50.

Firma (usuario de GitHub y fecha): TheKelex - 27/11/2026

## 7. Uso de inteligencia artificial

| Herramienta | Para qué la usé | Qué verifiqué o corregí yo |
|---|---|---|
| Copilot | Preparar borradores de las plantillas, los modelos, la trazabilidad y el README a partir del SRS y los casos de uso. | Revisé el contenido, adapté los casos a SIGRES, comprobé la trazabilidad y los modelos, y mantuve los datos y cambios registrados. |

## 8. Autocontrol antes del push

- [x] Hice commits propios en al menos 2 días distintos entre la publicación de la guía (sábado 21 de noviembre de 2026) y la entrega.
- [x] No subí el marco documental del instructor, plantillas sin diligenciar ni archivos de otras personas.
- [x] Hice `push` y comprobé en GitHub que la carpeta está publicada, incluido este README actualizado.