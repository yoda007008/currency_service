package migration

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigrations(connStr string, migrationsPath string) error {
	// Получаем абсолютный путь до папки с миграциями
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %w", err)
	}
	log.Println("Applying migrations from:", absPath)

	// Создаём подключение к базе данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening db: %w", err)
	}
	defer db.Close()

	// Создаём драйвер для миграций
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("error creating driver: %w", err)
	}

	// Формируем путь с префиксом "file://", указывая абсолютный путь
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+absPath,
		"postgres", driver,
	)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %w", err)
	}

	// Применяем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration up failed: %w", err)
	}

	log.Println("Migrations applied successfully.")
	return nil
}
