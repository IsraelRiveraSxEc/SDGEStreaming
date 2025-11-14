/* @titulo: modulo de gestion de contenido
   @descripcion: Define la estructura de datos para contenido de streaming y proporciona funciones para agregar, listar, y calificar contenido en memoria. */
package content // Declara que este archivo pertenece al paquete content.

import (
	"fmt"     // Importa el paquete fmt para imprimir texto en la consola.
	"strconv" // Importa el paquete strconv para realizar conversiones de string a otros tipos de datos (como números).
	"strings" // Importa el paquete strings para manipular cadenas de texto (por ejemplo, para reemplazar caracteres).
)

// UserRating representa la estructura de datos para una calificación individual de un usuario específico a un contenido.
type UserRating struct {
	UserID int     `json:"user_id"` // Campo UserID de tipo entero que identifica al usuario que realizó la calificación.
	Rating float64 `json:"rating"`  // Campo Rating de tipo número decimal (float64) que almacena la puntuación otorgada (por ejemplo, 7.5).
}

// Content representa la estructura de datos para un ítem de contenido en el sistema.
type Content struct {
	ID            int          `json:"id"`              // Campo ID de tipo entero para identificar al contenido de forma única.
	Title         string       `json:"title"`           // Campo Title de tipo string para almacenar el título del contenido.
	Type          string       `json:"type"`            // Campo Type de tipo string para almacenar el tipo de contenido (por ejemplo, "Audiovisual" o "Audio").
	Duration      int          `json:"duration"`        // Campo Duration de tipo entero para almacenar la duración del contenido en minutos.
	AverageRating float64      `json:"average_rating"`  // Campo AverageRating de tipo número decimal (float64) que almacena la calificación promedio calculada de todos los usuarios.
	Ratings       []UserRating `json:"-"`               // Campo Ratings que es una lista (slice) de estructuras UserRating. Almacena todas las calificaciones individuales recibidas. La etiqueta `json:"-"` indica que este campo no se debe serializar a JSON si se usara.
}

// contentInMemory es una variable global (dentro del paquete 'content') para almacenar el contenido temporalmente en memoria RAM.
// En una versión futura del proyecto (AA2), esta variable se reemplazaría probablemente por una base de datos.
var contentInMemory []Content

// nextContentID es una variable global (dentro del paquete 'content') para asignar IDs únicos a los nuevos contenidos registrados.
// Se inicializa en 1, y se incrementa cada vez que se agrega un nuevo contenido.
var nextContentID = 1

// AgregarContenido crea un nuevo ítem de contenido y lo añade a la lista en memoria (contentInMemory).
// Recibe como parámetros el título, el tipo y la duración del nuevo contenido.
func AgregarContenido(titulo string, tipo string, duracion int) {
	// Crea una nueva instancia de la estructura 'Content' con los datos proporcionados.
	// El campo ID se asigna usando el valor actual de la variable 'nextContentID'.
	// El campo AverageRating se inicializa en 0.0 porque recién se agrega y no tiene calificaciones.
	// El campo Ratings se inicializa como una lista vacía []UserRating.
	nuevoContenido := Content{
		ID:            nextContentID, // Asigna el siguiente ID disponible a este nuevo contenido.
		Title:         titulo,        // Asigna el título recibido como parámetro al campo Title del nuevo contenido.
		Type:          tipo,          // Asigna el tipo recibido como parámetro al campo Type del nuevo contenido.
		Duration:      duracion,      // Asigna la duración recibida como parámetro al campo Duration del nuevo contenido.
		AverageRating: 0.0,           // Inicializa la calificación promedio en 0.0.
		Ratings:       []UserRating{}, // Inicializa la lista de calificaciones individuales vacía.
	}

	// Añade el nuevo contenido a la lista global de contenido en memoria (contentInMemory) usando la función 'append'.
	contentInMemory = append(contentInMemory, nuevoContenido)

	// Incrementa la variable 'nextContentID' para que el próximo contenido registrado tenga un ID único consecutivo.
	nextContentID++
}

// ListarContenido imprime en la consola todos los ítems de contenido almacenados en la lista 'contentInMemory'.
func ListarContenido() {
	// Verifica si la lista de contenido en memoria está vacía.
	if len(contentInMemory) == 0 {
		// Si la longitud de la lista es 0, imprime un mensaje indicando que no hay contenido registrado.
		fmt.Println("No hay contenido registrado.")
		// Retorna de la función inmediatamente, ya que no hay nada más que hacer.
		return
	}

	// Si la lista no está vacía, imprime un encabezado para la lista de contenido.
	fmt.Println("--- Lista de Contenido ---")

	// Itera sobre cada ítem de contenido en la lista 'contentInMemory'.
	// La variable 'item' contendrá temporalmente cada valor de la lista en cada iteración.
	for _, item := range contentInMemory {
		// Imprime los detalles del ítem actual (ID, Título, Tipo, Duración, Calificación Promedio) en un formato legible.
		// Se usa 'fmt.Printf' para formatear la cadena de texto con los valores del ítem.
		// %.2f imprime el promedio con dos decimales de precisión.
		fmt.Printf("ID: %d, Título: %s, Tipo: %s, Duración: %d min, Calificación Promedio: %.2f\n", item.ID, item.Title, item.Type, item.Duration, item.AverageRating)
	}
}

// CalificarContenido permite a un usuario (identificado por userID) calificar un contenido específico (identificado por contentID).
// Recibe como parámetros el ID del contenido, el ID del usuario que califica, y la calificación como una cadena de texto (string).
func CalificarContenido(contentID int, userID int, ratingStr string) error {
	// Buscar el índice del contenido en la lista 'contentInMemory' usando el 'contentID' proporcionado.
	// Inicializa el índice en -1, asumiendo que no se encontró.
	contentIndex := -1
	// Itera sobre la lista de contenido.
	for i, item := range contentInMemory {
		// Verifica si el ID del ítem actual coincide con el ID del contenido buscado.
		if item.ID == contentID {
			// Si coincide, almacena el índice (posición) del contenido en la lista.
			contentIndex = i
			// Rompe el bucle una vez encontrado para no seguir iterando innecesariamente.
			break
		}
	}

	// Verifica si el contenido no fue encontrado en la lista.
	if contentIndex == -1 {
		// Retorna un error si no se encontró el contenido con el ID especificado.
		return fmt.Errorf("contenido con ID %d no encontrado", contentID)
	}

	// Convierte la calificación de string a float64, manejando punto y coma como separador decimal.
	// Reemplaza la primera coma (',') por un punto ('.') en la cadena de texto para estandarizar el formato decimal.
	// Esto permite que entradas como "7,5" se conviertan correctamente a 7.5.
	ratingStr = strings.Replace(ratingStr, ",", ".", 1)
	// Convierte la cadena de texto estandarizada (ratingStr) a un número decimal (float64).
	rating, err := strconv.ParseFloat(ratingStr, 64)
	// Verifica si ocurrió un error durante la conversión (por ejemplo, si la cadena no era un número válido).
	if err != nil {
		// Retorna un error si la calificación no es un número decimal válido.
		return fmt.Errorf("calificación '%s' no es válida. Debe ser un número entre 1.0 y 10.0", ratingStr)
	}

	// Validar que la calificación esté dentro del rango permitido (1.0 a 10.0).
	if rating < 1.0 || rating > 10.0 {
		// Retorna un error si la calificación está fuera del rango permitido.
		return fmt.Errorf("la calificación debe estar entre 1.0 y 10.0, recibido: %.2f", rating)
	}

	// Verificar si el usuario (userID) ya ha calificado este contenido anteriormente.
	// Inicializa el índice de la calificación existente en -1.
	existingRatingIndex := -1
	// Itera sobre la lista de calificaciones individuales del contenido encontrado (en contentInMemory[contentIndex].Ratings).
	for i, ur := range contentInMemory[contentIndex].Ratings {
		// Verifica si el UserID de la calificación actual en la iteración coincide con el UserID del usuario que está calificando ahora.
		if ur.UserID == userID {
			// Si coincide, almacena el índice de la calificación existente.
			existingRatingIndex = i
			// Rompe el bucle una vez encontrado para no seguir iterando innecesariamente.
			break
		}
	}

	// Crea una nueva instancia de la estructura 'UserRating' con los datos del usuario y la calificación.
	newUserRating := UserRating{UserID: userID, Rating: rating}

	// Verifica si el usuario ya había calificado este contenido (si se encontró un índice existente).
	if existingRatingIndex != -1 {
		// Si ya existía una calificación del mismo usuario, actualiza el valor de la calificación en la lista existente.
		// Sobrescribe la calificación en la posición 'existingRatingIndex'.
		contentInMemory[contentIndex].Ratings[existingRatingIndex] = newUserRating
	} else {
		// Si no existía una calificación previa del mismo usuario, agrega la nueva calificación a la lista de calificaciones del contenido.
		contentInMemory[contentIndex].Ratings = append(contentInMemory[contentIndex].Ratings, newUserRating)
	}

	// Recalcular el promedio de calificaciones para el contenido después de agregar o actualizar una calificación.
	recalcularPromedio(contentIndex)
	// Imprime un mensaje de confirmación indicando que la calificación se realizó correctamente.
	fmt.Printf("Contenido '%s' calificado exitosamente con %.2f por el usuario %d.\n", contentInMemory[contentIndex].Title, rating, userID)
	// Retorna nil indicando que no hubo error durante el proceso de calificación.
	return nil
}

// recalcularPromedio recalcula el AverageRating de un contenido basado en sus calificaciones almacenadas en la lista 'Ratings'.
// Recibe como parámetro el índice del contenido en la lista 'contentInMemory'.
func recalcularPromedio(index int) {
	// Inicializa una variable 'sum' de tipo float64 para acumular la suma total de todas las calificaciones individuales.
	sum := 0.0
	// Itera sobre la lista de calificaciones individuales (contentInMemory[index].Ratings) del contenido en la posición 'index'.
	for _, ur := range contentInMemory[index].Ratings {
		// Suma cada calificación individual (ur.Rating) al acumulador 'sum'.
		sum += ur.Rating
	}

	// Verifica si hay al menos una calificación registrada en la lista.
	if len(contentInMemory[index].Ratings) > 0 {
		// Si hay calificaciones, calcula el promedio dividiendo la suma total ('sum') por la cantidad de calificaciones (la longitud de la lista).
		// Se convierte la longitud de la lista (int) a float64 para la división.
		contentInMemory[index].AverageRating = sum / float64(len(contentInMemory[index].Ratings))
	} else {
		// Si no hay calificaciones registradas (lo cual es raro después de calificar, pero se contempla), el promedio se establece en 0.0.
		contentInMemory[index].AverageRating = 0.0
	}
}

// BuscarContenidoPorTitulo es una función adicional que podría ser útil en fases posteriores del proyecto (por ejemplo, AA2).
// Esta función busca un ítem de contenido específico en la lista 'contentInMemory' por su título exacto.
// Devuelve un puntero al contenido encontrado y un valor booleano que indica si se encontró o no.
// func BuscarContenidoPorTitulo(titulo string) (*Content, bool) {
// 	// Itera sobre la lista de contenido.
// 	for i, item := range contentInMemory {
// 		// Compara el título del ítem actual en la iteración con el título buscado.
// 		// strings.EqualFold(item.Title, titulo) compara sin distinguir mayúsculas de minúsculas.
// 		if item.Title == titulo { // Comparación simple por título (distingue mayúsculas/minúsculas en este ejemplo)
// 			// Si se encuentra, devuelve un puntero al contenido (usando &contentInMemory[i]) y true.
// 			return &contentInMemory[i], true
// 		}
// 	}
// 	// Si el bucle termina y no se encontró el contenido, devuelve nil (ningún contenido) y false.
// 	return nil, false
// }