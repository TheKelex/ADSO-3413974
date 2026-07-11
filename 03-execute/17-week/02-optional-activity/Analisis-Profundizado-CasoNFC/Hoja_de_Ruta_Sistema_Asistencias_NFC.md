# SERVICIO NACIONAL DE APRENDIZAJE - SENA
### Centro de la Industria, la Empresa y los Servicios
### Neiva, Huila

# HOJA DE RUTA DEL PROYECTO FORMATIVO
## Sistema de Asistencias mediante Tecnología NFC

| Campo                 | Detalle                                               |
| --------------------- | ----------------------------------------------------- |
| Programa de formación | Tecnólogo en Análisis y Desarrollo de Software (ADSO) |
| Ficha                 | 3413974                                               |
| Aprendiz              | Martín Julián Torres Fierro    
| Aprendiz              | Juan David Gomez Martinez
| Aprendiz              | Kevin Santiago Saavedra Chantris    
| Etapa del proyecto    | Análisis y planeación                                 |
| Fecha de elaboración  | Julio de 2026                                         |

---

## 1. Información General del Proyecto

### 1.1 Nombre del proyecto
Sistema de Asistencias mediante Tecnología NFC.

### 1.2 Planteamiento del problema
El instructor manifiesta inconformidad con el tiempo que se pierde registrando la asistencia mediante el aplicativo SENA Sofía Plus, lo cual reduce el tiempo disponible para el desarrollo de la formación.

### 1.3 Objetivo general
Desarrollar un sistema que permita registrar la asistencia de forma rápida y segura mediante tecnología NFC, permitiendo que cada aprendiz gestione su propio registro mientras el instructor supervisa el proceso.

### 1.4 Objetivos específicos
- Diseñar el módulo de enrolamiento NFC que asocie el carnet institucional o el dispositivo móvil de cada aprendiz a su cuenta dentro del sistema.
- Implementar la gestión de sesiones de asistencia, incluyendo el inicio, la configuración del tiempo de tolerancia y el cierre automático o manual.
- Desarrollar el módulo de registro de asistencia mediante lectura NFC, clasificando automáticamente los estados Presente, Retardo y Ausente.
- Implementar el registro manual de asistencia para casos excepcionales, garantizando la trazabilidad de cada acción del instructor.
- Construir el módulo de reportes e historiales de asistencia por ficha, aprendiz y fecha.
- Aplicar mecanismos de seguridad que impidan registros duplicados y garanticen la auditoría de las operaciones del sistema.

---

## 2. Metodología de Desarrollo

El proyecto se desarrollará siguiendo el ciclo de un proyecto formativo SENA, alineado con una metodología ágil de trabajo por incrementos. Cada fase corresponde a un conjunto de resultados de aprendizaje y productos verificables, permitiendo el seguimiento por parte del instructor y la trazabilidad de las evidencias de formación.

Las fases contempladas son: Análisis, Diseño, Desarrollo (dividido en incrementos funcionales), Pruebas, e Implementación y Cierre. El desarrollo se organiza en incrementos para permitir entregas parciales funcionales del sistema (enrolamiento y autenticación, gestión de sesiones y registro NFC, y registro manual, reportes y seguridad).

---

## 3. Hoja de Ruta del Proyecto

### Fase 1. Análisis y Levantamiento de Requisitos (Semanas 1 a 3)

**Objetivo:** Definir y validar los requisitos funcionales y no funcionales del sistema.

**Actividades:**
- Reunión con el instructor para validar el problema y alcance del proyecto.
- Levantamiento de requisitos funcionales (RF-01 a RF-29) y no funcionales (RNF-01 a RNF-11).
- Definición de los estados de asistencia (Presente, Retardo, Ausente) y sus reglas de negocio.
- Elaboración del documento de planteamiento del problema y objetivos.

**Entregables / evidencias:**
- Documento de requisitos funcionales y no funcionales.
- Planteamiento del problema y objetivos aprobado por el instructor.

### Fase 2. Diseño del Sistema (Semanas 4 a 6)

**Objetivo:** Diseñar la arquitectura, el modelo de datos y los flujos de interacción del sistema.

**Actividades:**
- Diseño de la arquitectura general (aplicación web, módulo de registro NFC y base de datos).
- Modelado de la base de datos: aprendices, instructores, fichas, competencias, sesiones, registros de asistencia, identificadores NFC y auditoría.
- Diseño de los flujos de enrolamiento, inicio de sesión de asistencia y registro NFC.
- Diseño de las interfaces principales del instructor y del módulo de autoservicio del aprendiz.

**Entregables / evidencias:**
- Diagrama de arquitectura del sistema.
- Modelo entidad-relación de la base de datos.
- Prototipos o mockups de las interfaces principales.

### Fase 3. Desarrollo - Incremento 1: Enrolamiento y Autenticación (Semanas 7 y 8)

**Objetivo:** Construir el módulo de autenticación de instructores y el proceso de enrolamiento NFC de los aprendices.

**Actividades:**
- Desarrollo de la autenticación de instructores (RNF-03).
- Desarrollo del registro inicial del aprendiz (nombre completo y documento de identidad) (RF-21, RF-22).
- Desarrollo de la asociación del identificador NFC al aprendiz, validando unicidad (RF-23, RF-24, RNF-09).

**Entregables / evidencias:**
- Módulo funcional de autenticación de instructores.
- Módulo funcional de enrolamiento NFC.

### Fase 4. Desarrollo - Incremento 2: Sesiones y Registro NFC (Semanas 9 a 11)

**Objetivo:** Construir la gestión de sesiones de asistencia y el registro de asistencia mediante NFC.

**Actividades:**
- Desarrollo del inicio y configuración de sesión (ficha, competencia, tiempo de tolerancia) (RF-01, RF-15).
- Habilitación del lector NFC o del dispositivo móvil como escáner durante la sesión activa (RF-02).
- Desarrollo del registro de asistencia por NFC y la identificación automática del aprendiz (RF-03, RF-04, RF-05).
- Clasificación automática de los estados Presente, Retardo y Ausente, y cálculo del tiempo de retraso (RF-16, RF-17, RF-26, RF-27).
- Implementación del cierre automático de sesión al cumplirse la jornada configurada (RF-20, RF-29).

**Entregables / evidencias:**
- Módulo funcional de gestión de sesiones.
- Módulo funcional de registro de asistencia por NFC con clasificación automática de estados.

### Fase 5. Desarrollo - Incremento 3: Registro Manual, Reportes y Seguridad (Semanas 12 y 13)

**Objetivo:** Construir el registro manual de excepciones, los reportes de asistencia y los mecanismos de seguridad y auditoría.

**Actividades:**
- Desarrollo del registro manual de asistencia y captura del motivo (olvido, daño o falla del lector) (RF-10, RF-11, RF-12).
- Desarrollo de la modificación o anulación de registros con justificación (RF-08, RF-19).
- Desarrollo del cálculo de horas de formación perdidas según políticas institucionales (RF-18, RF-28).
- Desarrollo de reportes de asistencia por ficha, aprendiz y fecha (RF-09).
- Implementación de la validación de registro único por sesión y el control de sesiones activas (RF-06, RF-14).
- Implementación del historial de auditoría de registros manuales y asociaciones NFC (RNF-06, RNF-10).

**Entregables / evidencias:**
- Módulo funcional de registro manual y trazabilidad.
- Módulo de reportes e historiales de asistencia.
- Registro de auditoría del sistema.

### Fase 6. Pruebas (Semanas 14 y 15)

**Objetivo:** Verificar el cumplimiento de los requisitos funcionales y no funcionales del sistema.

**Actividades:**
- Pruebas funcionales de enrolamiento, inicio de sesión, registro NFC y registro manual.
- Pruebas de duplicidad de registros y de identificadores NFC ya asociados.
- Pruebas de carga con grupos de hasta 40 aprendices (RNF-05).
- Verificación del tiempo de lectura NFC y la tasa de éxito de lecturas (RNF-01, RNF-02).
- Corrección de errores identificados durante las pruebas.

**Entregables / evidencias:**
- Plan de pruebas y matriz de casos de prueba.
- Informe de resultados de pruebas y correcciones aplicadas.

### Fase 7. Implementación y Cierre (Semana 16)

**Objetivo:** Realizar la puesta en marcha del sistema y el cierre formal del proyecto.

**Actividades:**
- Despliegue del sistema en el ambiente definido para la prueba piloto.
- Capacitación al instructor sobre el uso de la aplicación web.
- Socialización del proceso de enrolamiento a los aprendices de la ficha.
- Entrega final de la documentación técnica y de usuario.
- Sustentación del proyecto ante el instructor.

**Entregables / evidencias:**
- Sistema desplegado en ambiente de prueba piloto.
- Manual de usuario e informe final del proyecto.
- Acta de entrega y sustentación.

---

## 4. Cronograma General (16 semanas)

El siguiente cronograma resume la distribución de las fases a lo largo de un trimestre de formación (16 semanas). Las fechas exactas deben ajustarse al calendario académico vigente de la ficha 3413974.

| Fase                       | S1-3 | S4-6 | S7-8 | S9-11 | S12-13 | S14-15 | S16 |
| -------------------------- | ---- | ---- | ---- | ----- | ------ | ------ | --- |
| 1. Análisis                | X    |      |      |       |        |        |     |
| 2. Diseño                  |      | X    |      |       |        |        |     |
| 3. Desarrollo Inc. 1       |      |      | X    |       |        |        |     |
| 4. Desarrollo Inc. 2       |      |      |      | X     |        |        |     |
| 5. Desarrollo Inc. 3       |      |      |      |       | X      |        |     |
| 6. Pruebas                 |      |      |      |       |        | X      |     |
| 7. Implementación y cierre |      |      |      |       |        |        | X   |

*Nota: la duración de cada fase es una estimación inicial y puede ajustarse según el avance real del proyecto y las indicaciones del instructor.*

---

## 5. Recursos Necesarios

### 5.1 Recurso humano
- Aprendiz(es) desarrollador(es) del proyecto.
- Instructor asesor del proyecto formativo.

### 5.2 Recurso tecnológico
- Equipo de cómputo para desarrollo.
- Lector NFC dedicado o dispositivo móvil compatible con NFC.
- Carnet(s) institucional(es) con tecnología NFC para pruebas.
- Motor de base de datos para el almacenamiento de la información.
- Ambiente de pruebas para la validación del sistema.

---

## 6. Riesgos y Mitigación

| Riesgo identificado                                   | Mitigación                                                                        |
| ----------------------------------------------------- | --------------------------------------------------------------------------------- |
| Olvido del carnet NFC.                                | Registro manual autorizado por el instructor.                                     |
| Daño del carnet NFC.                                  | Registro manual autorizado por el instructor.                                     |
| Falla del lector NFC o del dispositivo móvil.         | Continuar el registro mediante el mecanismo manual.                               |
| Intento de registrar dos veces la asistencia.         | El sistema rechaza registros duplicados.                                          |
| Intento de asociar un identificador NFC ya vinculado. | El sistema rechaza el registro e informa la situación.                            |
| Retrasos en el desarrollo de un incremento.           | Reasignación de tiempo entre incrementos y seguimiento semanal con el instructor. |
