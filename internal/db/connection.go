/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDatabase initializes the database connection using environment variables.
func InitDatabase(ctx context.Context) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err := DB.PingContext(ctx); err != nil {
		return err
	}
	return nil
}

// EnsureSchema creates necessary tables if they do not exist.
func EnsureSchema(ctx context.Context) error {
	schema := `
	CREATE TABLE IF NOT EXISTS applications (
		id SERIAL PRIMARY KEY,
		company VARCHAR(255) NOT NULL,
		position VARCHAR(255) NOT NULL,
		status VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := DB.ExecContext(ctx, schema)
	return err
}

// CloseDatabaseConnection closes the database connection.
func CloseDatabaseConnection() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

// CheckTableExists checks if a table exists in the database.
func CheckTableExists(ctx context.Context, tableName string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (
		       SELECT 1 FROM information_schema.tables 
		       WHERE table_schema = 'public' AND table_name = $1
	       )`
	if err := DB.QueryRowContext(ctx, query, tableName).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
