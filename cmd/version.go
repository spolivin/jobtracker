/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/v2/internal/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display program version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("JobTracker version %s\n", version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
