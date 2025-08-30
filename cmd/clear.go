/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spolivin/jobtracker/jobs"
)

var force bool

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all job applications",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Not forcing deletion
		if !force {
			// Asking for confirmation to clear all job applications
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Are you sure you want to delete all job applications? (y/N): ")
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))

			// Cancelling if the answer is not yes
			if answer != "y" && answer != "yes" {
				fmt.Println("Operation cancelled.")
				return nil
			}
		}

		return jobs.ClearAllJobApplications()
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	clearCmd.Flags().BoolVarP(&force, "force", "f", false, "Skip confirmation")
}
