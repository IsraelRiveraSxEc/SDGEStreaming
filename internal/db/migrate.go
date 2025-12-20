// internal/db/migrations.go
package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// RunMigrations ejecuta todos los archivos SQL contenidos en la carpeta
// de migraciones indicada. Los archivos se ejecutan en orden alfabético.
func RunMigrations(db *sql.DB, migrationsPath string) error {
	files, err := os.ReadDir(migrationsPath)
	if err != nil {
		return fmt.Errorf("error leyendo migraciones: %w", err)
	}

	var migrationFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			migrationFiles = append(migrationFiles, filepath.Join(migrationsPath, file.Name()))
		}
	}

	sort.Strings(migrationFiles)

	for _, filePath := range migrationFiles {
		sqlBytes, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("error leyendo migración %s: %w", filePath, err)
		}

		if _, err := db.Exec(string(sqlBytes)); err != nil {
			return fmt.Errorf("error ejecutando migración %s: %w", filePath, err)
		}
	}

	return nil
}
