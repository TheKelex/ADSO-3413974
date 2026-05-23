# Entendimiento inicial - Sistema de Horarios SENA

## Control de scope

- Framework de referencia: `C:\www\code-dev-projects\automatization-develop` en modo solo lectura.
- Workspace del producto: `C:\www\code-dev-projects\design-software`.
- Proyecto operativo asumido: `PRJ-SENA-HORARIOS`.
- Fase activa: `DISCOVERY_MINIMO_PRODUCTO`.
- Agente activado: `A05 Product Owner`.
- Capacidad aplicada: `cap-product-discovery`.

## Contexto del problema

La programacion de formacion se gestiona actualmente en Excel. Ese archivo funciona como registro operativo por dia, pero mezcla datos academicos, disponibilidad, asignaciones, observaciones y posibles novedades en una misma superficie. Para el MVP, el Excel debe tratarse como fuente de entendimiento del dominio y de lenguaje operativo, no como modelo definitivo de datos ni como diseno de interfaz.

El problema central es organizar la programacion de formacion de manera consultable y validable, relacionando instructores, fichas, ambientes de formacion, horarios, programas, competencias, RAP y observaciones. La primera version debe ayudar a registrar y consultar horarios basicos, identificar dias sin programacion y advertir posibles cruces entre instructor, ficha y ambiente.

## Necesidad del producto

El producto necesita reemplazar una gestion dispersa en hoja de calculo por una aplicacion web con API centralizada y persistencia en PostgreSQL. La necesidad inicial no es automatizar todo el ecosistema SENA, sino construir una base confiable para programacion de horarios con reglas de negocio concentradas en backend y preparada para evolucionar hacia otros modulos institucionales.

## Datos observados en el Excel descrito

No se analizo un archivo Excel fisico en esta fase. Se toma como entrada la estructura funcional descrita por el solicitante:

- Programa de formacion.
- Nivel.
- Ficha.
- Ambiente de formacion.
- Competencia.
- Resultado de aprendizaje / RAP.
- Fecha de inicio de actividad.
- Fecha fin de actividad.
- Instructor asignado.
- Observaciones o novedades.
- Dias sin programacion.

Estos elementos se consideran senales del dominio. No implican, por si solos, que cada columna deba convertirse en tabla, atributo obligatorio o pantalla independiente.

## Actores iniciales

- Coordinador o responsable de programacion: registra, ajusta y consulta la programacion.
- Instructor: persona asignada a actividades de formacion; puede ser de planta o contratista.
- Aprendiz: beneficiario de la formacion asociado a una ficha; para el MVP no se gestiona individualmente.
- Administrador funcional: mantiene catalogos basicos como ambientes, instructores, fichas y programas.
- Consultor institucional: revisa horarios, novedades y posibles conflictos sin necesariamente editar.

## Entidades candidatas

Estas entidades son candidatas y no incluyen todavia atributos detallados:

- Instructor.
- Tipo de instructor.
- Ficha.
- Programa de formacion.
- Nivel.
- Ambiente de formacion.
- Horario.
- Bloque o franja de programacion.
- Competencia.
- RAP.
- Actividad de aprendizaje.
- Observacion o novedad.
- Dia sin programacion.
- Asignacion de formacion.

## Relaciones candidatas

- Un instructor puede tener varias asignaciones de horario.
- Una ficha puede tener varias asignaciones en diferentes dias o franjas.
- Un ambiente de formacion puede ser usado por varias fichas en momentos distintos.
- Una asignacion relaciona instructor, ficha, ambiente, fecha o rango de fechas, franja horaria y contenido formativo.
- Una competencia puede relacionarse con uno o varios RAP.
- Una actividad de aprendizaje puede asociarse a un RAP cuando aplique.
- Una observacion o novedad puede asociarse a una asignacion, ficha, instructor, ambiente o dia.
- Un dia sin programacion puede asociarse a una ficha, instructor o periodo, segun confirmacion de negocio.

## Reglas de negocio candidatas

- Un instructor no deberia estar asignado a dos horarios que se crucen en la misma fecha y franja.
- Una ficha no deberia tener dos actividades que se crucen en la misma fecha y franja.
- Un ambiente de formacion no deberia estar reservado para dos asignaciones cruzadas.
- Toda asignacion MVP deberia relacionar al menos ficha, instructor, ambiente y franja de horario.
- El tipo de instructor debe permitir distinguir planta y contratista.
- Las observaciones o novedades deben conservar contexto de la programacion afectada.
- Los dias sin programacion deben registrarse de forma explicita cuando sean relevantes para seguimiento.
- Las competencias, RAP y actividades de aprendizaje deben quedar abiertas a confirmacion institucional antes de cerrar atributos obligatorios.

## Limites de entendimiento

- No se inventan fichas, codigos, competencias, RAP, programas, centros, sedes ni regionales.
- No se define una taxonomia oficial SENA sin fuente institucional.
- No se copia literalmente el Excel como modelo de base de datos.
- No se define todavia un flujo completo de asistencia, matricula, inventario o evaluacion.
- No se decide una arquitectura de microservicios final.

## Pendientes por confirmar

- Archivo Excel real y ejemplos anonimizados de filas representativas.
- Regla institucional para fechas de inicio y fin de actividad frente a horarios por dia.
- Definicion operativa de "dia sin programacion".
- Si un horario se modela por rango de fechas, por dia individual o por recurrencia.
- Catalogos oficiales existentes para programas, competencias, RAP y ambientes.
- Roles reales de usuarios que editaran, aprobaran o solo consultaran.
- Politica de auditoria para cambios de programacion y novedades.

## Riesgos de modelado

- Convertir columnas del Excel en tablas sin entender el proceso real.
- Mezclar planeacion academica, programacion operativa y seguimiento de novedades en una sola entidad.
- Subestimar cruces de horario por no definir correctamente fecha, hora y recurrencia.
- Hacer obligatorios datos institucionales que pueden no estar disponibles al inicio.
- Acoplar validaciones al frontend y duplicar reglas fuera de la API.

## Recomendacion de siguiente activacion

Activar A05 nuevamente en capacidad de definicion de producto solo despues de validar este discovery minimo con un ejemplo real o anonimizacion del Excel. Si se decide avanzar a diseno tecnico sin mas insumos, activar A06 para patrones y arquitectura inicial, usando estos artefactos como entrada.
