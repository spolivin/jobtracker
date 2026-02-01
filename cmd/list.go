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
	"github.com/spolivin/jobtracker/v2/internal/display"
)

var sortBy string
var descending bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all job applications",
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
			return fmt.Errorf("List cannot proceed: table 'applications' does not exist.")
		}

		store := db.NewJobApplicationStore(dbase)
		rows, err := store.Read(ctx, sortBy, descending)
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
