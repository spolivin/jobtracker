/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"context"
	"time"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/db"
)

var company string
var position string
var status string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new job application",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Ensure the database schema is set up
		if schemaerr := db.EnsureSchema(ctx); schemaerr != nil {
			return schemaerr
		}
		// Add the job application to the database
		if err := db.AddJobApplication(ctx, company, position, status); err != nil {
			return err
		}
		cmd.Println("Job application added successfully")
		return nil
	},
}

func init() {

	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&company, "company", "c", "", "Company name")
	addCmd.Flags().StringVarP(&position, "position", "p", "", "Job position")
	addCmd.Flags().StringVarP(&status, "status", "s", "Applied", "Job status")

	addCmd.MarkFlagRequired("company")
	addCmd.MarkFlagRequired("position")
}
