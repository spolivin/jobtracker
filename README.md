# JobTracker CLI

A simple command-line tool written in Go to help you keep track of your job applications.  

You can add, update, sort, search, list, delete and clear job application records stored in a Postgres database.

## ✨ Features

- **Add** new job applications with company, position and status.
- **List** all saved job applications in a clean tabular format.
- **Update** existing job applications.
- **Search** job applications.
- **Delete** specific applications by ID.
- **Clear** all stored applications.
- **Sort** job applications.
- **Export** saved job applications to CSV and JSON.
- Starting from *v2* CLI interacts with **PostgreSQL** database for more convenient data storing

## 🚀 Installation

Make sure you have Go installed (1.24.5+). You can install the tool directly from the repository:

```bash
go install github.com/spolivin/jobtracker/v2@latest
```

After installation, you can make sure that the latest version is installed (v2.0.0) by running:

```bash
jobtracker version
```

## Usage 

General command structure:

```bash
jobtracker [command] [flags]
```

### *New* in **Version 2.1.0**

The newest version of the CLI (v2.1.0) enables avoiding having `.env` file set up in the same directory for the tool to work, thus removing the inconvenience. It is enough to run the command below to save Postgres connection parameters in a default config directory:

```bash
jobtracker configure
```

Afterwards we run migration command to make sure that all schemas are up-to-date:

```bash
jobtracker migrate
```
> Having run migration scripts enables being able to run primary commands without any problems. For security considerations, Postgres password is not stored in config set up by `jobtracker configure` and is required for running each command.

### Setting up Postgres database

In case of absence of locally running Postgres database, one can run one in *Docker*:

```bash
git clone https://github.com/spolivin/jobtracker.git
cd jobtracker
make start-db
```
> Database can be stopped by running `make stop-db`.

### Development mode

CLI can also be tested locally without standard installation by building and then running basic commands specified in *Makefile*:

```bash
make build
make populate-db
```
> Database should be up and running.

### Available commands

* `add` - Add a new job application.
* `clear` - Clear all job applications at once.
* `configure` - Configure database connection
* `delete` - Delete a specific job application by its ID.
* `export` - Export all job applications to a CSV or JSON file.
* `list` - List all saved job applications.
* `migrate` - Run database migrations.
* `search` - Search a job application based on a keyword.
* `update` - Update an existing job application.
* `version` - Display CLI version.

### Global flags

* `-h, --help` - Help for a command. 

## Examples

### Add a new job application

```bash
jobtracker add --company "OpenAI" --position "ML Engineer" --status "Applied"
```

Specifying `--status` is not strictly required, in case this flag is missing, it will be replaced with `"Applied"` for `status`:

```bash
jobtracker add -c "OpenAI" -p "ML Engineer"
```
> Running this command will create `applications` table in PostgreSQL database for storing the job applications history.

### List all job applications

```bash
jobtracker list
```

Running the above command will output the saved job applications in a convenient and easy-to-read format:

```
┌────┬───────────┬─────────────────────────────┬───────────┬───────────────────────────┬───────────────────────────┐
│ ID │  COMPANY  │          POSITION           │  STATUS   │        CREATED AT         │        UPDATED AT         │
├────┼───────────┼─────────────────────────────┼───────────┼───────────────────────────┼───────────────────────────┤ 
│ 1  │ Facebook  │ Software Engineer           │ Applied   │ 2026-01-04T11:56:50+01:00 │ 2026-01-04T11:56:50+01:00 │ 
│ 2  │ Google    │ Data Scientist              │ Interview │ 2026-01-04T11:56:55+01:00 │ 2026-01-04T11:56:55+01:00 │ 
│ 3  │ Apple     │ Machine Learning Engineer   │ Applied   │ 2026-01-04T11:57:00+01:00 │ 2026-01-04T11:57:00+01:00 │ 
│ 4  │ Microsoft │ Machine Learning Specialist │ Applied   │ 2026-01-04T11:57:09+01:00 │ 2026-01-04T11:57:09+01:00 │ 
│ 5  │ Huawei    │ Frontend Developer          │ Applied   │ 2026-01-04T11:57:16+01:00 │ 2026-01-04T11:57:16+01:00 │ 
│ 6  │ Luxoft    │ Backend Developer           │ Applied   │ 2026-01-04T11:57:22+01:00 │ 2026-01-04T11:57:22+01:00 │ 
│ 7  │ NCR       │ Devops                      │ Applied   │ 2026-01-04T11:57:27+01:00 │ 2026-01-04T11:57:27+01:00 │ 
└────┴───────────┴─────────────────────────────┴───────────┴───────────────────────────┴───────────────────────────┘
```

You can optionally sort the job applications by the above columns and display the result in the same convenient format in ascending or descending order (`--desc` flag):

```bash
jobtracker list --sort status --desc
```

### Search job applications

One can find the job applications by running the search query in this way:

```bash
jobtracker search --keyword Applied
```
> Matches for company name, position name and status are searched.

### Update an existing applications

If at some point we need to update the information on some applications, we can run:

```bash
jobtracker update --id 3 --status Interview
```
> This command will update entry with ID=3 and modify `status` to `"Interview"`.

### Delete an existing application

```bash
jobtracker delete --id 3
```
> This command will delete entry with ID=3 from the Postgres table.

### Clear all applications

```bash
jobtracker clear
```
> This command will firstly prompt for the user's confirmation and then delete all available applications. One can optionally set `--force` flag to skip prompting.

### Export all applications to CSV or JSON

```bash
jobtracker export
```
> This command will export all applications stored in `applications` table to `exported_data.json`, while setting flag `--format csv` will save database data to `exported_data.csv`.

## Data storage

All job applications are stored in `applications` table in the PostgreSQL database. Each entry includes:

* ID – Auto-incremented unique identifier

* Company – Company name

* Position – Job title/role

* Status – Application status (Applied, Interviewing, Offer, Rejected, etc.)

* CreatedAt – Stored in ISO 8601 format

* UpdatedAt – Stored in ISO 8601 format

## License

MIT License. See [LICENSE](./LICENSE) for details.
