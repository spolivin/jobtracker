/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package migrate

import (
	"context"
	"database/sql"
	"embed"
	"sort"
)

//go:embed migrations/*.sql
var fs embed.FS

// Run executes the defined migration scripts.
func Run(ctx context.Context, db *sql.DB) error {
	// Making sure schema_migrations table exists
	if _, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY
		)
	`); err != nil {
		return err
	}

	// Reading applied migrations
	rows, err := db.QueryContext(ctx,
		`SELECT version FROM schema_migrations`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	applied := map[string]bool{}
	for rows.Next() {
		var v string
		if err := rows.Scan(&v); err != nil {
			return err
		}
		applied[v] = true
	}

	// Reading embedded files
	entries, err := fs.ReadDir("migrations")
	if err != nil {
		return err
	}

	var files []string
	for _, e := range entries {
		files = append(files, e.Name())
	}
	sort.Strings(files)

	// Applying missing migrations
	for _, name := range files {
		if applied[name] {
			continue
		}

		sqlBytes, err := fs.ReadFile("migrations/" + name)
		if err != nil {
			return err
		}

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		if _, err := tx.ExecContext(ctx, string(sqlBytes)); err != nil {
			tx.Rollback()
			return err
		}

		if _, err := tx.ExecContext(ctx,
			`INSERT INTO schema_migrations (version) VALUES ($1)`,
			name,
		); err != nil {
			tx.Rollback()
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}
