## Entendimiento General del PRD

La aerolinea necesita un sistema que le permita administrar el ciclo completo de un tiquete aereo, tomando desde que un pasajero realiza una reserva hasta que aborda el avion, ademas necesita controlar recursos como los asientos, equipaje, pagos y poder generar reportes operativos como los pasajeros que no viajaron (pagaron pero no abordaron).

---------------------------------------------------------------------

### Que Probelma Busca Resolver

Actualmente la aerolinea necesita controlar estos aspectos clave:

- Quien compra un tiquete
- Que vuelo utilizara (No. Vuelo)
- Que asiento tiene asignado (Para ese vuelo)
- Si pago o no (Un vuelo puede no estar pagado)
- Cuanto equipaje registro
- Si abordo el avion

Ademas de que necesita detectar facilmente a los pasajeros que compraron un tiquete pero nunca abordaron el avion para el vuelo.

---------------------------------------------------------------------

### Flujo Principal (Para El Unico Usuario)

El flujo planteado en el [case-1.md](./case-1.md) es el siguiente:

```text
Agente inicia sesión
        │
        ▼
Crear pasajero
        │
        ▼
Crear reserva
        │
        ▼
Emitir tiquete
        │
        ▼
Crear vuelo
        │
        ▼
Asignar asiento
        │
        ▼
Registrar pago (opcional)
        │
        ▼
Registrar equipaje (opcional)
        │
        ▼
Registrar embarque (opcional)
        │
        ▼
Consultar reporte No-Show
```

---------------------------------------------------------------------

### Actor Del Sistema

En el MVP solamente existe un actor del sistema, el cual cumple el rol de cliente / administrador (El sistema arranca con un agente administrador sembrado):

**Agente de la Aerolinea**

Este puede realizar las siguiente acciones (siguiendo el flujograma):

- Iniciar sesión
- Registrar pasajeros
- Crear reservas
- Emitir tiquetes
- Crear vuelos
- Asignar asientos
- Registrar pagos
- Registrar equipaje
- Registrar embarques
- Consultar reportes

No existen pasajeros usando directamente el sistema planteado

---------------------------------------------------------------------

### Entidades Principales

Las entidades principales se toman como tablas individuales:

| Entidad         | Propósito                                      |
| --------------- | ---------------------------------------------- |
| Passenger       | Persona que viaja                              |
| Reservation     | Solicitud previa a la compra                   |
| Ticket          | Documento que autoriza el viaje                |
| Flight          | Vuelo específico de un día                     |
| Airport         | Origen y destino                               |
| Aircraft        | Avión asignado                                 |
| Seat            | Asiento físico del avión                       |
| Seat Assignment | Asignación del asiento al pasajero en un vuelo |
| Payment         | Pago del tiquete                               |
| Baggage         | Equipaje registrado                            |
| Boarding        | Registro de que el pasajero abordó             |

---------------------------------------------------------------------

### Reglas Importantes del Negocio

Hay varias reglas que se tienen que seguir para desarrollar el sistema:

#### 1. Un vuelo se identifica por

Número de vuelo  
+  
Fecha de salida

No solo por el numero.

#### 2. Un aeropuerto cumple dos roles

La entidad Airport participa 2 veces

Vuelo

Origen ------ Airport

Destino ----- Airport

No son 2 tablas distintas

#### 3. Los asientos pertenecen a una aeronave

No existen de forma independiente

#### 4. La asignacion del asiento depende del vuelo

Un mismo asiento puede utilizarse muchas veces.

Pero:

Vuelo A

12A → Juan

Vuelo B

12A → Pedro

Eso es válido.

Lo que NO puede ocurrir es:

Vuelo A

12A → Juan

12A → María

#### 5. Boarding es opcional

Si existe: El pasajeo viajo

Si no existe: No-show

Regla importante para el reporte solicitado

---------------------------------------------------------------------