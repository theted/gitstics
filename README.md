# Gitstics

A command-line tool written in Go that analyzes Git repositories and provides contribution statistics per author in an ASCII table.

## Features

- Displays commit count per author
- Calculates lines changed per author
- Shows percentage of total commits and lines changed
- Provides weekly code frequency statistics (lines changed per week per user)
- Supports filtering by file extension (only counts commits that modify files of the specified extension)
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

# Show weekly code frequency statistics
gitstics -weekly
# or
gitstics -weekly /path/to/repo
# or
gitstics -weekly -ext=.js /path/to/repo
```

## Example Output

### Default Output

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

### Weekly Code Frequency Output

```
+------------+---------+---------------+------------+---------+
| Week       | Author  | Lines Changed | Lines/Week | Commits |
+------------+---------+---------------+------------+---------+
| 2025-04-01 | Fredrik | 500           | 500.0      | 4       |
|            | Mary    | 10            | 10.0       | 2       |
|            |         |               |            |         |
| 2025-04-08 | Fredrik | 734           | 734.0      | 6       |
|            | Mary    | 2             | 2.0        | 3       |
|            |         |               |            |         |
+------------+---------+---------------+------------+---------+
```

## How It Works

Gitstics analyzes the Git commit history to calculate:

1. The number of commits per author (only counting commits that modify files matching the filter criteria)
2. The number of lines changed (additions + deletions) per author
3. The percentage of total commits and lines changed per author
4. Weekly code frequency statistics showing lines changed per week per author

When a file extension filter is specified (e.g., `.js`), the tool will only count commits that modify files with that extension. This provides accurate statistics for contributions to specific file types.

The tool respects `.gitignore` rules and automatically ignores common dependency files like `package-lock.json`, `yarn.lock`, etc.

The weekly code frequency feature groups commits by ISO week and shows how many lines each author changed during that week. This helps visualize development activity over time and identify periods of high productivity or code churn.

## Limitations

- The tool currently only analyzes the default branch
- Merge commits may skew the statistics
- Very large repositories may take longer to analyze
