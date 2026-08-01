## Design Thinking

Esto se hace para buscar una solucion al problema planteado centrandonos en las personas para realizar un sistema que resuelva las necesidades reales de los usuarios

### 1. Empatizar

El usuario principal es el agente de la aerolinea

**Necesidades**

- Registrar pasajeros rapidamente
- Encontrar informacion de un vuelo
- Evitar errores en la asignacion de asientos
- Registrar pagos y equipaje de forma sencilla
- Saber quien abordo
- Identificar pasajeros no presentados (no-show)

**Dolores (Pain Points)**

- Duplicacion de asientos
- Errores manuales
- Dificultad para encontrar pasajeros
- Procesos lentos durante el embarque
- Reportes manuales

Ademas de que al ser una aerolinea se debe hacer un sistema rapido, confiable, sostenible e intiutivo para no retrasar los ciclos completos de un tiquete aereo (definido en el archivo [Entendimiento.md](./Entendimiento.md))

### 2. Definir

El problema central es que el agente necesita gestionar de manera rapida, segura y confiable todo el ciclo de un pasajero, evitando inconsistencias (como asignar el mismo asiento dos veces) y permitiendo identificar facilmente los pasajeros que no abordaron el vuelo

### 3. Idear

Posibles soluciones:

- Flujo guiado de creacion de reserva -> tiquete -> vuelo
- Validaciones automaticas de reglas de negocio
- Consulta rapida por documento del pasajero
- Asignacion de asientos mostrando solo los disponibles
- Reporte automatico de no-show
- Mensajes de error claros
- Confirmaciones antes de operaciones criticas

(Las posibles soluciones planteadas no se desvian de lo propuesto en el [case-1.md](./case-1.md))

### 4. Prototipar

Pantallas minimas del MVP:

1. Login
2. Pasajeros
3. Reservas
4. Tiquetes
5. Vuelos
6. Asignacion de asientos
7. Pagos
8. Equipaje
9. Embarque
10. Reporte No-Show

### 5. Testear

Los escenarios clavess para verificar son:

- Inicio de sesion con credenciales validas e invalidas
- Creacion de un pasajero
- Creacion de una reserva
- Emision de un tiquete
- Creacion de un vuelo con origen y destino diferentes
- Intento de asignar el mismo asiento dos veces en el mismo vuelo (debe fallar)
- Registro de equipaje y pagos
- Registro de embarque
- Consulta del reporte de no-show, verificando que solo aparezcan los pasajeros con tiquete y sin registro de embarque
- Verificacion de que las operaciones criticas sean atomicas y mantengan la integridad de los datos