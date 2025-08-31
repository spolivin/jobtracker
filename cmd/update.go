/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a job application by its ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Making sure that one ID is passed
		if args_len := len(args); args_len == 0 {
			return fmt.Errorf("job ID is required")
		} else if args_len > 1 {
			return fmt.Errorf("too many arguments")
		}
		// Updating the job application
		id := args[0]
		if err := jobs.UpdateJobApplication(id, company, position, status, applied_on); err != nil {
			return fmt.Errorf("error updating job application: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&company, "company", "c", "", "Company name")
	updateCmd.Flags().StringVarP(&position, "position", "p", "", "Job position")
	updateCmd.Flags().StringVarP(&status, "status", "s", "", "Job status")
	updateCmd.Flags().StringVarP(&applied_on, "applied_on", "a", "", "Date applied")
}
