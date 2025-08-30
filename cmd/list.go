/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all job applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		return jobs.ListJobApplications()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
