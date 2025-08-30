/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package jobs

import (
	"encoding/json"
	"os"
)

const fileName = "jobs.json" // File to store job applications

// LoadJobs reads jobs.json into a slice of Job
func loadJobApplications() ([]JobApplication, error) {

	// Checking if the file exists
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []JobApplication{}, nil // If file doesn’t exist, then start empty
	}
	// Reading the file
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// Unmarshalling the JSON data
	var jobs []JobApplication
	if err := json.Unmarshal(data, &jobs); err != nil {
		return nil, err
	}
	return jobs, nil
}

// SaveJobs writes slice of Job into jobs.json
func saveJobApplications(jobs []JobApplication) error {

	// Marshalling the JSON data
	data, err := json.MarshalIndent(jobs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0644)
}
