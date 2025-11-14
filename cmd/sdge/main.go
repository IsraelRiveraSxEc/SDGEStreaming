/* @titulo: programa principal del proyecto SDGEStreaming
   @integrantes: Nelson Espinosa, Barbara
   @fecha: 15/11/2025
   @descripcion: Punto de entrada del sistema. Contiene el menú interactivo y la lógica de control principal que orquesta las interacciones con los módulos de usuarios y contenido. Utiliza programación funcional básica. */
package main // Declara que este archivo pertenece al paquete main, el punto de entrada del programa.

// Importa los paquetes necesarios para el funcionamiento del programa.
import (
	"SDGEStreaming/pkg/content" // Importa el paquete content del proyecto.
	"SDGEStreaming/pkg/users"   // Importa el paquete users del proyecto.
	"bufio"                     // Importa el paquete bufio para leer de stdin.
	"fmt"                       // Importa el paquete fmt para imprimir y leer texto.
	"os"                        // Importa el paquete os para interactuar con el sistema operativo.
	"strconv"                   // Importa el paquete strconv para conversiones de string.
	"strings"                   // Importa el paquete strings para manipular cadenas de texto.
)

// Declara una variable global 'scanner' de tipo bufio.Scanner.
// Esta variable se usará para leer líneas de texto desde la entrada estándar (stdin).
var scanner = bufio.NewScanner(os.Stdin)

// Función principal del programa. Es el punto de entrada cuando se ejecuta el binario.
func main() {
	// Imprime un mensaje de bienvenida al iniciar el programa.
	fmt.Println("--- Bienvenido al Sistema de Gestión de Streaming (SDGEStreaming) ---")

	// Inicia un bucle infinito para mantener el menú activo hasta que el usuario elija salir.
	for {
		// Llama a la función para mostrar el menú principal.
		mostrarMenu()
		// Llama a la función para leer la opción elegida por el usuario.
		opcion := leerOpcion()
		// Llama a la función para ejecutar la acción correspondiente a la opción elegida.
		ejecutarOpcion(opcion)
	}
}

// Función que imprime las opciones del menú principal en la consola.
func mostrarMenu() {
	// Imprime un encabezado para el menú.
	fmt.Println("\n--- Menú Principal ---")
	// Imprime cada opción del menú.
	fmt.Println("1. Gestionar Usuarios")
	fmt.Println("2. Gestionar Contenido")
	fmt.Println("3. Explorar Contenido")
	fmt.Println("4. Salir")
	// Solicita al usuario que seleccione una opción.
	fmt.Print("Seleccione una opción: ")
}

// Función que lee la entrada del usuario y la convierte a un número entero.
// Devuelve la opción seleccionada o 0 si la entrada no es válida.
func leerOpcion() int {
	// Lee la siguiente línea de texto desde stdin usando el scanner.
	scanner.Scan()
	// Obtiene el texto leído y elimina espacios en blanco al inicio y al final.
	entrada := strings.TrimSpace(scanner.Text())
	// Convierte la cadena de texto a un número entero.
	// Atoi convierte de string a int.
	opcion, err := strconv.Atoi(entrada)
	// Verifica si ocurrió un error durante la conversión (por ejemplo, si la entrada no era un número).
	if err != nil {
		// Imprime un mensaje de error si la entrada no es un número válido.
		fmt.Println("Entrada inválida. Por favor, ingrese un número.")
		// Devuelve 0 para indicar que la entrada fue incorrecta.
		return 0
	}
	// Devuelve el número entero correspondiente a la opción elegida.
	return opcion
}

// Función que toma la opción elegida por el usuario y llama a las funciones correspondientes.
func ejecutarOpcion(opcion int) {
	// Utiliza una sentencia switch para evaluar el valor de 'opcion'.
	switch opcion {
	// Caso 1: Gestionar Usuarios.
	case 1:
		// Llama a la función para manejar las opciones de usuarios.
		gestionarUsuarios()
	// Caso 2: Gestionar Contenido.
	case 2:
		// Llama a la función para manejar las opciones de contenido.
		gestionarContenido()
	// Caso 3: Explorar Contenido.
	case 3:
		// Llama a la función para explorar el catálogo de contenido.
		explorarContenido()
	// Caso 4: Salir del sistema.
	case 4:
		// Imprime un mensaje de despedida.
		fmt.Println("Saliendo del sistema...")
		// Finaliza la ejecución del programa con código de salida 0 (éxito).
		os.Exit(0)
	// Caso por defecto: Si la opción no coincide con ninguna de las anteriores.
	default:
		// Imprime un mensaje indicando que la opción no es válida.
		fmt.Println("Opción no válida. Intente de nuevo.")
	}
}

// Función que maneja las opciones del submenú de Usuarios.
func gestionarUsuarios() {
	// Imprime un encabezado para el submenú de usuarios.
	fmt.Println("\n--- Gestión de Usuarios ---")
	// Imprime las opciones disponibles dentro de este submenú.
	fmt.Println("1. Agregar Usuario")
	fmt.Println("2. Listar Usuarios")
	fmt.Println("3. Volver al Menú Principal")
	// Solicita al usuario que seleccione una opción dentro de este submenú.
	fmt.Print("Seleccione una opción: ")

	// Lee la opción elegida dentro del submenú de usuarios.
	subOpcion := leerOpcion()

	// Utiliza un switch para manejar la opción elegida en el submenú de usuarios.
	switch subOpcion {
	// Caso 1: Agregar un nuevo usuario.
	case 1:
		// Llama a la función para interactuar con el usuario y agregarlo.
		agregarUsuario()
	// Caso 2: Listar todos los usuarios registrados.
	case 2:
		// Llama a la función del paquete 'users' para listar usuarios.
		users.ListarUsuarios()
	// Caso 3: Volver al menú principal.
	case 3:
		// Retorna de la función, lo que lleva de vuelta al menú principal.
		return
	// Caso por defecto: Opción inválida en el submenú de usuarios.
	default:
		// Imprime un mensaje de error para opción inválida.
		fmt.Println("Opción no válida para Usuarios.")
	}
}

// Función que maneja las opciones del submenú de Contenido.
func gestionarContenido() {
	// Imprime un encabezado para el submenú de contenido.
	fmt.Println("\n--- Gestión de Contenido ---")
	// Imprime las opciones disponibles dentro de este submenú.
	fmt.Println("1. Agregar Contenido")
	fmt.Println("2. Listar Contenido")
	fmt.Println("3. Calificar Contenido") // Nueva opción
	fmt.Println("4. Volver al Menú Principal")
	// Solicita al usuario que seleccione una opción dentro de este submenú.
	fmt.Print("Seleccione una opción: ")

	// Lee la opción elegida dentro del submenú de contenido.
	subOpcion := leerOpcion()

	// Utiliza un switch para manejar la opción elegida en el submenú de contenido.
	switch subOpcion {
	// Caso 1: Agregar nuevo contenido.
	case 1:
		// Llama a la función para interactuar con el usuario y agregar contenido.
		agregarContenido()
	// Caso 2: Listar todo el contenido disponible.
	case 2:
		// Llama a la función del paquete 'content' para listar contenido.
		content.ListarContenido()
	// Caso 3: Calificar un contenido existente.
	case 3:
		// Llama a la función para calificar contenido.
		calificarContenido()
	// Caso 4: Volver al menú principal.
	case 4:
		// Retorna de la función, lo que lleva de vuelta al menú principal.
		return
	// Caso por defecto: Opción inválida en el submenú de contenido.
	default:
		// Imprime un mensaje de error para opción inválida.
		fmt.Println("Opción no válida para Contenido.")
	}
}

// Función que maneja las opciones del submenú de Explorar Contenido.
func explorarContenido() {
	// Imprime un encabezado para el submenú de explorar contenido.
	fmt.Println("\n--- Explorar Contenido ---")
	// Imprime las opciones disponibles dentro de este submenú.
	fmt.Println("1. Ver Catálogo Completo")
	fmt.Println("2. Volver al Menú Principal")
	// Solicita al usuario que seleccione una opción dentro de este submenú.
	fmt.Print("Seleccione una opción: ")

	// Lee la opción elegida dentro del submenú de explorar contenido.
	subOpcion := leerOpcion()

	// Utiliza un switch para manejar la opción elegida en el submenú de explorar contenido.
	switch subOpcion {
	// Caso 1: Ver el catálogo completo de contenido.
	case 1:
		// Llama a la función del paquete 'content' para listar contenido.
		content.ListarContenido()
	// Caso 2: Volver al menú principal.
	case 2:
		// Retorna de la función, lo que lleva de vuelta al menú principal.
		return
	// Caso por defecto: Opción inválida en el submenú de explorar contenido.
	default:
		// Imprime un mensaje de error para opción inválida.
		fmt.Println("Opción no válida para Explorar Contenido.")
	}
}

// Función que interactúa con el usuario para obtener datos y llama a la función del paquete users para agregarlo.
func agregarUsuario() {
	// Solicita al usuario que ingrese su nombre.
	fmt.Print("Ingrese el nombre del nuevo usuario (puede contener espacios): ")
	// Lee la siguiente línea de texto desde stdin.
	scanner.Scan()
	// Obtiene el texto leído y elimina espacios en blanco al inicio y al final.
	nombre := strings.TrimSpace(scanner.Text())

	// Solicita al usuario que ingrese su email.
	fmt.Print("Ingrese el email del nuevo usuario: ")
	// Lee la siguiente línea de texto desde stdin.
	scanner.Scan()
	// Obtiene el texto leído y elimina espacios en blanco al inicio y al final.
	email := strings.TrimSpace(scanner.Text())

	// Verifica si el nombre o el email están vacíos.
	if nombre == "" || email == "" {
		// Imprime un mensaje de error si algún campo está vacío.
		fmt.Println("Nombre y email no pueden estar vacíos.")
		// Retorna de la función si los datos son inválidos.
		return
	}

	// Llama a la función 'AgregarUsuario' del paquete 'users' para procesar los datos.
	users.AgregarUsuario(nombre, email)
	// Imprime un mensaje de confirmación.
	fmt.Println("Usuario agregado exitosamente.")
}

// Función que interactúa con el usuario para obtener datos y llama a la función del paquete content para agregarlo.
func agregarContenido() {
	// Solicita al usuario que ingrese el título del contenido.
	fmt.Print("Ingrese el título del contenido (puede contener espacios): ")
	// Lee la siguiente línea de texto desde stdin.
	scanner.Scan()
	// Obtiene el texto leído y elimina espacios en blanco al inicio y al final.
	titulo := strings.TrimSpace(scanner.Text())

	// Muestra las opciones de tipo de contenido.
	fmt.Println("Tipos de contenido: 1. Audiovisual, 2. Audio (Podcast)")
	// Solicita al usuario que seleccione el tipo.
	fmt.Print("Seleccione el tipo (1 o 2): ")
	// Lee la opción elegida para el tipo de contenido.
	tipoInput := leerOpcion()
	// Declara una variable para almacenar el tipo de contenido como string.
	var tipo string
	// Utiliza un switch para asignar el valor correcto a 'tipo' según la opción elegida.
	switch tipoInput {
	// Caso 1: Audiovisual.
	case 1:
		// Asigna "Audiovisual" a la variable tipo.
		tipo = "Audiovisual"
	// Caso 2: Audio.
	case 2:
		// Asigna "Audio" a la variable tipo.
		tipo = "Audio"
	// Caso por defecto: Opción inválida.
	default:
		// Imprime un mensaje de error si la opción de tipo es inválida.
		fmt.Println("Opción de tipo inválida.")
		// Retorna de la función si la opción es inválida.
		return
	}

	// Solicita al usuario que ingrese la duración del contenido.
	fmt.Print("Ingrese la duración del contenido (en minutos, solo número): ")
	// Lee la siguiente línea de texto desde stdin.
	scanner.Scan()
	// Obtiene el texto leído y elimina espacios en blanco al inicio y al final.
	duracionStr := strings.TrimSpace(scanner.Text())
	// Convierte la cadena de texto de la duración a un número entero.
	duracion, err := strconv.Atoi(duracionStr)
	// Verifica si ocurrió un error durante la conversión de la duración.
	if err != nil {
		// Imprime un mensaje de error si la duración no es un número válido.
		fmt.Println("Duración inválida. Debe ser un número.")
		// Retorna de la función si la duración es inválida.
		return
	}

	// Verifica si el título está vacío.
	if titulo == "" {
		// Imprime un mensaje de error si el título está vacío.
		fmt.Println("El título no puede estar vacío.")
		// Retorna de la función si el título es inválido.
		return
	}

	// Llama a la función 'AgregarContenido' del paquete 'content' para procesar los datos.
	content.AgregarContenido(titulo, tipo, duracion)
	// Imprime un mensaje de confirmación.
	fmt.Println("Contenido agregado exitosamente.")
}

// Función que interactúa con el usuario para obtener ID de contenido, ID de usuario y calificación, y llama a la función del paquete content.
func calificarContenido() {
	// Muestra la lista de contenido para que el usuario elija.
	content.ListarContenido()
	// Solicita al usuario que ingrese el ID del contenido a calificar.
	fmt.Print("Ingrese el ID del contenido que desea calificar: ")
	// Lee la siguiente línea de texto desde stdin.
	scanner.Scan()
	// Obtiene el texto leído (el ID como string) y elimina espacios en blanco al inicio y al final.
	contentIDStr := strings.TrimSpace(scanner.Text())
	// Convierte la cadena de texto del ID del contenido a un número entero.
	contentID, err := strconv.Atoi(contentIDStr)
	// Verifica si ocurrió un error durante la conversión del ID del contenido.
	if err != nil {
		// Imprime un mensaje de error si el ID del contenido no es válido.
		fmt.Println("ID de contenido inválido.")
		// Retorna de la función si el ID es inválido.
		return
	}

	// Solicita al usuario que ingrese su ID de usuario.
	fmt.Print("Ingrese su ID de usuario (número entero): ")
	// Lee la siguiente línea de texto desde stdin.
	scanner.Scan()
	// Obtiene el texto leído (el ID como string) y elimina espacios en blanco al inicio y al final.
	userIDStr := strings.TrimSpace(scanner.Text())
	// Convierte la cadena de texto del ID del usuario a un número entero.
	userID, err := strconv.Atoi(userIDStr)
	// Verifica si ocurrió un error durante la conversión del ID del usuario.
	if err != nil {
		// Imprime un mensaje de error si el ID del usuario no es válido.
		fmt.Println("ID de usuario inválido.")
		// Retorna de la función si el ID es inválido.
		return
	}

	// Solicita al usuario que ingrese la calificación (puede usar punto o coma).
	fmt.Print("Ingrese la calificación (1.0 a 10.0, ej: 7.5 o 7,5): ")
	// Lee la siguiente línea de texto desde stdin.
	scanner.Scan()
	// Obtiene el texto leído (la calificación como string) y elimina espacios en blanco al inicio y al final.
	ratingStr := strings.TrimSpace(scanner.Text())

	// Llama a la función 'CalificarContenido' del paquete 'content' para procesar los datos.
	err = content.CalificarContenido(contentID, userID, ratingStr)
	// Verifica si ocurrió un error durante el proceso de calificación.
	if err != nil {
		// Imprime un mensaje de error con los detalles del problema.
		fmt.Printf("Error al calificar: %v\n", err)
	}
}