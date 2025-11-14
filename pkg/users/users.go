/* @titulo: modulo de gestion de usuarios
   @descripcion: Define la estructura de datos para usuarios y proporciona funciones para agregar y listar usuarios en memoria. */
package users // Declara que este archivo pertenece al paquete users.

import "fmt" // Importa el paquete fmt para imprimir texto.

// Define una estructura llamada User que representa un usuario en el sistema.
type User struct {
	ID    int    `json:"id"`    // Campo para almacenar el ID único del usuario.
	Name  string `json:"name"`  // Campo para almacenar el nombre del usuario.
	Email string `json:"email"` // Campo para almacenar el email del usuario.
}

// Variable global (dentro del paquete) para almacenar los usuarios temporalmente en memoria.
var usersInMemory []User
// Variable global para asignar IDs únicos a los nuevos usuarios.
var nextUserID = 1

// Función para agregar un nuevo usuario a la lista en memoria.
// Recibe el nombre y el email del usuario como parámetros.
func AgregarUsuario(nombre string, email string) {
	// Crea una nueva instancia de la estructura User con los datos proporcionados.
	// El ID se asigna usando la variable nextUserID.
	nuevoUsuario := User{
		ID:    nextUserID, // Asigna el siguiente ID disponible.
		Name:  nombre,     // Asigna el nombre recibido.
		Email: email,      // Asigna el email recibido.
	}
	// Añade el nuevo usuario a la lista de usuarios en memoria.
	usersInMemory = append(usersInMemory, nuevoUsuario)
	// Incrementa la variable nextUserID para el próximo usuario.
	nextUserID++
}

// Función para imprimir en consola todos los usuarios almacenados en memoria.
func ListarUsuarios() {
	// Verifica si la lista de usuarios está vacía.
	if len(usersInMemory) == 0 {
		// Imprime un mensaje si no hay usuarios registrados.
		fmt.Println("No hay usuarios registrados.")
		// Retorna de la función si la lista está vacía.
		return
	}
	// Imprime un encabezado para la lista de usuarios.
	fmt.Println("--- Lista de Usuarios ---")
	// Itera sobre cada usuario en la lista usersInMemory.
	for _, usuario := range usersInMemory {
		// Imprime los detalles de cada usuario (ID, Nombre, Email).
		fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", usuario.ID, usuario.Name, usuario.Email)
	}
}
