/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import "context"

// AddJobApplication adds a new job application to the database.
func AddJobApplication(ctx context.Context, company, position, status string) error {
	query := `INSERT INTO applications (company, position, status) VALUES ($1, $2, $3)`
	if _, err := DB.ExecContext(ctx, query, company, position, status); err != nil {
		return err
	}
	return nil
}

// ReadJobApplications retrieves all job applications from the database with possible sorting by a specified field.
func ReadJobApplications(ctx context.Context, sortBy string, descending bool) ([]JobApplication, error) {
	query := `SELECT id, company, position, status, created_at, updated_at FROM applications`
	if sortBy != "" {
		query += ` ORDER BY ` + sortBy
		if descending {
			query += ` DESC`
		}
	}
	rows, err := DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []JobApplication
	for rows.Next() {
		var app JobApplication
		if err := rows.Scan(&app.ID, &app.Company, &app.Position, &app.Status, &app.CreatedAt, &app.UpdatedAt); err != nil {
			return nil, err
		}
		applications = append(applications, app)
	}
	return applications, rows.Err()
}

// UpdateJobApplicationStatus updates the status of a job application.
func UpdateJobApplicationStatus(ctx context.Context, id int, status string) (int64, error) {
	query := `UPDATE applications SET status=$1, updated_at=CURRENT_TIMESTAMP WHERE id=$2`
	res, err := DB.ExecContext(ctx, query, status, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// DeleteJobApplication deletes a job application from the database.
func DeleteJobApplication(ctx context.Context, id int) (int64, error) {
	query := `DELETE FROM applications WHERE id=$1`
	res, err := DB.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// DropJobApplicationsTable drops the applications table from the database.
func DropJobApplicationsTable(ctx context.Context) error {
	query := `DROP TABLE applications`
	_, err := DB.ExecContext(ctx, query)
	return err
}

// SearchJobApplications searches for job applications matching the given keyword in company, position, or status.
func SearchJobApplications(ctx context.Context, keyword string) ([]JobApplication, error) {
	query := `SELECT id, company, position, status, created_at, updated_at FROM applications WHERE company ILIKE $1 OR position ILIKE $1 OR status ILIKE $1`
	rows, err := DB.QueryContext(ctx, query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var applications []JobApplication
	for rows.Next() {
		var app JobApplication
		if err := rows.Scan(&app.ID, &app.Company, &app.Position, &app.Status, &app.CreatedAt, &app.UpdatedAt); err != nil {
			return nil, err
		}
		applications = append(applications, app)
	}
	return applications, rows.Err()
}
