### MockUp

Todo esto se realizo con un contexto previamente dado, a continuacion se listan las herramientas usadas para crear el MockUp:

- ChatGPT
- Figma (Create)

#### Prompt Dado

Necesito 4 paneles importante, se deben usar los terminos de apropiacion Sena (Instructor, aprendiz, ficha: 3413974, competencia: Instrumentos de recoleccion de informacion) evitar usar diseño generico de IA:

1. Apartado para crear la sesión:
- Seleccionar Ficha
- Seleccionar Competencia
- Indicar Tiempo Maximo para Registrarse
Indicaciones Visuales: El panel debe poder permitir ingresar todos estos datos uno detras de otro, osea, primero un menu de seleccioanr ficha, una vez seleccionada le da a siguiente y aparece el otro menu borrando el de la ficha y aparece el de la competencia y asi susecibamente (apartado vista computador)

2. Apartado vista QR:
- Una vez creada la sesion se abre este apartado
- Codigo QR gigante 
- Barra de carga bajo el codigo QR que indica cada cuanto se va a recargar
- Abajo timer que indica cuanto tiempo queda para registrarse
Indicaciones Visuales: El codigo QR debe ser muy grande, debe haber un boton para cancelar la asistencia y para pausarla, ademas debe haber un boton para registrar una asistencia manualmente en la que se despliega un menu al lado que no reduce el tamaño del codigo QR

3. Apartado para registrar aprendices:
- Seleccionar competencia
- Aparato (big text) para insertar los datos y que el instructor los separe por enters
Indicaciones Visuales: Debe ser facil de entender que es lo que se debe hacer, ademas de que debe haber un boton que lo redirija automaticamente al inicio de sesion de sofia plus, ademas de que se debe indicar que deben estar separados por enters los datos

4. Apartado registrar asistencia:
- Una vez escaneado el codigo QR se piden las credenciales del aprendiz (No. documento y contraseña)
- Se verifica que se encuentre conectado a la red
- Darle al boton de registrar asistencia
Indicaciones Visuales: FrontEnd para celulares, debe tener un flujo claro y en caso de que se recargue el codigo QR durante el registro debe haber un boton que diga "escanear nuevamente el QR" y que le permita escanearlo de nuevo para recargar el token @Razonamiento 

#### Prompt Entregado (ChatGPT)

Diseña un mockup de una plataforma web institucional para control de asistencia del SENA. Debe verse moderna, profesional, limpia y realista, evitando por completo el estilo genérico de IA. No usar estética futurista exagerada, ni fondos oscuros con neones, ni tarjetas infladas artificialmente. Quiero una interfaz sólida, usable y seria, con apariencia de producto real.

Usa estos términos exactamente como parte de la apropiación institucional:

* Instructor
* Aprendiz
* Ficha: 3413974
* Competencia: Instrumentos de recolección de información

Estilo visual:

* Diseño tipo SaaS moderno
* Mucho espacio en blanco
* Tarjetas con bordes redondeados suaves
* Sombras sutiles
* Tipografía clara y académica
* Botones bien jerarquizados
* Colores institucionales del SENA: verde, blanco y grises suaves
* Íconos simples y funcionales
* Nada recargado, nada caricaturesco, nada genérico de IA

La interfaz debe incluir 4 paneles importantes:

1. Apartado para crear la sesión
   Este panel es para vista computador.
   Debe permitir crear la sesión paso a paso, como un flujo tipo wizard o stepper.
   Primero se muestra un menú para seleccionar la Ficha.
   Cuando se selecciona la ficha, al presionar “Siguiente” desaparece ese campo y aparece el menú de Competencia.
   Luego, al continuar, aparece el campo para indicar el Tiempo Máximo para Registrarse.
   La idea es que los campos se ingresen uno detrás de otro, no todos al mismo tiempo.
   Debe verse claro que el instructor está creando una sesión de asistencia.
   Elementos visibles:

* Seleccionar Ficha
* Seleccionar Competencia
* Indicar Tiempo Máximo para Registrarse
* Botón “Siguiente”
* Botón “Atrás”
* Botón “Crear sesión”

2. Apartado vista QR
   Este panel se abre después de crear la sesión.
   Debe tener un código QR muy grande y protagonista al centro.
   Debajo del QR debe haber una barra de carga que indique visualmente el tiempo de recarga del QR.
   También debe aparecer un temporizador que indique cuánto tiempo queda para registrarse.
   Debe incluir:

* Botón para cancelar la asistencia
* Botón para pausar la asistencia
* Botón para registrar una asistencia manualmente
  Cuando se abre el menú de registro manual, debe aparecer al lado o como panel lateral sin reducir el tamaño del QR.
  El QR debe seguir viéndose gigante y dominante.
  Debe sentirse como una pantalla de proyección para el instructor.

3. Apartado para registrar aprendices
   Este panel también es para el Instructor.
   Debe permitir seleccionar la competencia y luego pegar una lista grande de aprendices.
   El área de texto debe ser muy visible y clara, con tipografía grande para que el instructor entienda que debe pegar los datos separados por enter.
   Debe ser obvio que el instructor debe copiar y pegar los aprendices uno por línea.
   Elementos visibles:

* Seleccionar Competencia
* Área de texto grande para pegar aprendices
* Texto de ayuda que diga claramente que los datos deben estar separados por enters
* Botón que redirija automáticamente al inicio de sesión de Sofia Plus
* Botón para guardar o importar la lista

4. Apartado registrar asistencia
   Este panel es para el Aprendiz y debe verse pensado para celulares.
   Debe tener un flujo claro, simple y rápido.
   Una vez escaneado el QR, se piden las credenciales del aprendiz:

* Número de documento
* Contraseña
  Debe verificarse visualmente que se encuentra conectado a la red institucional.
  Debe existir un botón principal para “Registrar asistencia”.
  Si el código QR se recarga mientras el aprendiz está en el proceso, debe aparecer un botón que diga “Escanear nuevamente el QR” para volver a leer el token.
  La pantalla debe verse muy clara, simple y centrada en la acción.

Requisitos de diseño general:

* La versión del Instructor debe verse para escritorio
* La versión del Aprendiz debe verse para celular
* Debe existir una identidad visual coherente entre todos los paneles
* Debe sentirse como una solución real para el SENA
* No usar elementos decorativos innecesarios
* No usar estética de plantilla genérica
* No usar texto en inglés salvo que sea necesario para la herramienta
* No usar estilos infantiles ni demasiado creativos

Quiero que el resultado se vea como un sistema institucional moderno, funcional y listo para presentar en sustentación.

#### Link Acceso a MockUp (Figma)

El MockUp se realizo mediante la herramienta "Create" de Figma:

[Clic Vista Previa Directa MockUp](https://www.figma.com/make/9g8F88RwQU5HXwTbZFVk8J/Mockup-plataforma-control-asistencia?t=PB8uQ3Gj1BtoLyQg-20&fullscreen=1)



**Usar el siguiente link en caso de que el primero no funcione:**

[Clic Acceso a Proyecto Figma](https://www.figma.com/make/9g8F88RwQU5HXwTbZFVk8J/Mockup-plataforma-control-asistencia?t=PB8uQ3Gj1BtoLyQg-1)