/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"context"
	"os"
	"time"

	"fmt"

	"github.com/spolivin/jobtracker/internal/db"

	"github.com/spf13/cobra"
)

var deleteId int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a job application by ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Check if 'applications' table exists in Postgres
		tableExists, err := db.CheckTableExists(ctx, "applications")
		if err != nil {
			return err
		}
		if !tableExists {
			return fmt.Errorf("Delete cannot proceed: table 'applications' does not exist")
		}
		// Delete the job application from the database
		rowsAffected, err := db.DeleteJobApplication(ctx, deleteId)
		if err != nil {
			return err
		}
		if rowsAffected == 0 {
			fmt.Fprintln(os.Stderr, "No job application found with the specified ID. No delete performed.")
			return nil
		}
		cmd.Println(fmt.Sprintf("Job application with ID %d deleted successfully", deleteId))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&deleteId, "id", "i", 0, "Job application ID to delete")
	deleteCmd.MarkFlagRequired("id")
}
