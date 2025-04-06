# Gitstics

A command-line tool written in Go that analyzes Git repositories and provides contribution statistics per author in an ASCII table.

## Features

- Displays commit count per author
- Calculates lines changed per author
- Shows percentage of total commits and lines changed
- Supports filtering by file extension
- Respects `.gitignore` rules
- Automatically ignores common dependency files (package-lock.json, yarn.lock, etc.)

## Installation

### Prerequisites

- Go 1.16 or higher
- Git

### Building from source

```bash
# Clone the repository
git clone https://github.com/fredrik/gitstics.git
cd gitstics

# Build the binary
go build -o gitstics

# Optionally, install the binary to your GOPATH
go install
```

## Usage

```bash
# Analyze the current repository
gitstics

# Analyze a specific repository
gitstics /path/to/repo

# Analyze only specific file types (e.g., JavaScript files)
gitstics /path/to/repo .js
# or
gitstics .js
# or
gitstics -ext=.js

# Ignore specific files (comma-separated)
gitstics -ignore="README.md,LICENSE" /path/to/repo

# Combine file type filtering and ignoring specific files
gitstics -ignore="test.js,mock.js" -ext=.js /path/to/repo
# or
gitstics -ignore="test.js,mock.js" /path/to/repo .js
```

## Example Output

```
+---------+---------+-----------------+-----------------+-----------+
| Author  | Commits | Lines Changed   | Lines Changed % | Commits % |
+---------+---------+-----------------+-----------------+-----------+
| Fredrik | 10      | 1234            | 90.0%           | 90.0%     |
| Mary    | 5       | 12              | 10.0%           | 10.0%     |
+---------+---------+-----------------+-----------------+-----------+
| TOTAL   | 14      | 1246            | 100%            | 100%      |
+---------+---------+-----------------+-----------------+-----------+
```

## How It Works

Gitstics analyzes the Git commit history to calculate:

1. The number of commits per author
2. The number of lines changed (additions + deletions) per author
3. The percentage of total commits and lines changed per author

The tool respects `.gitignore` rules and automatically ignores common dependency files like `package-lock.json`, `yarn.lock`, etc.

## Limitations

- The tool currently only analyzes the default branch
- Merge commits may skew the statistics
- Very large repositories may take longer to analyze
