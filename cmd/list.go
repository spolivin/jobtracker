/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/internal/db"
	"github.com/spolivin/jobtracker/internal/display"
)

var sortBy string
var descending bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all job applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Check if 'applications' table exists in Postgres
		tableExists, err := db.CheckTableExists(ctx, "applications")
		if err != nil {
			return err
		}
		if !tableExists {
			return fmt.Errorf("List cannot proceed: table 'applications' does not exist.")
		}

		rows, err := db.ReadJobApplications(ctx, sortBy, descending)
		if err != nil {
			return err
		}
		if len(rows) == 0 {
			fmt.Fprintln(os.Stderr, "Table is empty: no job applications found in the database.")
			return nil
		}
		return display.RenderTable(rows)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&sortBy, "sort", "s", "", "Sort job applications by field")
	listCmd.Flags().BoolVarP(&descending, "desc", "d", false, "Sort in descending order")
}
