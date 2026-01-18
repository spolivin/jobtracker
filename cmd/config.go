/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/db/config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show current connection config",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.LoadConfig()
		if err != nil {
			return fmt.Errorf("Config file not found. Run `jobtracker configure` first")
		}
		configInfo := fmt.Sprintf("host=%s\nport=%s\nuser=%s\ndbname=%s",
			cfg.DBHost,
			strconv.Itoa(cfg.DBPort),
			cfg.DBUser,
			cfg.DBName,
		)
		cmd.Println(configInfo)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
