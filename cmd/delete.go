/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a job application by its ID",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Making sure that one ID is passed
		if args_len := len(args); args_len == 0 {
			return fmt.Errorf("job ID is required")
		} else if args_len > 1 {
			return fmt.Errorf("too many arguments")
		}
		if err := jobs.DeleteJobApplication(args[0]); err != nil {
			return fmt.Errorf("error deleting job application: %w", err)
		}
		fmt.Printf("Deleted job application with id %s\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
