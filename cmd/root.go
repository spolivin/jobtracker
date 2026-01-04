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
)

var rootCmd = &cobra.Command{
	Use:   "jobtracker",
	Short: "Job tracker CLI for tracking job applications",
}

func init() {
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		// Skipping DB init for version or root (no subcommand)
		if cmd == rootCmd || cmd.Name() == "version" {
			return nil
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		return db.InitDatabase(ctx)
	}
}

func Execute() {
	err := rootCmd.Execute()
	cmd, _, e := rootCmd.Find(os.Args[1:])
	// Skip closing DB if no error and command is root or version
	skipDB := (e == nil && (cmd == rootCmd || cmd.Name() == "version"))
	var cerr error
	if !skipDB {
		cerr = db.CloseDatabaseConnection()
	}
	if err != nil || cerr != nil {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if cerr != nil {
			fmt.Fprintf(os.Stderr, "Error closing DB: %v\n", cerr)
		}
		os.Exit(1)
	}
}
