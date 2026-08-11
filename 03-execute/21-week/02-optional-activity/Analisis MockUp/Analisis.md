## Analisis Panel Aprendiz MockUp (Sistema de Gestion de Horarios)

## Indice de contenido

- [Rol: Auth y Shell](#rol-auth-y-shell)
  - [Alcance](#alcance)
  - [Flujograma Principal](#flujograma-principal)
  - [xEntendimiento de UI / UX](#xentendimiento-de-ui-ux)
  - [Comparacion con SIGA](#comparacion-con-siga)
  - [Reingenieria](#reingenieria)
- [Rol: Aprendiz](#rol-aprendiz)
  - [Alcance](#alcance)
  - [Flujograma Principal](#flujograma-principal)
  - [Entendimiento de UI / UX](#entendimiento-de-ui-ux)
  - [Comparacion con SIGA](#comparacion-con-siga)
  - [Reingenieria](#reingenieria)
- [Rol: Instructor](#rol-instructor)
  - [Alcance](#alcance)
  - [Flujograma Principal](#flujograma-principal)
  - [Entendimiento de UI / UX](#entendimiento-de-ui-ux)
  - [Comparacion con SIGA](#comparacion-con-siga)
  - [Reingenieria](#reingenieria)
- [Rol: Coordinador](#rol-coordinador)
  - [Alcance](#alcance)
  - [Flujograma Principal](#flujograma-principal)
  - [Entendimiento de UI / UX](#entendimiento-de-ui-ux)
  - [Comparacion con SIGA](#comparacion-con-siga)
  - [Reingenieria](#reingenieria)
- [Administrador](#administrador)
  - [Alcance](#alcance)
  - [Flujograma Principal](#flujograma-principal)
  - [Entendimiento de UI / UX](#entendimiento-de-ui-ux)
  - [Comparacion con SIGA](#comparacion-con-siga)
  - [Reingenieria](#reingenieria)
- [Rol: Back-office](#rol-back-office)
  - [Alcance](#alcance)
  - [Flujograma Principal](#flujograma-principal)
  - [Entendimiento de UI / UX](#entendimiento-de-ui-ux)
  - [Comparacion con SIGA](#comparacion-con-siga)
  - [Reingenieria](#reingenieria)
- [Rol: Parametrizacion](#rol-parametrizacion)
  - [Alcance](#alcance)
  - [Flujograma Principal](#flujograma-principal)
  - [Entendimiento de UI / UX](#entendimiento-de-ui-ux)
  - [Comparacion con SIGA](#comparacion-con-siga)
  - [Reingenieria](#reingenieria)
- [Cierre del analisis general del sistema](#cierre-del-analisis-general-del-sistema)
  - [Hallazgos generales](#hallazgos-generales)
  - [Matriz general de oportunidades](#matriz-general-de-oportunidades)
  - [Reingenieria general del sistema](#reingenieria-general-del-sistema)
  - [Propuesta de flujo general reingenierizado](#propuesta-de-flujo-general-reingenierizado)
  - [Parametrizacion](#parametrizacion)
    - [Planeacion](#planeacion)
    - [Validacion](#validacion)
    - [Publicacion](#publicacion)
    - [Comunicacion](#comunicacion)
    - [Ejecucion](#ejecucion)
    - [Seguimiento](#seguimiento)
    - [Monitoreo](#monitoreo)
  - [Relacion entre los roles](#relacion-entre-los-roles)
  - [Resultado esperado](#resultado-esperado)
  - [Conclusion](#conclusion)


El sistema busca complementar la gestion de horarios del SENA, centralizando la programacion academica y facilitando su consulta, creacion, modificacion y actualizacion.

Su objetivo es mejorar la organizacion y comunicacion entre los diferentes roles, permitiendo acceder a informacion actualizada de las formaciones de manera rapida y sencilla, sin reemplazar los sistemas institucionales existentes.

--------------------------------------------------------------------------------------

## Rol: Auth y Shell

Este rol tiene funciones transversales que afectan a todos los usuarios.

### Alcance

Este apartado cuenta con 6 pantallas:

1. Login
2. Recuperar contraseña
3. Nueva contraseña
4. App Shell por rol
5. Panel de notificaciones
6. Estados globales

Esto indica que su alcance comprende el acceso al sistema, recuperacion de credenciales y elementos generales de navegacion y comunicacion que utilizan los diferentes roles.

**Las pantallas son responsive**

### Flujograma Principal

1. Ingresar mediante las credenciales.
2. En caso de olvidar la contraseña, solicitar su recuperacion.
3. Establecer una nueva contraseña.
4. Acceder al sistema segun el rol correspondiente.
5. Consultar notificaciones desde el panel general.
6. Recibir informacio sobre los diferentes estados del sistema.

Esto demuestra que Auth y Shell funciona como la base de acceso e interaccion comun para los demas roles.

### xEntendimiento de UI / UX

1. **Que entiende rapido:** Donde iniciar sesion, recuperar la contraseña y acceder a las funciones correspondientes al rol.
2. **Que no queda claro:** Que informacion representan algunos estados globales y como afectan al usuario.
3. **Que botones o textos sobran:** Las funciones principales son concretas y necesarias para el funcionamiento general.
4. **Que informacion falta:** Seria util explicar claramente el significado de los estados globales cuando estos afecten una operacion (apartado de info).
5. **Que error podria cometer:** El usuario podria interpretar incorrectamente un estado del sistema o no identificar una notificacion importante.
6. **Que consecuencia tiene ese problema:** Podria realizar una accion incorrecta o ignorar informacion relevante para su proceso.

### Comparacion con SIGA

El MockUp complementa el acceso y la interaccion general con el sistema mediante un flujo centralizado de autenticacion, navegacion y notificaciones. Su estructura permite que cada usuario acceda a las funciones correspondientes segun su rol.

Frente a SIGA, se identifican oportunidades de mejora principalmente en la centralizacion de notificaciones y comunicacion de estados, buscando que la informacion importante sea visible y comprensible para cada usuario.

### Reingenieria

La reingenieria debe mantener un acceso sencillo y fortalecer la comunicacion transversal del sistema. Se propone:

* Explicar claramente los estados globales (aparato de info).
* Mantener mecanismos sencillos para la recuperacion de credenciales.

Con estos cambios, el sistema tendria una base comun mas clara y consistente para todos los roles, facilitando el acceso y la comunicacion entre los diferentes procesos.

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
* Opcion de mostrar el cronograma fijo.
* Vista del apartado "Mi Horario" usando la misma que la de instructor.
* Opcion para ver el avance de su proyecto formativo.
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

* Relacionar claramente la disponibilidad con el horario.
* Mostrar las excepciones de forma visible (en la pantalla de Mi horario).
* Identificar claramente la ficha y sesion al registrar un seguimiento.
* Mostrar una vista general de el seguimiento de una ficha.
* Opcion para ver el horario programado (horario fijo).
* Añadir una opcion para poder anexar los documentos validos en el apartado de excepciones.

Con estos cambios, el instructor podria gestionar su programacion y seguimiento de manera mas organizada, reduciendo errores, evitando duplicidad de procesos y validando la excepcion anexada.

--------------------------------------------------------------------------------------

## Rol: Coordinador

El coordinador es uno de los roles con mayor responsabilidad dentro del sistema, ya que interviene directamente en la gestion de horarios, sesiones, conflictos, ambientes y fichas. Su participacion permite centralizar la organizacion academica y controlar que la programacion se encuentre correctamente estructurada.

### Alcance

El rol de coordinador cuenta con 12 pantallas:

7. Dashboard / Inicio
8. Horarios - lista
9. Detalle de horario
10. Crear / editar horario
11. Modal agregar / editar sesion
12. Modal confirmar publicacion
13. Panel de conflictos
14. Modal resolver conflicto
15. Disponibilidad
16. Detalle de ambiente
17. Fichas - lista
18. Detalle de ficha

Esto indica que el alcance del coordinador comprende principalmente la creacion, organizacion, publicacion y control de horarios, ademas de la gestion de disponibilidad, ambientes y fichas.

**Las pantallas son responsive**

### Flujograma Principal

El flujo principal del coordinador puede resumirse de la siguiente manera:

1. Acceder al sistema mediante sus credenciales.
2. Revisar el estado general desde el dashboard (desde conflictos hasta borradores recientes).
3. Consultar o crear un horario.
4. Agregar o editar las sesiones correspondientes.
5. Verificar posibles conflictos en la programacion.
6. Resolver los conflictos encontrados.
7. Revisar la disponibilidad de los ambientes.
8. Publicar el horario una vez validado.
9. Consultar las fichas y su informacion relacionada.

Esto demuestra que el coordinador participa en gran parte del proceso de planificacion y validacion de la programacion academica.

### Entendimiento de UI / UX

1. **Que entiende rapido:** Donde consultar los horarios, crear o editar una programacion, revisar conflictos y acceder a fichas y ambientes.
2. **Que no queda claro:** La prioridad de cada conflicto y como los cambios realizados en un horario afectan a las demas partes de la programacion.
3. **Que botones o textos sobran:** La estructura presenta las funciones principales agrupadas segun las tareas del coordinador.
4. **Que informacion falta:** Un horario centralizado (como el apartado de "Mi Horario") con la informacion general del dia / semana.
5. **Que error podria cometer:** El coordinador podria publicar un horario sin resolver completamente un conflicto o sin comprobar la disponibilidad necesaria.
6. **Que consecuencia tiene ese problema:** Podrian generarse inconsistencias en la programacion, afectando ambientes, sesiones o usuarios involucrados.

### Comparacion con SIGA

El MockUp complementa los procesos de gestion academica relacionados con la programacion de horarios, fichas y ambientes. Su propuesta centraliza en un mismo flujo actividades que requieren planificacion, validacion y publicacion.

Frente a SIGA, se identifican oportunidades de mejora en la visualizacion de conflictos, disponibilidad y estado de la programacion, buscando que el coordinador pueda tomar decisiones con mayor rapidez antes de publicar un horario.

### Reingenieria

La reingenieria debe mantener la centralizacion del sistema, pero fortalecer el proceso de validacion y publicacion. Se propone:

* Mostrar claramente el estado del horario (mediante un apartado como "Mi Horario" en el formato de los instructores).
* Apartado para ver el horario fijo.
* Opcion para ver el documento que valida la excepcion de un instructor.
* Priorizar los conflictos pendientes.
* Relacionar la programacion con la disponibilidad de ambientes.
* Mostrar advertencias antes de publicar un horario con inconsistencias.
* Facilitar la navegacion entre horario -> sesion -> conflicto -> ambiente -> ficha.

Con estos cambios, el coordinador podria gestionar la programacion de manera mas organizada y reducir errores antes de que los horarios sean publicados.

--------------------------------------------------------------------------------------

## Administrador

El administrador tiene un rol orientado al control general del sistema, ya que puede consultar indicadores, administrar usuarios y gestionar los datos de referencia utilizados por las demas funciones.

### Alcance

El rol de administrador cuenta con 8 pantallas:

29. Panel de indicadores.
30. Drill-down de KPI.
31. Usuarios - lista.
32. Crear / editar usuario.
33. Detalle de usuario.
34. Modal asignar / revocar rol.
35. Datos de referencia.
36. Editar catalogo / valor / parametro.

Esto indica que su alcance se centra en monitorear el funcionamiento del sistema, administrar usuarios y controlar los parámetros y datos que sirven como base para los demas modulos.

**Las pantallas son responsive**

### Flujograma Principal

El flujo principal del administrador puede resumirse de la siguiente manera:

1. Acceder al sistema mediante sus credenciales.
2. Consultar los indicadores generales desde el panel.
3. Profundizar en un indicador para revisar informacion especifica.
4. Administrar los usuarios registrados.
5. Crear, editar o consultar la informacion de un usuario.
6. Asignar o revocar roles segun corresponda.
7. Gestionar los datos de referencia y parametros del sistema.

Esto demuestra que el administrador actua principalmente sobre la configuracion, control y supervision general del sistema, mas que sobre la gestion directa de horarios.

### Entendimiento de UI / UX

1. **Que entiende rapido:** Donde consultar indicadores, administrar usuarios y acceder a los datos de referencia.
2. **Que no queda claro:** Que impacto tienen algunos cambios administrativos sobre los demas modulos del sistema.
3. **Que botones o textos sobran:** Las funciones principales estan agrupadas segun las tareas administrativas.
4. **Que informacion falta:** Seria util mostrar el impacto o alcance de los cambios realizados sobre usuarios, roles y parametros ademas de que al desplazar el mouse por las barras de la grafica esta nos de mas informacion.
5. **Que error podria cometer:** El administrador podria modificar un rol o parametro incorrectamente.
6. **Que consecuencia tiene ese problema:** Podria afectar los permisos de un usuario o el funcionamiento de otros procesos del sistema.

### Comparacion con SIGA

El MockUp complementa las funciones administrativas necesarias para controlar usuarios, permisos y datos utilizados por el sistema. Su enfoque permite centralizar estas operaciones y separarlas de las funciones propias de la gestion academica.

Frente a SIGA, se identifican oportunidades de mejora principalmente en la visualizacion del impacto de los cambios administrativos y el control de los parametros, buscando reducir errores en configuraciones que afectan al resto del sistema.

### Reingenieria

La reingenieria debe enfocarse en fortalecer el control y seguridad de las operaciones administrativas. Se propone:

* Mostrar claramente el impacto de cada cambio.
* Incorporar confirmaciones para modificaciones importantes.
* Diferenciar claramente los roles y permisos.
* Mostrar el estado de los usuarios.
* Informacion extra en las barras de la grafica.
* Registrar los cambios realizados sobre usuarios y parametros.

Con estos cambios, el administrador podria controlar el sistema de manera mas segura y reducir el riesgo de afectar otros procesos por una configuracion incorrecta.

--------------------------------------------------------------------------------------

## Rol: Back-office

El usuario de Back-Office cumple una funcion de soporte y administracion operativa del sistema, enfocandose principalmente en la gestion documental, auditoria y parametrizacion de informacion que permite mantener el sistema organizado y funcionando correctamente.

### Alcance

El rol de Back-Office cuenta con 9 pantallas:

37. Documentos - lista.
38. Plantillas de documento.
39. Auditoria.
40. Parametrizacion / catalogos.
41. Detalle de documento + versiones.
42. Modal generar documento.
43. Editor / preview de plantilla.
44. Modal detalle de auditoria.
45. CRUD catalogo / valor / parametro.

Esto indica que su alcance se centra en la gestion de documentos, plantillas, auditoria y parametros del sistema, sirviendo como apoyo para mantener actualizada y controlada la informacion utilizada por los demas roles.

**Las pantallas son responsive**

### Flujograma Principal

El flujo principal del usuario de Back-Office puede resumirse de la siguiente manera:

1. Acceder al sistema mediante sus credenciales.
2. Consultar y gestionar los documentos existentes.
3. Crear o modificar plantillas de documentos.
4. Generar documentos segun las plantillas disponibles.
5. Consultar la auditoria y revisar los cambios realizados.
6. Gestionar los catalogos, valores y parametros del sistema.

Esto demuestra que el Back-Office se encarga principalmente de mantener, controlar y respaldar la informacion operativa del sistema.

### Entendimiento de UI / UX

1. **Que entiende rapido:** Donde gestionar documentos, plantillas, auditorias y parametros.
2. **Que no queda claro:** La relacion entre los documentos, sus versiones y las plantillas utilizadas para generarlos.
3. **Que botones o textos sobran:** Las funciones se encuentran separadas de acuerdo con las tareas principales del usuario.
4. **Que informacion falta:** Seria util identificar claramente el estado, version y ultima modificacion de los documentos y plantillas, ademas de tener subdivisiones de las plantillas para que sea facil ubicarlas.
5. **Que error podria cometer:** El usuario podria modificar una plantilla o parametro incorrectamente.
6. **Que consecuencia tiene ese problema:** Podrian generarse documentos incorrectos o afectar procesos que dependen de dichos parametros.

### Comparacion con SIGA

El MockUp complementa los procesos administrativos mediante herramientas para gestionar documentos, controlar cambios y administrar parametros. Esto permite centralizar tareas de soporte que pueden ser necesarias para mantener actualizada la informacion del sistema.

Frente a SIGA se identifican oportunidades de mejora en el control de versiones, trazabilidad de cambios y relacion entre documentos y parametros, permitiendo un mayor control sobre la informacion.

### Reingenieria

La reingenieria debe enfocarse en mejorar el control documental y la trazabilidad. Se propone:

* Mostrar claramente la version y estado de cada documento.
* Mantener un historial de cambios realizados (en cuestiones de plantillas / documentos).
* Relacionar las plantillas con los documentos generados.
* Solicitar confirmacion antes de modificar parametros importantes.
* Facilitar la consulta de quien realizo cada cambio y cuando.
* Permitir ingresar datos sin formato JSON (para usuarios no experimentados).

Con estos cambios, el usuario de Back-Office podria mantener la informacion del sistema de forma mas organizada, controlada y segura.

--------------------------------------------------------------------------------------

## Rol: Parametrizacion

El rol de parametrizacion se encarga de configurar los datos base y reglas que utiliza el sistema, permitiendo mantener actualizada la estructura academica, los horarios, los ambientes, los indicadores, los estados, la informacion geografica y los permisos.

**Este no es un rol en especifico, es mas un apartado del sistema al cual ciertos usuarios pueden acceder**

### Alcance

El rol de parametrizacion cuenta con 8 pantallas:

46. Hub de parametrizacion
47. Curriculo academico
48. Jornadas / franjas horarias
49. Tipos de ambiente e inventario
50. Catalogos de monitoreo (KPI/alertas)
51. Estados de actores
52. Geografia institucional
53. RBAC - roles y permisos

Esto indica que su alcance comprende la configuracion general del sistema, estableciendo los datos y reglas que posteriormente seran utilizados por los demas roles y procesos.

**Las pantallas son responsive**

### Flujograma Principal

El flujo principal del rol de parametrizacion puede resumirse de la siguiente manera:

1. Acceder al sistema mediante sus credenciales.
2. Ingresar al apartado de parametrizacion.
3. Seleccionar el componente que se desea configurar.
4. Gestionar la informacion correspondiente.
5. Actualizar los parametros necesarios para mantener el funcionamiento del sistema.

Esto demuestra que el rol actua principalmente como base de configuracion para el resto del sistema.

### Entendimiento de UI / UX

1. **Que entiende rapido:** Donde acceder a cada categoria de configuracion.
2. **Que no queda claro:** El impacto que puede tener modificar determinados parametros sobre los demas procesos.
3. **Que botones o textos sobran:** Las categorias estan separadas de acuerdo con su funcion, facilitando la navegacion.
4. **Que informacion falta:** Seria util indicar que procesos o modulos dependen de cada configuracion.
5. **Que error podria cometer:** Modificar incorrectamente un parametro que este siendo utilizado por otros procesos.
6. **Que consecuencia tiene ese problema:** Podrian generarse inconsistencias en horarios, ambientes, permisos o informacion academica.

### Comparacion con SIGA

El MockUp complementa la gestion institucional mediante la configuracion centralizada de informacion academica, administrativa y de seguridad. Esto permite establecer una base comin para los diferentes procesos del sistema.

Frente a SIGA, se identifican oportunidades de mejora en la visibilidad del impacto de los cambios y el control de las configuraciones, especialmente en parametros que afectan directamente otros modulos.

### Reingenieria

La reingenieria debe enfocarse en mejorar el control y trazabilidad de la parametrizacion. Se propone:

* Mostrar el impacto de cada configuracion antes de modificarla.
* Solicitar confirmacion para cambios criticos.
* Mantener un historial de modificaciones.
* Relacionar cada parametro con los modulos que lo utilizan.
* Diferenciar claramente las configuraciones activas e inactivas.

Con estos cambios, el rol podria realizar modificaciones de forma mas segura y reducir el riesgo de afectar otros procesos del sistema.

--------------------------------------------------------------------------------------

## Cierre del analisis general del sistema

### Hallazgos generales

Despues de analizar las 53 pantallas del sistema se puede evidenciar que el MockUp busca complementar la gestion de horarios del SENA mediante la centralizacion de diferentes procesos relacionados con la programacion academica.

Los principales hallazgos se concentran en:

- Informacion distribuida entre diferentes procesos y roles.
- Necesidad de priorizar informacion importante como proximas sesiones, conflictos y notificaciones.
- Relacion poco visible entre horarios, disponibilidad, ambientes y excepciones.
- Necesidad de mayor trazabilidad sobre cambios, documentos, parametros y acciones administrativas.
- Necesidad de validaciones que reduzcan errores antes de publicar o modificar informacion.
- Necesidad de mantener una comunicacion clara entre los diferentes roles del sistema.

### Matriz general de oportunidades

| Area | Oportunidad de mejora | Impacto |
| --- | --- | --- |
| Horarios | Centralizar y priorizar la informacion de las sesiones (Horario Fijo - Horario con Excepciones) | Facilita la consulta y reduce confusiones |
| Notificaciones | Diferenciar cambios y notificaciones pendientes | Evita ignorar informacion importante |
| Disponibilidad | Relacionarla directamente con la programacion | Reduce conflictos de horario |
| Conflictos | Priorizar y validar conflictos antes de publicar | Reduce errores en la programacion |
| Seguimiento | Mejorar la relacion entre ficha, sesion y seguimiento | Facilita el control del proceso formativo |
| Documentos | Incorporar versiones e historial de cambios | Mejora la trazabilidad |
| Administracion | Mostrar el impacto de modificaciones | Reduce errores de configuracion |
| Parametrizacion | Relacionar parametros con los modulos que utilizan | Facilita el control del sistema |
| Acceso y comunicacion | Centralizar autenticacion, notificaciones y estados | Mejora la experiencia general |

### Reingenieria general del sistema

La reingenieria debe mantener las funciones principales del MockUp, pero integrar mejor los procesos que actualmente se encuentran separados.

Se propone:

1. Centralizar la gestion de horarios como eje principal del sistema (Horario Fijo - Horario con Excepciones).
2. Relacionar directamente horarios, sesiones, disponibilidad, ambientes y conflictos.
3. Incorporar validaciones antes de publicar o modificar una programacion.
4. Priorizar informacion importante mediante alertas, estados y notificaciones.
5. Mantener trazabilidad de cambios realizados en documentos, usuarios y parametros.
6. Mostrar el impacto de cambios administrativos o de parametrizacion.
7. Facilitar la navegacion entre los procesos relacionados.
8. Mantener una interfaz sencilla y adaptada a cada rol.
9. Utilizar los terminos propios del SENA para evitar confusiones.
10. Complementar los sistemas institucionales existentes sin reemplazarlos.

### Propuesta de flujo general reingenierizado

El funcionamiento general propuesto puede representarse de la siguiente manera:

**Parametrizacion -> Planeacion -> Validacion -> Publicacion -> Comunicacion -> Ejecucion -> Seguimiento -> Monitoreo**

### Parametrizacion

Se establecen los datos base del sistema, como curriculo, jornadas, ambientes, estados, geografia, indicadores y permisos.

#### Planeacion

El coordinador crea y organiza los horarios teniendo en cuenta las sesiones, fichas, instructores y ambientes disponibles.

#### Validacion

El sistema identifica conflictos y permite revisar disponibilidad y excepciones antes de publicar.

#### Publicacion

Una vez validado el horario, este se publica y se generan las comunicaciones correspondientes.

#### Comunicacion

Los diferentes usuarios reciben la informacion que corresponde a su rol, incluyendo cambios, alertas y notificaciones.

#### Ejecucion

El aprendiz consulta su horario y el instructor gestiona sus sesiones, disponibilidad y seguimiento.

#### Seguimiento

El instructor registra el avance y las novedades relacionadas con las fichas y los procesos formativos.

#### Monitoreo

Los usuarios administrativos pueden revisar indicadores, auditoria, documentos y configuraciones para mantener el control general del sistema.

### Relacion entre los roles

La reingenieria propuesta debe permitir que los roles trabajen sobre un mismo flujo de informacion:

**Administrador / Parametrizacion**
-> Define usuarios, permisos, parametros y datos base.

**Coordinador**
-> Planifica, valida y publica los horarios.

**Instructor**
-> Consulta su programacion, gestiona disponibilidad y realiza seguimiento.

**Aprendiz**
-> Consulta su horario y recibe los cambios de su programacion.

**Back-Office**
-> Administra documentos, plantillas, auditoria y elementos de soporte.

Esta relacion permite que la informacion generada en un proceso sea aprovechada por los demas sin necesidad de duplicarla.

### Resultado esperado

Con la reingenieria propuesta se busca que el sistema:

- Reduzca errores en la programacion.
- Evite duplicidad de procesos e informacion.
- Facilite la consulta del horario.
- Mejore la comunicacion entre roles.
- Permita detectar conflictos antes de publicar.
- Aumente la trazabilidad de cambios.
- Facilite el control de la informacion.
- Mantenga una experiencia sencilla y adaptada a cada usuario.

### Conclusion

El analisis de las 53 pantallas permitio identificar que el MockUp presenta una propuesta integral para complementar la gestion de horarios y procesos relacionados del SENA. Cada rol participa en una parte diferente del proceso, pero existen oportunidades de mejora principalmente en la integracion, priorizacion, validacion y trazabilidad de la informacion.

La reingenieria propuesta no busca reemplazar los sistemas institucionales existentes, sino complementarlos mediante una herramienta que conecte los diferentes procesos y roles alrededor de la gestion de horarios ademas de que permite un uso intuitivo y rapido por parte de los diferentes usuarios que participan en este.

De esta manera, el sistema puede pasar de ser un conjunto de modulos independientes a una solucion integrada, donde la informacion fluye desde la parametrizacion y planeacion hasta la publicacion, comunicacion, seguimiento y monitoreo de las formaciones.