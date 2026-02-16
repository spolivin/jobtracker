/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"reflect"
	"testing"
	"time"
)

func TestJobApplication_ConvertToStringSlice(t *testing.T) {
	tests := []struct {
		name string
		app  JobApplication
		want []string
	}{
		{
			name: "normal application",
			app: JobApplication{
				ID:        1,
				Company:   "Google",
				Position:  "Software Engineer",
				Status:    "Applied",
				CreatedAt: time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC),
				UpdatedAt: time.Date(2026, 1, 15, 10, 30, 0, 0, time.UTC),
			},
			want: []string{
				"1",
				"Google",
				"Software Engineer",
				"Applied",
				"2026-01-15T10:30:00Z",
				"2026-01-15T10:30:00Z",
			},
		},
		{
			name: "application with zero ID",
			app: JobApplication{
				ID:        0,
				Company:   "Facebook",
				Position:  "Data Scientist",
				Status:    "Interview",
				CreatedAt: time.Date(2026, 2, 1, 14, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2026, 2, 1, 14, 0, 0, 0, time.UTC),
			},
			want: []string{
				"0",
				"Facebook",
				"Data Scientist",
				"Interview",
				"2026-02-01T14:00:00Z",
				"2026-02-01T14:00:00Z",
			},
		},
		{
			name: "application with empty strings",
			app: JobApplication{
				ID:        123,
				Company:   "",
				Position:  "",
				Status:    "",
				CreatedAt: time.Date(2026, 3, 10, 9, 15, 30, 0, time.UTC),
				UpdatedAt: time.Date(2026, 3, 10, 9, 15, 30, 0, time.UTC),
			},
			want: []string{
				"123",
				"",
				"",
				"",
				"2026-03-10T09:15:30Z",
				"2026-03-10T09:15:30Z",
			},
		},
		{
			name: "application with large ID",
			app: JobApplication{
				ID:        999999,
				Company:   "Tesla",
				Position:  "ML Engineer",
				Status:    "Offer",
				CreatedAt: time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
				UpdatedAt: time.Date(2026, 12, 31, 23, 59, 59, 0, time.UTC),
			},
			want: []string{
				"999999",
				"Tesla",
				"ML Engineer",
				"Offer",
				"2026-12-31T23:59:59Z",
				"2026-12-31T23:59:59Z",
			},
		},
		{
			name: "application with timezone offset",
			app: JobApplication{
				ID:        42,
				Company:   "Amazon",
				Position:  "DevOps Engineer",
				Status:    "Rejected",
				CreatedAt: time.Date(2026, 6, 15, 8, 0, 0, 0, time.FixedZone("EST", -5*3600)),
				UpdatedAt: time.Date(2026, 6, 15, 10, 0, 0, 0, time.FixedZone("EST", -5*3600)),
			},
			want: []string{
				"42",
				"Amazon",
				"DevOps Engineer",
				"Rejected",
				"2026-06-15T08:00:00-05:00",
				"2026-06-15T10:00:00-05:00",
			},
		},
		{
			name: "application with special characters in fields",
			app: JobApplication{
				ID:        7,
				Company:   "Company & Co.",
				Position:  "Senior Engineer (L5)",
				Status:    "Applied - In Review",
				CreatedAt: time.Date(2026, 4, 1, 12, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2026, 4, 1, 12, 0, 0, 0, time.UTC),
			},
			want: []string{
				"7",
				"Company & Co.",
				"Senior Engineer (L5)",
				"Applied - In Review",
				"2026-04-01T12:00:00Z",
				"2026-04-01T12:00:00Z",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.app.ConvertToStringSlice()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToStringSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJobApplication_ConvertToStringSlice_Length(t *testing.T) {
	app := JobApplication{
		ID:        1,
		Company:   "Test Corp",
		Position:  "Engineer",
		Status:    "Applied",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := app.ConvertToStringSlice()

	expectedLength := 6
	if len(result) != expectedLength {
		t.Errorf("ConvertToStringSlice() returned slice of length %d, want %d", len(result), expectedLength)
	}
}

func TestJobApplication_ConvertToStringSlice_Order(t *testing.T) {
	// Verify the order of fields in the slice
	app := JobApplication{
		ID:        100,
		Company:   "TestCompany",
		Position:  "TestPosition",
		Status:    "TestStatus",
		CreatedAt: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC),
	}

	result := app.ConvertToStringSlice()

	// Check each field is in the correct position
	if result[0] != "100" {
		t.Errorf("Expected ID at index 0, got %s", result[0])
	}
	if result[1] != "TestCompany" {
		t.Errorf("Expected Company at index 1, got %s", result[1])
	}
	if result[2] != "TestPosition" {
		t.Errorf("Expected Position at index 2, got %s", result[2])
	}
	if result[3] != "TestStatus" {
		t.Errorf("Expected Status at index 3, got %s", result[3])
	}
	if result[4] != "2026-01-01T00:00:00Z" {
		t.Errorf("Expected CreatedAt at index 4, got %s", result[4])
	}
	if result[5] != "2026-01-02T00:00:00Z" {
		t.Errorf("Expected UpdatedAt at index 5, got %s", result[5])
	}
}

func TestJobApplication_ConvertToStringSlice_TimeFormat(t *testing.T) {
	// Verify time is formatted as RFC3339
	testTime := time.Date(2026, 5, 20, 15, 45, 30, 0, time.UTC)
	app := JobApplication{
		ID:        1,
		Company:   "Test",
		Position:  "Test",
		Status:    "Test",
		CreatedAt: testTime,
		UpdatedAt: testTime,
	}

	result := app.ConvertToStringSlice()

	expectedTimeStr := testTime.Format(time.RFC3339)
	if result[4] != expectedTimeStr {
		t.Errorf("CreatedAt not formatted as RFC3339: got %s, want %s", result[4], expectedTimeStr)
	}
	if result[5] != expectedTimeStr {
		t.Errorf("UpdatedAt not formatted as RFC3339: got %s, want %s", result[5], expectedTimeStr)
	}
}

func TestJobApplication_StructFields(t *testing.T) {
	// Test that struct can be initialized with all fields
	app := JobApplication{
		ID:        42,
		Company:   "TestCorp",
		Position:  "Engineer",
		Status:    "Applied",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if app.ID != 42 {
		t.Errorf("Expected ID 42, got %d", app.ID)
	}
	if app.Company != "TestCorp" {
		t.Errorf("Expected Company 'TestCorp', got %s", app.Company)
	}
	if app.Position != "Engineer" {
		t.Errorf("Expected Position 'Engineer', got %s", app.Position)
	}
	if app.Status != "Applied" {
		t.Errorf("Expected Status 'Applied', got %s", app.Status)
	}
}

func TestJobApplication_JSONTags(t *testing.T) {
	// Verify struct has correct JSON tags
	app := JobApplication{}
	appType := reflect.TypeOf(app)

	expectedJSONTags := map[string]string{
		"ID":        "id",
		"Company":   "company",
		"Position":  "position",
		"Status":    "status",
		"CreatedAt": "created_at",
		"UpdatedAt": "updated_at",
	}

	for fieldName, expectedTag := range expectedJSONTags {
		field, found := appType.FieldByName(fieldName)
		if !found {
			t.Errorf("Field %s not found in JobApplication struct", fieldName)
			continue
		}

		jsonTag := field.Tag.Get("json")
		if jsonTag != expectedTag {
			t.Errorf("Field %s has JSON tag %q, want %q", fieldName, jsonTag, expectedTag)
		}
	}
}

// Benchmark tests
func BenchmarkConvertToStringSlice(b *testing.B) {
	app := JobApplication{
		ID:        123,
		Company:   "BenchCorp",
		Position:  "Engineer",
		Status:    "Applied",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	b.ResetTimer()
	for b.Loop() {
		_ = app.ConvertToStringSlice()
	}
}

func BenchmarkConvertToStringSlice_LargeID(b *testing.B) {
	app := JobApplication{
		ID:        999999999,
		Company:   "LargeIDCorp",
		Position:  "Senior Engineer",
		Status:    "Interview",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	b.ResetTimer()
	for b.Loop() {
		_ = app.ConvertToStringSlice()
	}
}
