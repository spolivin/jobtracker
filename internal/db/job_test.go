/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"context"
	"testing"
)

// TestReadSQLInjectionProtection tests that Read validates the sortBy parameter
// and rejects SQL injection attempts
func TestReadSQLInjectionProtection(t *testing.T) {
	tests := []struct {
		name              string
		sortBy            string
		expectValidation  bool // true if we expect a validation error
		validationMessage string
	}{
		{
			name:             "SQL injection: DROP TABLE",
			sortBy:           "id; DROP TABLE applications--",
			expectValidation: true,
			validationMessage: "invalid column name",
		},
		{
			name:             "SQL injection: UNION SELECT",
			sortBy:           "id UNION SELECT * FROM users",
			expectValidation: true,
			validationMessage: "invalid column name",
		},
		{
			name:             "SQL injection: OR 1=1",
			sortBy:           "id OR 1=1",
			expectValidation: true,
			validationMessage: "invalid column name",
		},
		{
			name:             "Invalid column name",
			sortBy:           "invalid_column",
			expectValidation: true,
			validationMessage: "invalid column name",
		},
		{
			name:             "Invalid column: spaces",
			sortBy:           "company name",
			expectValidation: true,
			validationMessage: "invalid column name",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &JobApplicationsStore{db: nil}
			ctx := context.Background()

			_, err := store.Read(ctx, tt.sortBy, false)

			if tt.expectValidation {
				if err == nil {
					t.Errorf("Read() with sortBy=%q should return validation error, got nil", tt.sortBy)
					return
				}
				if err.Error()[:len(tt.validationMessage)] != tt.validationMessage {
					t.Errorf("Read() error = %v, want error containing %q", err, tt.validationMessage)
				}
			}
		})
	}
}

// TestUpdateSQLInjectionProtection tests that Update validates field names
// and rejects SQL injection attempts
func TestUpdateSQLInjectionProtection(t *testing.T) {
	tests := []struct {
		name              string
		fields            map[string]string
		expectValidation  bool
		validationMessage string
	}{
		{
			name: "SQL injection: DROP TABLE in field name",
			fields: map[string]string{
				"company; DROP TABLE applications--": "evil",
			},
			expectValidation:  true,
			validationMessage: "invalid column name",
		},
		{
			name: "SQL injection: UNION in field name",
			fields: map[string]string{
				"status UNION SELECT password FROM users--": "malicious",
			},
			expectValidation:  true,
			validationMessage: "invalid column name",
		},
		{
			name: "Invalid field name",
			fields: map[string]string{
				"invalid_column": "value",
			},
			expectValidation:  true,
			validationMessage: "invalid column name",
		},
		{
			name: "Mix of valid and invalid fields",
			fields: map[string]string{
				"company":       "Microsoft",
				"invalid_field": "bad",
			},
			expectValidation:  true,
			validationMessage: "invalid column name",
		},
		{
			name: "SQL injection: semicolon attack",
			fields: map[string]string{
				"company; DELETE FROM applications WHERE 1=1--": "attack",
			},
			expectValidation:  true,
			validationMessage: "invalid column name",
		},
		{
			name:             "Empty fields map returns early",
			fields:           map[string]string{},
			expectValidation: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &JobApplicationsStore{db: nil}
			ctx := context.Background()

			_, err := store.Update(ctx, 1, tt.fields)

			if tt.expectValidation {
				if err == nil {
					t.Errorf("Update() with fields=%v should return validation error, got nil", tt.fields)
					return
				}
				if err.Error()[:len(tt.validationMessage)] != tt.validationMessage {
					t.Errorf("Update() error = %v, want error containing %q", err, tt.validationMessage)
				}
			} else if len(tt.fields) == 0 {
				// Empty fields should return 0, nil
				if err != nil {
					t.Errorf("Update() with empty fields should not error, got %v", err)
				}
			}
		})
	}
}

// BenchmarkReadValidation benchmarks the Read method with validation
func BenchmarkReadValidation(b *testing.B) {
	store := &JobApplicationsStore{db: nil}
	ctx := context.Background()

	b.ResetTimer()
	for b.Loop() {
		_, _ = store.Read(ctx, "company", false)
	}
}

// BenchmarkUpdateValidation benchmarks the Update method with validation
func BenchmarkUpdateValidation(b *testing.B) {
	store := &JobApplicationsStore{db: nil}
	ctx := context.Background()
	fields := map[string]string{
		"company":  "TestCorp",
		"position": "Engineer",
		"status":   "Applied",
	}

	b.ResetTimer()
	for b.Loop() {
		_, _ = store.Update(ctx, 1, fields)
	}
}
