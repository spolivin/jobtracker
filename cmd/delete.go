/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/db"
	"github.com/spolivin/jobtracker/v2/internal/db/config"
)

var deleteId int

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a job application by ID",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("Config file not found. Run `jobtracker configure` first")
		}

		password, err := config.GetPassword()
		if err != nil {
			return err
		}
		ctx := cmd.Context()

		dbase, err := db.Connect(ctx, cfg, password)
		if err != nil {
			return err
		}
		defer dbase.Close()
		// Check if 'applications' table exists in Postgres
		tableExists, err := db.CheckTableExists(ctx, dbase, "applications")
		if err != nil {
			return err
		}
		if !tableExists {
			return fmt.Errorf("Delete cannot proceed: table 'applications' does not exist")
		}
		// Delete the job application from the database
		store := db.NewJobApplicationStore(dbase)
		rowsAffected, err := store.Delete(ctx, deleteId)
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
