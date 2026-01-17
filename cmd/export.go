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
	"github.com/spolivin/jobtracker/v2/internal/exporter"
)

var exportFormat string

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export job applications to a specified format (e.g., CSV, JSON)",
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
			return fmt.Errorf("Export cannot proceed: table 'applications' does not exist.")
		}
		// Read all job applications from the database
		store := db.NewJobApplicationStore(dbase)
		rows, err := store.Read(ctx, "", false)
		if err != nil {
			return err
		}
		if len(rows) == 0 {
			fmt.Fprintln(os.Stderr, "Nothing to export: no job applications found in the database.")
			return nil
		}
		// Export data based on the specified format
		switch exportFormat {
		case "json":
			if err := exporter.ExportToJson(rows, "exported_data.json"); err != nil {
				return err
			}

		case "csv":
			if err := exporter.ExportToCsv(rows, "exported_data.csv"); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported export format: %s", exportFormat)
		}
		cmd.Println("Data exported successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	exportCmd.Flags().StringVarP(&exportFormat, "format", "f", "json", "Export format (json or csv)")

}
