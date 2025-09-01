/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/version"
)

var versionFlag bool

var rootCmd = &cobra.Command{
	Use:   "jobtracker",
	Short: "Job tracker CLI for tracking job applications",
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			fmt.Println("Job Tracker CLI", version.Version)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Show version information")
}
