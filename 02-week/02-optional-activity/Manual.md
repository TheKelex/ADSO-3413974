# Tutorial Paso a Paso: Usar Git desde la Terminal

Este tutorial explica **cómo usar Git desde cero**, desde instalarlo hasta **hacer un commit y subir archivos a GitHub**.

---

# 1. Instalar Git

## Paso 1: Descargar Git

1. Abre tu navegador (Chrome, Edge, Firefox, etc.).
2. Ve a la página oficial:

https://git-scm.com

3. Haz clic en **Download for Windows**.

---

## Paso 2: Instalar Git

1. Abre el archivo descargado.
2. Se abrirá el instalador de Git.
3. Haz clic en **Next** varias veces.
4. Deja todas las opciones **por defecto**.
5. Haz clic en **Install**.
6. Cuando termine, haz clic en **Finish**.

---

# 2. Abrir la terminal en la carpeta del proyecto

Para usar Git debes abrir la terminal **dentro de la carpeta de tu proyecto**.

## Método recomendado (Windows)

1. Crea una carpeta para tu proyecto.

Ejemplo:

```
MiProyecto
```

2. Entra a esa carpeta.

3. Haz **clic derecho en un espacio vacío dentro de la carpeta**.

4. Haz clic en:

```
Open in Terminal
```

o

```
Abrir en Terminal
```

Dependiendo del idioma de Windows.

Esto abrirá la terminal **directamente dentro de la carpeta**.

---

# 3. Verificar que Git funciona

En la terminal escribe:

```
git --version
```

Luego presiona **Enter**.

Ejemplo de salida:

```
git version 2.42.0
```

Si aparece la versión, significa que Git está instalado correctamente.

---

# 4. Configuración inicial de Git (primera vez)

La primera vez que usas Git debes configurar tu **nombre y correo**.

## Configurar nombre

En la terminal escribe:

```
git config --global user.name "Tu Nombre"
```

Ejemplo:

```
git config --global user.name "Kevin Saavedra"
```

Presiona **Enter**.

---

## Configurar correo

Ahora escribe:

```
git config --global user.email "tuemail@email.com"
```

Ejemplo:

```
git config --global user.email "kevin@email.com"
```

Presiona **Enter**.

---

## Ver la configuración

Para comprobar que se guardó correctamente:

```
git config --list
```

Esto mostrará todos los datos configurados.

---

# 5. Crear un repositorio Git

Ahora vamos a convertir la carpeta del proyecto en un repositorio Git.

En la terminal escribe:

```
git init
```

Luego presiona **Enter**.

Git creará una carpeta oculta llamada:

```
.git
```

Esa carpeta guarda **todo el historial del proyecto**.

---

# 6. Crear archivos en el proyecto

Ahora crea algunos archivos dentro de la carpeta.

Ejemplo:

```
index.html
style.css
script.js
```

Puedes crearlos con:

1. Clic derecho en la carpeta.
2. **Nuevo → Documento de texto**
3. Cambiar el nombre del archivo.

---

# 7. Ver el estado del proyecto

En la terminal escribe:

```
git status
```

Esto mostrará los archivos que Git detecta como nuevos.

Ejemplo:

```
Untracked files:
index.html
style.css
```

"Untracked" significa que Git **todavía no está siguiendo esos archivos**.

---

# 8. Agregar archivos al área de preparación (Staging)

Antes de guardar cambios, debes agregarlos al **staging area**.

## Agregar todos los archivos

En la terminal escribe:

```
git add .
```

Luego presiona **Enter**.

Esto agrega **todos los archivos de la carpeta**.

---

# 9. Crear el primer commit

Un **commit** guarda una versión del proyecto.

En la terminal escribe:

```
git commit -m "Primer commit del proyecto"
```

Luego presiona **Enter**.

Git guardará esa versión del proyecto.

---

# 10. Crear un repositorio en GitHub

1. Ve a:

https://github.com

2. Inicia sesión.

3. Haz clic en el botón **New Repository**.

4. Escribe el nombre del repositorio.

Ejemplo:

```
miProyecto
```

5. Haz clic en **Create Repository**.

---

# 11. Conectar el proyecto local con GitHub

GitHub mostrará una URL del repositorio.

Ejemplo:

```
https://github.com/usuario/miProyecto.git
```

Ahora en la terminal escribe:

```
git remote add origin https://github.com/usuario/miProyecto.git
```

Presiona **Enter**.

Esto conecta tu proyecto local con GitHub.

---

# 12. Subir el proyecto a GitHub

Primero establece la rama principal:

```
git branch -M main
```

Luego sube el proyecto:

```
git push -u origin main
```

Presiona **Enter**.

Git enviará todos los archivos al repositorio en GitHub.

---

# 13. Flujo normal de trabajo con Git

Cada vez que hagas cambios en el proyecto debes hacer lo siguiente.

## 1. Ver cambios

```
git status
```

---

## 2. Agregar cambios

```
git add .
```

---

## 3. Crear commit

```
git commit -m "Descripción del cambio"
```

Ejemplo:

```
git commit -m "Se agregó la página de login"
```

---

## 4. Subir cambios

```
git push
```

Esto actualizará el repositorio en GitHub.

---

# 14. Descargar cambios del repositorio

Si alguien más modificó el repositorio puedes descargar los cambios con:

```
git pull
```

---

# 15. Clonar un repositorio existente

Si quieres descargar un proyecto desde GitHub:

1. Copia la URL del repositorio.
2. Abre la terminal en la carpeta donde quieres guardarlo.

Luego ejecuta:

```
git clone https://github.com/usuario/proyecto.git
```

Esto descargará el proyecto completo en tu computadora.

---

# Conclusión

Git permite:

* Guardar versiones de un proyecto
* Trabajar en equipo
* Recuperar versiones anteriores
* Mantener historial de cambios

El flujo básico de trabajo en Git es:

```
git status
git add .
git commit -m "mensaje"
git push
```

Siguiendo estos pasos podrás manejar un proyecto usando Git desde la terminal.
