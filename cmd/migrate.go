/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/db"
	"github.com/spolivin/jobtracker/v2/internal/db/config"
	"github.com/spolivin/jobtracker/v2/internal/db/migrate"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
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

		if err := migrate.Run(ctx, dbase); err != nil {
			return err
		}
		cmd.Println("Migrations run successfully")
		return nil

	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
