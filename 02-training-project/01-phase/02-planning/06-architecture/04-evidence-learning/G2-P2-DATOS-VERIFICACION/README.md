# README de entrega · Guía 2: Diseño de la solución

## 1. Datos de la entrega

| Campo | Valor |
|---|---|
| Producto | P2 (modelo de datos y verificación) |
| Usuario de GitHub del autor | TheKelex |
| Ficha | 3413974 |
| Caso elegido (nombre del proyecto) | Sistema Integral de Gestión para Restaurantes (SIGRES) |
| Enlace al informe de análisis o SRS del caso (ruta en tu fork) | `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx` |
| Cuenta de GitHub con la que haces los commits | TheKelex |

## 2. Origen del insumo y trabajo en equipo

| Pregunta | Respuesta |
|---|---|
| ¿El SRS / informe de análisis lo elaboró un equipo? ¿Cuántas personas? | Sí. El SRS registra cuatro integrantes. |
| ¿Qué parte del SRS es tuya? | El SRS se reutiliza como insumo y fue elaborado en conjunto con el equipo; no declaro autoría individual de ese documento. |
| ¿Los artefactos de esta entrega los hiciste tú solo? | Sí, son productos individuales de diseño. Se usó IA como herramienta y se declara en la sección 6; el autor debe revisar y poder explicar cada decisión. |
| ¿Otro aprendiz usó o te compartió archivos de esta entrega? | No. Se reutilizan únicamente el SRS y artefactos SIGRES del equipo indicados en la sección 3. |

## 3. Qué reutilicé, qué ajusté y qué hice nuevo

| Artefacto | Archivo en esta carpeta | Origen (ruta en mi fork o «nuevo») | Qué ajusté o agregué |
|---|---|---|---|
| B1 Modelo conceptual | [01-B1-Modelo-Conceptual.md](./01-B1-Modelo-Conceptual.md) | SRS SIGRES y MER de G1-P2 | Trazabilidad RF, cardinalidades mínimas/máximas y receta producto-insumo para RF-10; se explicita el supuesto de un rol por usuario. |
| B2 Modelo lógico | [02-B2-Modelo-Logico.md](./02-B2-Modelo-Logico.md) | MER G1-P2 y clases G2-P1 | Resolví relaciones N:M, claves, datos de anulación y campos personales cifrados. |
| B3 Normalización | [03-B3-Normalizacion.md](./03-B3-Normalizacion.md) | Nuevo, derivado del modelo lógico B2 | Normalicé detalle de pedido con llave compuesta y factura con dependencia transitiva, mostrando DF. |
| B4 Diccionario de datos | [04-B4-Diccionario-de-Datos.md](./04-B4-Diccionario-de-Datos.md) | Nuevo, derivado de B2 | Documenté las 16 tablas y 96 columnas con tipos, tamaños, restricciones y dominios; se comparó con MySQL. |
| B5 Script de creación | [05-B5-Crear-Base-de-Datos.sql](./05-B5-Crear-Base-de-Datos.sql) | Nuevo | Creé el esquema MySQL, vista de reportes, restricciones, FK y protección de facturas pagadas. |
| Evidencia de ejecución B5 | [07-Evidencia-Ejecucion-MySQL.md](./07-Evidencia-Ejecucion-MySQL.md), [07-Ejecucion-MySQL.log](./07-Ejecucion-MySQL.log) | Nuevo, ejecución local con MySQL Community 8.0.46 | Registré tablas, restricciones y pruebas de rechazo/aceptación en un contenedor temporal con datos sintéticos. |
| B6 Políticas de seguridad | [06-B6-Politicas-de-Seguridad.md](./06-B6-Politicas-de-Seguridad.md) | Nuevo, a partir de RF/RNF del SRS | Definí políticas de integridad, confidencialidad y disponibilidad, matriz de roles, cifrado, bcrypt, TLS y respaldos. |
| B7 Lista de chequeo y hallazgos | [08-B7-Lista-de-Chequeo-y-Hallazgos.md](./08-B7-Lista-de-Chequeo-y-Hallazgos.md) | Nuevo, revisión de los artefactos G2-P1 y P2 | Apliqué 27 criterios medibles a ambos productos y registré cuatro hallazgos con ajuste y estado. |

**Qué reutilicé (resumen en 3 a 5 líneas):**
- El SRS SIGRES, que fue elaborado por el equipo de cuatro integrantes.
- El MER y el modelo de dominio de G1-P2, además del diseño publicado en G2-P1.
- Los RF/RNF y reglas de negocio del caso; no reutilicé plantillas de ejemplo como producto final.

**Qué ajusté para llevarlo al estándar de la rúbrica (resumen en 3 a 5 líneas):**
- Resolví en lógico las relaciones N:M de roles-permisos, pedido-productos y producto-insumos.
- Añadí trazabilidad a RF-10, control de anulación de factura y protección de datos personales.
- Alineé columnas, tipos y restricciones del modelo, diccionario y SQL; contrasté sus 96 nombres de columnas.

**Qué hice completamente nuevo (resumen en 3 a 5 líneas):**
- La normalización paso a paso, el diccionario completo y el script MySQL ejecutable.
- Las políticas de seguridad, la matriz de acceso y la lista de chequeo sobre P1 y P2.
- Las pruebas de ejecución del esquema en MySQL 8.0.46 y el registro de hallazgos.

## 4. Autoverificación de mínimos

- [x] B1 Modelo conceptual: ≥ 6 entidades, cardinalidades, ≥ 1 N:M y ≥ 1 1:N.
- [x] B2 Modelo lógico: ≥ 8 tablas con PK/FK, tabla intermedia para cada N:M y reglas de transformación.
- [x] B3 Normalización: ≥ 2 tablas de 0FN a 3FN con dependencias funcionales, llave compuesta y dependencia transitiva.
- [x] B4 Diccionario de datos completo.
- [x] B5 Script SQL ejecutado sin errores, evidencia adjunta y `ON DELETE` en ≥ 3 FK.
- [x] B6 ≥ 6 políticas de seguridad, matriz de ≥ 3 roles, bcrypt con sal y TLS.
- [x] B7 ≥ 20 criterios medibles, diligenciados para P1 y P2 con evidencia.

## 5. Historial de commits

La entrega P2 fue publicada el 28/11/2026. El último commit informado para esta carpeta es `0769a476e1e03b63d2109f98021594b034770c3d`. Junto con el avance de P1 del 27/11/2026, queda registrado trabajo en dos días distintos.

| Día | Qué avancé (mensaje del commit) |
|---|---|
| 27/11/2026 | G2-P1: diseño SIGRES publicado; commit registrado en el README de P1. |
| 28/11/2026 | Publicación de los artefactos G2-P2 (commit `0769a476e1e03b63d2109f98021594b034770c3d`). |

## 6. Uso de inteligencia artificial (declaración obligatoria)

| Herramienta | Para qué la usé | Qué verifiqué o cambié después |
|---|---|---|
| Copilot en VS Code | Apoyo en borradores de modelo, diccionario, políticas, lista de chequeo y script SQL; revisión de consistencia entre artefactos. | Ejecuté el DDL y pruebas funcionales en MySQL 8.0.46, comparé las 96 columnas del diccionario con el esquema y corregí inconsistencias. El autor debe revisar las decisiones de negocio y poder explicarlas en la verificación oral. |

## 7. Enlaces

- Informe de análisis / SRS: `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx`.
- Carpeta de esta entrega en GitHub: https://github.com/TheKelex/ADSO-3413974/tree/0769a476e1e03b63d2109f98021594b034770c3d/02-training-project/01-phase/02-planning/06-architecture/04-evidence-learning/G2-P2-DATOS-VERIFICACION/.
- Otros enlaces de apoyo: ninguno.

## 8. Declaración de autoría

Declaro que los artefactos de esta carpeta son de mi autoría, que lo que reutilicé de mi trabajo del año o de otras fuentes está indicado en la sección 3, que no entregué como propio material del instructor ni de otros aprendices, y que puedo explicar cada artefacto en la verificación oral.

Usuario de GitHub: TheKelex  Fecha: 28/11/2026.
