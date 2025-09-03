/*
Copyright © 2025 Sergey Polivin <s.polivin@gmail.com>
*/
package jobs

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// headerColumns defines the header columns for the job applications table
var headerColumns = []string{"ID", "Company", "Position", "Status", "AppliedOn"}

// renderTable renders job applications in a table format
func renderTable(results []JobApplication) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(headerColumns)

	for _, app := range results {
		if err := table.Append(app.convertToStringSlice()); err != nil {
			return fmt.Errorf("failed to append row: %w", err)
		}
	}

	if err := table.Render(); err != nil {
		return fmt.Errorf("failed to render table: %w", err)
	}
	return nil
}
