# JobTracker CLI

> A PostgreSQL-backed command-line application for managing job applications with precision and efficiency.

[![Go Version](https://img.shields.io/badge/Go-1.24.5+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Latest Release](https://img.shields.io/github/v/release/spolivin/jobtracker)](https://github.com/spolivin/jobtracker/releases)

## Overview

JobTracker is a professional-grade CLI tool designed to streamline the job application tracking process. Built with Go and PostgreSQL, it provides a robust, database-backed solution for managing application workflows, eliminating the chaos of spreadsheet-based tracking.

## Key features

### Core functionality

* **Application Management** - Complete CRUD operations for job applications
* **Advanced Search** - Keyword-based queries across company names, positions, and statuses
* **Flexible Sorting** - Sort by any column in ascending or descending order
* **Data Export** - Export application data to CSV or JSON formats
* **Clean Interface** - Formatted tabular display with automatic timestamp tracking

### Technical highlights

* PostgreSQL backend for reliable data persistence
* Automated database migrations
* Streamlined configuration management
* Docker Compose support for development
* Cross-platform compatibility

## Installation

### Prerequisistes

* Go 1.24.5 or higher
* PostgreSQL database (local or Docker)

### Quick install

Install directly from GitHub:

```bash
go install github.com/spolivin/jobtracker/v2@latest
```

Verify installation:

```bash
jobtracker version
```
> Expected output: `JobTracker version v2.2.0`

## Getting started

### Initial Setup

1. **Configure database connection**

Run the configuration wizard:

```bash
jobtracker configure
```

You'll be prompted to enter:

* Database host (default: localhost)
* Database port (default: 5432)
* Database name (default: postgres)
* Database user (default: postgres)

Configuration is saved to your system's default config directory:

* Linux/macOS: `~/.config/jobtracker/config.json`
* Windows: `%APPDATA%\jobtracker\config.json`

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
> **Security Note:** Database passwords are not stored in configuration files. You'll be prompted to enter your password when executing commands.

**Updating configuration:**

To change configuration, run `jobtracker configure` again. This will overwrite existing settings.

**Manual configuration:**

* *Linux/macOS*

```bash
nano ~/.config/jobtracker/config.json
```

* *Windows*

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

* Creates the applications table if it doesn't exist
* Applies any pending schema updates
* Can be run safely multiple times (idempotent)

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

| Command | Description |
|---|---|
| `add` | Create a new job application entry |
| `list` | Display all applications in tabular format |
| `update` | Modify an existing application |
| `search` | Find applications by keyword |
| `delete` | Remove a specific application by ID |
| `clear` | Delete all applications (with confirmation) |
| `export` | Export data to CSV or JSON |
| `configure` | Set up database connection |
| `config` | Display current database configuration |
| `migrate` | Execute database migrations |
| `version` | Display CLI version information |

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

---

#### Exporting data

Export to JSON (default):

```bash
jobtracker export
```

Export to CSV:

```bash
jobtracker export --format csv
```

Output files: `exported_data.json` or `exported_data.csv`.

## Data schema

Applications are stored in the `applications` table with the following structure:

| Field | Type | Description | 
|---|---|---|
| `id` | Integer |Auto-incremented primary key |
| `company` | String | Company name|
| `position` | String | Job title/role |
| `status` | String | Application status |
| `created_at` | Timestamp | Record creation time (ISO 8601) |
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

Current Version: v2.2.0 (Latest):

* Simplified configuration workflow
* Enhanced user experience

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Support

For bugs, feature requests, or questions:

* Open an [issue](https://github.com/spolivin/jobtracker/issues)
* Check existing [releases](https://github.com/spolivin/jobtracker/releases)
