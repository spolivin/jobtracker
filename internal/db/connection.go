/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/spolivin/jobtracker/v2/internal/db/config"
)

// Connect connects to the Postgres database using config and password.
func Connect(ctx context.Context, cfg *config.ConnectionConfig, password string) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable connect_timeout=15",
		cfg.DBHost,
		strconv.Itoa(cfg.DBPort),
		cfg.DBUser,
		password,
		cfg.DBName,
	)
	var err error
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}

// CheckTableExists checks if a table exists in the database.
func CheckTableExists(ctx context.Context, db *sql.DB, tableName string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (
		       SELECT 1 FROM information_schema.tables 
		       WHERE table_schema = 'public' AND table_name = $1
	       )`
	if err := db.QueryRowContext(ctx, query, tableName).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}
