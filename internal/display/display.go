/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package display

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spolivin/jobtracker/internal/db"
)

// RenderTable renders the job applications data in a table format
func RenderTable(data []db.JobApplication) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"ID", "Company", "Position", "Status", "Created At", "Updated At"})
	for _, row := range data {
		row := row.ConvertToStringSlice()
		table.Append(row)
	}
	return table.Render()
}
