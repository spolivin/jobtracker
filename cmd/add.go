/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var company string
var position string
var status string
var applied_on string

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new job application",
	RunE: func(cmd *cobra.Command, args []string) error {
		return jobs.AddJobApplication(company, position, status, applied_on)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&company, "company", "c", "", "Company name")
	addCmd.Flags().StringVarP(&position, "position", "p", "", "Job position")
	addCmd.Flags().StringVarP(&status, "status", "s", "", "Job status")
	addCmd.Flags().StringVarP(&applied_on, "applied_on", "a", "", "Date applied")

	addCmd.MarkFlagRequired("company")
	addCmd.MarkFlagRequired("position")
}
