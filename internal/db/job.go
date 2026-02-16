/*
Copyright © 2026 Sergey Polivin <s.polivin@gmail.com>
*/
package db

import (
	"context"
	"database/sql"
	"strconv"
)

// Wrapper around SQL-connection.
type JobApplicationsStore struct {
	db *sql.DB
}

// Constructor for JobApplicationsStore.
func NewJobApplicationStore(db *sql.DB) *JobApplicationsStore {
	return &JobApplicationsStore{db: db}
}

// Add adds a new job application to the database.
func (s *JobApplicationsStore) Add(ctx context.Context, company, position, status string) error {
	query := `INSERT INTO applications (company, position, status) VALUES ($1, $2, $3)`
	if _, err := s.db.ExecContext(ctx, query, company, position, status); err != nil {
		return err
	}
	return nil
}

// Read retrieves all job applications from the database with possible sorting by a specified field.
func (s *JobApplicationsStore) Read(ctx context.Context, sortBy string, descending bool) ([]JobApplication, error) {
	query := `SELECT * FROM applications`
	if sortBy != "" {
		// Validate column name to prevent SQL injection
		if err := ValidateColumnName(sortBy); err != nil {
			return nil, err
		}
		query += ` ORDER BY ` + sortBy
		if descending {
			query += ` DESC`
		}
	}
	rows, err := s.db.QueryContext(ctx, query)
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

// Update updates fields of a job application. Only provided fields are updated.
func (s *JobApplicationsStore) Update(ctx context.Context, id int, fields map[string]string) (int64, error) {
       if len(fields) == 0 {
	       return 0, nil
       }

       // Validate all column names to prevent SQL injection
       fieldNames := make([]string, 0, len(fields))
       for k := range fields {
	       fieldNames = append(fieldNames, k)
       }
       if err := ValidateColumnNames(fieldNames); err != nil {
	       return 0, err
       }

       setClause := ""
       args := make([]any, 0, len(fields)+1)
       i := 1
       for k, v := range fields {
	       if setClause != "" {
		       setClause += ", "
	       }
	       setClause += k + "=$" + strconv.Itoa(i)
	       args = append(args, v)
	       i++
       }

       setClause += ", updated_at=CURRENT_TIMESTAMP"
       query := "UPDATE applications SET " + setClause + " WHERE id=$" + strconv.Itoa(i)
       args = append(args, id)
       res, err := s.db.ExecContext(ctx, query, args...)
       if err != nil {
	       return 0, err
       }
       rowsAffected, err := res.RowsAffected()
       if err != nil {
	       return 0, err
       }
       return rowsAffected, nil
}

// Delete deletes a job application from the database.
func (s *JobApplicationsStore) Delete(ctx context.Context, id int) (int64, error) {
	query := `DELETE FROM applications WHERE id=$1`
	res, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// Clear clears all job applications from the table and resets the ID sequence.
func (s *JobApplicationsStore) Clear(ctx context.Context) error {
	query := `TRUNCATE TABLE applications RESTART IDENTITY`
	_, err := s.db.ExecContext(ctx, query)
	return err
}

// Search searches for job applications matching the given keyword in company, position, or status.
func (s *JobApplicationsStore) Search(ctx context.Context, keyword string) ([]JobApplication, error) {
	query := `SELECT id, company, position, status, created_at, updated_at FROM applications WHERE company ILIKE $1 OR position ILIKE $1 OR status ILIKE $1`
	rows, err := s.db.QueryContext(ctx, query, "%"+keyword+"%")
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
