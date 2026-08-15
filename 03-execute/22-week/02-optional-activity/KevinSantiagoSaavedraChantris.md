# Retrospectiva y Reingenieria - Sistema de Gestion de Horarios SENA
## Version final

> Este documento es una **version**. Su objetivo es presentar una seleccion depurada de oportunidades de reingenieria para que posteriormente el autor decida cuales conservar.
>
> El objeto analizado es un **MockUp/prototipo estatico de revision UX/UI**, no un sistema implementado.
>
> Fuente primaria: https://code-sena.github.io/design-software-mockup/
> Fuentes secundarias: los archivos `.md` del paquete de proyectos revisado.
> Referencia institucional: SIGA/MIPG y documentacion oficial del SENA.

---

# 1. Alcance y criterio de analisis

El mockup del Sistema de Gestion de Horarios declara que es un prototipo estatico, usa datos ficticios, no consume backend y presenta 53 pantallas y modales en 7 roles/areas.
Por esta razon, este documento **no evalua si el backend funciona** ni afirma que una accion inexistente en el prototipo sea un error de programacion.
La unidad de analisis es el comportamiento que la interfaz **representa o deberia representar**.
Se evalua si el flujo es comprensible, si una tarea esta bien planteada, si la informacion necesaria aparece en el momento correcto y si la reingenieria propuesta seria razonable de implementar.
Los archivos `.md` previos se utilizan como antecedentes y fuente de hallazgos.
No se copian automaticamente sus propuestas.
Cada propuesta debe superar un filtro de necesidad, evidencia, simplicidad y resultado.
El criterio central queda definido como:
**Usuario -> Flujo -> Problema observable -> Cambio minimo -> Resultado verificable.**

---

# 2. Template replanteado para el estudio

## 2.1 Usuario y flujo
Se identifica el rol y la tarea real.
Se describe solo el recorrido necesario para entender el problema.
No se incluyen historias extensas ni contexto que no cambie la decision.

## 2.2 Problema real
El problema debe estar relacionado con una pantalla, componente, contenido o secuencia del mockup.
No se acepta una necesidad inventada solo porque sea una buena practica general.
Cuando el problema dependa de una regla institucional no comprobada, se marca como pendiente.

## 2.3 Evidencia
Se registra la pantalla, flujo o patron donde se observa.
En un mockup estatico, la evidencia debe ser visual o de flujo representado.
No se debe presentar como hecho que un proceso de backend falla.

## 2.4 Propuesta
Debe plantear el cambio minimo necesario.
La solucion debe utilizar primero las pantallas, datos y patrones que ya existen.
Solo se agrega una pantalla nueva cuando el problema no puede resolverse dentro del flujo actual.

## 2.5 Funcionalidad minima
Se define que comportamiento tendria que existir si el mockup pasara a implementacion.
No se confunde la interfaz representada con funcionalidad actualmente ejecutada.

## 2.6 UX/UI
Se analizan jerarquia, lectura, navegacion, estados, feedback y consistencia.
La interfaz se modifica para mejorar una tarea, no para "modernizar" por gusto.

## 2.7 Control de sobreingenieria
Se indica expresamente que no se agrega.
El objetivo es que cada propuesta tenga una frontera.
Si una funcion no es necesaria para resolver el problema, queda fuera.

## 2.8 SIGA
SIGA se utiliza como marco de fundamento.
Se consideran especialmente enfoque por procesos, orientacion a las personas, resultados, eficiencia, eficacia, efectividad, seguimiento, control y mejora continua.
La comparacion no significa copiar otra plataforma.
Debe explicar por que una mejora contribuye al proceso.

## 2.9 Resultado
Se define una consecuencia observable.
El resultado debe poder comprobarse en una futura implementacion o prueba de usuario.

---

# 3. Estudio consolidado de los `.md`

Los 17 Markdown revisados aportan varios enfoques.
Algunos son auditorias de UI.
Otros reconstruyen flujos.
Otros proponen historias de usuario.
Otros conectan el mockup con SIGA.
El problema de reunirlos sin filtro es que mezclan tres niveles:
**observacion**, **requisito** y **solucion**.
La presente version los separa.
Una observacion describe lo que se ve.
Un requisito explica lo que el usuario necesita.
Una solucion propone como atenderlo.
No toda observacion necesita una nueva funcionalidad.
No todo requisito necesita una pantalla nueva.
No toda mejora visual requiere una reingenieria profunda.
Esta distincion reduce sesgos.

---

# 4. Hallazgos seleccionados

| ID | Rol | Oportunidad | Prioridad preliminar |
|---|---|---|---|
| H01 | Aprendiz | Priorizar la proxima actividad | Alta |
| H02 | Aprendiz | Hacer accionables las notificaciones | Alta |
| H03 | Instructor | Conectar excepciones con sesiones afectadas | Alta |
| H04 | Instructor | Mantener contexto al registrar seguimiento | Media |
| H05 | Coordinador | Hacer accionables los conflictos | Alta |
| H06 | Coordinador | Validar el estado antes de publicar | Alta |
| H07 | Coordinador | Mostrar disponibilidad en contexto | Media |
| H08 | Coordinador | Diferenciar crear y editar sesion | Media |
| H09 | Administrador | Resumir impacto de cambios sensibles | Media |
| H10 | Back-office | Hacer visible la vigencia de plantillas | Media |
| H11 | Parametrizacion | Mostrar contexto antes de cambiar configuracion | Media |
| H12 | Transversal | Mejorar mensajes y estados de sistema | Alta |
| H13 | Transversal | Mantener contexto entre pantallas relacionadas | Alta |
| H14 | Transversal | Unificar patrones de interaccion | Media |

---

# 5. H01 - Aprendiz: proxima actividad

## Usuario y flujo
Rol: Aprendiz.
Proceso: consulta del horario.
Pantallas relacionadas: Mi horario, detalle de clase.

El aprendiz entra al sistema.
Consulta su horario.
Identifica las actividades programadas.
Abre una actividad si necesita detalle.

## Problema
La vista semanal sirve para revisar el conjunto de actividades, pero puede no priorizar la pregunta inmediata:
**"Que tengo a continuacion?"**

## Evidencia
El propio analisis de Kevin identifica dificultad para reconocer cual es la proxima formacion.
Otros documentos tambien proponen priorizacion de informacion.
El problema es de jerarquia, no de ausencia total de datos.

## Propuesta
Agregar en la misma pantalla de horario un resumen compacto de la proxima actividad.
Debe utilizar datos que ya aparecen en el horario.
No requiere un nuevo modulo.

## Funcionalidad minima
Identificar la siguiente actividad futura.
Mostrar fecha, hora, actividad, instructor y lugar cuando esten disponibles.
Enlazar con el detalle existente.

## UX/UI
Colocar el resumen arriba del calendario.
Diferenciarlo del resto sin convertirlo en un dashboard.
Mantener el calendario como fuente principal de consulta.

## No sobreingenieria
No crear un asistente.
No crear un nuevo centro de actividades.
No duplicar el horario.

## SIGA
Se relaciona con orientacion a las personas y resultados: el sistema facilita que el usuario obtenga la informacion que necesita para ejecutar su actividad.

## Resultado
El aprendiz identifica la siguiente actividad sin recorrer visualmente toda la semana.

---

# 6. H02 - Aprendiz: notificacion accionable

## Usuario y flujo
Rol: Aprendiz.
Proceso: revisar una novedad del horario.

Recibe una notificacion.
Lee el mensaje.
Determina que cambio ocurrio.
Busca la actividad relacionada.

## Problema
Informar "el horario cambio" no es suficiente si el usuario debe buscar manualmente que actividad fue afectada.

## Evidencia
El mockup separa notificaciones y detalle de notificacion.
Los analisis del grupo identifican riesgo de perder cambios relevantes.

## Propuesta
Relacionar la notificacion con el objeto afectado.
La estructura minima debe ser:
**Que cambio -> cuando aplica -> que actividad afecta -> ver detalle.**

## Funcionalidad minima
La notificacion debe conservar una referencia al detalle correspondiente.
Debe diferenciar estado leida/no leida.
Debe mostrar el cambio de forma comprensible.

## UX/UI
Diferenciar visualmente novedades pendientes.
Usar una jerarquia especial para cancelaciones o reprogramaciones.
Mantener un enlace de entrada al detalle.

## No sobreingenieria
No exigir correo, SMS y otros canales para resolver este problema.
No crear una bandeja nueva.

## SIGA
Se relaciona con acceso oportuno a la informacion y orientacion al grupo de valor.

## Resultado
El usuario pasa de la notificacion al contenido afectado sin buscarlo de nuevo.

---

# 7. H03 - Instructor: excepcion conectada al horario

## Usuario y flujo
Rol: Instructor.
Proceso: informar una excepcion de disponibilidad.

Consulta disponibilidad.
Registra una excepcion.
El sistema debe indicar si existe una sesion en el periodo afectado.
El caso queda disponible para gestion del horario.

## Problema
Registrar disponibilidad y consultar horario son actividades relacionadas, pero el usuario puede no visualizar inmediatamente su relacion.

## Evidencia
La estructura del mockup separa Mi disponibilidad, Crear excepcion y Mi horario.
Los analisis previos identifican esta separacion como oportunidad.

## Propuesta
Despues de guardar una excepcion, mostrar si existen sesiones afectadas.
No se reprograma automaticamente.

## Funcionalidad minima
Guardar excepcion.
Comparar la franja con las sesiones existentes.
Listar las sesiones coincidentes.
Mostrar el estado de afectacion.

## UX/UI
Presentar el resultado inmediatamente.
Un mensaje positivo si no existe conflicto.
Un resumen de afectaciones si existe.

## No sobreingenieria
No crear sustitucion automatica.
No mover sesiones automaticamente.
No construir un optimizador de horarios.

## SIGA
Relaciona actividades del mismo proceso y mejora seguimiento y control.

## Resultado
El instructor comprende la consecuencia de su registro.

---

# 8. H04 - Instructor: contexto de seguimiento

## Usuario y flujo
Rol: Instructor.
Proceso: registrar seguimiento.

Selecciona una ficha.
Abre seguimiento.
Registra informacion.
Guarda.

## Problema
El registro requiere que el usuario conserve mentalmente el contexto al pasar a la pantalla de captura.

## Evidencia
El mockup contiene Seguimiento de ficha y Registrar seguimiento como pasos relacionados.

## Propuesta
Mantener visible:
**Ficha -> programa -> actividad/sesion**.

## Funcionalidad minima
El seguimiento debe quedar asociado al contexto seleccionado.

## UX/UI
Usar encabezado contextual.
No permitir que el usuario pierda la referencia principal durante el registro.

## No sobreingenieria
No duplicar la informacion en varias secciones.
No agregar un segundo modulo de contexto.

## SIGA
Apoya trazabilidad y reduccion de reprocesos.

## Resultado
El instructor registra informacion sabiendo exactamente a que elemento pertenece.

---

# 9. H05 - Coordinador: conflictos accionables

## Usuario y flujo
Rol: Coordinador academico.
Proceso: revisar un horario.

Crea o modifica una sesion.
El sistema representa un conflicto.
El coordinador debe entenderlo.
Corrige la configuracion.
Vuelve a validar.

## Problema
Un aviso de conflicto solo es util si explica su causa y orienta la correccion.

## Evidencia
Los analisis anteriores identifican conflictos de instructor y otros recursos.
El problema valido es de comprension y accion, no la existencia de un algoritmo especifico.

## Propuesta
Cada conflicto debe comunicar:
**que pasa + con que elemento + por que + que puede hacer el coordinador.**

## Funcionalidad minima
Detectar la condicion representada.
Relacionarla con la sesion afectada.
Mostrar una explicacion.
Permitir regresar al campo o elemento que debe corregirse.

## UX/UI
Severidad visual clara:
Critico.
Advertencia.
Informativo.

## No sobreingenieria
No agregar un Smart Assistant.
No generar automaticamente un nuevo horario.
No automatizar sustituciones sin reglas institucionales verificadas.

## SIGA
Favorece eficiencia y disminucion de reprocesos.
La solucion mantiene el control en el flujo del proceso.

## Resultado
El coordinador pasa de alerta a correccion con menos interpretacion manual.

---

# 10. H06 - Coordinador: validacion antes de publicar

## Usuario y flujo
Rol: Coordinador.
Proceso: publicar un horario.

Completa sesiones.
Revisa.
Corrige.
Publica.

## Problema
La publicacion es una decision importante y necesita un resumen claro del estado antes de confirmar.

## Evidencia
El mockup presenta funciones de publicacion y conflictos.
Los analisis recomiendan una comprobacion previa.

## Propuesta
Antes de publicar mostrar un resumen compacto:
Sesiones completas.
Conflictos criticos.
Advertencias.
Estado de validacion.

## Funcionalidad minima
Validar las condiciones que el proceso ya reconoce.
Mostrar resultado.
Permitir cancelar o publicar.

## UX/UI
El usuario debe poder responder:
**Que voy a publicar?**
**Hay problemas?**
**Puedo continuar?**

## No sobreingenieria
No crear un dashboard paralelo.
No agregar una etapa de aprobacion si no existe una regla institucional que la exija.

## SIGA
Relacion directa con seguimiento, control y resultados del proceso.

## Resultado
El coordinador publica con conocimiento del estado del horario.

---

# 11. H07 - Coordinador: disponibilidad en contexto

## Usuario y flujo
Rol: Coordinador.
Proceso: asignar instructor o ambiente.

Selecciona una sesion.
Define recursos.
