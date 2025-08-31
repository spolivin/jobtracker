/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var sortBy string
var descending bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all job applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := jobs.ListJobApplications(sortBy, descending); err != nil {
			return fmt.Errorf("error listing job applications: %w", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&sortBy, "sort", "s", "", "Sort job applications by field (company, position, status, applied_on)")
	listCmd.Flags().BoolVarP(&descending, "desc", "d", false, "Sort in descending order")
}
