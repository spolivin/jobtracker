/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package exporter

import (
	"encoding/csv"
	"encoding/json"
	"os"

	"github.com/spolivin/jobtracker/v2/internal/db"
)

// ExportToJson exports job application data to a JSON file.
func ExportToJson(data []db.JobApplication, filename string) error {
	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	if err = os.WriteFile(filename, encodedData, 0644); err != nil {
		return err
	}
	return nil
}

// ExportToCsv exports job application data to a CSV file.
func ExportToCsv(data []db.JobApplication, filename string) error {
	// Creating CSV file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// Creating a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing header
	var headerColumns = []string{"ID", "Company", "Position", "Status", "CreatedAt", "UpdatedAt"}
	if err := writer.Write(headerColumns); err != nil {
		return err
	}

	// Writing job entries
	for _, row := range data {
		row := row.ConvertToStringSlice()
		if err := writer.Write(row); err != nil {
			return err
		}
	}
	return nil
}
