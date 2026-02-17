# JobTracker CLI

> A PostgreSQL-backed command-line application for managing job applications with precision and efficiency.

[![Go Version](https://img.shields.io/badge/Go-1.24.5+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Latest Release](https://img.shields.io/github/v/release/spolivin/jobtracker)](https://github.com/spolivin/jobtracker/releases)

## Overview

JobTracker is a professional-grade CLI tool designed to streamline the job application tracking process. Built with Go and PostgreSQL, it provides a robust, database-backed solution for managing application workflows, eliminating the chaos of spreadsheet-based tracking.

## Key features

### Core functionality

- **Application Management** - Complete CRUD operations for job applications
- **Advanced Search** - Keyword-based queries across company names, positions, and statuses
- **Flexible Sorting** - Sort by any column in ascending or descending order
- **Data Export** - Export application data to CSV or JSON formats
- **Clean Interface** - Formatted tabular display with automatic timestamp tracking

### Technical highlights

- PostgreSQL backend for reliable data persistence
- SQL injection protection with input validation
- Automated database migrations
- Streamlined configuration management
- Docker Compose support for development
- Cross-platform compatibility

## Installation

### Prerequisites

- Go 1.24.5 or higher
- PostgreSQL database (local or Docker)

### Quick install

Install directly from GitHub:

```bash
go install github.com/spolivin/jobtracker/v2@latest
```

Verify installation:

```bash
jobtracker version
```

> Expected output: `JobTracker version v2.4.0`

## Getting started

### Initial Setup

1. **Configure database connection**

Run the configuration wizard:

```bash
jobtracker configure
```

You'll be prompted to enter:

- Database host (default: localhost)
- Database port (default: 5432)
- Database name (default: postgres)
- Database user (default: postgres)

Configuration is saved to your system's default config directory:

- Linux/macOS: `~/.config/jobtracker/config.json`
- Windows: `%APPDATA%\jobtracker\config.json`

**View current configuration:**

```bash
jobtracker config
```

Example output:

```
host=localhost
port=6432
user=appuser
dbname=appdb
```

**SECURITY NOTE:** Database passwords are not stored in configuration files. You'll be prompted to enter your password when executing commands. Alternatively, to avoid entering password each time, you can set password via environmental variables on Linux:

```bash
export DB_PASS=secret
```

or on Windows:

```powershell
set DB_PASS=secret
```

> Setting environmental variables in such a way is only valid for the current shell. After closing session, one will have to set the password again.

**Updating configuration:**

To change configuration, run `jobtracker configure` again. This will overwrite existing settings.

**Manual configuration:**

- _Linux/macOS_

```bash
nano ~/.config/jobtracker/config.json
```

- _Windows_

```powershell
notepad %APPDATA%\jobtracker\config.json
```

Configuration file format:

```json
{
  "db_host": "localhost",
  "db_port": 6432,
  "db_user": "appuser",
  "db_name": "appdb"
}
```

2. **Run database migrations**

```bash
jobtracker migrate
```

This command:

- Creates the applications table if it doesn't exist
- Applies any pending schema updates
- Can be run safely multiple times (idempotent)

### Setting Up PostgreSQL

#### Option 1: Docker (Recommended for Development)

Clone the repository and start the database:

```bash
git clone https://github.com/spolivin/jobtracker.git
cd jobtracker
make start-db
```

Stop the database when finished:

```bash
make stop-db
```

#### Option 2: Local PostgreSQL Installation

Ensure PostgreSQL is installed and running on your system. Configure connection details using `jobtracker configure`.

## Usage

### Command reference

| Command     | Description                                 |
| ----------- | ------------------------------------------- |
| `add`       | Create a new job application entry          |
| `list`      | Display all applications in tabular format  |
| `update`    | Modify an existing application              |
| `search`    | Find applications by keyword                |
| `delete`    | Remove a specific application by ID         |
| `clear`     | Delete all applications (with confirmation) |
| `export`    | Export data to CSV or JSON                  |
| `configure` | Set up database connection                  |
| `config`    | Display current database configuration      |
| `migrate`   | Execute database migrations                 |
| `version`   | Display CLI version information             |

### Common workflows

---

#### Adding applications

**Full specification:**

```bash
jobtracker add --company "OpenAI" --position "ML Engineer" --status "Applied"
```

**Quick add** (status defaults to "Applied"):

```bash
jobtracker add -c "Google" -p "Software Engineer"
```

---

#### Viewing applications

**List all:**

```bash
jobtracker list
```

**Sorted view:**

```bash
jobtracker list --sort status --desc
```

Example output:

```
┌────┬──────────┬───────────────────────────┬───────────┬───────────────────────────┬───────────────────────────┐
│ ID │ COMPANY  │         POSITION          │  STATUS   │        CREATED AT         │        UPDATED AT         │
├────┼──────────┼───────────────────────────┼───────────┼───────────────────────────┼───────────────────────────┤
│ 5  │ Facebook │ Software Engineer         │ Offer     │ 2026-01-18T18:55:53+01:00 │ 2026-01-18T18:57:00+01:00 │
│ 7  │ Apple    │ Machine Learning Engineer │ Interview │ 2026-01-18T18:56:29+01:00 │ 2026-01-18T18:57:09+01:00 │
│ 6  │ Google   │ Data Scientist            │ Applied   │ 2026-01-18T18:56:11+01:00 │ 2026-01-18T18:56:11+01:00 │
└────┴──────────┴───────────────────────────┴───────────┴───────────────────────────┴───────────────────────────┘
```

---

#### Searching applications

Search across company, position, and status fields:

```bash
jobtracker search --keyword "Engineer"
```

#### Updating applications

Update application status:

```bash
jobtracker update --id 3 --status "Interview"
```

or company and position names:

```bash
jobtracker update --id 3 --company "Tesla" --position "MLOps Engineer"
```

---

#### Deleting applications

Single deletion:

```bash
jobtracker delete --id 3
```

Clear all (with confirmation prompt):

```bash
jobtracker clear
```

Force clear (skip confirmation):

```bash
jobtracker clear --force
```

The `clear` command will clear all job applications from the SQL-table and reset the ID counter.

---

#### Exporting data

Export to JSON:

```bash
jobtracker export --format json
```

Export to CSV:

```bash
jobtracker export --format csv
```

Output files: `exported_data.json` or `exported_data.csv`. Alternatively, one can set filenames for exported files:

```bash
jobtracker export --format json --output applications
jobtracker export --format csv --output applications
```

Output files: `applications.json` and `applications.csv`.

## Data schema

Applications are stored in the `applications` table with the following structure:

| Field        | Type      | Description                       |
| ------------ | --------- | --------------------------------- |
| `id`         | Integer   | Auto-incremented primary key      |
| `company`    | String    | Company name                      |
| `position`   | String    | Job title/role                    |
| `status`     | String    | Application status                |
| `created_at` | Timestamp | Record creation time (ISO 8601)   |
| `updated_at` | Timestamp | Last modification time (ISO 8601) |

## Development

### Building from source

```bash
git clone https://github.com/spolivin/jobtracker.git
cd jobtracker
make build
```

### Development commands

```bash
# Start PostgreSQL in Docker
make start-db

# Stop PostgreSQL
make stop-db

# Populate database with sample data
make populate-db
```

## Running Tests

JobTracker includes a comprehensive test suite covering unit tests, security tests, and performance benchmarks.

### Test Coverage

The project includes tests for:

- **Data Models** (`internal/db/models_test.go`) - JobApplication struct and conversion methods
- **Security Validation** (`internal/db/validator_test.go`) - SQL injection protection and column name validation
- **SQL Operations** (`internal/db/job_test.go`) - Database CRUD operations with security validation
- **Version Management** (`internal/version/version_test.go`) - Semantic versioning compliance

### Running Tests

**Run all tests:**

```bash
go test ./...
```

**Run tests with verbose output:**

```bash
go test -v ./...
```

**Run tests for a specific package:**

```bash
go test ./internal/db
go test ./internal/version
```

**Run specific test functions:**

```bash
go test -run TestValidateColumnName ./internal/db
go test -run TestReadSQLInjectionProtection ./internal/db
```

### Code Coverage

**Generate coverage report:**

```bash
go test -cover ./...
```

**Generate detailed coverage profile:**

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

This opens an HTML coverage report in your browser, highlighting tested and untested code.

### Benchmark Tests

**Run all benchmarks:**

```bash
go test -bench=. ./...
```

**Run specific benchmarks:**

```bash
go test -bench=BenchmarkValidateColumnName ./internal/db
go test -bench=BenchmarkConvertToStringSlice ./internal/db
```

**Benchmarks with memory statistics:**

```bash
go test -bench=. -benchmem ./...
```

### Race Detection

**Run tests with race detector** (recommended before commits):

```bash
go test -race ./...
```

This helps identify potential concurrency issues.

### Test Categories

- **Unit Tests** - Test individual functions and methods in isolation
- **Security Tests** - Validate SQL injection protection and input validation
- **Benchmark Tests** - Measure performance of critical operations
- **Integration Tests** - Validate database operations (requires nil db mock for current tests)

## Technology Stack

- **Language:** Go 1.24.5+
- **Database:** PostgreSQL
- **CLI Framework:** [Cobra](https://github.com/spf13/cobra)
- **Table Formatting:** [tablewriter](https://github.com/olekukonko/tablewriter)
- **Containerization:** Docker, Docker Compose

## Project Structure

```
jobtracker/
├── cmd/                  # CLI command implementations
├── internal/             # Internal packages
│   ├── db/               # Database management (connection, migrations, CRUD, data models)
│   ├── display/          # Data display
|   ├── exporter/         # Data export to JSON or CSV
│   └── version/          # CLI version tracking
├── docker-compose.yml    # PostgreSQL container definition
├── Makefile              # Development automation
└── main.go               # Application entry point
```

## Version history

See [RELEASES](https://github.com/spolivin/jobtracker/releases) for detailed changelog.

Current Version: v2.4.0 (Latest):

- Added SQL injection protection to Read and Update operations
- Comprehensive test coverage for security vulnerabilities
- Performance benchmarks for validation operations
- Enhanced security with column name validation

Previous Version (v2.3.0):

- Simplified configuration workflow
- Enhanced user experience
- Improved applications updating logic
- Simplified password recognition for commands

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Support

For bugs, feature requests, or questions:

- Open an [issue](https://github.com/spolivin/jobtracker/issues)
- Check existing [releases](https://github.com/spolivin/jobtracker/releases)
