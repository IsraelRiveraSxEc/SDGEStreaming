/* @titulo: modulo de gestion de contenido
   @descripcion: Define la estructura de datos para contenido de streaming y proporciona funciones para agregar, listar, y calificar contenido en memoria. */
package content // Declara que este archivo pertenece al paquete content.

import (
	"fmt"     // Importa el paquete fmt para imprimir texto.
	"strconv" // Importa el paquete strconv para conversiones de string.
	"strings" // Importa el paquete strings para manipular cadenas de texto.
)

// Define una estructura llamada UserRating que representa una calificación de un usuario específico.
type UserRating struct {
	UserID int     `json:"user_id"` // Campo para almacenar el ID del usuario que calificó.
	Rating float64 `json:"rating"`  // Campo para almacenar la calificación (del 1.0 al 10.0).
}

// Define una estructura llamada Content que representa un ítem de contenido en el sistema.
type Content struct {
	ID            int          `json:"id"`              // Campo para almacenar el ID único del contenido.
	Title         string       `json:"title"`           // Campo para almacenar el título del contenido.
	Type          string       `json:"type"`            // Campo para almacenar el tipo (Audiovisual/Audio).
	Duration      int          `json:"duration"`        // Campo para almacenar la duración en minutos.
	AverageRating float64      `json:"average_rating"`  // Campo para almacenar el promedio de calificaciones.
	Ratings       []UserRating `json:"-"`               // Campo para almacenar la lista de calificaciones individuales (no se imprime en ListarContenido).
}

// Variable global (dentro del paquete) para almacenar el contenido temporalmente en memoria.
var contentInMemory []Content
// Variable global para asignar IDs únicos a los nuevos contenidos.
var nextContentID = 1

// Función para agregar un nuevo ítem de contenido a la lista en memoria.
// Recibe el título, el tipo y la duración como parámetros.
func AgregarContenido(titulo string, tipo string, duracion int) {
	// Crea una nueva instancia de la estructura Content con los datos proporcionados.
	// El ID se asigna usando la variable nextContentID.
	// El promedio de calificaciones se inicializa en 0.0.
	// La lista de calificaciones individuales se inicializa vacía.
	nuevoContenido := Content{
		ID:            nextContentID, // Asigna el siguiente ID disponible.
		Title:         titulo,        // Asigna el título recibido.
		Type:          tipo,          // Asigna el tipo recibido.
		Duration:      duracion,      // Asigna la duración recibida.
		AverageRating: 0.0,           // Inicializa el promedio en 0.
		Ratings:       []UserRating{}, // Inicializa la lista de calificaciones vacía.
	}
	// Añade el nuevo contenido a la lista de contenido en memoria.
	contentInMemory = append(contentInMemory, nuevoContenido)
	// Incrementa la variable nextContentID para el próximo contenido.
	nextContentID++
}

// Función para imprimir en consola todos los ítems de contenido almacenados en memoria.
func ListarContenido() {
	// Verifica si la lista de contenido está vacía.
	if len(contentInMemory) == 0 {
		// Imprime un mensaje si no hay contenido registrado.
		fmt.Println("No hay contenido registrado.")
		// Retorna de la función si la lista está vacía.
		return
	}
	// Imprime un encabezado para la lista de contenido.
	fmt.Println("--- Lista de Contenido ---")
	// Itera sobre cada ítem de contenido en la lista contentInMemory.
	for _, item := range contentInMemory {
		// Imprime los detalles de cada ítem (ID, Título, Tipo, Duración, Calificación Promedio).
		// %.2f imprime el promedio con dos decimales.
		fmt.Printf("ID: %d, Título: %s, Tipo: %s, Duración: %d min, Calificación Promedio: %.2f\n", item.ID, item.Title, item.Type, item.Duration, item.AverageRating)
	}
}

// Función para permitir a un usuario calificar un contenido específico.
// Recibe el ID del contenido, el ID del usuario y la calificación como string.
func CalificarContenido(contentID int, userID int, ratingStr string) error {
	// Buscar el contenido en la lista usando el ID proporcionado.
	contentIndex := -1 // Inicializa el índice en -1, asumiendo que no se encontró.
	// Itera sobre la lista de contenido.
	for i, item := range contentInMemory {
		// Verifica si el ID del ítem actual coincide con el ID buscado.
		if item.ID == contentID {
			// Si coincide, almacena el índice.
			contentIndex = i
			// Rompe el bucle una vez encontrado.
			break
		}
	}

	// Verifica si el contenido no fue encontrado.
	if contentIndex == -1 {
		// Retorna un error si no se encontró el contenido.
		return fmt.Errorf("contenido con ID %d no encontrado", contentID)
	}

	// Convierte la calificación de string a float64, manejando punto y coma.
	// Reemplaza la primera coma por punto para estandarizar la notación decimal.
	ratingStr = strings.Replace(ratingStr, ",", ".", 1)
	// Convierte la cadena de texto estandarizada a un número de punto flotante.
	rating, err := strconv.ParseFloat(ratingStr, 64)
	// Verifica si ocurrió un error durante la conversión.
	if err != nil {
		// Retorna un error si la calificación no es un número válido.
		return fmt.Errorf("calificación '%s' no es válida. Debe ser un número entre 1.0 y 10.0", ratingStr)
	}

	// Validar que la calificación esté dentro del rango permitido (1.0 a 10.0).
	if rating < 1.0 || rating > 10.0 {
		// Retorna un error si la calificación está fuera del rango.
		return fmt.Errorf("la calificación debe estar entre 1.0 y 10.0, recibido: %.2f", rating)
	}

	// Verificar si el usuario ya ha calificado este contenido (opcional, aquí lo sobrescribe).
	// Inicializa el índice de la calificación existente en -1.
	existingRatingIndex := -1
	// Itera sobre las calificaciones del contenido encontrado.
	for i, ur := range contentInMemory[contentIndex].Ratings {
		// Verifica si el UserID de la calificación actual coincide con el UserID del usuario que califica.
		if ur.UserID == userID {
			// Si coincide, almacena el índice de la calificación existente.
			existingRatingIndex = i
			// Rompe el bucle una vez encontrado.
			break
		}
	}

	// Crea una nueva instancia de UserRating con el UserID y la Rating.
	newUserRating := UserRating{UserID: userID, Rating: rating}

	// Verifica si el usuario ya había calificado este contenido.
	if existingRatingIndex != -1 {
		// Si ya existía una calificación, actualiza el valor en la lista.
		contentInMemory[contentIndex].Ratings[existingRatingIndex] = newUserRating
	} else {
		// Si no existía una calificación, agrega la nueva a la lista.
		contentInMemory[contentIndex].Ratings = append(contentInMemory[contentIndex].Ratings, newUserRating)
	}

	// Recalcular el promedio de calificaciones para el contenido.
	recalcularPromedio(contentIndex)
	// Imprime un mensaje de confirmación.
	fmt.Printf("Contenido '%s' calificado exitosamente con %.2f por el usuario %d.\n", contentInMemory[contentIndex].Title, rating, userID)
	// Retorna nil indicando que no hubo error.
	return nil
}

// Función para recalcular el AverageRating de un contenido basado en sus calificaciones almacenadas.
// Recibe el índice del contenido en la lista contentInMemory.
func recalcularPromedio(index int) {
	// Inicializa una variable para acumular la suma de todas las calificaciones.
	sum := 0.0
	// Itera sobre la lista de calificaciones del contenido en la posición 'index'.
	for _, ur := range contentInMemory[index].Ratings {
		// Suma cada calificación individual al acumulador.
		sum += ur.Rating
	}
	// Verifica si hay al menos una calificación registrada.
	if len(contentInMemory[index].Ratings) > 0 {
		// Calcula el promedio dividiendo la suma por la cantidad de calificaciones.
		contentInMemory[index].AverageRating = sum / float64(len(contentInMemory[index].Ratings))
	} else {
		// Si no hay calificaciones, el promedio se establece en 0.0.
		contentInMemory[index].AverageRating = 0.0
	}
}
