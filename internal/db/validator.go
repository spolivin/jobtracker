/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"fmt"
	"strings"
)

// validColumns defines the allowed column names for the applications table
var validColumns = map[string]bool{
	"id":         true,
	"company":    true,
	"position":   true,
	"status":     true,
	"created_at": true,
	"updated_at": true,
}

// ValidateColumnName checks if a column name is valid for SQL operations.
// It prevents SQL injection by ensuring only whitelisted columns are used.
func ValidateColumnName(column string) error {
	// Normalize to lowercase for case-insensitive comparison
	normalized := strings.ToLower(strings.TrimSpace(column))

	// Check for empty string after trimming
	if normalized == "" {
		return fmt.Errorf("column name cannot be empty")
	}

	if !validColumns[normalized] {
		return fmt.Errorf("invalid column name: %q (allowed: id, company, position, status, created_at, updated_at)", column)
	}

	return nil
}

// ValidateColumnNames validates multiple column names at once.
func ValidateColumnNames(columns []string) error {
	for _, col := range columns {
		if err := ValidateColumnName(col); err != nil {
			return err
		}
	}
	return nil
}
