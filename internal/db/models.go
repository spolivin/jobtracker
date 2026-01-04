/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"strconv"
	"time"
)

// JobApplication represents a job application record in the database.
type JobApplication struct {
	ID        int       `json:"id"`
	Company   string    `json:"company"`
	Position  string    `json:"position"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConvertToStringSlice converts a JobApplication to a slice of strings for display.
func (app JobApplication) ConvertToStringSlice() []string {
	return []string{
		strconv.Itoa(app.ID),
		app.Company,
		app.Position,
		app.Status,
		app.CreatedAt.Format(time.RFC3339),
		app.UpdatedAt.Format(time.RFC3339),
	}
}
