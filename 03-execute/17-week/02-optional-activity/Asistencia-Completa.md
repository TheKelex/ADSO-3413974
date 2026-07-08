### Proyecto Asistencia Sena

**Planteamiento general:** Se requiere un sistema que permita registrar asistencia de la manera mas eficiente (Sin tener en cuenta la integracion con Sofia Plus).

**Solucion propuesta:** Se propone una solucion mediante una pagina web que agiliza el proceso de toma de asistencia y mejora el control de esta mediante los siguientes items:

1. Autentificacion de usuarios.
2. Generacion de un QR dinamico cambiante cada (5 minutos) para la sesion de la clase previamente registrada.
3. Restriccion de la red institucional del Sena.
4. Un registro unico por aprendiz para cada clase.
5. Control de historial de asistencia.

### Flujos

#### Funcionamiento General

##### Creacion de Sesion / Inicio Sesion Instructor

El intructor:

1. Inicia sesion.
2. Selecciona la ficha o grupo (previamente registrada).
3. Inicia sesion de asistencia:
    1. Selecciona la competencia (previamente registrada).
    2. Ingresa un tiempo maximo para reigistrar la asistencia (en minutos).

El sistema:

1. Crea la sesion de asistencia.
2. Crea un QR temporar (cambio cada 7 min).
3. Habilita el registro por el tiempo indicado por el instructor.

##### Registro de Asistencia Aprendiz

El aprendiz:

1. Se conecta al internet institucional.
2. Escanea el QR proyectado por el instructor.
3. Accede con sus credenciales.
4. El sistema valida:
    - Que la sesion este activa.
    - Que se encuentre en la red institucional (verifica IP publica).
    - Que no haya registrado asistencia para esa formacion anteriormente.