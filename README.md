# Todo App CLI
A CLI application for managing tasks in terminal.

```bash
$ todo-app-cli list # add, delete, update, etc
```

## Available Commands
```bash
$ todo-app-cli list [flags]
$ todo-app-cli add [arg] [flag]
$ todo-app-cli update [arg] [flags]
$ todo-app-cli delete [arg]

```

Each command supports specific flags that control how the command operates.

Add `-h` or `--help` at the end of the command to show detailed information of the command.

## Installing
> [!NOTE]
> Create a `.env` file and set the `DB_URL` value before proceeding with installation.

In order to be able to run the app using `todo-app-cli` command, follow these steps:
```bash
$ git clone {url_to_this_repo}
$ cd /path/to/todo-app-cli
$ go mod tidy # Checks go.mod file and download missing dependencies, or removed unneeded ones.
$ go install . # Compile and install the `todo-app-cli` binary to $GOBIN directory (or $GOPATH/bin). 
$ export PATH=$PATH:$(go env GOBIN)  # Or export PATH=$PATH:$(go env GOPATH)/bin if GOBIN is empty
$ echo 'export PATH=$PATH:$(go env GOBIN)' >> ~/.zshrc  # Or ~/.bashrc, etc. to make PATH persistence
$ source ~/.zshrc  # Or restart your terminal
```

Verify the installation by running `which todo-app-cli`, it should show the path to the binary. Then, try running `todo-app-cli list`.

## Running Locally
> [!NOTE]
> Create a `.env` file and set the `DB_URL` value before proceeding with installation.

Running the application locally is pretty straightforward:
```bash
$ git clone {url_to_this_repo}
$ cd /path/to/todo-app-cli
$ go mod tidy # Checks go.mod file and download missing dependencies, or removed unneeded ones.
$ go run main.go list # Show list of todos in the database
```
## Database Schema
The application uses a PostgreSQL database with the following schema:

| Column Name | Data Type   | Nullable | Default                                        | Description                    |
|------------|-------------|----------|------------------------------------------------|--------------------------------|
| id         | int4        | NO       | nextval('untitled_table_1_id_seq'::regclass)  | Primary key                    |
| title      | varchar     | NO       | NULL                                           | Task title/description         |
| completed  | bool        | NO       | false                                          | Task completion status         |
| created_at | timestamptz | NO       | now()                                          | Timestamp of task creation     |
| updated_at | timestamptz | NO       | now()                                          | Timestamp of last modification |

### Schema Notes
- The `id` field auto-increments using a sequence
- All fields are required (non-nullable)
- New tasks are automatically marked as not completed (`completed = false`)
- Both `created_at` and `updated_at` are automatically set to the current timestamp when creating a task
- The `updated_at` field should be updated whenever the task is modified

## Folders/Files Structure
```
todo-app-cli/
├── .git/                # Git repository files (not detailed here)
├── cmd/                 # Contains the command definitions
│   ├── add.go           # Implementation for the 'add' command
│   ├── delete.go        # Implementation for the 'delete' command
│   ├── list.go          # Implementation for the 'list' command
│   ├── root.go          # Root command and initialization logic
│   └── update.go        # Implementation for the 'update' command
├── config/              # Configuration related files
│   └── config.go        # Configuration loading and handling
├── internal/            # Internal packages for the application
│   └── database/        # Database interaction logic
│       └── database.go  # Database operations and connection management
├── .env.example         # Example environment variables file
├── .gitignore           # Specifies intentionally untracked files that Git should ignore
├── go.mod               # Go module definition file
├── go.sum               # Checksums of Go modules for integrity
├── main.go              # Entry point of the application
├── Makefile             # Makefile for build and install automation
└── README.md            # This file (project documentation)
```

### Key Directories and Files

* `cmd/`: Contains all CLI command definitions
  * `add.go`: Implements the 'add' command
  * `delete.go`: Implements the 'delete' command
  * `list.go`: Implements the 'list' command
  * `root.go`: Contains root command and initialization logic
  * `update.go`: Implements the 'update' command

* `config/`: Contains configuration-related code
  * `config.go`: Handles configuration loading and management

* `internal/`: Contains private application code
  * `database/`: Manages database interactions
    * `database.go`: Implements database operations and connection handling

* Core Files:
  * `.env.example`: Template for required environment variables
  * `go.mod` & `go.sum`: Go module definition and checksums
  * `main.go`: Application entry point
  * `Makefile`: Build and development automation
