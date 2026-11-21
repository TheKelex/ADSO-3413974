# Rúbricas · Guía 1 · Modelo conceptual (análisis)

**Ficha 3413974 · Análisis y Desarrollo de Software (228118) · Competencia 220501093 - Análisis de la especificación de requisitos del software**
RAP evaluados: **592375** (planear actividades de análisis) · **592373** (modelar las funciones del software) · **592374** (verificar los modelos del análisis). Instructor: Jesús Ariel González Bonilla (CIES).

Este documento es público: lo lees antes de empezar, lo usas para autoevaluarte y el instructor lo aplica tal cual para calificar. No hay criterios ocultos: lo único que no se publica son los casos de la prueba de desempeño, sus errores sembrados y su solución.

## 1. Cómo se calcula la nota de la guía

| Entregable | Peso | Cuándo | Rúbrica |
|---|---|---|---|
| Prueba de conocimiento | 25 % | Sábado 28 de noviembre de 2026, presencial, en cuestionario de la plataforma Moodle (Zajuna) con calificación automática (20 preguntas, 40 min) | Cada pregunta vale 5 puntos y la plataforma califica sola; la nota sobre 100 es la nota de la evidencia. Se estudia con `Banco_Estudio_G1 (.html o .pdf)` |
| Producto 1: plan de análisis y casos de uso | 25 % | Entrega a más tardar el sábado 28 de noviembre de 2026 a las 6:00 a. m. | Sección 4 |
| Producto 2: modelo de análisis y verificación | 25 % | Entrega a más tardar el sábado 28 de noviembre de 2026 a las 6:00 a. m. | Sección 5 |
| Prueba de desempeño | 25 % | Sábado 28 de noviembre de 2026, presencial, individual, en computador, 90 min | Sección 6 |

- **Nota de la guía = promedio de los cuatro entregables (0 a 100).** Se aprueba con **70 o más** y además se cumple la regla por RAP de la sección 1.1.
- No hay más entregables ni puntos extra.
- **Productos 1 y 2:** cada criterio se aplica con **cuatro niveles estrictos: 100, 70, 40 y 0 % del peso del criterio**, sin niveles intermedios. Si el instructor duda entre dos niveles, califica el inferior y lo registra en la retroalimentación.
- **Prueba de desempeño:** rige la **hoja de puntos de la sección 6** (los puntos de cada ítem se suman). Los niveles de esa tabla son solo una guía de lo que se espera; si hay diferencia, vale la hoja de puntos.

### 1.1 Juicio por RAP (cómo se pasa de la nota al juicio en Sofía)

La nota de la guía no basta: el juicio se registra **por resultado de aprendizaje**. Cada RAP reúne los criterios y preguntas siguientes y se calcula como el porcentaje de puntos obtenidos sobre los puntos posibles de esos elementos:

| RAP | Qué lo evidencia |
|---|---|
| 592375 · Planear actividades de análisis | Producto 1 (C1, C2) y las preguntas de la prueba de conocimiento con etiqueta 592375 |
| 592373 · Modelar las funciones del software | Producto 1 (C3, C4), producto 2 (C1, C2, C3, C4, C8), prueba de desempeño (D1, D2, D3) y las preguntas con etiqueta 592373 |
| 592374 · Verificar los modelos del análisis | Producto 2 (C5, C6, C7), prueba de desempeño (D4, D5) y las preguntas con etiqueta 592374 |

**Se registra Aprobado en un RAP solo si la nota de la guía es 70 o más y ese RAP alcanza al menos 70 % de sus puntos.** El criterio C5 del producto 1 (README y autoría) es un control de autoría y no pertenece a ningún RAP. Esta regla impide aprobar la guía con dos productos incompletos (tope 60) compensados por las pruebas. Si un RAP no alcanza el 70 %, queda **por evaluar** y se concreta un plan de mejoramiento con fecha de recuperación; nunca se registra como reprobado. La misma salida aplica a quien no se presente el sábado 28: el instructor fija la fecha de recuperación de las pruebas.

## 2. Qué se reutiliza del año y qué se construye nuevo

Los productos 1 y 2 **no se hacen desde cero**: reutilizas y ajustas lo que ya está en tu fork, si existe. Cada tema tiene un punto de partida posible; revisa qué tienes antes de empezar y escríbelo en el README:

| Artefacto | Dónde buscar en tu fork | Si no lo tienes |
|---|---|---|
| Informe de requisitos (SRS) | `03-execute` (semanas de requisitos, p. ej. 09 y 11) | Toma el SRS de tu proyecto del año o el caso del instructor |
| Casos de uso y diagrama de clases | `03-execute` (semanas 15 y 16, documentos de casos de uso, `.drawio`, Word) | Construye con la plantilla y los anexos |
| Mockups y pitch | `03-execute` (semanas 13, 14 y 17) | Construye 3 pantallas mínimas |
| Caracterización, priorización y MER | Cualquier `02-optional-activity` o `04-evidence-learning` donde hayas trabajado Design Thinking, MoSCoW o MER | Construye con la guía; el MER es nuevo para ti |
| Comparación y selección de metodología, plan y cronograma de análisis | No hay base: se construye con la guía | Construye |
| Diagrama de actividades | Si tienes flujos BPMN de procesos, se adaptan a notación UML de actividades | Construye |
| Plantilla extendida de caso de uso | Descripciones narrativas de tu SRS o de tus casos de uso, reestructuradas con el Anexo A | Construye |
| Modelo de dominio | El diagrama de clases del año es la base para abstraer conceptos del negocio (sin métodos ni claves) | Construye |
| Informe de análisis, lista de chequeo, matriz de trazabilidad, registro de ajustes | No hay base: usa los anexos de la plantilla de entrega | Construye |

Lo que se califica es **tu ajuste individual** y lo que construyes nuevo, no el archivo base. Si tu artefacto base lo hiciste en equipo, puedes reutilizarlo, pero lo declaras en el README y **el ajuste, la selección de metodología, la lista de chequeo, el informe de análisis y el registro de ajustes son individuales**.

### Presupuesto de tiempo sugerido (supuesto de diseño)

Las 6 horas de la actividad de apropiación (3 de trabajo directo y 3 independientes) se reparten en **un bloque por producto de 3 horas (180 minutos) como máximo**. El tiempo estimado, con reutilización, es:

| Producto | Tarea | Minutos |
|---|---|---|
| 1 | Caracterizar, comparar metodologías, seleccionar y justificar | 50 |
| 1 | Plan de actividades de análisis con cronograma y entregables | 40 |
| 1 | Matriz requisito → caso de uso | 25 |
| 1 | Ajustar el diagrama de casos de uso (reutilizado) y README | 55 |
| | **Subtotal producto 1** | **170** |
| 2 | 2 diagramas de actividades (adaptando flujos o texto) | 25 |
| 2 | 3 plantillas extendidas (adaptando texto narrativo del SRS) | 35 |
| 2 | Modelo de dominio (a partir del diagrama de clases) | 20 |
| 2 | Ajuste del MER y del mockup reutilizados | 30 |
| 2 | Lista de chequeo (Anexo B), aplicación, ajustes y segunda verificación | 35 |
| 2 | Informe de análisis y matriz de trazabilidad (Anexo C) | 30 |
| 2 | README | 5 |
| | **Subtotal producto 2** | **180** |
| | **Total aproximado** | **350 min (5 h 50 min) dentro de las 6 h** |

El estudio del banco de preguntas es independiente y cuenta dentro de las horas de la actividad de conceptualización, no de estas 6 horas. Si tu caso base está incompleto, apóyate en el acompañamiento del instructor y en los anexos de la plantilla. Las cantidades mínimas no se reducen, pero la profundidad de cada artefacto debe ser proporcional: un artefacto correcto y completo vale más que uno extenso.

## 3. Reglas generales de aceptación y autoría

### 3.1 Requisitos mínimos de aceptación (tope 60)

Cada producto tiene una lista de artefactos obligatorios con cantidad mínima (secciones 4 y 5). **Si falta un artefacto obligatorio, o no alcanza la cantidad mínima, el producto tiene un tope de 60 puntos**, sin importar lo bien hecho que esté el resto.

### 3.2 Control de autoría individual (verificable)

1. **Historial de commits.** Debe haber commits tuyos en **al menos 2 días calendario distintos** dentro de la ventana de trabajo (entre la publicación de la guía, el sábado 21 de noviembre de 2026, y la entrega, el sábado 28 de noviembre de 2026 a las 6:00 a. m.), en la carpeta de cada producto. Con menos de 2 días el producto tiene tope de 70, salvo que el instructor verifique de otro modo tu autoría en la sustentación.
2. **Commits masivos.** Un commit de 40 o más archivos se revisa: si los archivos son el marco documental del instructor, plantillas o trabajo de otra persona, **no cuenta como trabajo tuyo** y no se califica.
3. **README de reutilización.** Cada carpeta lleva su `README.md` (plantilla `Plantilla_README_Entrega_G1.md`) con lo que reutilizaste, lo que ajustaste y los enlaces. Sin README no hay calificación del criterio de autoría (C5 del producto 1) y el producto queda con tope de 60.
4. **Verificación oral el sábado 28 de noviembre.** El instructor hace **3 preguntas** sobre tus productos: una sobre el producto 1, una sobre el producto 2 y una tercera a elección del instructor sobre cualquiera de los dos (por ejemplo: por qué una relación es «include» y no «extend», qué cambiaste respecto a tu versión original y en qué commit, qué artefactos tocarías si cambia un requisito). Respondes bien **2 de 3 o más**: sin efecto. Respondes bien **1 de 3 o ninguna**: **queda con tope de 50 el producto (o los productos) sobre el que fallaste**. La verificación se hace de a un aprendiz, durante la prueba de desempeño y en la franja de cierre; su tiempo **no se descuenta** de los 90 minutos de la prueba. Si el grupo no se alcanza a verificar en esas franjas, el instructor cita a sustentación en una fecha que anuncia por el canal oficial de la ficha.
5. **Copia entre aprendices.** Si dos productos tienen contenido idéntico sin que el README lo explique como trabajo en equipo, ambos pasan a revisión; mientras no se aclare quién es el autor, el tope es 50. El trabajo en equipo del año es válido como base, pero el ajuste, la selección de metodología, la lista de chequeo, el informe de análisis y el registro de ajustes son individuales.
6. **Uso de inteligencia artificial.** Se permite, pero se declara en el README (para qué la usaste y qué verificaste). Si no la declaras y no puedes explicar tu producto en la verificación oral, aplica el tope de 50.

### 3.3 Qué versión se califica

Se califica el último commit anterior al sábado 28 de noviembre de 2026 a las 6:00 a. m. Los commits posteriores no se tienen en cuenta.

## 4. Rúbrica del Producto 1: plan de análisis y casos de uso

Carpeta: `02-training-project/01-phase/01-analysis/03-conceptual-model/04-evidence-learning/G1-P1-<nombre-corto>/`

### 4.1 Artefactos obligatorios y cantidades mínimas

| # | Artefacto | Mínimo exigible |
|---|---|---|
| a | Caracterización de la solución | Tipo de proyecto, tamaño del equipo, estabilidad de requisitos y disponibilidad del cliente |
| b | Comparación de metodologías | Al menos **3** de RUP, Scrum, XP y RAD, con una seleccionada y justificada |
| c | Plan de actividades de análisis | Al menos **6 actividades** con responsable, fecha y entregable cada una |
| d | Diagrama de casos de uso UML | Al menos **8 casos de uso**, **2 actores** primarios, al menos **1 «include» y 1 «extend»**, cada relación con una línea de justificación (si tu caso no admite «extend», lo justificas por escrito) |
| e | Matriz requisito → caso de uso | Cubre al menos el **90 %** de los requisitos funcionales del SRS; los no cubiertos van justificados |
| f | Fuente editable e imagen del diagrama | `.drawio` (u otra herramienta) y `.png` o `.svg` |
| g | `README.md` de la carpeta | Según la plantilla |

### 4.2 Criterios (suman 100)

Los criterios C1 a C4 son los criterios de evaluación del programa; C5 es un control de autoría de esta guía.

| # | Criterio | Peso | 100 % | 70 % | 40 % | 0 % |
|---|---|---|---|---|---|---|
| C1 | **Identifica metodologías de desarrollo de software** de acuerdo con las características del software a desarrollar (RAP 592375) | 15 | Caracteriza la solución con los 4 aspectos y compara 3 o más metodologías con al menos 2 criterios tomados de esas características | Caracteriza con 3 aspectos y compara 3 metodologías, con criterios generales | Menciona metodologías sin comparar con las características del proyecto | No hay caracterización ni comparación |
| C2 | **Establece las actividades de análisis** de acuerdo con la metodología seleccionada (RAP 592375) | 20 | Metodología seleccionada y justificada con las características; 6 o más actividades priorizadas, con responsable, fecha y entregable, coherentes con la metodología | Selecciona y justifica, pero 1 o 2 actividades carecen de responsable, fecha o entregable | Selecciona sin justificar, o el cronograma no responde a la metodología (por ejemplo, XP con una sola entrega final) | No selecciona metodología o no hay plan |
| C3 | **Interpreta el informe de requisitos** para modelar las funciones del software (RAP 592373) | 15 | Matriz con 90 % o más de los requisitos funcionales cubiertos y los restantes justificados; distingue requisitos funcionales de no funcionales | Cobertura de 75 % a 89 %, con la distinción funcional y no funcional | Cobertura de 50 % a 74 %, o confunde funcionales con no funcionales | Cobertura inferior al 50 % o no hay matriz |
| C4 | **Elabora diagramas de casos de uso** de acuerdo con los estándares actuales mediante Lenguaje Unificado de Modelado (UML) (RAP 592373). Incluye el uso de «include», «extend» y generalización con criterio | 40 | 8 o más casos con nombre de verbo más objeto, frontera del sistema, actores fuera de ella, asociaciones correctas, sin ciclos ni casos duplicados, y cada relación «include», «extend» o de generalización justificada en una línea y conforme a la sección 7 | Cumple los mínimos con 1 o 2 errores de notación (por ejemplo, un nombre que no es función) o con justificaciones incompletas de las relaciones | **De 5 a 7 casos** (además aplica el tope de 60 de la sección 3.1), o 8 o más casos con 3 o más errores, o confunde «include» y «extend» | Menos de 5 casos de uso, sin frontera o sin actores |
| C5 | **Control de autoría y trazabilidad del ajuste**: README de reutilización (no es criterio del programa) | 10 | README completo: qué se reutilizó, qué se ajustó (con al menos 3 cambios concretos y enlaces a commits), declaración de autoría e IA | README completo, con 1 o 2 cambios concretos | README parcial o los ajustes no se pueden comprobar | Sin README, o entrega copiada sin ajustes |

## 5. Rúbrica del Producto 2: modelo de análisis y verificación

Carpeta: `02-training-project/01-phase/01-analysis/03-conceptual-model/04-evidence-learning/G1-P2-<nombre-corto>/`

### 5.1 Artefactos obligatorios y cantidades mínimas

| # | Artefacto | Mínimo exigible |
|---|---|---|
| a | Diagramas de actividades | Al menos **2**, de casos de uso críticos, con inicio, fin, calles por actor y **decisiones con guardas entre corchetes que cubren todas las salidas**; cada rama alterna o de excepción aparece también en la plantilla del mismo caso. Si hay concurrencia usa bifurcación y unión; si no la hay, lo dices en una línea |
| b | Plantillas extendidas de caso de uso | Al menos **3**, cada una con precondiciones, flujo principal de 6 o más pasos, **1 flujo alterno y 1 excepción** como mínimo, y postcondiciones |
| c | Modelo de dominio | Al menos **8 clases conceptuales**, con asociaciones y multiplicidades, sin métodos; en al menos 5 clases, **2 o más atributos de negocio**; al menos **1 generalización o clase de asociación** (o una línea que justifique por qué el caso no la tiene) |
| d | Modelo entidad-relación | Notación declarada (Chen o pata de gallo). Al menos **6 entidades**, con identificadores, cardinalidad y participación; al menos **1 relación N:M** resuelta con entidad asociativa que lleve atributos propios; al menos **1 entidad débil o relación recursiva** (o una línea que justifique por qué el caso no la tiene) |
| e | Prototipo inicial (mockup) | Al menos **3 pantallas**, una por cada caso de uso crítico, rotuladas con el caso de uso que atienden |
| f | Lista de chequeo diligenciada | Al menos **15 criterios verificables** (respuesta sí/no con evidencia), de los cuales mínimo **3 de cobertura o trazabilidad** |
| g | Registro de ajustes | Al menos **5 hallazgos** con antes y después (o una justificación escrita de por qué no hay), y una segunda verificación de los criterios corregidos |
| h | Informe de análisis | De 1 a 2 páginas: requisitos interpretados (funcionales y no funcionales), enlace a cada modelo, resultado de la verificación (criterios que cumplen y no cumplen) y ajustes realizados |
| i | Matriz de trazabilidad (Anexo C) | Al menos el **90 %** de los requisitos funcionales trazados hasta caso de uso, plantilla o diagrama, clase de dominio y entidad; la pantalla se exige para los casos críticos |
| j | `README.md` de la carpeta | Según la plantilla |

### 5.2 Criterios (suman 100)

Todos los criterios son criterios de evaluación del programa.

| # | Criterio | Peso | 100 % | 70 % | 40 % | 0 % |
|---|---|---|---|---|---|---|
| C1 | **Realiza diagramas de actividades** exponiendo detalles de los casos de uso (RAP 592373) | 15 | 2 o más diagramas con decisiones completas y con guardas, flujos alternos, calles y bifurcación y unión donde hay concurrencia (o su justificación) | 2 diagramas correctos, con 1 o 2 omisiones (por ejemplo, una decisión con una sola salida) | 1 diagrama, o 2 con errores de notación frecuentes | Ningún diagrama o ilegibles |
| C2 | **Genera plantillas extendidas de casos de uso** expresando la intención de las acciones a desarrollar (RAP 592373) | 15 | 3 o más plantillas completas, con flujos alternos y excepciones coherentes con el diagrama de actividades y sin detalles de implementación | 3 plantillas, con 1 o 2 omisiones (un campo vacío o un flujo sin excepción) | 2 plantillas, o incluyen detalles de implementación (SQL, nombres de tablas, tecnología) | Menos de 2 plantillas |
| C3 | **Representa el negocio en término de clases abstractas** generando un modelo de dominio consistente (RAP 592373) | 10 | 8 o más clases conceptuales con atributos de negocio, multiplicidades correctas, generalización o clase de asociación (o su justificación), sin métodos y coherentes con las plantillas | Cumple los mínimos con 1 o 2 inconsistencias o multiplicidades dudosas | Menos de 8 clases, o mezcla métodos y tablas | Menos de 5 clases o no es un modelo de dominio |
| C4 | **Elabora el modelo entidad relación** de acuerdo con los requisitos del software (RAP 592373) | 15 | 6 o más entidades con identificadores, cardinalidad y participación en notación declarada; N:M resuelta con entidad asociativa con atributos; entidad débil o relación recursiva (o su justificación); cada requisito de datos soportado | Cumple los mínimos con 1 o 2 errores de cardinalidad o de identificador | Menos de 6 entidades, o la N:M sin resolver | Menos de 4 entidades o contradice los requisitos |
| C5 | **Realiza prototipo inicial del software** de acuerdo con los casos de uso identificados (RAP 592374) | 10 | 3 o más pantallas, cada una rotulada con su caso de uso, sin pantallas huérfanas | 3 pantallas, con 1 sin rotular | Menos de 3 pantallas o sin relación con los casos de uso | Sin prototipo |
| C6 | **Elabora listas de chequeo para validación de la documentación de análisis** y las aplica (RAP 592374) | 10 | 15 o más criterios verificables (3 o más de cobertura), cada uno con evidencia y resultado | 15 criterios con 1 o 2 ambiguos (por ejemplo, «está bien hecho») | Menos de 15 criterios, o la mayoría sin evidencia | Sin lista o sin diligenciar |
| C7 | **Realiza mejoras a la documentación de análisis** de acuerdo con los resultados de la evaluación (RAP 592374) | 10 | 5 o más hallazgos con antes y después, corregidos y verificados de nuevo | 3 o 4 hallazgos corregidos y verificados | Hallazgos registrados pero no corregidos, o corregidos sin nueva verificación | Sin registro de ajustes |
| C8 | **Documenta las actividades de análisis a través de un informe** (RAP 592373) | 15 | Informe con requisitos interpretados, enlaces a todos los modelos, resultado de la verificación y ajustes, y matriz de trazabilidad con 90 % o más de los requisitos funcionales trazados hasta clase y entidad | Informe completo y trazabilidad entre 75 % y 89 %, o informe al que le falta una de las cuatro partes | Informe incompleto (faltan dos partes o más), o trazabilidad entre 50 % y 74 % | Sin informe o sin trazabilidad |

## 6. Prueba de desempeño (sábado 28 de noviembre de 2026)

Individual, presencial, **en computador** (equipo del ambiente de formación o portátil propio), 90 minutos, con un caso nuevo que se entrega ese día en formato digital. Se trabaja con una herramienta de diagramación UML y ofimática, y los archivos resultantes se suben al terminar a una tarea de Moodle (Zajuna) que cierra a los 90 minutos. El caso se asigna entre tres variantes de igual estructura y dificultad (9 requisitos funcionales, 4 o 5 actores y 5 errores sembrados de los mismos tipos). El instructor tiene los casos, los errores sembrados y la solución; aquí se publica la estructura y la hoja de puntos.

**Recursos permitidos:** tus propios productos 1 y 2 (copia de tu fork, en solo lectura), la plantilla de caso de uso extendido y el formato de lista de chequeo de la plantilla de entrega, una herramienta de diagramación UML y ofimática. **No se permiten** asistentes de IA, mensajería ni compartir archivos con otros aprendices.

**Entregables:** (1) diagrama de casos de uso, (2) una plantilla extendida de un caso de uso asignado, (3) modelo de dominio, (4) lista de chequeo diligenciada sobre un modelo con errores y registro de hallazgos y correcciones. Se entregan como archivos (imagen o PDF del diagrama, más la fuente editable si la herramienta la genera, y documento para la plantilla y la lista de chequeo) subidos a la tarea de Moodle; el instructor califica esos archivos con la hoja de puntos de la sección 6.1.

**Requisito mínimo de aceptación:** si falta alguno de los cuatro entregables, la prueba tiene tope de 60.

### 6.1 Hoja de puntos (suma 100)

La nota de la prueba es la suma de los puntos de los ítems.

| Criterio | Ítem | Puntos |
|---|---|---|
| **D1 · Elabora diagramas de casos de uso** de acuerdo con los estándares UML, a partir de un informe de requisitos nuevo (RAP 592373) · **25** | Frontera del sistema y 2 actores primarios correctos, fuera de la frontera | 2 |
| | Sistema externo identificado como actor secundario | 2 |
| | Proceso automático con iniciador (temporizador o evento como actor, o caso que lo incluye o extiende) | 2 |
| | Cobertura de requisitos funcionales (1 punto por requisito cubierto, 9 en total) | 9 |
| | Nombres con verbo y objeto, un objetivo por caso (3 puntos con 90 % o más; 1 con 50 % a 89 %) | 3 |
| | Relaciones «include» correctas (lo obligatorio) | 3 |
| | Relaciones «extend» correctas (lo opcional o condicional) o justificación escrita de por qué el caso no la admite | 3 |
| | Sin ciclos, sin casos duplicados ni CRUD fragmentado sin justificación | 1 |
| **D2 · Genera plantillas extendidas de casos de uso** expresando la intención de las acciones (RAP 592373) · **20** | Precondiciones correctas (solo lo asumido antes de iniciar) | 2 |
| | Postcondiciones correctas (estado final del sistema) | 2 |
| | Flujo principal de 6 pasos o más, coherente y en orden actor-sistema | 6 |
| | Al menos un flujo alterno ligado a un paso del principal | 4 |
| | Al menos una excepción con su resultado | 3 |
| | Sin detalle de implementación (3 puntos si no hay ninguno; 1 si hay uno) | 3 |
| **D3 · Representa el negocio en término de clases abstractas** generando un modelo de dominio consistente (RAP 592373) · **20** | Clases conceptuales sin duplicados (6 puntos con 8 o más; 4 con 6 o 7; 2 con 4 o 5) | 6 |
| | Cuatro asociaciones clave con multiplicidad correcta (2 puntos cada una) | 8 |
| | Sin métodos, sin claves foráneas ni atributos técnicos | 3 |
| | Coherencia con la plantilla (cada concepto de la plantilla existe en el dominio) | 3 |
| **D4 · Elabora listas de chequeo para validación de la documentación de análisis y realiza mejoras a la documentación de análisis** de acuerdo con los resultados de la evaluación (RAP 592374) · **25** | Por cada error sembrado, detección citando el criterio incumplido (5 errores × 3 puntos) | 15 |
| | Por cada error, corrección correcta (5 errores × 2 puntos) | 10 |
| | Penalización: por cada elemento correcto marcado como error | -2 (máximo -6) |
| **D5 · Evalúa el informe de análisis** teniendo en cuenta la calidad de los artefactos generados y la respuesta al cumplimiento de requisitos (RAP 592374) · **10** | Casos de uso, plantilla y dominio coinciden entre sí (sin contradicciones de nombres o conceptos) | 6 |
| | Los cuatro entregables subidos a la tarea de Moodle dentro de los 90 minutos (sin contar el tiempo de la verificación oral) | 4 |

Niveles de referencia (solo orientan; el puntaje sale de la hoja): en D4, detectar 5 de 5 errores con su corrección ronda el 100 %; 4 de 5, el 80 %; 2 o 3, entre el 40 % y el 60 %; 1 o ninguno, menos del 20 %. Las penalizaciones restan sobre la suma. El puntaje de un criterio no baja de 0.

Se descuenta por faltas de integridad: usar herramientas no permitidas durante la prueba anula la prueba (nota 0 en este entregable).

## 7. Criterios de arbitraje (cómo decide el instructor en los puntos discutibles)

Estos son los puntos en que los expertos suelen discrepar. Para esta guía se fija una regla única, de modo que la calificación sea la misma sin importar quién califique:

1. **CRUD como casos de uso.** Un caso de uso es un objetivo del actor, no una operación técnica. «Gestionar clientes» es un caso; Crear, Leer, Actualizar y Eliminar se agrupan, salvo que una operación tenga reglas, actor o flujo propios que justifiquen separarla (por ejemplo, «Dar de baja cliente» con aprobación del coordinador). Separar sin esa justificación baja un nivel en C4.
2. **«include» y «extend».** «include» es lo que **siempre** ocurre dentro del caso base; la flecha va del caso base hacia el caso incluido. «extend» es lo **opcional o condicional**, con un punto de extensión indicado; la flecha va del caso que extiende hacia el caso base. Se califica primero el significado y después la dirección: la relación correcta con la flecha invertida baja un nivel en la justificación de relaciones del C4, no en el resto del criterio.
3. **Precondición o flujo alterno.** Si la comprueba el sistema durante el caso (por ejemplo, el saldo), va como flujo alterno o excepción. Es precondición solo lo que el sistema asume cierto **antes** de que empiece el caso (por ejemplo, que el usuario inició sesión).
4. **Flujo alterno o excepción.** Un **flujo alterno** es una variante válida del negocio que continúa o termina el caso con un resultado exitoso distinto (por ejemplo, el horario elegido ya fue tomado y el sistema propone otros). Una **excepción** es una falla o violación de una regla que impide lograr el objetivo (por ejemplo, el pago es rechazado o vence el tiempo de retención). Si la clasificación es defendible con este criterio, se acepta.
5. **Dominio o MER.** El modelo de dominio muestra conceptos del negocio sin métodos ni claves; el MER muestra los datos que se guardan con identificadores, cardinalidad y participación. No se aceptan atributos tipo clave foránea en el dominio ni métodos en el MER. En el MER, una relación N:M se muestra resuelta con una entidad asociativa cuando lleva atributos propios (fecha, nota); la participación total o parcial se indica con la notación que declaraste.
6. **Iniciador de un caso de uso.** Todo caso de uso tiene un iniciador: un actor asociado a él, o un caso base que lo incluye o lo extiende. Un sistema externo que solo recibe información (por ejemplo, un servicio de mensajería) es actor secundario y no cuenta como iniciador.
7. **Metodología única o híbrida.** Se acepta una combinación solo si se declara y se justifica con las características del proyecto. Una mezcla sin justificar se califica como no seleccionada.
8. **Cantidades mínimas.** Se cuentan los elementos correctos y distintos: casos duplicados, clases repetidas con otro nombre o pantallas sin rótulo no suman.
9. **Fuente editable.** Un diagrama solo en imagen no cuenta como entregado en el artefacto (f del producto 1); se requiere la fuente editable.
10. **Detalle de implementación en el análisis.** Nombres de tablas, SQL, frameworks o endpoints dentro de plantillas o del dominio se descuentan en C2 y C3 (y en D2 y D3 de la prueba).

## 8. Autoevaluación antes de entregar

- [ ] Cada artefacto obligatorio de mi producto existe y alcanza la cantidad mínima.
- [ ] Mi README dice qué reutilicé, qué ajusté y trae los enlaces a los commits.
- [ ] Tengo commits míos en al menos 2 días distintos entre la publicación de la guía (sábado 21 de noviembre) y la entrega.
- [ ] Puedo explicar cada decisión de mis diagramas sin mirar el archivo.
- [ ] Mi producto 2 incluye el informe de análisis y la matriz de trazabilidad.
- [ ] Hice `push` y verifiqué en GitHub que las dos carpetas quedaron publicadas antes del sábado 28 de noviembre de 2026 a las 6:00 a. m.
