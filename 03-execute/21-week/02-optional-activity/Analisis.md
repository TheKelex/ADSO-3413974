### Analisis Panel Aprendiz MockUp (Sistema de Gestión de Horarios)

El sistema busca complementar la gestion de horarios del SENA, centralizando la programacion academica y facilitando su consulta, creacion, modificacion y actualizacion.

Su objetivo es mejorar la organizacion y comunicacion entre los diferentes roles, permitiendo acceder a informacion actualizada de las formaciones de manera rapida y sencilla, sin reemplazar los sistemas institucionales existentes.

--------------------------------------------------------------------------------------

## Rol: Aprendiz

Para realizar el entendimiento general del sistema se tienen que sentar las bases en uno de los actores principales de este, el rol de aprendiz nos permite comprender en su mayoria que es lo que busca solucionar el sistema.

### Alcance

El rol de aprendiz cuenta con 4 pantallas siendo estas:

25. Mi horario - semana
26. Notificaciones
27. Detalle de clase
28. Detalle de notificacion

Esto nos indica que el alcance de los aprendices dentro del sistema es el de consultar su horario de manera rapida, sencilla y centralizada, cubriendo incluso las actualizaciones de estas (notificaciones).

**Todas las pantallas son responsive**

### Flujograma Principal

El paso a paso por parte del aprendiz para consultar su horario es el siguiente:

1. Acceder mediante sus credenciales.
2. Revisar sus formaciones programadas (dashboard).
    - Puede consultar de manera clara los detalles de la sesion como competencia, instructor, ambiente, ubicacion, fecha y la franja (horario de esta), ademas de una nota de la sesion actual.
3. En el apartado de notificaciones o mediante la barra superior puede consultar los cambios de estas.

Esto nos demuestra que el flujo principal para trabajar el software es rapido e intuitivo para el aprendiz.

### Entendimiento de UI / UX

1. **Que entiende rapido:** Donde obtener rapidamente la informacion que busca, como consultar su horario (dashboard), donde consultar los cambios y como obtener mas informacion de cada uno de estos aspectos.
2. **Que no queda claro:** Cual es la proxima formacion (orden dependiente de la fecha) y que notificacion le falta por revisar.
3. **Que botones o textos sobran:** La informacion se encuentra clara y consisa.
4. **Que informacion falta:** Falta informacion acerca del dia exacto en el que esta y informacion de quien es el lider de su ficha.
5. **Que error podria cometer:** El aprendiz podria confundir una notificacion no vista con una que si y perder informacion clave sobre esa sesion como una posible reprogramacion.
6. **Que consecuencia tiene ese problema para el aprendiz:** El aprendiz podria ignorar una modificacion importante de su horario (como la cancelacion de una formacion).

### Comparacion con SIGA

El MockUp comparte con SIGA el objetivo de permitir al aprendiz consultar su informacion academica, especialmente su horario y las novedades de este. Sin embargo, el MockUp presenta la informacion de una forma mas centralizada, visual y rapida, facilitando la consulta de las sesiones.

Aun asi, frente a SIGA, se identifican oportunidades de mejora en la priorizacion de la proxima formacion, el estado de las notificaciones y la informacion relacionada con la ficha del aprendiz.

### Reingenieria

La reingenieria debe mantener la simplicidad del MockUp, pero mejorar la organizacion y priorizacion de la informacion. Se propone:

* Destacar la proxima formacion.
* Diferenciar claramente las notificaciones leidas y no leidas.
* Mostrar el dia y fecha actual.
* Incorporar informacion del lider de la ficha.
* Resaltar cambios importantes como reprogramaciones o cancelaciones.
* Apropiarse de los terminos usados en el Sena.

Con estos cambios, el aprendiz podria consultar y comprender su programacion con mayor rapidez, reduciendo posibles confusiones.

--------------------------------------------------------------------------------------

## Rol: Instructor

El instructor es uno de los actores principales del sistema, ya que ademas de consultar su programacion, participa directamente en la gestion de sus sesiones, disponibilidad y seguimiento de las fichas.

### Alcance

El rol de instructor cuenta con 6 pantallas:

19. Mi horario - semana
20. Detalle de sesion
21. Mi disponibilidad
22. Modal crear excepcion
23. Seguimiento de ficha
24. Registrar seguimiento

Esto indica que el alcance del instructor comprende tanto la consulta y gestion de su horario como la administracion de su disponibilidad y el seguimiento del proceso formativo de sus fichas.

**Las pantallas son responsive**

### Flujograma Principal

El flujo principal del instructor puede resumirse de la siguiente manera:

1. Acceder mediante sus credenciales.
2. Consultar sus formaciones programadas desde su horario.
3. Revisar el detalle de una sesion (click).
4. Gestionar su disponibilidad y registrar excepciones cuando sea necesario.
5. Consultar el seguimiento de sus fichas.
6. Registrar el seguimiento correspondiente a la ficha (Academico, bienestar, proyecto, etapa productiva).

Esto muestra que el instructor no solamente consulta informacion, sino que tambien interviene directamente en la gestion y seguimiento de las actividades formativas.

### Entendimiento de UI / UX

1. **Que entiende rapido:** Donde consultar su horario, revisar una sesion, gestionar su disponibilidad y acceder al seguimiento de sus fichas.
2. **Que no queda claro:** La relacion entre la disponibilidad del instructor y las sesiones programadas, asi como el impacto de crear una excepcion sobre su horario.
3. **Que botones o textos sobran:** La interfaz presenta las opciones principales de manera directa.
4. **Que informacion falta:** Podria ser util mostrar de forma mas evidente el estado de las sesiones, las excepciones registradas deberian de permitir anexar un documento valido y un resumen general de el avance de una ficha.
5. **Que error podria cometer:** El instructor podria registrar incorrectamente una excepcion o realizar un seguimiento sin identificar claramente la ficha o sesion correspondiente.
6. **Que consecuencia tiene ese problema:** Podrian generarse inconsistencias en la programacion o en el seguimiento del proceso formativo.

### Comparacion con SIGA

El MockUp complementa las funciones relacionadas con la programacion y seguimiento academico que maneja el SENA. Su principal diferencia esta en presentar estas actividades de manera mas centralizada, permitiendo al instructor consultar su horario, gestionar su disponibilidad y registrar seguimientos desde un mismo sistema.

Se identifican oportunidades de mejora principalmente en la relacion entre disponibilidad, programacion y seguimiento, buscando que los cambios realizados por el instructor sean claros y tengan un impacto visible sobre su planificacion.

### Reingenieria

La reingenieria debe mantener la centralizacion del sistema y mejorar la relacion entre sus diferentes funciones. Se propone:

Relacionar claramente la disponibilidad con el horario.
Mostrar las excepciones de forma visible.
Identificar claramente la ficha y sesion al registrar un seguimiento.
Mostrar el estado de las sesiones y seguimientos.
Facilitar la navegacion entre horario -> sesion -> ficha -> seguimiento.

Con estos cambios, el instructor podria gestionar su programacion y seguimiento de manera mas organizada, reduciendo errores y evitando duplicidad de procesos.