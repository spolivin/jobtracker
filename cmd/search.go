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

var keyword string

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search job applications by keyword",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Check if 'applications' table exists in Postgres
		tableExists, err := db.CheckTableExists(ctx, "applications")
		if err != nil {
			return err
		}
		if !tableExists {
			return fmt.Errorf("Search cannot proceed: table 'applications' does not exist.")
		}
		// Search job applications in the database
		rows, err := db.SearchJobApplications(ctx, keyword)
		if err != nil {
			return err
		}
		if len(rows) == 0 {
			fmt.Fprintln(os.Stderr, "No data found matching the keyword.")
			return nil
		}
		return display.RenderTable(rows)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringVarP(&keyword, "keyword", "k", "", "Keyword to search for")
	searchCmd.MarkFlagRequired("keyword")
}
