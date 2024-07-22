# Lista de Tareas API

¡Bienvenido al proyecto de API de Lista de Tareas! Este proyecto está construido utilizando **Golang** y la librería **Fiber**. Actualmente, se ha completado todo el backend y el diseño del frontend aún no se ha iniciado. El proyecto está en proceso.

## Descripción

Esta API permite la gestión de una lista de tareas, incluyendo la creación, lectura, actualización y eliminación de tareas. Es una aplicación sencilla pero poderosa que sirve como base para futuros desarrollos.

## Características

- **Crear una tarea**: Permite añadir nuevas tareas a la lista.
- **Leer tareas**: Permite recuperar todas las tareas o una tarea específica.
- **Actualizar tareas**: Permite modificar los detalles de una tarea existente.
- **Eliminar tareas**: Permite eliminar una tarea de la lista.

## Tecnologías Utilizadas

- **Lenguaje**: Golang
- **Framework**: Fiber

## Instalación

Sigue estos pasos para configurar y ejecutar el proyecto en tu entorno local.

1. **Clonar el repositorio**

    ```bash
    git clone https://github.com/tu-usuario/lista-de-tareas-api.git
    ```

2. **Navegar al directorio del proyecto**

    ```bash
    cd lista-de-tareas-api
    ```

3. **Instalar las dependencias**

    Asegúrate de tener **GO** y **Fiber** instalados en tu máquina.

    ```bash
    go get -u github.com/gofiber/fiber/v2
    ```

4. **Ejecutar la aplicación**

    ```bash
    go run main.go
    ```

    La aplicación se ejecutará en `http://localhost:3000`.

## Uso de la API

Aquí están los endpoints disponibles en la API:

### Crear una tarea

- **URL**: `/api/todos`
- **Método**: `POST`
- **Cuerpo de la solicitud**:
    ```json
    {
      "_id": "Nuevo título de la tarea",
      "completed": "Estado de la tarea",
      "body":"Descripción del proyecto"
    }
    ```

### Obtener todas las tareas

- **URL**: `/api/todos`
- **Método**: `GET`

### Actualizar una tarea

- **URL**: `/api/todos/:id`
- **Método**: `PATCH`
- **Cuerpo de la solicitud**:
    ```json
    {
      "_id": "Nuevo título de la tarea",
      "completed": "Estado de la tarea",
      "body":"Descripción del proyecto"
    }
    ```

### Eliminar una tarea

- **URL**: `/api/todos/:id`
- **Método**: `DELETE`


## Estado del Proyecto

El backend está completamente funcional, pero el frontend aún no se ha iniciado. Este proyecto está en proceso y se actualizará continuamente.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.

