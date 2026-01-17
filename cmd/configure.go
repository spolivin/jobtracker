/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/db/config"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure database connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		ask := func(label, def string) string {
			fmt.Printf("%s [%s]: ", label, def)
			v, _ := reader.ReadString('\n')
			v = strings.TrimSpace(v)
			if v == "" {
				return def
			}
			return v
		}

		portStr := ask("Postgres port", "5432")
		port, _ := strconv.Atoi(portStr)

		cfg := &config.ConnectionConfig{
			DBHost: ask("Postgres host", "localhost"),
			DBPort: port,
			DBUser: ask("Postgres user", "postgres"),
			DBName: ask("Database name", "postgres"),
		}

		path, err := config.SaveConfig(cfg)
		if err != nil {
			return err
		}

		cmd.Printf("Configuration saved to %s\n\n", path)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
