# Rúbricas · Guía 3 · Propuesta técnica (v2)

**Ficha 3413974 · Análisis y Desarrollo de Software · Competencia 220501094** (Estructurar propuesta técnica de servicio de tecnología de la información según requisitos técnicos y normativa)

Este documento es público: lo lee el aprendiz antes de empezar. Dice qué se entrega, con qué mínimos, cómo se califica cada criterio y qué topes se aplican. No contiene respuestas de las pruebas.

## 1. Cómo se calcula la nota de la guía

| Entregable | Peso | Cuándo |
|---|---|---|
| 1. Prueba de conocimiento (cuestionario de 20 preguntas, 40 min, en la plataforma Moodle con calificación automática) | 25 % | Sábado 5 de diciembre de 2026, presencial |
| 2. Producto 1: especificaciones técnicas | 25 % | A más tardar el sábado 5 de diciembre de 2026 a las 6:00 a. m. (antes de iniciar la sesión presencial) |
| 3. Producto 2: propuesta técnica y validación | 25 % | A más tardar el sábado 5 de diciembre de 2026 a las 6:00 a. m. (antes de iniciar la sesión presencial) |
| 4. Prueba de desempeño (individual, en computador, unos 90 min; los archivos se suben a una tarea de Moodle) | 25 % | Sábado 5 de diciembre de 2026, presencial |

- La nota de la guía es el promedio de los cuatro entregables. Se aprueba con **mínimo 70 sobre 100**.
- Si el promedio no llega a 70, el resultado no se registra como «No aprobado» en la plataforma de gestión académica: el resultado de aprendizaje queda por evaluar y se acuerda un plan de mejoramiento con fecha. El instructor lo comunica por la plataforma de formación.
- La **prueba de conocimiento** es un cuestionario de 20 preguntas de selección múltiple en la plataforma Moodle (Zajuna), presencial en el ambiente de formación, con **calificación automática**: cada pregunta vale 5 puntos, no hay puntos negativos y la nota sobre 100 es la nota de la evidencia de conocimiento (25 % de la guía). Las preguntas de cálculo traen cifras fijas y opciones de respuesta; en Moodle el orden de las preguntas y de las opciones cambia por aprendiz.
- No hay más entregables: el material de estudio, el banco de preguntas y esta rúbrica no se califican.
- Los criterios de cada rúbrica salen de la columna «Criterios de evaluación» y de los conocimientos de proceso y del saber de la competencia 220501094 del programa de formación. Cada fila indica la línea del programa que la sostiene.

### Escala por criterio

Cada criterio se califica en solo cuatro niveles: **100, 70, 40 o 0**. La nota del entregable es la suma de (peso del criterio × nivel) / 100. Después se aplican los topes de la sección 2 y se toma el menor.

| Nivel | Significado |
|---|---|
| 100 | Cumple todo lo descrito, con evidencia verificable en el repositorio o en la prueba |
| 70 | Cumple lo esencial; faltan detalles puntuales o hay errores menores que no cambian la conclusión |
| 40 | Cumple de forma parcial: hay intento real, pero faltan elementos que impiden usar el artefacto en una negociación o en un contrato |
| 0 | No está, no se puede verificar o contradice lo pedido |

**Regla de desempate (cómo se asigna el nivel).** El nivel de un criterio es el **más bajo** cuyo requisito cumple el producto en **todos** los elementos listados en ese nivel; si un elemento queda entre dos niveles, se asigna el inferior. El instructor llena una hoja de calificación por criterio con el conteo observado (por ejemplo «referentes completos: 4 de 4; fichas con 6 características medibles: 3 de 4») para que otro instructor llegue a la misma nota. Un artefacto obligatorio que quede por debajo de la mitad de su cantidad mínima activa el tope de 60 de la sección 2.

**Control de exactitud (aplica a todos los criterios de producto y de desempeño).** Las cantidades miden cumplimiento; la exactitud mide calidad. Un **error conceptual material** (norma inexistente o mal aplicada, afirmación legal falsa, ficha que exige una marca, matriz cuyos pesos no suman 100 %, IVA o contingencia mal aplicados, cantidad de licencias sin relación con la variante) limita ese criterio a **70**; **dos o más** errores materiales en el mismo criterio lo limitan a **40**. El instructor verifica por muestreo **3 cifras y 2 citas normativas** de cada producto con una hoja de verificación y comprueba la **coherencia entre el Producto 1, el presupuesto y la propuesta**: una contradicción entre ellos (cifra distinta, componente sin costo) cuenta como un error material en el criterio de la propuesta documentada.

**Definiciones usadas en las tablas.** *Característica medible*: valor, unidad y forma de verificarla. *Vigencia verificada*: enlace a la fuente oficial y fecha de consulta en la fila del referente. *Alternativa real*: producto o servicio que existe en el mercado, con enlace o cotización y fecha de consulta. *TCO parcial*: falta algún rubro (adquisición, licencias, soporte, operación o capacitación) o no cubre los 36 meses. *Contingencia*: suma de los valores esperados de los riesgos que se decide cubrir con reserva.

## 2. Topes y control de autoría (aplican a los productos 1 y 2)

| Situación | Efecto |
|---|---|
| Falta un **artefacto obligatorio** (lista en la sección 3) o está por debajo de la mitad de su cantidad mínima | La nota del producto no pasa de **60** |
| El aprendiz **no puede explicar su propio producto** en la verificación oral (sección 7, con derecho de réplica) | La nota del producto no pasa de **50** |
| Contenido idéntico a otro fork o al material del instructor, sin ajustes y sin declararlo en el README (el trabajo conjunto declarado y diferenciado no es copia) | **0** en ese producto; el instructor revisa el caso con el aprendiz antes de registrar la nota |

### Control de autoría individual (verificable)

1. **Historial de commits.** Los commits de cada producto son de la cuenta del aprendiz y están repartidos en **al menos 2 días calendario distintos** entre la publicación de la guía (sábado 21 de noviembre de 2026) y la entrega (**el sábado 5 de diciembre de 2026 a las 6:00 a. m.**), con mínimo 3 commits por producto. Un único commit grande con todo no cumple. Los productos se trabajan en las semanas del 21 de noviembre al 5 de diciembre; la reunión de validación con la contraparte se hace **antes de la entrega** y el acta no puede estar fechada después de ella.
2. **Commits sospechosos.** Todo commit que agregue 40 o más archivos, o que traiga completo el marco documental del instructor o de otro fork, se revisa: debe existir un ajuste posterior propio sobre esos archivos y el README debe decir qué se ajustó.
3. **README de cada carpeta** (plantilla `Plantilla_README_Entrega_G3.md`): qué reutilizó, qué ajustó o construyó desde cero, enlaces, declaración de autoría y uso de IA declarado.
4. **Variante del escenario.** Cada aprendiz recibe del instructor una variante (A, B, C o D) con parámetros distintos (anexo A). Los números de sus fichas, licencias, costos y ANS deben salir de su variante. Así, dos productos con los mismos números son evidencia de copia.
5. **Verificación oral** el sábado 5 de diciembre (sección 7).

El trabajo en equipo del año se aprovecha: el **caso** puede ser el mismo del equipo, pero cada aprendiz entrega su carpeta con sus propias decisiones, cálculos y redacción.

## 3. Mínimos de aceptación (cantidades exigibles)

### Producto 1 · Especificaciones técnicas (carpeta `G3-P1-<nombre-corto>/`)

| Artefacto obligatorio | Mínimo |
|---|---|
| A1. Formato de ficha técnica (plantilla propia) | 1 formato con 12 o más campos |
| A2. Fichas técnicas diligenciadas | 4 fichas: 2 de software o servicio (una es su propio software) y 2 de hardware o infraestructura; 6 o más características medibles por ficha |
| A3. Referente técnico | 4 referentes obligatorios en tabla (norma o estándar, qué exige, a qué parte aplica, cómo se cumple, cómo se verifica, fecha y enlace de verificación): Ley 1581 de 2012 con su reglamentación, Ley 23 de 1982, Decisión Andina 351 de 1993 y un estándar técnico reconocido; un quinto referente pertinente es opcional y suma |
| A4. Dimensionamiento de licencias | 6 componentes con tipo, modelo de cobro, cantidad calculada, costo y obligación; inventario de 5 dependencias del código con su licencia; decisión sobre la licencia de su propio software |
| A5. Estimación de condiciones técnicas | 5 parámetros de capacidad y servicio con supuesto o fórmula |
| A6. Estimación preliminar de condiciones económicas | Hoja de cálculo con fórmulas vivas, 12 y 36 meses, método de estimación declarado, 8 o más ítems de esfuerzo |
| README | Plantilla completa |

### Producto 2 · Propuesta técnica y validación (carpeta `G3-P2-<nombre-corto>/`)

| Artefacto obligatorio | Mínimo |
|---|---|
| B1. Análisis comparativo de proveedores | 2 categorías, 3 alternativas reales por categoría, 5 criterios ponderados que suman 100 %, TCO a 36 meses en la alternativa recomendada y en la mejor descartada de cada categoría |
| B2. Presupuesto de software y hardware | Hoja de cálculo con fórmulas, 6 categorías, directos e indirectos, únicos y recurrentes, contingencia (suma de valores esperados), IVA, una columna que liga cada rubro con el nodo o componente del diagrama (o lo declara excluido) y una fila de utilidad o margen con el precio de venta |
| B3. Propuesta técnica documentada | 6 a 12 páginas, 12 secciones, 1 diagrama de arquitectura o despliegue propio ajustado (nodos, componentes, protocolos), matriz de riesgos de 6 filas, EDT de 10 tareas con duración, predecesoras y ruta crítica calculada |
| B4. Presentación a la contraparte | 8 a 12 diapositivas o guion de 10 minutos |
| B5. Acta de validación con sugerencias y ajustes | 4 sugerencias clasificadas, 2 ajustes de parámetros con antes y después |
| B6. Lista de verificación para contrato, matriz de cumplimiento de términos de referencia y ANS | 10 elementos de contrato, 6 requisitos de términos de referencia, ANS con 3 métricas |
| B7. Propuesta ajustada (v1.1) con control de cambios | Versión, fecha y lista de cambios |
| README | Plantilla completa |

### Presupuesto de tiempo (referencia de diseño)

Los mínimos están estimados en unos **170 minutos de trabajo concentrado por producto** (335 en total) sobre lo que ya existe en el fork. El momento 3.3 dispone de 6 horas (360 minutos): 2 de acompañamiento del instructor y 4 de trabajo independiente que tú agendas entre el 21 de noviembre y el 5 de diciembre. El margen es de unos 25 minutos en total (7 %).

Esta cifra es una **estimación de diseño, no una medición** (queda pendiente un piloto con un caso real; ver el acta de debate). Por eso no hay válvula de escape: los mínimos se exigen, y para cumplirlos se reparte el avance en las dos semanas, sin dejarlo para la última noche. Si va atrasado, priorice en este orden: A2, A4, A6 en el Producto 1 y B2, B3, B5 en el Producto 2, y no deje de dejar commits.

| Producto 1 | Minutos | Producto 2 | Minutos |
|---|---|---|---|
| A1 Formato de ficha | 12 | B1 Comparativo (6 alternativas y matriz) | 26 |
| A2 Cuatro fichas | 40 | B2 Presupuesto | 24 |
| A3 Cuatro referentes | 32 | B3 Propuesta (12 secciones breves) | 48 |
| A4 Seis componentes y 5 dependencias | 40 | B4 Presentación | 12 |
| A5 Cinco parámetros | 12 | B5 Reunión y acta | 22 |
| A6 Hoja de costos preliminar | 25 | B6 Contrato, términos de referencia y ANS | 20 |
| README | 8 | B7 Versión 1.1 y README | 14 |
| **Total** | **169** | **Total** | **166** |

**Datos personales.** Todo lo que suba al fork es público. Use datos ficticios o de dominio público; no incluya nombres reales de personas, documentos, teléfonos, correos ni contratos reales de terceros. A la contraparte se la identifica por su usuario de GitHub o por su rol.

## 4. Rúbrica del Producto 1 · Especificaciones técnicas (RAP 593060)

Referentes del programa: criterios de evaluación y conocimientos de proceso del RAP 01 «Definir especificaciones técnicas del software de acuerdo con las características del software a construir».

### C1 · Formato de ficha técnica (peso 15)

Línea del programa: «Diseña formatos de fichas técnicas de acuerdo con estándares»; conocimiento de proceso «Diseñar formatos de fichas técnicas».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Formato propio de 12 o más campos: identificación, descripción, característica con valor, unidad y método de verificación, cantidad y unidad de medida, normas o estándares, garantía, compatibilidad y observaciones. Indica el estándar en que se basa y se usó en todas las fichas. | Formato de 8 a 11 campos, con valor y unidad pero sin método de verificación o sin normas. Se usó en todas las fichas. | Formato de 5 a 7 campos, copiado sin adaptar, sin unidad de medida o que no usan todas las fichas. | No hay formato, o tiene 4 campos o menos. |

### C2 · Fichas técnicas diligenciadas (peso 20)

Línea del programa: «Diseña formatos de fichas técnicas de acuerdo con estándares»; conocimientos de proceso «Diligenciar formatos de fichas técnicas»; saber «Ficha técnica: elementos mínimos para describir los referentes técnicos».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| 4 o más fichas (2 de software o servicio, una de ellas de su propio software, y 2 de hardware o infraestructura), con 6 o más características medibles cada una (valor, unidad y forma de verificar). Ninguna pide una marca o modelo como requisito, salvo compatibilidad justificada. Los valores salen de su caso y de su variante. | 4 o más fichas, pero alguna tiene menos de 6 características medibles, o una ficha pide una marca, o hay adjetivos no medibles en 3 líneas o menos. | 2 o 3 fichas, o 4 o más con más de la tercera parte de las características sin medida, o valores que no corresponden al caso. | Menos de 2 fichas, o fichas sin valores. |

### C3 · Referente técnico con normatividad vigente (peso 20)

Línea del programa: «Elabora referentes técnicos de acuerdo con normatividad legal vigente»; conocimiento de proceso «Elaborar referentes técnicos»; saber «Soluciones informáticas: diagnóstico de necesidades y elaboración de referentes técnicos».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| 4 o más referentes en tabla (norma o estándar, qué exige, a qué parte del software aplica, cómo se cumple, cómo se verifica). Incluye la Ley 1581 de 2012 y su reglamentación (Decreto 1074 de 2015), la Ley 23 de 1982, la Decisión Andina 351 de 1993 y un estándar técnico reconocido (por ejemplo ISO/IEC 25010, ISO/IEC 27001:2022 o WCAG 2.2). Un quinto referente pertinente (por ejemplo la Ley 1273 de 2009 sobre delitos informáticos, o una norma sectorial como la de historia clínica si el caso es de salud) suma pero no es obligatorio. Separa las reglas internas de su proyecto (su «Normativa del Proyecto») de la normatividad legal. Cita nombre, número y año, con vigencia verificada (enlace y fecha de consulta). | 3 referentes completos que incluyen la Ley 1581 de 2012 y una norma de derechos de autor, o 4 referentes a los que les falta una de las normas obligatorias; falta «cómo se verifica» o la fecha de verificación en 2 o menos; mezcla una regla interna sin marcarla. | 2 referentes, o normas sin relación con las partes del software, o no separa norma interna de norma legal. | Sin referentes, o presenta solo reglas internas del proyecto como si fueran normatividad legal, o cita normas inexistentes. |

### C4 · Dimensionamiento de licencias (peso 20)

Línea del programa: «Dimensiona las licencias de software de acuerdo con las características de la solución y las necesidades de la empresa»; conocimiento de proceso «Determinar licenciamientos tecnológicos»; saber «Licenciamiento de software: concepto, tipos, características, diferencias de carácter técnico y de condiciones de cumplimiento» y «Propiedad intelectual: derechos de autor, propiedad industrial, registro de software».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Tabla de 6 o más componentes (sistema operativo, base de datos, lenguaje o framework, librerías clave, herramientas de desarrollo y diseño, hosting o nube, control de versiones o integración continua) con tipo de licencia, modelo de cobro (nominal, concurrente, por núcleo, suscripción, gratuita), cantidad calculada con fórmula a partir de su variante, costo y periodo, y obligación de cumplimiento (atribución, copyleft, uso no comercial). Inventario de 5 o más dependencias del código con su licencia y alerta de compatibilidad. Decide y justifica la licencia de su propio software. | 5 componentes con tipo, modelo y cantidad, pero sin fórmula o sin obligaciones en 3 filas o menos; inventario de 3 o más dependencias; sin decisión sobre la licencia propia. | 3 o 4 componentes, cantidades sin cálculo, ignora restricciones de uso comercial o educativo, o no hay inventario. | Menos de 3 componentes, no distingue software gratuito de licenciado o propone software sin licencia. |

### C5 · Condiciones técnicas (peso 12)

Línea del programa: «Estima condiciones técnicas y económicas de acuerdo con las características de la solución de software y las necesidades de la empresa» (parte técnica); saber «Tecnología informática: características de hardware y software».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| 5 o más parámetros de capacidad y servicio (usuarios totales y concurrentes, almacenamiento inicial y crecimiento anual, ancho de banda, disponibilidad objetivo, tiempo de respuesta, RPO y RTO de las copias de seguridad), derivados de su variante y de su caso, con fórmula o supuesto declarado y coherentes con las fichas. | 3 o 4 parámetros con supuestos declarados pero sin fórmulas, o una incoherencia con las fichas. | 2 parámetros, o valores sin justificación. | Sin estimación técnica o con valores contradictorios. |

### C6 · Condiciones económicas (peso 13)

Línea del programa: «Estima condiciones técnicas y económicas de acuerdo con las características de la solución de software y las necesidades de la empresa» (parte económica); saber «Costos: conceptos, fundamentos, clasificación».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Hoja de cálculo con fórmulas vivas: talento humano con método declarado (analogía, tres puntos, puntos de función u horas por módulo con 8 o más ítems), licencias, infraestructura y operación a 12 y 36 meses; clasifica costos directos e indirectos y únicos y recurrentes; indica rango de error y lista de supuestos y exclusiones; respeta el tope de su variante o explica el exceso. | Hoja con fórmulas, talento, licencias e infraestructura a 12 meses; clasificación parcial; supuestos listados. | Totales sin desglose, valores pegados sin fórmula o sin horizonte de tiempo. | Sin estimación económica o con cifras sin sustento. |

## 5. Rúbrica del Producto 2 · Propuesta técnica y validación (RAP 593062 y 593061)

Referentes del programa: criterios de evaluación y conocimientos del RAP 02 «Elaborar propuesta técnica del software de acuerdo con las especificaciones técnicas definidas» y del RAP 03 «Validar las condiciones de la propuesta técnica del software de acuerdo con los intereses de las partes».

### C1 · Análisis comparativo de proveedores (peso 15)

Línea del programa: «Elabora análisis comparativo de proveedores considerando costos, cantidad y especificaciones técnicas definidas»; conocimiento de proceso «Elaborar análisis comparativo de proveedores y presupuestos».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| 2 o más categorías (por ejemplo nube o hosting, equipos, licencias) con 3 o más alternativas reales cada una y evidencia (enlace o captura con fecha de consulta, cotización si la hay). Todas se comparan contra las mismas fichas del Producto 1 (misma especificación y cantidad). Matriz con 5 o más criterios cuyos pesos suman 100 %, puntajes justificados, TCO a 36 meses con fórmulas en la alternativa recomendada y en la mejor descartada de cada categoría, y recomendación que explica lo descartado. | 2 o más categorías con 3 o más alternativas y matriz ponderada, pero TCO parcial o evidencia sin fecha en 2 alternativas o menos. | 1 categoría o menos de 3 alternativas, comparación solo por precio, o alternativas con especificaciones distintas sin normalizar. | Sin comparativo, o alternativas inventadas sin evidencia. |

### C2 · Costos de software y hardware (peso 15)

Línea del programa: «Estima costos de software y hardware de acuerdo con las características del negocio»; conocimiento de proceso «Estimar costos y presupuestos del software».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Presupuesto en hoja de cálculo con fórmulas, 6 o más categorías (talento humano, licencias, hardware o infraestructura, nube y operación, capacitación y soporte, contingencia) con software y hardware separados; directos e indirectos; únicos y recurrentes; IVA aplicado y supuesto de moneda si hay dólares; contingencia igual a la suma de los valores esperados de los riesgos que decide cubrir con reserva (si el porcentaje resultante sobre el subtotal queda fuera de 5 % a 15 %, lo explica); fila de utilidad o margen y precio de venta con la fórmula (declara si es margen sobre el precio, costo/(1 − margen), o recargo sobre el costo); total consistente con el comparativo y con la arquitectura, con cada nodo o componente del diagrama ligado a un rubro o a una exclusión justificada; flujo a 36 meses; respeta el tope de su variante o lo discute. | Fórmulas y 5 o más categorías; contingencia justificada sin ligarla a riesgos; una inconsistencia menor con la arquitectura. | Menos de 4 categorías, valores pegados, contingencia sin justificar, o 2 o más componentes de la arquitectura sin costo. | Sin presupuesto o sin separar software y hardware. |

### C3 · Propuesta técnica documentada (peso 20)

Línea del programa: «Documenta la propuesta técnica según las especificaciones de la solución de software»; conocimiento de proceso «Construir la propuesta técnica teniendo en cuenta la arquitectura a implementar para el software».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Documento de 6 a 12 páginas con las 12 secciones: resumen ejecutivo; entendimiento de la necesidad; alcance con inclusiones, exclusiones y supuestos; solución y arquitectura con 1 diagrama de arquitectura o despliegue propio, reutilizado y ajustado, con nodos, componentes y protocolos (los de casos de uso, clases, BPMN, MER y mockups se pueden anexar, pero no lo reemplazan); especificaciones técnicas (remite al Producto 1); metodología y cronograma con EDT de 10 o más tareas con duración, predecesoras, hitos y ruta crítica calculada (holgura por tarea); equipo y roles; costos (remite al presupuesto); matriz de 6 o más riesgos con probabilidad, impacto, valor esperado y respuesta; garantías y soporte; condiciones comerciales y de propiedad intelectual; vigencia de la oferta y anexos. Coherente con el Producto 1 y el presupuesto. | 10 u 11 secciones, un diagrama de arquitectura propio ajustado, 4 o 5 riesgos, cronograma sin ruta crítica calculada; una inconsistencia menor. | 7 a 9 secciones, diagramas copiados sin ajuste, menos de 4 riesgos, o costos que no concuerdan con las especificaciones. | Menos de 7 secciones o copia sustancial de otro documento. |

### C4 · Presentación de la propuesta (peso 10)

Línea del programa: «Realiza la presentación de la propuesta según propósitos de la negociación»; conocimiento de proceso «Presentar la propuesta técnica y costos del software».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| 8 a 12 diapositivas o guion de 10 minutos que declaran el propósito de la negociación (objetivo, posición ideal, posición mínima y mejor alternativa a un acuerdo), siguen el orden del interés del cliente (problema, beneficio, solución, costo, plazo, riesgos, siguientes pasos), usan cifras coherentes con la propuesta y anticipan 3 o más objeciones con su respuesta. Se presentó a la contraparte, según consta en el acta. | Presentación completa y coherente, con propósito declarado pero sin mejor alternativa o con menos de 3 objeciones anticipadas. | Presentación centrada en la tecnología, sin beneficio para el negocio o sin propósito de negociación. | Sin presentación, o no se presentó a una contraparte. |

### C5 · Sugerencias del cliente documentadas (peso 10)

Línea del programa: «Documenta las sugerencias del cliente según la necesidad del negocio»; saber «Comunicación: técnicas, tipos y habilidades».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Acta con fecha (anterior a la entrega; nunca posterior a ella), medio, participantes (usuario de GitHub o rol; la contraparte es distinta del autor), 4 o más sugerencias u objeciones convertidas en requisito verificable y clasificadas (aceptada, negociada, rechazada) con justificación, prioridad y efecto en costo y plazo; confirmada por la contraparte por escrito. | 3 sugerencias clasificadas; algunas sin efecto en costo y plazo o sin confirmación de la contraparte. | 2 sugerencias ambiguas, sin clasificar. | Sin acta, o sugerencias sin contraparte real. |

### C6 · Ajuste de parámetros técnicos (peso 15)

Línea del programa: «Ajusta los parámetros técnicos de acuerdo con los resultados de la negociación»; conocimiento de proceso «Formular ajustes sobre los parámetros técnicos»; saber «Negociación de tecnología: estilos de negociación, principios de negociación, identificación y solución de conflictos».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Tabla de 2 o más parámetros técnicos modificados (antes, después, motivo) con impacto cuantificado en costo, plazo y riesgo. Los cambios están reflejados en la propuesta v1.1, el presupuesto y el ANS, con control de cambios (versión y fecha). Hay al menos una sugerencia rechazada con argumento técnico. | 2 ajustes cuantificados y reflejados en propuesta y presupuesto, pero no en el ANS, o sin control de cambios, o sin sugerencia rechazada. | 1 ajuste, o ajustes descritos sin cuantificar, o solo anotados en el acta y no en los documentos. | Sin ajustes ni versión ajustada. |

### C7 · Información para el contrato, términos de referencia y ANS (peso 15)

Línea del programa: conocimientos de proceso «Verificar que la propuesta técnica contenga la información necesaria para la elaboración del contrato» y «Verificar los términos técnicos de referencia»; saber «Contrato: concepto, tipos, partes, normatividad», «Acuerdos de niveles de servicio: concepto, formulación, métricas» y «Términos de referencia: concepto, elementos y construcción».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| (a) Lista de verificación para el contrato con 10 o más elementos (partes, objeto, alcance, valor, forma e hitos de pago, plazos, entregables, criterios de aceptación, garantías, propiedad intelectual y titularidad, confidencialidad y datos personales, ANS, terminación, vigencia), cada uno con la sección donde se cumple. (b) Matriz de cumplimiento contra 6 o más requisitos de términos de referencia (los entregó la contraparte o los simuló), con cumple, parcial o no cumple y la evidencia. (c) ANS con 3 o más métricas (por ejemplo disponibilidad, tiempo de respuesta, RPO y RTO) con valor, ventana, fuente de medición y compensación. | Contrato con 8 o más elementos, matriz de 4 o más requisitos y ANS con 2 métricas medibles. | Lista con menos de 8 elementos, ANS con adjetivos («rápido», «alta disponibilidad») o matriz de menos de 4 requisitos. | No presenta ninguno de los tres. |

## 6. Rúbrica de la prueba de desempeño (RAP 593060, 593062 y 593061)

El caso es nuevo e inédito y no se publica. Se publican las tareas tipo y la rúbrica para que el aprendiz sepa qué se observa.

**Formato.** Individual, presencial y **en computador** (PC del ambiente de formación o portátil del aprendiz), 90 minutos (se aplica después de la prueba de conocimiento de 40 minutos, el mismo sábado 5 de diciembre): 5 de lectura, 20 de ficha y referentes, 15 de licencias, 25 de costos, 15 de ajustes ante dos cambios del cliente y 10 de respuesta escrita. **Recursos permitidos:** hoja de cálculo, procesador de texto, calculadora (la del computador), su plantilla de ficha técnica, su guía y esta rúbrica, y consulta de fuentes oficiales de normas y de páginas públicas de precios. **No permitido:** asistentes de IA, mensajería, repositorios de otros compañeros. **Entrega:** al terminar, el aprendiz sube sus archivos (un libro de hoja de cálculo con las hojas pedidas y un archivo de texto con la respuesta al cliente) a una tarea de Moodle (Zajuna); no se sube al fork.

### D1 · Ficha técnica y referentes aplicables al caso (peso 20)

Línea del programa: «Diseña formatos de fichas técnicas de acuerdo con estándares» y «Elabora referentes técnicos de acuerdo con normatividad legal vigente».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Ficha del componente pedido con 6 o más características medibles (valor, unidad y forma de verificar), sin marca como requisito; identifica 2 normas aplicables al caso, las cita bien y explica qué exigen. | 5 características medibles o una norma bien aplicada y la otra solo nombrada. | 3 o 4 características medibles, o normas genéricas sin relación con el caso. | Ficha sin valores o sin normas. |

### D2 · Dimensionamiento de licencias (peso 20)

Línea del programa: «Dimensiona las licencias de software de acuerdo con las características de la solución y las necesidades de la empresa».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Calcula cantidades con fórmula y margen, elige el modelo de cobro más conveniente comparando el costo a 36 meses y detecta las restricciones de licencia presentes en el caso, proponiendo una acción para cada una. | Cantidad y modelo correctos, pero sin comparar a 36 meses o sin detectar una de las restricciones. | Error de cantidad o modelo sin justificación. | No dimensiona o propone usar software sin licencia. |

### D3 · Estimación de costos y condiciones (peso 25)

Línea del programa: «Estima costos de software y hardware de acuerdo con las características del negocio» y «Estima condiciones técnicas y económicas de acuerdo con las características de la solución de software y las necesidades de la empresa».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Hoja de cálculo con fórmulas (no valores pegados), 5 o más rubros, contingencia e IVA bien aplicados, total exacto (diferencia de 1 % o menos frente al total que resulta de las decisiones de licencia que el propio aprendiz declaró en su hoja de licencias), supuestos escritos e indica si supera el tope del cliente. | Total con error de hasta 5 % por un solo rubro mal calculado; fórmulas en lo esencial. | Error mayor de 5 %, valores pegados u omite 2 rubros o más. | Sin hoja de cálculo o sin total. |

### D4 · Ajuste ante los cambios del cliente (peso 25)

Línea del programa: «Ajusta los parámetros técnicos de acuerdo con los resultados de la negociación»; conocimiento de proceso «Formular ajustes sobre los parámetros técnicos».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Aplica los dos cambios, recalcula costo, plazo y licencias con cifras, propaga el cambio en la tabla «Propagación» de la hoja 4-Ajustes (valor antes y después de la métrica en la ficha y en el ANS) y propone una contrapropuesta con criterio objetivo. | Recalcula los dos cambios con cifras, pero no los propaga a la ficha o al ANS. | Recalcula solo uno de los cambios. | No ajusta o ajusta sin cifras. |

### D5 · Documentación de sugerencias y justificación (peso 10)

Línea del programa: «Documenta las sugerencias del cliente según la necesidad del negocio».

| 100 | 70 | 40 | 0 |
|---|---|---|---|
| Registra las dos sugerencias del cliente como requisitos verificables y redacta una respuesta de 150 palabras o menos que explica el impacto y argumenta con beneficio para el negocio. | Registra las dos sugerencias, pero una queda ambigua o la respuesta no argumenta el beneficio. | Registra una sugerencia o la respuesta solo repite lo pedido. | No documenta. |

## 7. Verificación oral de autoría (sábado 5 de diciembre)

**No es un quinto entregable ni tiene nota propia: solo puede activar los topes de la sección 2.** Es una conversación corta y registrada para comprobar que el aprendiz entiende su propio trabajo.

**Cómo se hace.** Mientras los demás hacen la prueba de conocimiento (40 min) y la de desempeño (90 min), el instructor llama a cada aprendiz a su puesto y le hace **3 preguntas** de unos 3 minutos en total: una del Producto 1, una del Producto 2 y una transversal (qué reutilizó y qué cambió). Con unos 33 aprendices son unos 100 minutos, que caben en los 130 de pruebas y en los 30 de retroalimentación del momento 3.4. El aprendiz puede abrir su repositorio. Al aprendiz que se llama durante una prueba se le suman 3 minutos de compensación y se anota la hora.

**Cómo se califica cada respuesta.** *Explica* (2 puntos): ubica el artefacto en su repositorio, reproduce la cifra, fórmula o decisión y la justifica. *Explica en parte* (1 punto): cumple dos de esas tres cosas. *No explica* (0 puntos): cumple una o ninguna.

**Efectos.**
- Total de 5 o 6 puntos: sin efecto.
- Total de 3 o 4 puntos: sin tope; el instructor revisa el historial de commits y deja constancia en la ficha de registro.
- Si la pregunta de un producto queda en 0, ese producto queda con tope de 50.
- Si **dos o más** respuestas quedan en 0, los dos productos quedan con tope de 50.
- «Explica en parte» nunca genera tope por sí sola.

**Ficha de registro.** El instructor anota, por aprendiz: fecha y hora, las tres preguntas, un resumen de cada respuesta, el puntaje y observaciones; el aprendiz la firma. Es el soporte de cualquier tope.

**Derecho de réplica antes de aplicar un tope de 50 o un 0 por copia.** El aprendiz puede presentar por escrito, en el plazo que defina el instructor (por definir por el instructor), una explicación o evidencia adicional (por ejemplo, commits o archivos de origen). El instructor la revisa y deja la decisión motivada. Por eso **la nota de los productos 1 y 2 es provisional y se consolida después de la verificación oral del 5 de diciembre y del plazo de réplica**, cuando termina este proceso.

Tipos de pregunta (ejemplos): de dónde salió una cifra de su hoja de cálculo y qué pasa si cambia; qué cambió en una ficha respecto de su versión inicial del año; por qué eligió esa licencia y qué obligación le impone; por qué descartó una alternativa del comparativo; qué sugerencia rechazó y con qué argumento; qué diagrama es suyo y qué ajustó.

## 8. Uso de herramientas de IA

Se permite usar IA como apoyo, con tres condiciones: (1) declararlo en el README (herramienta, para qué, qué parte); (2) verificar normas, cifras y precios en fuentes oficiales y no copiar textos legales de memoria del asistente; (3) poder explicar todo el trabajo en la verificación oral. Lo que no se pueda explicar o verificar no suma.

## 9. De dónde reutilizar (y qué construir desde cero)

Los productos parten de lo que el aprendiz hizo durante el año en su fork. **No todos los forks tienen todo**: usa lo que tu fork tenga y, si falta, la alternativa indicada.

| Artefacto de la guía | Reutiliza de tu fork | Si no lo tiene |
|---|---|---|
| Fichas técnicas (P1 A2) | El documento «Normativa del Proyecto» o la ficha de proyecto (`03-execute/13-week/02-optional-activity/` y `03-execute/14-week/02-optional-activity/`), el SRS y los requisitos no funcionales | `03-execute/12-week/01-session/design-software/docs/00-discovery/estructura-tecnica-inicial.md` y `alcance-mvp.md` (carpeta del instructor) |
| Condiciones técnicas (P1 A5) | Metas de rendimiento, disponibilidad y accesibilidad que ya definió en su SRS o en la ficha de proyecto; `01-context/constraints.md` del marco, si lo adaptó | Los valores de la variante del anexo A |
| Dimensionamiento de licencias (P1 A4) | `package.json`, `go.mod`, `requirements.txt` o `docker-compose.yml` del proyecto de la semana 24 (`03-execute/24-week/02-optional-activity/`), si lo tiene | `03-execute/12-week/01-session/design-software/api/go.mod`, `web/package.json` y `docker-compose.yml` (dependencias e imágenes con su licencia); complete con las tecnologías declaradas en su ficha de proyecto |
| Esfuerzo y costos (P1 A6, P2 B2) | La hoja de ruta (`03-execute/17-week/02-optional-activity/`), el cronograma, la lista de módulos del SRS | La lista de casos de uso o de pantallas del mockup |
| Arquitectura y diagramas (P2 B3) | Casos de uso y clases (`03-execute/15-week` y `16-week`), BPMN y modelo de datos (`22-week` y `23-week`), MER (`20-week` y `21-week`), mockups (`13-week` y `17-week`), todos en `02-optional-activity/` | El esquema `web, api y db` de `design-software` solo como referencia, redibujado y ajustado a su caso |
| Riesgos (P2 B3) | `15-project-control/risks.md` del marco `Laboratorio_Sena` en tu fork y el plan de contingencia (`03-execute/18-week` y `19-week`) | Construye la matriz desde cero con la plantilla del anexo B |
| Presentación (P2 B4) | El pitch y las presentaciones (`03-execute/11-week`, `13-week`, `14-week` y `17-week`) | Construya una presentación nueva de 8 a 12 diapositivas |

**Lo que no se vio en clase y se construye desde cero:** estimación de esfuerzo, costos y presupuesto, licenciamiento y propiedad intelectual, normatividad legal (datos personales y derechos de autor), matriz de riesgos con valor esperado, cronograma formal con ruta crítica, comparativo de proveedores y validación formal con acta. La «Normativa del Proyecto» del año recoge normas **internas** del proyecto: sirve de insumo, pero no cuenta como referente legal.

**Lo que no se acepta sin ajustes:** archivos del marco `Laboratorio_Sena` o de `design-software` copiados tal cual, y documentos idénticos a los de otro fork.

## Anexo A · Variantes del escenario (las asigna el instructor)

Aplica los parámetros de tu variante al caso de tu fork; si no coinciden con la realidad de tu proyecto, decláralos como supuestos del ejercicio.

| Variante | Usuarios totales | Usuarios concurrentes (pico) | Sedes | Disponibilidad objetivo | Tope de presupuesto (primer año) |
|---|---|---|---|---|---|
| A | 60 | 15 | 1 | 99,0 % mensual | $45.000.000 |
| B | 120 | 30 | 2 | 99,5 % mensual | $70.000.000 |
| C | 250 | 60 | 3 | 99,5 % mensual | $110.000.000 |
| D | 500 | 120 | 4 | 99,9 % mensual | $160.000.000 |

Los valores se usan en pesos colombianos, con IVA del 19 % sobre el total cuando la hoja de cálculo lo pida y TRM de referencia de $4.000 por dólar si hay precios en dólares.

## Anexo B · Plantilla de matriz de riesgos

| N.º | Riesgo | Causa | Probabilidad (%) | Impacto ($) | Valor esperado ($) | Respuesta (evitar, mitigar, transferir, aceptar) | Responsable |
|---|---|---|---|---|---|---|---|
| R1 | | | | | =probabilidad × impacto | | |

La reserva de contingencia sale de la suma de los valores esperados de los riesgos que decida cubrir con reserva.

## Anexo C · Lista de chequeo antes del push

- [ ] Las carpetas `G3-P1-<nombre-corto>` y `G3-P2-<nombre-corto>` están en `02-training-project/01-phase/02-planning/07-technical-proposal/04-evidence-learning/` de mi fork.
- [ ] Cada carpeta tiene su `README.md` con la plantilla completa.
- [ ] Cumplí los mínimos de la sección 3 (conté fichas, referentes, componentes, alternativas, riesgos, sugerencias).
- [ ] Mis hojas de cálculo tienen fórmulas, no solo valores.
- [ ] La reunión de validación con la contraparte fue antes de la entrega y el acta no está fechada después de ella.
- [ ] Verifiqué en fuente oficial, con fecha, las normas que cité y al menos 3 cifras de mis hojas.
- [ ] Mis commits están repartidos en al menos 2 días distintos entre el 21 de noviembre y el 5 de diciembre.
- [ ] No hay datos personales reales ni contratos reales en el repositorio.
- [ ] Verifiqué en GitHub, desde otro navegador, que el push quedó publicado antes de el sábado 5 de diciembre de 2026 a las 6:00 a. m.

## Anexo D · Convenciones de cálculo (únicas para toda la guía)

- Valores en pesos colombianos y sin IVA, salvo que el enunciado diga otra cosa. El IVA del 19 % se aplica al final, sobre el total con contingencia (subtotal, luego contingencia, luego IVA).
- Contingencia (reserva) = suma de los valores esperados de los riesgos que se decide cubrir con reserva. Se muestra como porcentaje del subtotal; si queda fuera de 5 % a 15 %, se explica. No se ajustan las probabilidades para que el resultado «entre» en el rango.
- Cada precio indica su periodicidad (único, mensual o anual). Los precios en dólares se convierten con la TRM de la variante.
- Precio de venta con margen sobre el precio: costo / (1 − margen). Con recargo sobre el costo: costo × (1 + recargo). Declare cuál usa; no son equivalentes (un margen del 20 % equivale a un recargo del 25 %).
- El costo de reemplazar un componente (por ejemplo, cambiar una librería) se calcula como horas de desarrollo × tarifa y se registra junto a las licencias cuando la decisión es de licenciamiento.

## Anexo E · Trazabilidad criterio × evidencia (los 10 criterios de evaluación del programa)

| Criterio literal del programa (RAP) | Prueba de conocimiento | Producto 1 | Producto 2 | Prueba de desempeño |
|---|---|---|---|---|
| Diseña formatos de fichas técnicas de acuerdo con estándares (01) | Sí | C1, C2 | | D1 |
| Elabora referentes técnicos de acuerdo con normatividad legal vigente (01) | Sí | C3 | | D1 |
| Dimensiona las licencias de software (01) | Sí | C4 | | D2 |
| Estima condiciones técnicas y económicas (01) | Sí | C5, C6 | | D3 |
| Elabora análisis comparativo de proveedores (02) | Sí | | C1 | |
| Estima costos de software y hardware (02) | Sí | | C2 | D3 |
| Documenta la propuesta técnica (02) | | | C3 | |
| Realiza la presentación de la propuesta (03) | | | C4 | |
| Documenta las sugerencias del cliente (03) | Sí | | C5 | D5 |
| Ajusta los parámetros técnicos (03) | Sí | | C6, C7 | D4 |

Los conocimientos de proceso «Verificar la información para el contrato» y «Verificar los términos técnicos de referencia» se evalúan en el Producto 2, C7, y en la prueba de conocimiento.
