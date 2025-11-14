
# SDGEStreaming - Aprendizaje Autónomo 1 (AA1)

## Descripción del Proyecto

SDGEStreaming es un Sistema de Gestión de Streaming desarrollado como parte de un proyecto académico. La primera fase (AA1) se centra en aplicar los fundamentos de programación en Go, específicamente programación funcional, uso de funciones, paquetes y estructura de proyecto, para crear una base funcional del sistema.

## Objetivo de la Fase AA1

El objetivo principal de esta fase es establecer una base funcional y modular del sistema SDGEStreaming mediante una interfaz de consola interactiva. Se implementaron funcionalidades básicas de gestión de usuarios y contenido, aplicando buenas prácticas de programación.

## Módulos Desarrollados (AA1)

*   **`cmd/sdge/main.go`**: Contiene la función `main`, el menú interactivo principal y la lógica de control que orquesta las interacciones con los otros módulos.
*   **`pkg/users/users.go`**: Gestiona el registro y listado básico de usuarios.
*   **`pkg/content/content.go`**: Permite agregar, listar y calificar contenido.

## Funcionalidades Implementadas (AA1)

*   **Menú Interactivo:** Menú principal y submenús para navegar entre las opciones del sistema.
*   **Gestión de Usuarios:**
    *   Agregar nuevo usuario (nombre, email).
    *   Listar todos los usuarios registrados.
*   **Gestión de Contenido:**
    *   Agregar nuevo ítem de contenido (título, tipo, duración).
    *   Listar todo el contenido disponible, incluyendo su calificación promedio.
    *   Calificar contenido por parte del usuario con puntuaciones del 1.0 al 10.0, aceptando notación con punto (5.5) o coma (5,5).
*   **Manejo de Datos:** Almacenamiento temporal de usuarios y contenido en memoria RAM.
*   **Modularidad y Documentación:** Código organizado en paquetes con sus comentarios respectivos.

## Tecnologías y Paquetes Utilizados (AA1)

*   **Lenguaje:** Go (Golang).
*   **Paquetes Estándar de Go:** `fmt`, `bufio`, `os`, `strings`, `strconv`.
*   **Control de Versiones:** Git.

## Estructura del Proyecto (AA1)

El proyecto SDGEStreaming se estructura de la siguiente manera:

```
├── cmd
│   └── sdge
│       └── main.go
├── pkg
│   ├── content
│   │   └── content.go
│   └── users
│       └── users.go
└── README.md
```

Disfrutando de la experiencia de programación en Go.