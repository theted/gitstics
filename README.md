# Gitstics

A command-line tool written in Go that analyzes Git repositories and provides contribution statistics per author in an ASCII table. Gitstics helps you understand who contributed what to your codebase and when, making it easier to track team productivity and code ownership.

## Features

- Displays commit count per author
- Calculates lines changed per author (additions + deletions)
- Shows percentage of total commits and lines changed
- Provides weekly code frequency statistics (lines changed per week per user)
- Supports filtering by file extension (only counts commits that modify files of the specified extension)
- Respects `.gitignore` rules
- Automatically ignores common dependency files (package-lock.json, yarn.lock, go.sum, etc.)

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
+-----------------+---------+-----------------+-----------------+-----------+
| Author          | Commits | Lines Changed   | Lines Changed % | Commits % |
+-----------------+---------+-----------------+-----------------+-----------+
| Test User       | 7       | 1250            | 42.4%           | 38.9%     |
| Charlie         | 5       | 663             | 22.5%           | 27.8%     |
| Joe             | 5       | 625             | 21.2%           | 27.8%     |
| Alice Developer | 2       | 412             | 14.0%           | 11.1%     |
+-----------------+---------+-----------------+-----------------+-----------+
| TOTAL           | 18      | 2950            | 100%            | 100%      |
+-----------------+---------+-----------------+-----------------+-----------+
```

### Weekly Code Frequency Output

```
+------------+----------------+---------------+------------+---------+
| Week       | Author         | Lines Changed | Lines/Week | Commits |
+------------+----------------+---------------+------------+---------+
| 2025-04-06 | Test User      | 850           | 850.0      | 4       |
|            |                |               |            |         |
| 2025-04-12 | Test User      | 400           | 400.0      | 3       |
|            |                |               |            |         |
| 2025-04-13 | Test User      | 150           | 150.0      | 1       |
|            | Alice Developer| 412           | 412.0      | 2       |
|            | Charlie        | 663           | 663.0      | 5       |
|            |                |               |            |         |
| 2025-04-20 | Joe            | 625           | 625.0      | 5       |
|            |                |               |            |         |
+------------+----------------+---------------+------------+---------+
```

## How It Works

Gitstics analyzes the Git commit history to calculate:

1. The number of commits per author (only counting commits that modify files matching the filter criteria)
2. The number of lines changed (additions + deletions) per author
3. The percentage of total commits and lines changed per author
4. Weekly code frequency statistics showing lines changed per week per author

When a file extension filter is specified (e.g., `.js`), the tool will only count commits that modify files with that extension. This provides accurate statistics for contributions to specific file types.

The tool respects `.gitignore` rules and automatically ignores common dependency files like `package-lock.json`, `yarn.lock`, `go.sum`, etc.

The weekly code frequency feature groups commits by ISO week and shows how many lines each author changed during that week. This helps visualize development activity over time and identify periods of high productivity or code churn.

### Use Cases

- **Team Performance Analysis**: Track team member contributions over time
- **Code Ownership**: Identify who has the most knowledge about specific parts of the codebase
- **Project Handover**: Understand who contributed to which parts of the project before transitioning responsibilities
- **Language-Specific Analysis**: Focus on contributions to specific file types (e.g., only JavaScript files)
- **Activity Tracking**: Visualize development activity patterns over time with weekly statistics

## Limitations

- The tool currently only analyzes the default branch
- Merge commits may skew the statistics
- Very large repositories may take longer to analyze
