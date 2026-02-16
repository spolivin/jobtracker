/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"strings"
	"testing"
)

func TestValidateColumnName(t *testing.T) {
	tests := []struct {
		name    string
		column  string
		wantErr bool
	}{
		// Valid columns
		{"valid id", "id", false},
		{"valid company", "company", false},
		{"valid position", "position", false},
		{"valid status", "status", false},
		{"valid created_at", "created_at", false},
		{"valid updated_at", "updated_at", false},

		// Case variations (should be accepted)
		{"uppercase ID", "ID", false},
		{"uppercase COMPANY", "COMPANY", false},
		{"mixed case Company", "Company", false},
		{"mixed case CreaTed_At", "CreaTed_At", false},

		// Whitespace (should be trimmed and accepted)
		{"leading space", " company", false},
		{"trailing space", "company ", false},
		{"both spaces", " position ", false},
		{"multiple spaces", "  status  ", false},

		// Invalid columns
		{"invalid column", "invalid_column", true},
		{"non-existent field", "email", true},
		{"sql injection attempt 1", "id; DROP TABLE applications", true},
		{"sql injection attempt 2", "id--", true},
		{"sql injection attempt 3", "1=1 OR", true},
		{"sql injection with comment", "id/* comment */", true},

		// Edge cases
		{"empty string", "", true},
		{"only spaces", "   ", true},
		{"special characters", "@#$%", true},
		{"sql keyword", "SELECT", true},
		{"partial match", "comp", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateColumnName(tt.column)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateColumnName(%q) error = %v, wantErr %v", tt.column, err, tt.wantErr)
			}
		})
	}
}

func TestValidateColumnNames(t *testing.T) {
	tests := []struct {
		name    string
		columns []string
		wantErr bool
	}{
		{
			name:    "all valid columns",
			columns: []string{"id", "company", "position"},
			wantErr: false,
		},
		{
			name:    "valid with mixed case",
			columns: []string{"ID", "Company", "STATUS"},
			wantErr: false,
		},
		{
			name:    "valid with whitespace",
			columns: []string{" id ", "company", "  position  "},
			wantErr: false,
		},
		{
			name:    "all valid columns (all fields)",
			columns: []string{"id", "company", "position", "status", "created_at", "updated_at"},
			wantErr: false,
		},
		{
			name:    "one invalid column",
			columns: []string{"id", "invalid_field", "company"},
			wantErr: true,
		},
		{
			name:    "all invalid columns",
			columns: []string{"invalid1", "invalid2"},
			wantErr: true,
		},
		{
			name:    "sql injection in list",
			columns: []string{"id", "company; DROP TABLE applications"},
			wantErr: true,
		},
		{
			name:    "empty slice",
			columns: []string{},
			wantErr: false,
		},
		{
			name:    "contains empty string",
			columns: []string{"id", "", "company"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateColumnNames(tt.columns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateColumnNames(%v) error = %v, wantErr %v", tt.columns, err, tt.wantErr)
			}
		})
	}
}

func TestValidateColumnNameErrorMessage(t *testing.T) {
	err := ValidateColumnName("invalid_field")
	if err == nil {
		t.Fatal("expected error but got nil")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "invalid column name") {
		t.Errorf("error message should contain 'invalid column name', got: %s", errMsg)
	}

	if !strings.Contains(errMsg, "invalid_field") {
		t.Errorf("error message should contain the invalid field name, got: %s", errMsg)
	}

	// Check that error message lists allowed columns
	allowedColumns := []string{"id", "company", "position", "status", "created_at", "updated_at"}
	for _, col := range allowedColumns {
		if !strings.Contains(errMsg, col) {
			t.Errorf("error message should list allowed column %q, got: %s", col, errMsg)
		}
	}
}

func TestValidateColumnNameEmptyError(t *testing.T) {
	err := ValidateColumnName("")
	if err == nil {
		t.Fatal("expected error for empty column name but got nil")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "cannot be empty") {
		t.Errorf("error message should indicate empty column, got: %s", errMsg)
	}
}

// Benchmark tests
func BenchmarkValidateColumnName(b *testing.B) {
	for b.Loop() {
		ValidateColumnName("company")
	}
}

func BenchmarkValidateColumnNameInvalid(b *testing.B) {
	for b.Loop() {
		ValidateColumnName("invalid_column")
	}
}

func BenchmarkValidateColumnNames(b *testing.B) {
	columns := []string{"id", "company", "position", "status"}
	b.ResetTimer()
	for b.Loop() {
		ValidateColumnNames(columns)
	}
}
