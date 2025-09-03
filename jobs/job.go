/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package jobs

import "strconv"

// JobApplication represents a job application entry
type JobApplication struct {
	ID        int    `json:"id"`
	Company   string `json:"company"`
	Position  string `json:"position"`
	Status    string `json:"status"`
	AppliedOn string `json:"applied_on"`
}

// getNextID returns the next available ID for a new job application
func getNextID(jobs []JobApplication) int {
	maxID := 0
	for _, job := range jobs {
		if job.ID > maxID {
			maxID = job.ID
		}
	}
	return maxID + 1
}

// convertToStringSlice converts a JobApplication to a slice of strings
func (app *JobApplication) convertToStringSlice() []string {
	return []string{
		strconv.Itoa(app.ID),
		app.Company,
		app.Position,
		app.Status,
		app.AppliedOn,
	}
}
