/* @titulo: modulo de gestion de usuarios
   @descripcion: Define la estructura de datos para usuarios y proporciona funciones para agregar y listar usuarios en memoria. */
package users // Declara que este archivo pertenece al paquete users.

import "fmt" // Importa el paquete fmt para imprimir texto en la consola.

// User representa la estructura de datos para un usuario en el sistema.
type User struct {
	ID    int    `json:"id"`    // Campo ID de tipo entero para identificar al usuario de forma única. La etiqueta `json:"id"` indica cómo se debería serializar a JSON si se usara.
	Name  string `json:"name"`  // Campo Name de tipo string para almacenar el nombre del usuario. La etiqueta `json:"name"` indica cómo se debería serializar a JSON si se usara.
	Email string `json:"email"` // Campo Email de tipo string para almacenar el email del usuario. La etiqueta `json:"email"` indica cómo se debería serializar a JSON si se usara.
}

// usersInMemory es una variable global (dentro del paquete 'users') para almacenar los usuarios temporalmente en memoria RAM.
// En una versión futura del proyecto (AA2), esta variable se reemplazaría probablemente por una base de datos.
var usersInMemory []User

// nextUserID es una variable global (dentro del paquete 'users') para asignar IDs únicos a los nuevos usuarios registrados.
// Se inicializa en 1, y se incrementa cada vez que se agrega un nuevo usuario.
var nextUserID = 1

// AgregarUsuario crea un nuevo usuario y lo añade a la lista en memoria (usersInMemory).
// Recibe como parámetros el nombre y el email del nuevo usuario.
func AgregarUsuario(nombre string, email string) {
	// Crea una nueva instancia de la estructura 'User' con los datos proporcionados.
	// El campo ID se asigna usando el valor actual de la variable 'nextUserID'.
	nuevoUsuario := User{
		ID:    nextUserID, // Asigna el siguiente ID disponible a este nuevo usuario.
		Name:  nombre,     // Asigna el nombre recibido como parámetro al campo Name del nuevo usuario.
		Email: email,      // Asigna el email recibido como parámetro al campo Email del nuevo usuario.
	}

	// Añade el nuevo usuario a la lista global de usuarios en memoria (usersInMemory) usando la función 'append'.
	usersInMemory = append(usersInMemory, nuevoUsuario)

	// Incrementa la variable 'nextUserID' para que el próximo usuario registrado tenga un ID único consecutivo.
	nextUserID++
}

// ListarUsuarios imprime en la consola todos los usuarios almacenados en la lista 'usersInMemory'.
func ListarUsuarios() {
	// Verifica si la lista de usuarios en memoria está vacía.
	if len(usersInMemory) == 0 {
		// Si la longitud de la lista es 0, imprime un mensaje indicando que no hay usuarios registrados.
		fmt.Println("No hay usuarios registrados.")
		// Retorna de la función inmediatamente, ya que no hay nada más que hacer.
		return
	}

	// Si la lista no está vacía, imprime un encabezado para la lista de usuarios.
	fmt.Println("--- Lista de Usuarios ---")

	// Itera sobre cada usuario en la lista 'usersInMemory'.
	// La variable 'usuario' contendrá temporalmente cada valor de la lista en cada iteración.
	for _, usuario := range usersInMemory {
		// Imprime los detalles del usuario actual (ID, Nombre, Email) en un formato legible.
		// Se usa 'fmt.Printf' para formatear la cadena de texto con los valores del usuario.
		fmt.Printf("ID: %d, Nombre: %s, Email: %s\n", usuario.ID, usuario.Name, usuario.Email)
	}
}

// BuscarUsuarioPorNombre es una función adicional que podría ser útil en fases posteriores del proyecto (por ejemplo, AA2).
// Esta función busca un usuario específico en la lista 'usersInMemory' por su nombre exacto.
// Devuelve un puntero al usuario encontrado y un valor booleano que indica si se encontró o no.
// func BuscarUsuarioPorNombre(nombre string) (*User, bool) {
// 	// Itera sobre la lista de usuarios.
// 	for i, usuario := range usersInMemory {
// 		// Compara el nombre del usuario actual en la iteración con el nombre buscado.
// 		// strings.EqualFold(usuario.Name, nombre) compara sin distinguir mayúsculas de minúsculas.
// 		if usuario.Name == nombre { // Comparación simple por nombre (distingue mayúsculas/minúsculas en este ejemplo)
// 			// Si se encuentra, devuelve un puntero al usuario (usando &usersInMemory[i]) y true.
// 			return &usersInMemory[i], true
// 		}
// 	}
// 	// Si el bucle termina y no se encontró el usuario, devuelve nil (ningún usuario) y false.
// 	return nil, false
// }
