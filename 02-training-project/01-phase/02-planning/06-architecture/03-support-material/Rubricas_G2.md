# Rúbricas y condiciones de evaluación · Guía 2: Diseño de la solución (v2)

**Ficha 3413974 · Análisis y Desarrollo de Software · Competencia 220501095 (Modelado de los artefactos del software)**  
**RAP de la guía:** 593103 (01) Elaborar los artefactos de diseño · 593101 (02) Estructurar el modelo de datos · 593102 (04) Verificar los entregables de la fase de diseño. El RAP 03 (interfaz gráfica) no se desarrolla en esta guía.

Este documento es público: lo puedes leer desde el primer día. Los criterios de evaluación están tomados **literalmente** del programa de formación; los niveles están redactados para que se pueda comprobar con el archivo en la mano.

---

## 1. Reglas generales

### 1.1 Los cuatro entregables

| # | Entregable | Peso | Fecha | Forma |
|---|---|---|---|---|
| 1 | Prueba de conocimiento | 25 % | Sábado 28 de noviembre de 2026 | Presencial, individual, 40 min, cuestionario de 20 preguntas (14 del banco y 6 inéditas) en la plataforma Moodle (Zajuna), con calificación automática |
| 2 | Producto 1: artefactos de diseño | 25 % | Hasta el sábado 28 de noviembre de 2026 a las 6:00 a. m. | Carpeta `G2-P1-<nombre-corto>/` en tu fork |
| 3 | Producto 2: modelo de datos y verificación | 25 % | Hasta el sábado 28 de noviembre de 2026 a las 6:00 a. m. | Carpeta `G2-P2-<nombre-corto>/` en tu fork |
| 4 | Prueba de desempeño | 25 % | Sábado 28 de noviembre de 2026 | Presencial, individual, en computador, aproximadamente 90 min; los archivos se entregan en una tarea de Moodle (Zajuna) |

**Nota de la guía = promedio de los cuatro entregables. Se aprueba con 70/100 o más.** No hay más entregables ni puntos extra.

Aclaraciones de diseño de la evaluación:

- **No hay nota mínima por entregable**: cuenta el promedio. La excepción es la infracción de recursos en una prueba (0 en esa prueba), que sí puede dejar el promedio por debajo de 70.
- **La escala por saltos (100, 70, 40, 0) es deliberada**: se califica con el archivo en la mano y sin valores intermedios discutibles. Un producto con todos los criterios en nivel 70 obtiene 70, que es la nota mínima de aprobación; cualquier tope por debajo de 70 (secciones 1.4 y 1.5) hace que ese producto quede por debajo del mínimo.
- **Si el promedio no llega a 70**, el resultado no se registra como «No aprobado» en la plataforma de gestión académica: el resultado de aprendizaje queda por evaluar y se acuerda un plan de mejoramiento con fecha. El instructor lo comunica por la plataforma de formación.
- **Alcance frente a la guía 1:** el modelo conceptual (B1) se evalúa aquí por su **ajuste al tipo de base de datos y a las especificaciones del análisis**, y su transformación a lógico, normalización, diccionario, script y seguridad. Si ya fue calificado en la guía 1 como parte del análisis, aquí se reutiliza y se ajusta, y no se califica dos veces el mismo hecho: se evalúa lo que cambiaste.

### 1.2 Cómo se calcula un producto

Cada rúbrica tiene criterios ponderados que suman 100. En cada criterio se ubica el trabajo en uno de cuatro niveles: **100**, **70**, **40** o **0** (no hay valores intermedios). La nota del producto es la suma de `peso × nivel / 100` de todos los criterios. Después se aplican los topes de las secciones 1.4 y 1.5, si corresponden.

**Regla de resolución cuando el trabajo cumple solo una parte de un nivel (aplica a todas las tablas de niveles):**

1. El nivel 100 exige **todas** las condiciones de su celda. Si no las cumples todas, bajas a 70.
2. En los niveles 70 y 40, las condiciones unidas por «o» son **defectos que ubican el trabajo en ese nivel**: basta uno de ellos. No hay que cumplirlos todos.
3. Si tienes defectos de dos niveles distintos, **rige el nivel más bajo**. Si el trabajo no cumple ni siquiera una condición del 40, es 0.
4. Los porcentajes se cuentan sobre lo que **debía** existir. Por ejemplo, «trazabilidad en ≥ 70 % de las clases» se cuenta sobre todas tus clases; con menos de 70 % no llegas al 70 y bajas a 40.
5. «Responsabilidades de áreas distintas» significa que una misma clase mezcla dos o más de estas áreas: interfaz de usuario, lógica de negocio, persistencia, cobro, notificaciones y reportes.

### 1.3 Ruta y fecha de entrega

Entrega con **commit y push en tu fork**, en:

```
02-training-project/01-phase/02-planning/06-architecture/04-evidence-learning/G2-P1-<nombre-corto>/
02-training-project/01-phase/02-planning/06-architecture/04-evidence-learning/G2-P2-<nombre-corto>/
```

Cada carpeta lleva un `README.md` hecho con **Plantilla_README_Entrega_G2.md**. Cuenta lo que esté publicado en tu fork con fecha de push hasta el sábado 28 de noviembre de 2026 a las 6:00 a. m. (antes de iniciar la sesión presencial). Lo que se suba después, o por otro medio (correo, mensajería, USB), no se evalúa. El instructor toma una **copia de tu fork a esa hora límite** y evalúa lo que contenga esa copia; no reescribas el historial (sin `git push --force` sobre las ramas de la entrega).

### 1.4 Requisitos mínimos de aceptación (RMA) y tope de 60

Cada producto tiene una lista de **artefactos obligatorios** con una cantidad mínima (secciones 2 y 3). **Si falta un artefacto obligatorio, el producto tiene un tope de 60 puntos**, sin importar lo bien hecho que esté el resto. Un artefacto cuenta como entregado solo si es tuyo, corresponde a tu caso y se puede abrir. Quedarse corto en la cantidad mínima de un artefacto que sí entregaste **no activa el tope**: se refleja en el nivel del criterio (el nivel 100 exige alcanzar el mínimo).

### 1.5 Control de autoría (individual y verificable)

Los productos se trabajan entre la publicación de la guía (sábado 21 de noviembre) y la entrega (el sábado 28 de noviembre de 2026 a las 6:00 a. m.). Esa misma semana trabajas también la guía 1, cuyos productos vencen el mismo 28 de noviembre: planifica tu tiempo para las dos. El insumo (el informe de análisis o SRS de tu proyecto) puede ser del equipo; **los productos de diseño son individuales**. Se verifican así:

| Control | Qué se revisa | Si no se cumple |
|---|---|---|
| Historial de commits | Commits con **tu cuenta** en al menos **2 días distintos** entre la publicación de la guía (sábado 21 de noviembre de 2026) y la entrega (el sábado 28 de noviembre de 2026 a las 6:00 a. m., hora de Colombia), con mensajes que describan el cambio. El día se verifica con la fecha en que se **publicó** (push) en GitHub, no solo con la fecha escrita en el commit, y cada día debe traer **cambios sustantivos** en archivos de las carpetas G2-P1 o G2-P2 (no cuentan espacios, saltos de línea ni renombrados sin contenido). Está prohibido reescribir el historial con force-push | 1 solo día: tope 70 · ningún día con cambios sustantivos propios: tope 60 |
| README de la entrega | Plantilla completa: qué reutilizaste y de qué ruta, qué ajustaste, qué hiciste nuevo, origen del insumo, uso de IA declarado y declaración de autoría | README ausente o incompleto: tope 70 |
| Verificación oral de autoría (sábado 28 de noviembre o fecha de respaldo que el instructor comunique por la plataforma de formación) | **Una pregunta por producto** (P1 y P2) sobre **tus** archivos, con tu entrega abierta: por qué una clase está relacionada de cierto modo, cómo llegaste de una tabla a 3FN, qué cambia en tu diseño si cambia un requisito, dónde está en tu diagrama un patrón, etc. Apruebas el producto si **explicas con tus palabras** qué hiciste, por qué y qué cambiarías, sin leer un texto preparado. Si hay duda, se hace una tercera pregunta sobre ese producto. Cada oral dura unos 2 minutos y se hace aparte, en un bloque propio | Si no sabes explicar **tu propio** producto (ni la pregunta ni la tercera): ese producto queda con **tope 50**. Antes de aplicarlo puedes pedir una segunda oportunidad el mismo día |
| Contenido idéntico a otro aprendiz | **Artefactos de diseño** (P1 o P2) iguales o casi iguales entre aprendices (mismo texto, mismas clases o tablas, mismos errores) **sin declararlo**. Compartir el insumo del equipo (SRS, casos de uso, MER de equipo) **no** activa este tope si lo declaras en el README | Tope 40 en el producto de cada uno, **después de una aclaración con derecho a réplica** (el instructor te muestra la coincidencia y puedes explicar tu proceso con tu historial) |
| Material del instructor como propio | `design-software/` de la semana 12, el laboratorio MER y su `ejemplo/` de la semana 20, y plantillas del framework sin diligenciar | Ese artefacto se cuenta como **no entregado** (activa el tope 60) |
| Commits masivos | Un commit que agrega decenas de archivos que no son tuyos (por ejemplo, el framework completo) | No cuenta como aporte propio; se revisa el resto del historial |

**Insumo del equipo, producto individual.** Si tu SRS, tus casos de uso o tu MER son del equipo, declara en el README qué subconjunto de requisitos (≥ 8 RF) diseñas tú. Si dos compañeros diseñan el mismo subconjunto, sus artefactos deben diferir en decisiones propias (clases, patrón elegido, tablas, política de seguridad) y cada uno explica las suyas en la oral.

Si se aplican varios topes, **rige el más bajo** (no se suman). El uso de IA está permitido **solo si lo declaras** en el README (para qué la usaste y qué cambiaste después); lo que no puedas explicar en la verificación oral cuenta como no propio.

### 1.6 De dónde sale el insumo (reutilizar y ajustar)

Elige **un solo caso** y úsalo en los dos productos: el de tu fork donde tengas un informe de análisis o SRS (por ejemplo, el de tu proyecto del año). Reutiliza lo que ya hiciste (SRS, casos de uso, diagrama de clases, mockups, MER, BPMN, scripts SQL de las carpetas `03-execute/NN-week/02-optional-activity` y de `04-evidence-learning`) y **ajústalo al estándar de estas rúbricas**. Debes declarar en el README qué reutilizaste y qué ajustaste.

**Si no tienes alguno de los artefactos de partida** (lo más común: MER propio, diagrama de clases, diccionario de datos), lo construyes desde tu SRS usando **Plantillas_G2.md** (plantillas con ejemplos pequeños de otro dominio). En la práctica casi nadie tiene diagrama de componentes, vista de despliegue ni normalización practicados: son trabajo nuevo y se declaran como «nuevo» en el README. Reutilizar es **parcial**: de los 14 artefactos obligatorios, espera tener antecedente en unos pocos.

**Si no tienes un SRS utilizable**, avisa al instructor a más tardar el miércoles 25 de noviembre y entrega en `G2-P1-.../00-insumo-analisis.md` un informe de análisis mínimo de tu proyecto: al menos 8 requisitos funcionales con identificador, 3 no funcionales, los actores y 4 casos de uso descritos. Reconstruirlo es tu trabajo y no resta puntos, pero no cuenta como parte de la nota.

### 1.7 Presupuesto de tiempo (cada producto está dimensionado para unas 3 horas)

El presupuesto supone que **reutilizas** lo que ya tienes y **no incluye holgura**. Si partes de cero o tu SRS está incompleto, no alcanza: empieza el mismo sábado 21 o el domingo 22, haz primero lo nuevo para casi todos (componentes, despliegue, normalización, patrones, políticas) y deja un commit de avance a mitad de semana (te sirve también para el control de autoría). Recuerda que la guía 1 también vence el 28 de noviembre: el presupuesto de esta guía no la incluye. Los artefactos sin antecedente toman más que lo presupuestado; reparte esa carga en el trabajo independiente de la semana.

| Producto 1 | Min | Producto 2 | Min |
|---|---|---|---|
| A1 ficha de interpretación y plataforma | 15 | B1 modelo conceptual (ajuste de tu MER) | 15 |
| A2 diagrama de clases (ajuste del tuyo) | 40 | B2 modelo lógico y cuadro de reglas | 20 |
| A3 dos patrones aplicados | 20 | B3 normalización de 2 tablas | 25 |
| A4 arquitectura y tabla de RNF | 25 | B4 diccionario (sale del modelo y del script) | 15 |
| A5 vista de componentes | 25 | B5 script SQL y evidencia de ejecución | 20 |
| A6 vista de despliegue | 15 | B6 políticas y matriz de roles | 25 |
| A7 matriz de trazabilidad | 15 | B7 lista de 20 criterios aplicada a P1 y P2 | 35 |
| README y exportación de imágenes | 10 | README | 10 |
| **Total** | **165** | **Total** | **165** |

---

## 2. Producto 1: artefactos de diseño (RAP 593103 · 25 %)

### 2.1 Artefactos obligatorios y mínimos

| Artefacto | Mínimo exigible |
|---|---|
| A1. Ficha de interpretación del análisis y plataforma tecnológica | ≥ 8 requisitos funcionales y ≥ 3 no funcionales identificados con su ID; lenguaje orientado a objetos y tipo de base de datos con justificación |
| A2. Diagrama de clases | ≥ 8 clases con atributos tipados y métodos; multiplicidades en todas las asociaciones; ≥ 1 composición o agregación; ≥ 1 herencia o interfaz (si es interfaz, con ≥ 2 implementaciones, para evidenciar el polimorfismo) |
| A3. Patrones de diseño | ≥ 2 patrones GoF distintos (estrategia, comando, iterador, singleton, fábrica abstracta o fachada) aplicados sobre clases del diagrama; **al menos uno distinto de singleton** |
| A4. Arquitectura del software | Estilo definido con diagrama y tabla de decisiones: ≥ 3 requisitos no funcionales con la decisión que los atiende |
| A5. Vista de componentes | ≥ 5 componentes con interfaces y dependencias |
| A6. Vista de despliegue | ≥ 3 nodos con los artefactos que alojan y los protocolos de comunicación |
| A7. Matriz de trazabilidad (análisis → diseño) | Cada requisito funcional mapeado a clase(s) y componente(s) |

Formato: fuente editable (draw.io, PlantUML, Mermaid o equivalente) **y** imagen exportada (PNG o SVG) de cada diagrama, más un `.md` con las tablas. El `README.md` de la carpeta sirve de índice del conjunto (enlace a A1-A7 y una frase de la decisión principal de cada uno); no se pide un documento adicional. Las plantillas están en **Plantillas_G2.md**.

### 2.2 Rúbrica

| # | Criterio del programa (texto literal) | Aspecto que se observa | Peso |
|---|---|---|---|
| 1 | Interpreta el informe de análisis identificando las características del software a diseñar. Define las plataformas tecnológicas a emplear en el desarrollo de acuerdo con las condiciones del software a desarrollar. | Ficha de interpretación y plataforma | 10 |
| 2 | Define los entregables de diseño siguiendo los conceptos y principios de orientación a objetos. Crea el diagrama de clases de acuerdo con los requisitos, aplicando buenas prácticas de diseño orientado a objetos. | Diagrama de clases | 25 |
| 3 | Incorpora patrones de diseño propendiendo en mejores prácticas para la codificación y mantenibilidad del software. | Patrones aplicados | 15 |
| 4 | Define la arquitectura del software dando cumplimiento a los requisitos funcionales y no funcionales. | Arquitectura y decisiones frente a los RNF | 15 |
| 5 | Crea la vista de componentes para visualizar el software en fases avanzadas del ciclo de vida. | Vista de componentes | 15 |
| 6 | Crea la vista de despliegue del software para determinar condiciones de la implantación de la solución informática. | Vista de despliegue | 10 |
| 7 | Interpreta el informe de análisis identificando las características del software a diseñar. | Trazabilidad entre el análisis y el diseño y calidad documental del conjunto (matriz A7 y README) | 10 |

| Criterio | 100 | 70 | 40 | 0 |
|---|---|---|---|---|
| **1. Interpretación y plataforma** | Ficha con ≥ 8 RF y ≥ 3 RNF con su ID del SRS, actores y restricciones; lenguaje y tipo de base de datos justificados con ≥ 3 condiciones concretas del caso; lista de entregables de diseño | Identifica los requisitos pero con < 8 RF o < 3 RNF, o justifica la plataforma con 1 o 2 condiciones | Copia la lista de requisitos sin seleccionar o elige plataforma sin justificar | No hay ficha |
| **2. Entregables y diagrama de clases** | ≥ 8 clases con atributos tipados, visibilidad y métodos con parámetros; multiplicidades en todas las asociaciones; ≥ 1 composición o agregación y ≥ 1 herencia o interfaz (con ≥ 2 implementaciones si es interfaz) usadas con criterio («es un» / «es parte de»); ninguna clase reúne responsabilidades de áreas distintas (regla 5 de la sección 1.2); cada clase traza a ≥ 1 requisito | 6 o 7 clases, o faltan tipos o visibilidad en hasta el 25 % de los atributos, o 1 o 2 errores de relación o multiplicidad; trazabilidad en ≥ 70 % de las clases | 4 o 5 clases, o solo atributos sin métodos, o asociaciones sin multiplicidad, o no se puede rastrear al SRS | Ausente, menos de 4 clases o corresponde a otro caso |
| **3. Patrones** | ≥ 2 patrones distintos (al menos uno diferente de singleton) aplicados en el diagrama; para cada uno: problema concreto del caso, por qué ese patrón, qué clases lo implementan y qué mejora en mantenibilidad | 2 patrones aplicados pero con justificación genérica (definición de libro), o 1 patrón con justificación sólida | Patrones nombrados sin reflejarse en el diagrama, o aplicados de forma incorrecta | Ninguno |
| **4. Arquitectura** | Estilo definido con diagrama; tabla con ≥ 3 RNF y la decisión de cada uno; estilo de comunicación entre componentes; una alternativa descartada con su razón; sin capas saltadas | Estilo, diagrama y tabla con 2 RNF, sin alternativa descartada, o un salto de capa justificado | Nombra el estilo sin diagrama ni relación con los RNF | Ausente |
| **5. Vista de componentes** | ≥ 5 componentes con interfaces provistas y requeridas nombradas; dependencias con la dirección correcta; cada componente indica las clases que agrupa y los requisitos que atiende | ≥ 5 componentes con dependencias, pero con interfaces faltantes o mapeo parcial | 3 o 4 componentes, o cajas sin interfaces | Ausente, menos de 3 componentes, o entrega un despliegue en su lugar |
| **6. Vista de despliegue** | ≥ 3 nodos (por ejemplo cliente, aplicación, datos) con artefactos alojados; protocolo o puerto en cada conexión; ≥ 1 decisión que responde a un RNF (disponibilidad, seguridad de red o desempeño) | ≥ 3 nodos con artefactos, sin protocolos o sin relación con un RNF | 2 nodos, o solo un archivo de contenedores copiado sin diagrama | Ausente |
| **7. Trazabilidad y calidad documental** | Matriz que cubre el 100 % de los RF y ninguna clase queda sin requisito; README completo (que sirve de índice); diagramas legibles con título y fuente editable | Cubre ≥ 80 % de los RF, hasta 2 clases sin requisito | Matriz parcial (< 80 %) o diagramas ilegibles | Sin matriz |

---

## 3. Producto 2: modelo de datos y verificación (RAP 593101 y 593102 · 25 %)

### 3.1 Artefactos obligatorios y mínimos

| Artefacto | Mínimo exigible |
|---|---|
| B1. Modelo conceptual | ≥ 6 entidades con atributos e identificador; cardinalidad mínima y máxima en cada relación; ≥ 1 relación N:M y ≥ 1 relación 1:N |
| B2. Modelo lógico relacional | ≥ 8 tablas con PK y FK; tabla intermedia para cada N:M; cuadro de reglas de transformación aplicadas (incluye el atributo multivaluado en tabla aparte y la relación recursiva o la generalización, si tu caso las tiene) |
| B3. Evidencia de normalización | ≥ 2 tablas de tu propio modelo llevadas paso a paso de 0FN a 3FN, con las dependencias funcionales escritas; entre las dos, una con **llave compuesta** (dependencia parcial, 2FN) y una con **dependencia transitiva** (3FN) |
| B4. Diccionario de datos | 100 % de las tablas y columnas: nombre, tipo, tamaño, restricciones y dominio |
| B5. Script SQL de creación | Crea el 100 % de las tablas con PK, FK, `NOT NULL`, `UNIQUE` y ≥ 3 `CHECK`; acción referencial (`ON DELETE`/`ON UPDATE`) declarada en ≥ 3 FK; evidencia de ejecución sin errores (captura o log) |
| B6. Políticas de seguridad de los datos | ≥ 6 políticas (≥ 2 por cada principio: integridad, confidencialidad, disponibilidad); matriz de ≥ 3 roles contra tablas; contraseñas con **hash y sal** (no «cifradas») y datos en tránsito por TLS |
| B7. Lista de chequeo diligenciada | ≥ 20 criterios de aceptación medibles, aplicados a los artefactos de **ambos** productos, con evidencia por criterio |

### 3.2 Rúbrica

| # | Criterio del programa (texto literal) | Aspecto que se observa | Peso |
|---|---|---|---|
| 1 | Genera el modelo conceptual de acuerdo con el tipo de base de datos seleccionada y las especificaciones del análisis. | Modelo conceptual | 15 |
| 2 | Genera el modelo lógico de acuerdo con la técnica seleccionada. | Modelo lógico relacional | 15 |
| 3 | Normaliza el modelo lógico de acuerdo con el tipo de base de datos. | Evidencia de normalización | 15 |
| 4 | Crea el diccionario de datos de acuerdo con el modelo lógico. | Diccionario de datos | 10 |
| 5 | Define políticas de seguridad para garantizar integridad, confidencialidad y disponibilidad de los datos. | Políticas y matriz de roles | 15 |
| 6 | Elabora listas de chequeo para verificación de la documentación de diseño. Evalúa los artefactos de diseño teniendo en cuenta el cumplimiento de requisitos y la calidad de los entregables generados. | Lista de chequeo diligenciada sobre P1 y P2 | 20 |
| 7 | Genera el modelo lógico de acuerdo con la técnica seleccionada. Crea el diccionario de datos de acuerdo con el modelo lógico. | Script SQL y consistencia entre modelo, diccionario y script | 10 |

| Criterio | 100 | 70 | 40 | 0 |
|---|---|---|---|---|
| **1. Modelo conceptual** | ≥ 6 entidades con atributos e identificador; cardinalidad mín/máx y verbo en cada relación; ≥ 1 N:M y ≥ 1 1:N; tipo de base de datos justificado; cada entidad traza a un requisito; notación consistente | ≥ 6 entidades con cardinalidades parciales o hasta 2 entidades sin atributos | 4 o 5 entidades, o sin cardinalidades | Ausente, menos de 4 entidades o igual al ejercicio del instructor |
| **2. Modelo lógico** | Todas las entidades transformadas con cuadro de reglas aplicadas (N:M a tabla intermedia, 1:N a FK, 1:1 a FK única); PK y FK en todas las tablas; ≥ 8 tablas; diagrama lógico | 1 o 2 errores (una FK ausente, una N:M sin resolver) | 3 o más errores, o tablas sin llaves | Ausente |
| **3. Normalización** | Evidencia paso a paso 0FN → 1FN → 2FN → 3FN de ≥ 2 tablas propias (una con llave compuesta y una con dependencia transitiva), con dependencias funcionales escritas y estado antes y después; el modelo final queda sin dependencias parciales ni transitivas; si desnormalizas algo, lo justificas | 1 tabla con los tres pasos, o pasos sin dependencias funcionales escritas | Afirma «está en 3FN» sin evidencia, o solo aplica 1FN | Ausente |
| **4. Diccionario de datos** | 100 % de tablas y columnas con nombre (estándar de nombrado declarado), tipo, tamaño, restricciones y dominio; coincide con el modelo y el script | ≥ 80 % de las columnas completas | ≥ 50 % de las columnas, o sin dominio | < 50 % o ausente |
| **5. Políticas de seguridad** | ≥ 2 políticas por principio (≥ 6), con medida concreta sobre el caso (columnas que se cifran, qué rol ve qué); matriz de ≥ 3 roles contra tablas con permisos; tratamiento de datos personales (finalidad y autorización); contraseñas con hash y sal y tráfico por TLS; respaldo con frecuencia y prueba de restauración | ≥ 1 política por principio y matriz de roles, sin respaldo o sin datos personales detallados | Políticas genéricas copiadas de una definición, sin vincularlas a tablas del caso | Ausente |
| **6. Lista de chequeo y evaluación** | ≥ 20 criterios medibles agrupados por artefacto (clases, patrones, arquitectura, componentes, despliegue, modelo, diccionario, seguridad), con severidad (crítico, mayor, menor); diligenciada sobre los artefactos de P1 y P2 con evidencia por criterio; ≥ 3 hallazgos o ajustes registrados con su corrección | ≥ 15 criterios, o sin severidad, o sin evidencia en más del 30 % de los criterios | < 15 criterios, criterios vagos («el diseño es bueno») o sin diligenciar | Ausente |
| **7. Script SQL y consistencia** | Ejecuta sin errores en un motor real (evidencia); crea el 100 % de las tablas con PK, FK (con `ON DELETE` en ≥ 3), `NOT NULL`, `UNIQUE` y ≥ 3 `CHECK`; nombres y tipos idénticos al diccionario | Ejecuta, ≥ 90 % de las tablas y hasta 3 diferencias con el diccionario | Ejecuta con errores corregibles, o faltan FK | Ausente o no ejecuta |

**Sobre los criterios 7 de los dos productos.** El programa no trae un criterio específico para la trazabilidad documental (P1-7) ni para el script SQL (P2-7). Por eso se usan oraciones literales del programa que son lo más cercano (la interpretación del análisis; el modelo lógico y el diccionario) y se declara que **son evidencia derivada**: lo que se observa en ellos es la consistencia entre artefactos (matriz y README; modelo, diccionario y script), no una competencia nueva. Para P2-6 y la prueba de desempeño: en el Producto 2 se evalúa el **diseño de tu instrumento y su aplicación a tu propio trabajo**; en la prueba, **su uso sobre trabajo ajeno con defectos**. Son desempeños distintos y no se califican dos veces el mismo hecho.

---

## 4. Prueba de desempeño (RAP 593102 · 25 %)

**Qué es.** El sábado 28 de noviembre de 2026, de forma presencial e individual y **en computador** (el del ambiente de formación o tu portátil), recibes como archivos digitales un paquete de artefactos de diseño de un caso que no conoces, con su informe de análisis. Tienes aproximadamente 90 minutos para **verificarlo** y **corregirlo** con tu herramienta de diagramación, tu ofimática y un motor SQL. Recibes una **plantilla de registro de hallazgos ya formateada** (archivo digital): ve anotando cada hallazgo en el momento en que lo encuentras. Hay dos versiones del caso, repartidas por turno, para que no se copien resultados.

**Qué se te pide entregar (los tres son obligatorios):**
1. Una **lista de chequeo** diligenciada.
2. Un **registro de hallazgos** con trazabilidad al informe de análisis.
3. Los artefactos **corregidos**: modelo de datos normalizado y diagrama de clases ajustado.

**Recursos permitidos:** computador (del ambiente o portátil tuyo) con herramienta de diagramación, ofimática (editor de texto y hoja de cálculo) y un motor SQL relacional (SQLite, MySQL o PostgreSQL); apuntes digitales, esta guía, el banco de preguntas, Plantillas_G2.md y **tu propia** lista de chequeo del Producto 2 (copia sin conexión; no se abre tu fork durante la prueba). Solo se permite la conexión a Moodle (Zajuna), para descargar el paquete y subir tus archivos. No se permite internet para otra cosa, mensajería, asistentes de IA ni ayuda de otras personas. Una infracción anula la prueba (0).

**Cómo entregas.** Al terminar los 90 minutos subes **un solo archivo comprimido** (`G2-desempeno-<variante>.zip`) a la tarea de Moodle de la prueba, con los tres productos: la lista de chequeo y el registro de hallazgos (`.md`, `.xlsx`, `.ods` o `.docx`), el modelo de datos normalizado (script `.sql` o tabla en `.md`) y el diagrama de clases ajustado (fuente editable e imagen exportada). Plantillas_G2.md, sección 14, da el detalle. Lo que no esté en la tarea de Moodle al cierre no se evalúa.

**Topes:** si falta uno de los tres entregables, la prueba tiene tope de 60.

| # | Criterio del programa (texto literal) | Aspecto que se observa | Peso |
|---|---|---|---|
| 1 | Elabora listas de chequeo para verificación de la documentación de diseño. | Lista de chequeo | 25 |
| 2 | Evalúa los artefactos de diseño teniendo en cuenta el cumplimiento de requisitos y la calidad de los entregables generados. | Cumplimiento de requisitos y trazabilidad | 20 |
| 3 | Normaliza el modelo lógico de acuerdo con el tipo de base de datos. | Normalización del modelo de datos | 20 |
| 4 | Evalúa los artefactos de diseño teniendo en cuenta el cumplimiento de requisitos y la calidad de los entregables generados. | Calidad del diagrama de clases: bajo acoplamiento y alta cohesión (práctica de diseño del programa) | 15 |
| 5 | Evalúa los artefactos de diseño teniendo en cuenta el cumplimiento de requisitos y la calidad de los entregables generados. | Arquitectura y vistas frente a los requisitos no funcionales | 10 |
| 6 | Elabora listas de chequeo para verificación de la documentación de diseño. | Registro de hallazgos completo y priorizado | 10 |

| Criterio | 100 | 70 | 40 | 0 |
|---|---|---|---|---|
| **1. Lista de chequeo** | ≥ 15 criterios medibles, agrupados por artefacto y con severidad; diligenciada completa con referencia al elemento del paquete que sustenta cada marca | 10 a 14 criterios, o diligenciada en ≥ 80 % | Menos de 10 criterios o criterios genéricos | No entrega |
| **2. Cumplimiento de requisitos y trazabilidad** | Matriz completa requisito ↔ artefacto; detecta ≥ 90 % de los requisitos sin cubrir y de los artefactos sin requisito que traiga el paquete | Detecta entre el 60 % y el 89 % | Detecta entre el 30 % y el 59 % | Menos del 30 % |
| **3. Normalización** | Identifica todas las violaciones de 1FN, 2FN y 3FN presentes; escribe las dependencias funcionales; entrega el modelo final con PK, FK, tipos y tablas intermedias correctos | Identifica dos de las tres categorías, o el modelo final tiene un error de llaves o tipos | Identifica una categoría, o el modelo final tiene varios errores | No normaliza |
| **4. Diagrama de clases** | Detecta y corrige los problemas de cohesión, de acoplamiento y de relaciones mal usadas que traiga el paquete, y entrega el diagrama ajustado | Detecta y corrige dos de las tres categorías | Detecta una categoría | No detecta ni corrige |
| **5. Arquitectura y vistas** | Detecta los incumplimientos de capas y de requisitos no funcionales en las vistas y propone un ajuste concreto para cada uno | Detecta y ajusta uno | Detecta sin proponer ajuste | No detecta |
| **6. Registro de hallazgos** | Cada hallazgo con ID, artefacto, descripción, requisito afectado, severidad coherente y corrección propuesta | Faltan campos en hasta el 30 % de los hallazgos | Faltan campos en la mayoría o no hay severidad | Sin registro |

El instructor califica con una **solución de referencia privada** del caso de cada turno. Los defectos que el paquete trae son los que el instructor sembró; si reportas otra brecha real del paquete, se acepta como válida pero no sustituye a las sembradas. La verificación oral de autoría de los Productos 1 y 2 se hace en un bloque aparte (sección 1.5).

---

## 5. Prueba de conocimiento (RAP 593103, 593101 y 593102 · 25 %)

Se presenta como **cuestionario en la plataforma Moodle (Zajuna)**, de forma presencial en el ambiente de formación y con **calificación automática**: la nota sobre 100 que calcula Moodle es la nota de la evidencia de conocimiento (25 % de la guía). 20 preguntas de selección múltiple (una sola respuesta correcta), 5 puntos cada una, 40 minutos, sin recursos. Moodle puede barajar el orden de las opciones. **14 salen del banco** (Banco_Estudio_G2 (.html o .pdf)) **y 6 son inéditas** del nivel avanzado. Las preguntas del banco pueden aparecer con las opciones en otro orden. El examen y su clave no se publican; el banco público es un documento de estudio sin respuestas.

**Tabla de especificaciones de la prueba (preguntas por RAP y criterio del programa que se verifica):**

| RAP | Preguntas del examen (20) | Preguntas del banco (60) | Qué se verifica (criterios del programa, resumidos) |
|---|---|---|---|
| 593103 (01) Artefactos de diseño | 9 | 26 | Interpretar el análisis; conceptos y principios de orientación a objetos (cohesión, acoplamiento, encapsulamiento, polimorfismo, herencia, composición); patrones GoF; arquitectura y sus estilos; vistas de componentes y de despliegue |
| 593101 (02) Modelo de datos | 6 | 20 | Modelo conceptual y lógico; reglas de transformación; normalización (1FN a 3FN); diccionario de datos; políticas de seguridad y tratamiento de datos personales |
| 593102 (04) Verificación del diseño | 5 | 14 | Listas de chequeo y criterios de aceptación; evaluación de artefactos frente a los requisitos; trazabilidad; severidad y priorización de hallazgos |

La proporción del examen sigue la del banco (aprox. 45 % / 30 % / 25 %). Los niveles del examen son: 14 preguntas del banco (de nivel intermedio y avanzado) y 6 inéditas avanzadas.

---

## 6. Resumen de topes

| Situación | Tope |
|---|---|
| Falta un artefacto obligatorio (por producto); no alcanzar la cantidad mínima solo baja el nivel del criterio | 60 |
| Sin ningún día con commits sustantivos propios en el periodo (por fecha de push); force-push sobre la entrega: se revisa el caso | 60 |
| Commits sustantivos en un solo día del periodo | 70 |
| README ausente o incompleto | 70 |
| Artefactos de diseño idénticos a otro aprendiz sin declarar (tras la aclaración con réplica) | 40 |
| No explicas tu propio producto en la verificación oral (la pregunta y la tercera de apoyo) | 50 (ese producto) |
| Falta uno de los tres entregables de la prueba de desempeño | 60 |
| Infracción de los recursos en la prueba de desempeño o de conocimiento | 0 en esa prueba |
