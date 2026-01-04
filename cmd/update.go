/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/internal/db"
)

var updateId int
var newStatus string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the status of a job application",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Check if 'applications' table exists in Postgres
		tableExists, err := db.CheckTableExists(ctx, "applications")
		if err != nil {
			return err
		}
		if !tableExists {
			return fmt.Errorf("Update cannot proceed: table 'applications' does not exist.")
		}
		// Update the job application status in the database
		rowsAffected, err := db.UpdateJobApplicationStatus(ctx, updateId, newStatus)
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			fmt.Fprintln(os.Stderr, "No job application found with the specified ID. No update performed.")
			return nil
		}
		cmd.Println("Job application status updated successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateId, "id", "i", 0, "Job application ID")
	updateCmd.Flags().StringVarP(&newStatus, "status", "s", "", "Job status")

	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("status")
}
