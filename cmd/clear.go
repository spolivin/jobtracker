/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/internal/db"
)

var force bool

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all job applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Check if 'applications' table exists in Postgres
		tableExists, err := db.CheckTableExists(ctx, "applications")
		if err != nil {
			return err
		}
		if !tableExists {
			return fmt.Errorf("Drop cannot proceed: table 'applications' does not exist")
		}
		// Prompt user for confirmation
		if !force {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Are you sure you want to delete all job applications? (y/N): ")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))
			if answer != "y" && answer != "yes" {
				fmt.Fprintln(os.Stderr, "Drop operation cancelled.")
				return nil
			}
		}
		// Drop the 'applications' table
		if err = db.DropJobApplicationsTable(ctx); err != nil {
			return err
		}
		cmd.Println("Table 'applications' dropped successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation")
}
