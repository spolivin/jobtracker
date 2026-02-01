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

var updateId int
var updateCompany string
var updatePosition string
var updateStatus string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
       Use:   "update",
       Short: "Update fields of a job application",
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
		       return fmt.Errorf("Update cannot proceed: table 'applications' does not exist.")
	       }
	       // Build fields to update
	       fields := make(map[string]string)
	       if updateCompany != "" {
		       fields["company"] = updateCompany
	       }
	       if updatePosition != "" {
		       fields["position"] = updatePosition
	       }
	       if updateStatus != "" {
		       fields["status"] = updateStatus
	       }
	       if len(fields) == 0 {
		       return fmt.Errorf("No fields specified to update. Use --company, --position, or --status.")
	       }
	       // Update the job application in the database
	       store := db.NewJobApplicationStore(dbase)
	       rowsAffected, err := store.Update(ctx, updateId, fields)
	       if err != nil {
		       return err
	       }
	       if rowsAffected == 0 {
		       fmt.Fprintln(os.Stderr, "No job application found with the specified ID. No update performed.")
		       return nil
	       }
	       cmd.Println("Job application updated successfully")
	       return nil
       },
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVarP(&updateId, "id", "i", 0, "Job application ID")
	updateCmd.Flags().StringVarP(&updateCompany, "company", "c", "", "Job company")
	updateCmd.Flags().StringVarP(&updatePosition, "position", "p", "", "Job position")
	updateCmd.Flags().StringVarP(&updateStatus, "status", "s", "", "Job status")

	updateCmd.MarkFlagRequired("id")
}
