package main

import (
    "fmt"
    "os"
    "strings"
    "time"
    "SDGEStreaming/internal/admin"
    "SDGEStreaming/internal/audio"
    "SDGEStreaming/internal/audiovisual"
    "SDGEStreaming/internal/categories"
    "SDGEStreaming/internal/contentclass"
    "SDGEStreaming/internal/errors"
    "SDGEStreaming/internal/genres"
    "SDGEStreaming/internal/profiles"
    "SDGEStreaming/internal/ratings"
    "SDGEStreaming/internal/utils"
    "SDGEStreaming/pkg/cli"
)

// Variables globales para la sesión
var (
    currentUser *categories.User
    currentSessionID string
    lastActivity time.Time
    sessionTimeout = 5 * time.Minute // 5 minutos de inactividad
)

func main() {
    utils.ClearScreen()
    cli.ShowHeader()
    
    // Inicializo la última actividad
    lastActivity = time.Now()
    
    // Bucle principal de la aplicación
    for {
        // Verifico si la sesión ha expirado por inactividad
        if currentUser != nil && time.Since(lastActivity) > sessionTimeout {
            cli.PrintWarning("Sesión expirada por inactividad. Por favor inicie sesión nuevamente.")
            currentUser = nil
            utils.WaitForEnter()
            continue
        }
        
        if currentUser == nil {
            showAuthMenu()
        } else {
            showMainMenu()
        }
    }
}

// Muestro el menú de autenticación
func showAuthMenu() {
    items := []cli.MenuItem{
        {"1", "Iniciar Sesión"},
        {"2", "Registrarse"},
        {"3", "Explorar como Invitado"},
        {"4", "Salir"},
    }
    
    option := cli.ShowMenu("Bienvenido a SDGEStreaming", items)
    updateActivity()
    
    switch option {
    case "1":
        login()
    case "2":
        register()
    case "3":
        guestMode()
    case "4":
        utils.ClearScreen()
        if currentUser != nil {
            cli.ShowGoodbye(currentUser.Name, false)
        } else {
            cli.ShowGoodbye("Invitado", true)
        }
        os.Exit(0)
    default:
        if option != "" {
            cli.PrintError("Opción inválida. Por favor seleccione una opción del menú.")
        }
    }
}

// Manejo el inicio de sesión
func login() {
    utils.ClearScreen()
    cli.PrintTitle("Iniciar Sesión")
    
    email, back := cli.ReadEmail("Email: ")
    if back {
        return
    }
    
    password, back := cli.ReadPassword("Contraseña: ", false)
    if back {
        return
    }
    
    // Busco el usuario
    user, err := profiles.FindByEmail(email)
    if err != nil {
        cli.PrintError("Usuario no encontrado")
        utils.WaitForEnter()
        return
    }
    
    // Valido la contraseña
    if user.Password != password {
        cli.PrintError("Contraseña incorrecta")
        utils.WaitForEnter()
        return
    }
    
    // Actualizo el último inicio de sesión
    profiles.UpdateLastLogin(user.ID)
    
    // Establezco la sesión actual
    currentUser = user
    currentSessionID = fmt.Sprintf("sess_%d_%d", user.ID, time.Now().Unix())
    lastActivity = time.Now()
    
    cli.ShowWelcome(user.Name, false)
    utils.WaitForEnter()
}

// Manejo el registro de nuevos usuarios
func register() {
    utils.ClearScreen()
    cli.PrintTitle("Registro de Nuevo Usuario")
    
    // Leo los datos del usuario
    name, back := cli.ReadName("Nombre completo: ")
    if back {
        return
    }
    
    age, back := cli.ReadAge("Edad: ")
    if back {
        return
    }
    
    email, back := cli.ReadEmail("Email: ")
    if back {
        return
    }
    
    password, back := cli.ReadPassword("Contraseña (6-32 caracteres): ", true)
    if back {
        return
    }
    
    // Muestro las clasificaciones por edad disponibles
    cli.PrintSubtitle("Clasificación por Edad")
    for i, rating := range contentclass.GetAllRatings() {
        fmt.Printf("%d. %s - %s\n", i+1, rating.Name, rating.Description)
    }
    
    ageRatingID, back := cli.ReadInt("Seleccione su clasificación (1-3): ")
    if back {
        return
    }
    
    if ageRatingID < 1 || ageRatingID > len(contentclass.GetAllRatings()) {
        cli.PrintError("Opción inválida")
        utils.WaitForEnter()
        return
    }
    
    ageRating := contentclass.GetAllRatings()[ageRatingID-1].Name
    
    // Registro el usuario con plan Free por defecto
    _, err := profiles.AddUser(name, age, email, password, "Free", ageRating, false)
    if err != nil {
        errors.HandleAppError(err)
        utils.WaitForEnter()
        return
    }
    
    cli.PrintSuccess("¡Usuario registrado exitosamente!")
    cli.PrintWarning("Su cuenta tiene el plan Free. Para acceder a funciones premium, contacte al administrador.")
    utils.WaitForEnter()
}

// Modo invitado con acceso limitado
func guestMode() {
    currentUser = nil
    showContentMenu(true)
}
