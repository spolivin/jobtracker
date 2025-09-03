/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var id string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for job applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := jobs.SearchJobApplications(id, company, position, status, applied_on); err != nil {
			return fmt.Errorf("failed to search job applications: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&id, "id", "i", "", "Job ID")
	searchCmd.Flags().StringVarP(&company, "company", "c", "", "Company name")
	searchCmd.Flags().StringVarP(&position, "position", "p", "", "Job position")
	searchCmd.Flags().StringVarP(&status, "status", "s", "", "Job status")
	searchCmd.Flags().StringVarP(&applied_on, "applied_on", "a", "", "Date applied")
}
