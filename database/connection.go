package database

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func extractDatabaseName(dsn string) string {
	parts := strings.Split(dsn, "/")
	if len(parts) == 0 {
		return ""
	}

	last := parts[len(parts)-1]
	if before, _, ok := strings.Cut(last, "?"); ok {
		return before
	}

	return last
}

var Pool *pgxpool.Pool

func Connect() {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatalln("DATABASE_URL not set")
	}

	dbName := extractDatabaseName(url)
	if dbName == "" {
		log.Fatalln("Could not parse database name from DATABASE_URL")
	}

	adminURL := strings.Replace(url, "/"+dbName, "/postgres", 1)
	conn, err := pgx.Connect(context.Background(), adminURL)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL server: %v\n", err)
	}
	defer conn.Close(context.Background())

	var databaseExists bool

	err = conn.QueryRow(context.Background(), `SELECT EXISTS(SELECT FROM pg_database WHERE datname = $1);`, dbName).Scan(&databaseExists)
	if err == nil {
		if !databaseExists {
			_, err = conn.Exec(context.Background(), fmt.Sprintf(`CREATE DATABASE "%s"`, dbName))
			if err != nil {
				log.Fatalf("Failed to create database: %v\n", err)
			}
		}
	} else {
		log.Printf("Failed to check if database exists: %v\n", err)
	}

	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	Pool = pool

	fmt.Println("Connected to PostgreSQL")
}

func RunMigrations(migrationsFS embed.FS) {
	sourceDriver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		log.Fatalf("Migration source error: %v\n", err)
	}

	migration, err := migrate.NewWithSourceInstance(
		"iofs",
		sourceDriver,
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatalf("Migration init error: %v\n", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration error: %v\n", err)
	}

	fmt.Println("Migrations applied successfully.")
}
