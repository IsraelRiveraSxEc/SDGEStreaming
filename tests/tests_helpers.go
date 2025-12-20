// tests/tests_helpers.go
package tests

import (
	"os"
	"path/filepath"
	"testing"

	"SDGEStreaming/internal/db"
)

var testDBPath string

// SetupTestDB inicializa una base de datos SQLite temporal
// y ejecuta las migraciones reales del proyecto.
func SetupTestDB(t *testing.T) {
	t.Helper()

	tmpDir := t.TempDir()
	testDBPath = filepath.Join(tmpDir, "test.db")

	// Ruta correcta a las migraciones desde /tests
	migrationsPath := filepath.Join("..", "internal", "db", "migrations")

	if err := db.InitDB(testDBPath, migrationsPath); err != nil {
		t.Fatalf("error inicializando DB de test: %v", err)
	}
}

// TeardownTestDB cierra y elimina la base de datos de test.
func TeardownTestDB() {
	_ = db.Close()
	if testDBPath != "" {
		_ = os.Remove(testDBPath)
	}
}
