// internal/db/db.go
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB es la conexión global a la base de datos SQLite
var DB *sql.DB

// InitDB inicializa la base de datos, activa foreign keys y ejecuta migraciones.
// dbPath: ruta al archivo .db
// migrationsPath: ruta al directorio de migraciones SQL
func InitDB(dbPath string, migrationsPath string) error {
	var err error

	DB, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("error al abrir la base de datos: %w", err)
	}

	// SQLite no es concurrente por defecto, se limita a 1 conexión
	DB.SetMaxOpenConns(1)

	// Activar claves foráneas
	if _, err := DB.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return fmt.Errorf("error al activar foreign keys: %w", err)
	}

	// Ejecutar migraciones
	if err := RunMigrations(DB, migrationsPath); err != nil {
		return fmt.Errorf("error ejecutando migraciones: %w", err)
	}

	return nil
}

// GetDB devuelve la conexión activa a la base de datos
func GetDB() *sql.DB {
	return DB
}

// Close cierra la conexión a la base de datos
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
