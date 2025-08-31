/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export job applications to a CSV file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := jobs.ExportToCsv(); err != nil {
			return fmt.Errorf("error exporting job applications to CSV: %w", err)
		}
		fmt.Println("Job applications exported to CSV successfully")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
