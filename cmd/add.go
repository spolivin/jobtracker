/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/db"
	"github.com/spolivin/jobtracker/v2/internal/db/config"
)

var company string
var position string
var status string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new job application",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("Config file not found. Run `jobtracker configure` first")
		}

		password, err := config.PromptPassword()
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
			return fmt.Errorf("Add cannot proceed: table 'applications' does not exist. Run `jobtracker migrate` to create one.")
		}

		// Add the job application to the database
		store := db.NewJobApplicationStore(dbase)
		if err := store.Add(ctx, company, position, status); err != nil {
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
