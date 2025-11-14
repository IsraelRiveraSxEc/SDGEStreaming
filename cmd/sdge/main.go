/* @titulo: programa principal del proyecto SDGEStreaming
   @autores: Nelson Espinosa, Barbara Peñaherrera
   @fecha: 15/11/2025
   @descripcion: Punto de entrada del sistema. Contiene el menú interactivo y la lógica de control principal que orquesta las interacciones con los módulos de usuarios y contenido. Utiliza programación funcional básica. */
package main // Paquete principal del programa.

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
var scanner = bufio.NewScanner(os.Stdin)

// Variable global para simular un usuario logueado.
var usuarioActual string = ""

// Función principal del programa. Es el punto de entrada cuando se ejecuta el binario.
func main() {
	fmt.Println("--- Bienvenido al Sistema de Gestión de Streaming (SDGEStreaming) ---")

	for {
		if usuarioActual == "" {
			mostrarMenuInicio()
			opcion := leerOpcion()
			ejecutarOpcionInicio(opcion)
		} else {
			mostrarMenuPrincipal()
			opcion := leerOpcion()
			ejecutarOpcionPrincipal(opcion)
		}
	}
}

// Función que imprime las opciones del menú de inicio.
func mostrarMenuInicio() {
	fmt.Println("\n--- Bienvenido a SDGEStreaming ---")
	fmt.Println("1. Iniciar Sesión")
	fmt.Println("2. Registrarse")
	fmt.Println("3. Explorar Contenido (Invitado)")
	fmt.Println("4. Salir")
	fmt.Print("Seleccione una opción: ")
}

// Función que maneja las opciones del menú de inicio.
func ejecutarOpcionInicio(opcion int) {
	switch opcion {
	case 1:
		iniciarSesion()
	case 2:
		registrarse()
	case 3:
		fmt.Println("Explorando como invitado. Algunas funciones están limitadas.")
		explorarContenido()
	case 4:
		fmt.Println("Saliendo del sistema...")
		os.Exit(0)
	default:
		fmt.Println("Opción no válida. Intente de nuevo.")
	}
}

// Función que imprime las opciones del menú principal para usuarios logueados.
func mostrarMenuPrincipal() {
	fmt.Printf("\n--- Menú Principal (Usuario: %s) ---\n", usuarioActual)
	fmt.Println("1. Mi Perfil")
	fmt.Println("2. Explorar Contenido")
	fmt.Println("3. Mi Lista")
	fmt.Println("4. Historial de Reproducción")
	fmt.Println("5. Gestionar Usuarios (Admin)")
	fmt.Println("6. Gestionar Contenido (Admin)")
	fmt.Println("7. Cerrar Sesión")
	fmt.Println("8. Salir")
	fmt.Print("Seleccione una opción: ")
}

// Lee la entrada del usuario y la convierte a un número entero.
func leerOpcion() int {
	scanner.Scan()
	entrada := strings.TrimSpace(scanner.Text())
	opcion, err := strconv.Atoi(entrada)
	if err != nil {
		fmt.Println("Entrada inválida. Por favor, ingrese un número.")
		return 0 // Indica entrada incorrecta.
	}
	return opcion
}

// Función que toma la opción elegida por el usuario logueado y llama a las funciones correspondientes.
func ejecutarOpcionPrincipal(opcion int) {
	switch opcion {
	case 1:
		miPerfil()
	case 2:
		explorarContenido()
	case 3:
		miLista()
	case 4:
		historialReproduccion()
	case 5:
		gestionarUsuarios()
	case 6:
		gestionarContenido()
	case 7:
		cerrarSesion()
	case 8:
		fmt.Println("Saliendo del sistema...")
		os.Exit(0)
	default:
		fmt.Println("Opción no válida. Intente de nuevo.")
	}
}

// Maneja el inicio de sesión del usuario (simulación para AA1).
func iniciarSesion() {
	fmt.Print("Ingrese su nombre de usuario: ")
	scanner.Scan()
	usuario := strings.TrimSpace(scanner.Text())

	// Simulación de inicio de sesión: se asigna el nombre a la variable global.
	usuarioActual = usuario
	fmt.Printf("Sesión iniciada como: %s\n", usuarioActual)
}

// Maneja el registro de un nuevo usuario (funcionalidad de AA1).
func registrarse() {
	fmt.Print("Ingrese el nombre del nuevo usuario (puede contener espacios): ")
	scanner.Scan()
	nombre := strings.TrimSpace(scanner.Text())

	fmt.Print("Ingrese el email del nuevo usuario: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	if nombre == "" || email == "" {
		fmt.Println("Nombre y email no pueden estar vacíos.")
		return
	}

	users.AgregarUsuario(nombre, email)
	fmt.Println("Usuario registrado exitosamente.")
}

// Maneja la visualización y edición del perfil del usuario (simulación para AA1).
func miPerfil() {
	fmt.Println("\n--- Mi Perfil ---")
	fmt.Printf("Nombre de Usuario: %s\n", usuarioActual)
	fmt.Println("Email: usuario@ejemplo.com (Simulado)")
	fmt.Println("Plan de Suscripción: Básico (Simulado)")
	fmt.Println("1. Editar Perfil (No implementado en AA1)")
	fmt.Println("2. Configuración (No implementado en AA1)")
	fmt.Println("3. Volver al Menú Principal")
	fmt.Print("Seleccione una opción: ")

	subOpcion := leerOpcion()

	switch subOpcion {
	case 1, 2:
		fmt.Println("Funcionalidad aún no implementada en AA1.")
	case 3:
		return
	default:
		fmt.Println("Opción no válida.")
	}
}

// Maneja la visualización de la lista personal de favoritos (simulación para AA1).
func miLista() {
	fmt.Println("\n--- Mi Lista ---")
	fmt.Println("Funcionalidad de listas personales aún no implementada en AA1.")
	fmt.Println("Aquí se mostrarían los ítems que el usuario marcó como 'Mi Lista'.")
}

// Maneja la visualización del historial de reproducción (simulación para AA1).
func historialReproduccion() {
	fmt.Println("\n--- Historial de Reproducción ---")
	fmt.Println("Funcionalidad de historial de vistas/escuchas aún no implementada en AA1.")
	fmt.Println("Aquí se mostrarían los ítems que el usuario ha reproducido.")
}

// Maneja las opciones del submenú de Usuarios (Admin).
func gestionarUsuarios() {
	fmt.Println("\n--- Gestión de Usuarios (Admin) ---")
	fmt.Println("1. Agregar Usuario")
	fmt.Println("2. Listar Usuarios")
	fmt.Println("3. Volver al Menú Principal")
	fmt.Print("Seleccione una opción: ")

	subOpcion := leerOpcion()

	switch subOpcion {
	case 1:
		agregarUsuario()
	case 2:
		users.ListarUsuarios()
	case 3:
		return
	default:
		fmt.Println("Opción no válida para Usuarios.")
	}
}

// Maneja las opciones del submenú de Contenido (Admin).
func gestionarContenido() {
	fmt.Println("\n--- Gestión de Contenido (Admin) ---")
	fmt.Println("1. Agregar Contenido")
	fmt.Println("2. Listar Contenido")
	fmt.Println("3. Calificar Contenido")
	fmt.Println("4. Volver al Menú Principal")
	fmt.Print("Seleccione una opción: ")

	subOpcion := leerOpcion()

	switch subOpcion {
	case 1:
		agregarContenido()
	case 2:
		content.ListarContenido()
	case 3:
		calificarContenido()
	case 4:
		return
	default:
		fmt.Println("Opción no válida para Contenido.")
	}
}

// Maneja las opciones del submenú de Explorar Contenido.
func explorarContenido() {
	fmt.Println("\n--- Explorar Contenido ---")
	fmt.Println("1. Ver Catálogo Completo")
	fmt.Println("2. Ver Contenido por Género (Simulado)")
	fmt.Println("3. Ver Contenido por Tipo (Simulado)")
	fmt.Println("4. Buscar Contenido (Simulado)")
	fmt.Println("5. Volver al Menú Principal")
	fmt.Print("Seleccione una opción: ")

	subOpcion := leerOpcion()

	switch subOpcion {
	case 1:
		content.ListarContenido()
	case 2:
		fmt.Println("Funcionalidad de filtrado por género aún no implementada en AA1.")
	case 3:
		fmt.Println("Funcionalidad de filtrado por tipo aún no implementada en AA1.")
	case 4:
		fmt.Println("Funcionalidad de búsqueda por título aún no implementada en AA1.")
	case 5:
		return
	default:
		fmt.Println("Opción no válida para Explorar Contenido.")
	}
}

// Cierra la sesión del usuario actual.
func cerrarSesion() {
	fmt.Printf("Cerrando sesión de %s...\n", usuarioActual)
	usuarioActual = ""
}

// Interactúa con el usuario para obtener datos y llama a la función del paquete users para agregarlo.
func agregarUsuario() {
	fmt.Print("Ingrese el nombre del nuevo usuario (puede contener espacios): ")
	scanner.Scan()
	nombre := strings.TrimSpace(scanner.Text())

	fmt.Print("Ingrese el email del nuevo usuario: ")
	scanner.Scan()
	email := strings.TrimSpace(scanner.Text())

	if nombre == "" || email == "" {
		fmt.Println("Nombre y email no pueden estar vacíos.")
		return
	}

	users.AgregarUsuario(nombre, email)
	fmt.Println("Usuario agregado exitosamente.")
}

// Interactúa con el usuario para obtener datos y llama a la función del paquete content para agregarlo.
func agregarContenido() {
	fmt.Print("Ingrese el título del contenido (puede contener espacios): ")
	scanner.Scan()
	titulo := strings.TrimSpace(scanner.Text())

	fmt.Println("Tipos de contenido: 1. Audiovisual, 2. Audio (Podcast)")
	fmt.Print("Seleccione el tipo (1 o 2): ")
	tipoInput := leerOpcion()
	var tipo string
	switch tipoInput {
	case 1:
		tipo = "Audiovisual"
	case 2:
		tipo = "Audio"
	default:
		fmt.Println("Opción de tipo inválida.")
		return
	}

	fmt.Print("Ingrese la duración del contenido (en minutos, solo número): ")
	scanner.Scan()
	duracionStr := strings.TrimSpace(scanner.Text())
	duracion, err := strconv.Atoi(duracionStr)
	if err != nil {
		fmt.Println("Duración inválida. Debe ser un número.")
		return
	}

	if titulo == "" {
		fmt.Println("El título no puede estar vacío.")
		return
	}

	content.AgregarContenido(titulo, tipo, duracion)
	fmt.Println("Contenido agregado exitosamente.")
}

// Interactúa con el usuario para obtener ID de contenido, ID de usuario y calificación, y llama a la función del paquete content.
// Ahora usa el ID del usuario actual logueado.
func calificarContenido() {
	content.ListarContenido()
	fmt.Print("Ingrese el ID del contenido que desea calificar: ")
	scanner.Scan()
	contentIDStr := strings.TrimSpace(scanner.Text())
	contentID, err := strconv.Atoi(contentIDStr)
	if err != nil {
		fmt.Println("ID de contenido inválido.")
		return
	}

	// Simulamos un ID de usuario basado en el nombre de usuario actual.
	userID := 0
	for _, char := range usuarioActual {
		userID += int(char)
	}
	if userID == 0 {
		userID = 999
	}

	fmt.Print("Ingrese la calificación (1.0 a 10.0, ej: 7.5 o 7,5): ")
	scanner.Scan()
	ratingStr := strings.TrimSpace(scanner.Text())

	err = content.CalificarContenido(contentID, userID, ratingStr)
	if err != nil {
		fmt.Printf("Error al calificar: %v\n", err)
	} else {
		fmt.Printf("Usuario '%s' (ID: %d) ha calificado el contenido.\n", usuarioActual, userID)
	}
}

// Ejemplo de función anónima utilizada como callback para mostrar un mensaje de bienvenida personalizado.
// Esta función se define dentro de main y se llama inmediatamente.
func ejemploFuncionAnonima() {
	bienvenida := func(nombre string) {
		fmt.Printf("¡Hola %s, bienvenido a SDGEStreaming!\n", nombre)
	}
	bienvenida(usuarioActual)
}

// Ejemplo de función variádica que imprime una lista de opciones del menú principal.
// Recibe un número variable de strings (opciones).
func imprimirOpciones(opciones ...string) {
	for i, opcion := range opciones {
		fmt.Printf("%d. %s\n", i+1, opcion)
	}
}
