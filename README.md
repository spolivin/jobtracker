# JobTracker CLI

A simple command-line tool written in Go to help you keep track of your job applications.  

You can add, update, sort, list, delete and clear job application records stored in a local JSON file.  

## Features

- **Add** new job applications with company, position, status, and date.
- **List** all saved job applications in a clean tabular format.
- **Update** existing job applications.
- **Delete** specific applications by ID.
- **Clear** all stored applications.
- **Sort** job applications.
- **Export** saved job applications to CSV.
- Data stored locally in a JSON file for persistence.

## Installation

Make sure you have Go installed (1.18+ recommended). You can install the tool directly from the repository:

```bash
go install github.com/spolivin/jobtracker@latest
```

## Usage 

General command structure:

```bash
jobtracker [command] [flags]
```

### Available commands

* `add` - Add a new job application.
* `update` - Update an existing job application.
* `list` - List all saved job applications.
* `delete` - Delete a specific job application by its ID.
* `clear` - Clear all job applications at once.
* `export` - Export all job applications to a CSV file.

### Global flags

* `-h, --help` - Help for a command. 

## Examples

### Add a new job application

```bash
jobtracker add --company "OpenAI" --position "ML Engineer" --status "Applied" --applied_on "2025-08-26"
```

Specifying `--status` and `--applied_on` is not strictly required, in case these are missing, they will be replaced with `"Applied"` for `status` and today's date for `applied_on`:

```bash
jobtracker add -c "OpenAI" -p "ML Engineer"
```
> Running this command will create `jobs.json` file in the current directory for storing the job applications history.

### List all job applications

```bash
jobtracker list
```

Running the above command will output the saved job applications in a convenient and easy-to-read format:

```
ID  Company    Position                     Status     AppliedOn
--  -------    --------                     ------     ---------
1   Facebook   Software Engineer            Applied    2023-10-01
2   Google     Data Scientist               Interview  2023-09-15
3   Apple      Machine Learning Engineer    Applied    2023-10-01
4   Microsoft  Machine Learning Specialist  Applied    2025-08-30
5   Huawei     Frontend Developer           Applied    2025-08-30
6   Luxoft     Backend Developer            Applied    2025-08-30
7   NCR        Devops                       Applied    2025-08-30
```

You can optionally sort the job applications by the above columns and display the result in the same convenient format in ascending or descending order (`--desc` flag):

```bash
jobtracker list --sort=applied_on --desc
```

### Update an existing applications

If at some point we need to update the information on some applications, we can run:

```bash
jobtracker update 3 --status="Interview"
```
> This command will update entry with ID=3 and modify `status` to `"Interview"`.

### Delete an existing application

```bash
jobtracker delete 3
```
> This command will delete entry with ID=3 from the history.

### Clear all applications

```bash
jobtracker clear
```
> This command will firstly prompt for the user's confirmation and then delete all available applications. One can optionally set `--force` flag to skip prompting.

### Export all applications to CSV

```bash
jobtracker export
```
> This command will export all applications stored in `jobs.json` to `jobs.csv`

## Data storage

All job applications are stored locally in `jobs.json` in the current directory. Each entry includes:

* ID – Auto-incremented unique identifier

* Company – Company name

* Position – Job title/role

* Status – Application status (Applied, Interviewing, Offer, Rejected, etc.)

* DateApplied – Stored in ISO 8601 format (YYYY-MM-DD)

## License

MIT License. See [LICENSE](./LICENSE) for details.
