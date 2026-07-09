### Proyecto Asistencia Sena

**Planteamiento general:** Se requiere un sistema que permita registrar asistencia de la manera mas eficiente (Sin tener en cuenta la integracion con Sofia Plus).

**Solucion propuesta:** Se propone una solucion mediante una pagina web que agiliza el proceso de toma de asistencia y mejora el control de esta mediante los siguientes items:

1. Autentificacion de usuarios.
2. Generacion de un QR dinamico cambiante cada (5 minutos) para la sesion de la clase previamente registrada.
3. Restriccion de la red institucional del Sena.
4. Un registro unico por aprendiz para cada clase.
5. Control de historial de asistencia.

### Funcionamiento General del Sistema

#### Creacion de Sesion / Inicio Sesion Instructor

El intructor:

1. Inicia sesion.
2. Selecciona la ficha (previamente registrada).
3. Inicia sesion de asistencia:
    1. Selecciona la competencia (previamente registrada).
    2. Ingresa un tiempo maximo para reigistrar la asistencia (en minutos).

El sistema:

1. Crea la sesion de asistencia.
2. Crea un QR temporar (cambio cada 5 min).
3. Habilita el registro por el tiempo indicado por el instructor.

#### Registro de Asistencia Aprendiz

El aprendiz:

1. Se conecta al internet institucional.
2. Escanea el QR proyectado por el instructor.
3. Accede con sus credenciales.
4. El sistema valida:
    - Que la sesion este activa.
    - Que se encuentre en la red institucional (verifica IP publica).
    - Que no haya registrado asistencia para esa formacion anteriormente.

#### Registro de Aprendices

El instructor:

1. Copia los listados desde Sofia Plus
2. Los separa mediante enters.
3. Los copia y los pega en el sistema.

El sistema:

- Los registra con su numero de documento

-------------------------------------------------------------------------------------

### Seguridad

#### QR Dinamico

Cada sesion genera un codigo QR diferente y cambiante (cada 5 min).

Eso evita que:

- Reutilicen los codigos antiguos.
- Manipulen el proceso de validacion.

#### Restriccion por Red del Sena

Solo los dispositivos conectados a la red del Sena pueden reigstrar la asistencia.

Eso evita que:

- Se realicen registros desde casa.
- Se realicen registros desde otras ubicacion.

#### Registro Unico 

Cada aprendiz solo puede registrar una unica asistencia por formacion (sesion).

Eso evita que:

- Problemas de duplicacion de registros.
- La manipulacion de las estadisticas.

-------------------------------------------------------------------------------------

### Riesgos Encontrados

Bajo todos estos riegos se cuenta con la supervision directa del instructor.

#### Compartir el Codigo QR Con Otra Persona del Centro de Formacion.

Solucion:

- Tiempo de expiracion del codigo QR corto.
- Supervision directa del instructor.

#### Caida de la Red del Sena

Solucion:

- Metodo alternativo de registro manual.
- Opcion para desactivar la verificacion de la red del Sena.

#### Posible Aprendiz Sin Celular

Solucion:

- Posibilidad de registrar una asistencia manualmente por parte del instructor.
- Unicamente ingresar su numero de documento.

#### QR Refrescado Durante un Registro

Solucion:

- Alerta avisando sobre el QR refrescado
- Opcion para volver a scanear.