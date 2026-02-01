/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/db"
	"github.com/spolivin/jobtracker/v2/internal/db/config"
)

var force bool

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all job applications",
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
			return fmt.Errorf("Clear cannot proceed: table 'applications' does not exist.")
		}
		// Prompt user for confirmation
		if !force {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Are you sure you want to delete all job applications? (y/N): ")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))
			if answer != "y" && answer != "yes" {
				fmt.Fprintln(os.Stderr, "Clear operation cancelled.")
				return nil
			}
		}
		// Clearing all job applications
		store := db.NewJobApplicationStore(dbase)
		rows, err := store.Read(ctx, sortBy, descending)
		if err != nil {
			return err
		}
		if len(rows) == 0 {
			fmt.Fprintln(os.Stderr, "Table is empty: no job applications found in the database.")
			return nil
		}
		if err = store.Clear(ctx); err != nil {
			return err
		}
		cmd.Println("All job applications have been cleared successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation")
}
