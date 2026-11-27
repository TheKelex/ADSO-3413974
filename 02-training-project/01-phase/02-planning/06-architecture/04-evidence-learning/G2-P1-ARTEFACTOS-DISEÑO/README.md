# README de entrega · Guía 2: Diseño de la solución

## 1. Datos de la entrega

| Campo | Valor |
|---|---|
| Producto | P1 (artefactos de diseño) |
| Usuario de GitHub del autor | TheKelex |
| Ficha | 3413974 |
| Caso elegido | Sistema Integral de Gestión para Restaurantes (SIGRES) |
| Enlace al informe de análisis o SRS del caso | `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx` |
| Cuenta de GitHub con la que haces los commits | TheKelex |

## 2. Origen del insumo y trabajo en equipo

| Pregunta | Respuesta |
|---|---|
| ¿El SRS lo elaboró un equipo? ¿Cuántas personas? | Sí. El SRS registra cuatro integrantes. |
| ¿Qué parte del SRS es tuya? | El SRS se reutiliza como insumo y fue realizado en su totalidad en conjunto con el equipo. |
| ¿Los artefactos de esta entrega los hiciste tú solo? | Fueron elaborados con apoyo de IA declarado en la sección 6. |
| ¿Otro aprendiz usó o te compartió archivos de esta entrega? | No, se usan el SRS y los artefactos SIGRES del equipo declarados en la sección 3. |

## 3. Qué reutilicé, qué ajusté y qué hice nuevo

| Artefacto | Archivo en esta carpeta | Origen (ruta en mi fork o «nuevo») | Qué ajusté o agregué |
|---|---|---|---|
| A1 Ficha de interpretación | [01-Ficha-de-Interpretacion.md](./01-Ficha-de-Interpretacion.md) | SRS SIGRES: `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx` | Seleccioné requisitos, restricciones y stack permitidos por el SRS. |
| A2 Diagrama de clases | [02-Diagrama-de-Clases.puml](./02-Diagrama-de-Clases.puml), [02-Diagrama-de-Clases.jpeg](./02-Diagrama-de-Clases.jpeg) (anterior al ajuste de disposición) | Modelo de dominio G1-P2 y SRS SIGRES | Transformé conceptos del dominio en clases con tipos, visibilidad, métodos, relaciones y multiplicidades. |
| A3 Patrones | [03-Patrones-de-Diseno.md](./03-Patrones-de-Diseno.md) | Nuevo, a partir de RF-04, RF-05 y RF-13 del SRS | Apliqué Estrategia a los métodos de pago y Fachada al flujo de pedido. |
| A4 Arquitectura | [04-Arquitectura.puml](./04-Arquitectura.puml), [04-Arquitectura-vista-previa.png](./04-Arquitectura-vista-previa.png), [04-Decisiones-de-Arquitectura.md](./04-Decisiones-de-Arquitectura.md) | Nuevo, a partir de los RNF del SRS | Definí capas, comunicaciones, decisiones para RNF y una alternativa descartada. |
| A5 Vista de componentes | [05-Vista-de-Componentes.puml](./05-Vista-de-Componentes.puml), [05-Vista-de-Componentes-vista-previa.png](./05-Vista-de-Componentes-vista-previa.png), [05-Mapeo-de-Componentes.md](./05-Mapeo-de-Componentes.md) | Nuevo, a partir de A2 y el SRS | Separé componentes, interfaces, dependencias, clases y RF atendidos. |
| A6 Vista de despliegue | [06-Vista-de-Despliegue.puml](./06-Vista-de-Despliegue.puml), [06-Vista-de-Despliegue-vista-previa.png](./06-Vista-de-Despliegue-vista-previa.png) | Nuevo, a partir de RNF-03, RNF-20 y RNF-30 | Definí nodos, artefactos, protocolos y el respaldo en ubicación separada. |
| A7 Matriz de trazabilidad | [07-Matriz-de-Trazabilidad.md](./07-Matriz-de-Trazabilidad.md) | Matriz G1-P1 y nuevo diseño G2-P1 | Relacioné los 17 RF con clases, componentes y decisiones. |

**Qué reutilicé:**
- El SRS de SIGRES y sus 17 RF y 31 RNF.
- Los casos de uso y el modelo conceptual elaborados durante G1.
- El stack permitido y los actores ya definidos para el caso.

**Qué ajusté:**
- Convertí conceptos del análisis en clases, interfaces y relaciones de diseño.
- Alineé los componentes y el despliegue con los RNF.
- Dejé explícitos los supuestos que el SRS no resuelve.

**Qué hice completamente nuevo:**
- Los patrones, la arquitectura y sus decisiones.
- Las vistas de componentes y despliegue.
- La matriz de trazabilidad entre requisitos y diseño.

## 4. Autoverificación de mínimos

- [x] A1 Ficha de interpretación: ≥ 8 RF y ≥ 3 RNF con ID; plataforma justificada.
- [x] A2 Diagrama de clases: ≥ 8 clases, tipos, métodos, multiplicidades, composición e interfaz con implementaciones.
- [x] A3 ≥ 2 patrones GoF aplicados; uno distinto de Singleton.
- [x] A4 Arquitectura con diagrama y decisiones para ≥ 3 RNF.
- [x] A5 Vista de componentes con interfaces y dependencias.
- [x] A6 Vista de despliegue con ≥ 3 nodos y protocolos.
- [x] A7 Matriz de trazabilidad de los 17 RF.
- [x] Imágenes finales exportadas desde las fuentes PlantUML; hay PNG de vista previa local para cada diagrama.

## 5. Historial de commits

| Día | Qué avancé (mensaje del commit) |
|---|---|
| 27/11/2026 | Se realizo todo el contenido entendido en G2-P1 |
| Pendiente | Registrar un segundo avance sustantivo en otro día. |

## 6. Uso de inteligencia artificial

| Herramienta | Para qué la usé | Qué verifiqué o cambié después |
|---|---|---|
| Copilot | Preparar borradores de fichas, clases, patrones, arquitectura y trazabilidad a partir del SRS y los materiales SIGRES. | Revisaré las decisiones, corregiré los artefactos y confirmaré que puedo explicar cada diseño antes de entregar. |

## 7. Enlaces

- Informe de análisis / SRS: `03-execute/28-week/02-optional-activity/SRS - SIGRES.docx`.
- Carpeta de esta entrega en GitHub: https://github.com/TheKelex/ADSO-3413974/tree/ef965b4babb7d1f18cb483f8489cc0034f3d546d/02-training-project/01-phase/02-planning/06-architecture/04-evidence-learning/G2-P1-ARTEFACTOS-DISE%C3%91O.
- Otros enlaces de apoyo: ninguno.

## 8. Declaración de autoría

Declaro que los artefactos de esta carpeta son de mi autoría, que lo reutilizado de SIGRES está indicado en la sección 3, que no entregué como propio material del instructor ni de otros aprendices, y que puedo explicar cada artefacto en la verificación oral.

Usuario de GitHub: TheKelex  
Fecha y confirmación final: 27/11/2026.
