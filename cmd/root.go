/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "jobtracker",
	Short: "Job tracker CLI for tracking job applications",
}

func Execute() {
	rootCmd.Execute()
}
