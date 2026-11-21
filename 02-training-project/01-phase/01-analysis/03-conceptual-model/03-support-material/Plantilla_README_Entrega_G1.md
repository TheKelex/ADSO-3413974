# Plantilla de README de entrega · Guía 1 · Modelo conceptual

**Cómo usarla:** copia este archivo como `README.md` dentro de cada carpeta de entrega de tu fork y completa **todos** los campos. Se usa una vez en `G1-P1-<nombre-corto>/` (producto 1) y otra en `G1-P2-<nombre-corto>/` (producto 2). Borra las instrucciones entre paréntesis cuando termines.

Ruta de entrega en tu fork:
`02-training-project/01-phase/01-analysis/03-conceptual-model/04-evidence-learning/`

Fecha límite: **el sábado 28 de noviembre de 2026 a las 6:00 a. m.** (último commit antes de esa hora, antes de iniciar la sesión presencial). Criterios y mínimos: `Rubricas_G1.md`.

---

> Este repositorio es público. No escribas tu nombre completo, tu número de documento, teléfonos ni correos, ni datos reales de terceros. Identifícate con tu usuario de GitHub.

## 1. Datos de la entrega

- **Producto:** (P1 plan de análisis y casos de uso / P2 modelo de análisis y verificación)
- **Usuario de GitHub del autor:**
- **Caso de estudio:** (por ejemplo, el nombre del proyecto de tu equipo del año)
- **Trabajo en equipo del año:** (sí / no. Si sí, indica el nombre del equipo, no los datos personales de los integrantes)

## 2. Qué reutilicé

(Una fila por artefacto que tomaste de tu trabajo del año. La ruta es la de tu fork.)

| Artefacto | Ruta original en mi fork (semana) | ¿Lo hice solo o en equipo? |
|---|---|---|
| | | |

## 3. Qué ajusté

(Al menos 3 cambios concretos por producto. Cada uno con el enlace al commit donde lo hiciste y qué criterio de `Rubricas_G1.md` mejora.)

| # | Artefacto | Qué cambié y por qué | Enlace al commit |
|---|---|---|---|
| 1 | | | |
| 2 | | | |
| 3 | | | |

## 4. Qué construí nuevo

(Lo que no existía en tu trabajo del año y creaste para esta guía.)

-

## 5. Artefactos entregados

(Marca cada uno que entregas, con enlace relativo al archivo. Las cantidades mínimas están en `Rubricas_G1.md`.)

**Producto 1**

- [ ] Caracterización de la solución: [enlace]()
- [ ] Comparación de metodologías y selección justificada: [enlace]()
- [ ] Plan de actividades, cronograma y entregables: [enlace]()
- [ ] Diagrama de casos de uso (fuente editable): [enlace]()
- [ ] Diagrama de casos de uso (imagen): [enlace]()
- [ ] Matriz requisito → caso de uso: [enlace]()

**Producto 2**

- [ ] Diagramas de actividades (2 o más): [enlaces]()
- [ ] Plantillas extendidas de casos de uso (3 o más): [enlaces]()
- [ ] Modelo de dominio: [enlace]()
- [ ] Modelo entidad-relación: [enlace]()
- [ ] Prototipo inicial (3 o más pantallas): [enlace]()
- [ ] Lista de chequeo diligenciada (15 o más criterios): [enlace]()
- [ ] Registro de ajustes y segunda verificación: [enlace]()
- [ ] Informe de análisis (Anexo E): [enlace]()
- [ ] Matriz de trazabilidad (Anexo C): [enlace]()

## 6. Declaración de autoría

Marca y completa:

- [ ] Declaro que el ajuste y lo construido nuevo en esta carpeta es **mi trabajo individual**.
- [ ] Declaro que lo que tomé de mi equipo o de otra fuente está identificado en la sección 2.
- [ ] Entiendo que el sábado 28 de noviembre de 2026 el instructor me hará **3 preguntas orales** sobre mis productos (una del producto 1, una del producto 2 y una tercera a su elección) y que, si respondo bien menos de 2, el producto sobre el que fallé queda con tope de 50.

Firma (usuario de GitHub y fecha):

## 7. Uso de inteligencia artificial

(Es válido usarla; hay que declararla. Si no la usaste, escribe «No usé IA».)

| Herramienta | Para qué la usé | Qué verifiqué o corregí yo |
|---|---|---|
| | | |

## 8. Autocontrol antes del push

- [ ] Hice commits propios en al menos 2 días distintos entre la publicación de la guía (sábado 21 de noviembre de 2026) y la entrega.
- [ ] No subí el marco documental del instructor, plantillas sin diligenciar ni archivos de otras personas.
- [ ] Hice `push` y comprobé en GitHub que la carpeta está publicada.

---

# Anexos de formato

Los siguientes formatos cubren artefactos que no se trabajaron durante el año. Cópialos en tu carpeta y diligéncialos.

## Anexo A. Plantilla de caso de uso extendido

| Campo | Contenido |
|---|---|
| Identificador y nombre | CU-00 · Verbo + objeto (por ejemplo, Registrar préstamo) |
| Requisitos que atiende | RF-00, RF-00 |
| Actor primario / secundarios | |
| Objetivo del actor | Lo que el actor quiere lograr |
| Precondiciones | Lo que el sistema asume cierto antes de empezar (por ejemplo, el usuario inició sesión) |
| Disparador | Qué hace el actor para iniciar el caso |
| Flujo principal | 1. El actor … 2. El sistema … (mínimo 6 pasos, redactados como intención, sin tecnología) |
| Flujos alternos | 3a. Si [variante válida del negocio, que continúa o cierra el caso con éxito alternativo], el sistema … (mínimo 1) |
| Excepciones | 4b. Si [falla o regla violada que impide el objetivo], el sistema … (mínimo 1) |
| Postcondiciones | Estado en que queda el sistema si el caso termina con éxito |
| Reglas de negocio | |
| Pantalla del prototipo | Enlace al mockup |

## Anexo B. Formato de lista de chequeo de validación

| # | Criterio verificable (responde sí o no) | Artefacto | ¿Cumple? | Evidencia (enlace o cita) | Hallazgo |
|---|---|---|---|---|---|
| 1 | Cada requisito funcional del SRS aparece en al menos un caso de uso | Matriz | | | |
| 2 | Cada caso de uso tiene plantilla extendida con flujo alterno y excepción | Plantillas | | | |
| 3 | Todo concepto usado en las plantillas existe en el modelo de dominio | Dominio | | | |
| 4 | Cada rama del diagrama de actividades aparece en la plantilla del mismo caso | Actividades y plantillas | | | |
| 5 | Toda pantalla del prototipo está rotulada con su caso de uso | Mockup | | | |

(Completa hasta 15 criterios como mínimo; incluye al menos 3 de cobertura o trazabilidad. Cada criterio debe poder responderse sí o no con una evidencia: «está bien hecho» no es un criterio verificable.)

## Anexo C. Matriz de trazabilidad requisito → artefactos

| Requisito | Caso de uso | Diagrama de actividades | Plantilla | Clase de dominio | Entidad del MER | Pantalla |
|---|---|---|---|---|---|---|
| RF-01 | | | | | | |

## Anexo D. Registro de ajustes

| # | Hallazgo (criterio que no cumplía) | Artefacto | Antes | Después | Enlace al commit | ¿Se verificó de nuevo? |
|---|---|---|---|---|---|---|
| 1 | | | | | | |
