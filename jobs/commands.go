/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package jobs

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/olekukonko/tablewriter"
	"github.com/spolivin/jobtracker/utils"
)

const dateLayout = "2006-01-02" // Date format for job applications
const csvFileName = "jobs.csv"  // Default CSV file name

// AddJobApplication adds a new job application
func AddJobApplication(company string, position string, status string, applied_on string) error {

	// Loading existing job applications
	jobApplicationsList, err := loadJobApplications()
	if err != nil {
		return fmt.Errorf("failed to load job applications: %w", err)
	}

	// Parsing the date
	appliedOnDate, err := utils.ParseDate(dateLayout, applied_on)
	if err != nil {
		return fmt.Errorf("failed to parse date of application: %w", err)
	}

	// Setting default status
	if status == "" {
		status = "Applied"
	}

	// Creating new job application
	newJobApplication := JobApplication{
		ID:        getNextID(jobApplicationsList),
		Company:   company,
		Position:  position,
		Status:    status,
		AppliedOn: appliedOnDate.Format(dateLayout),
	}
	// Adding a new job application
	jobApplicationsList = append(jobApplicationsList, newJobApplication)

	// Saving updated job applications list
	if err := saveJobApplications(jobApplicationsList); err != nil {
		return fmt.Errorf("error saving job application: %w", err)
	}

	return nil
}

// ListJobApplications lists all saved job applications (with sorting options)
func ListJobApplications(sortBy string, descending bool) error {

	// Loading existing job applications
	apps, err := loadJobApplications()
	if err != nil {
		return fmt.Errorf("failed to load job applications: %w", err)
	}
	// Checking if there are any job applications
	if len(apps) == 0 {
		return fmt.Errorf("no job applications found")
	}

	// Defining sorting functions
	sortFuncs := map[string]func(i, j int) bool{
		"company":    func(i, j int) bool { return apps[i].Company < apps[j].Company },
		"position":   func(i, j int) bool { return apps[i].Position < apps[j].Position },
		"status":     func(i, j int) bool { return apps[i].Status < apps[j].Status },
		"applied_on": func(i, j int) bool { return apps[i].AppliedOn < apps[j].AppliedOn },
	}

	// Sorting job applications if needed
	if sortBy != "" {
		sortFunc, present := sortFuncs[sortBy]
		if !present {
			return fmt.Errorf("invalid sort option: %s", sortBy)
		}
		if descending {
			sort.Slice(apps, func(i, j int) bool { return !sortFunc(i, j) })
		} else {
			sort.Slice(apps, sortFunc)
		}
	}

	// Displaying job applications
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Company", "Position", "Status", "AppliedOn"})

	for _, app := range apps {
		table.Append([]string{
			strconv.Itoa(app.ID),
			app.Company,
			app.Position,
			app.Status,
			app.AppliedOn,
		})
	}

	table.Render()

	return nil
}

// DeleteJobApplication deletes a job application by ID
func DeleteJobApplication(id string) error {

	// Converting ID from string to int
	id_num, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}

	// Loading existing job applications
	apps, err := loadJobApplications()
	if err != nil {
		return fmt.Errorf("failed to load job applications: %w", err)
	}
	// Checking if there any job applications
	if len(apps) == 0 {
		return fmt.Errorf("no job applications found")
	}

	// Iterating over available job applications and removing the needed one
	updatedJobApplications := make([]JobApplication, 0, len(apps))
	found := false
	for _, app := range apps {
		if app.ID != id_num {
			updatedJobApplications = append(updatedJobApplications, app)
		} else {
			found = true
		}
	}
	// Error if job application with ID is not found
	if !found {
		return fmt.Errorf("job application with id %d not found", id_num)
	}

	// Saving updated job applications
	if err := saveJobApplications(updatedJobApplications); err != nil {
		return fmt.Errorf("error saving job application: %w", err)
	}

	return nil

}

// ClearAllJobApplications clears all job applications
func ClearAllJobApplications() error {

	// Loading existing job applications
	apps, err := loadJobApplications()
	if err != nil {
		return fmt.Errorf("failed to load job applications: %w", err)
	}
	// Checking if there any job applications
	if len(apps) == 0 {
		return fmt.Errorf("no job applications found")
	}

	// Saving an empty list of job applications
	if err := saveJobApplications([]JobApplication{}); err != nil {
		return fmt.Errorf("error saving job applications: %w", err)
	}

	return nil
}

// UpdateJobApplication updates an existing job application
func UpdateJobApplication(id string, company string, position string, status string, applied_on string) error {
	// Converting ID from string to int
	id_num, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %w", err)
	}

	// Loading existing job applications
	apps, err := loadJobApplications()
	if err != nil {
		return fmt.Errorf("failed to load job applications: %w", err)
	}
	// Checking if there any job applications
	if len(apps) == 0 {
		return fmt.Errorf("no job applications found")
	}

	// Finding the job application to update
	var jobApplicationToUpdate *JobApplication
	for i := range apps {
		if apps[i].ID == id_num {
			jobApplicationToUpdate = &apps[i]
			break
		}
	}
	// Error job application is not found
	if jobApplicationToUpdate == nil {
		return fmt.Errorf("job application with id %d not found", id_num)
	}

	// Updating fields if new values are provided
	if company != "" {
		jobApplicationToUpdate.Company = company
	}
	if position != "" {
		jobApplicationToUpdate.Position = position
	}
	if status != "" {
		jobApplicationToUpdate.Status = status
	}
	if applied_on != "" {
		appliedOnDate, err := utils.ParseDate(dateLayout, applied_on)
		if err != nil {
			return fmt.Errorf("failed to parse date of application: %w", err)
		}
		jobApplicationToUpdate.AppliedOn = appliedOnDate.Format(dateLayout)
	}

	// Saving updated job applications
	if err := saveJobApplications(apps); err != nil {
		return fmt.Errorf("error saving job application: %w", err)
	}

	// Checking if any fields were updated
	if company == "" && position == "" && status == "" && applied_on == "" {
		fmt.Println("No changes made.")
	} else {
		fmt.Printf("Updated job application with id %d\n", id_num)
	}

	return nil
}

// ExportToCsv exports job applications to a CSV file
func ExportToCsv() error {

	// Loading existing job applications
	apps, err := loadJobApplications()
	if err != nil {
		return fmt.Errorf("failed to load job applications: %w", err)
	}
	// Checking if there any job applications
	if len(apps) == 0 {
		return fmt.Errorf("no job applications found")
	}

	// Creating CSV file
	file, err := os.Create(csvFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Creating a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Writing header
	if err := writer.Write([]string{"ID", "Company", "Position", "Status", "DateApplied"}); err != nil {
		return err
	}

	// Writing job entries
	for _, job := range apps {
		row := []string{
			strconv.Itoa(job.ID),
			job.Company,
			job.Position,
			job.Status,
			job.AppliedOn,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
