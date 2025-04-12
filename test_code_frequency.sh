#!/bin/bash

# Create a temporary directory for the test repository
TEST_REPO=$(mktemp -d)
echo "Creating test repository at $TEST_REPO"

# Initialize git repository
cd "$TEST_REPO"
git init

# Configure git to not require email verification
git config --global --add safe.directory "$TEST_REPO"

# Create some test files
echo "Initial content" > file1.txt
echo "Initial content" > file2.txt

# Function to make commits with a specific author and date
make_commit() {
    local author="$1"
    local email="$2"
    local date="$3"
    local message="$4"
    local file="$5"
    local content="$6"
    
    echo "$content" > "$file"
    git add "$file"
    GIT_AUTHOR_NAME="$author" GIT_AUTHOR_EMAIL="$email" GIT_AUTHOR_DATE="$date" \
    GIT_COMMITTER_NAME="$author" GIT_COMMITTER_EMAIL="$email" GIT_COMMITTER_DATE="$date" \
    git commit -m "$message"
}

# Add initial commit
git add .
GIT_AUTHOR_NAME="Setup" GIT_AUTHOR_EMAIL="setup@example.com" GIT_AUTHOR_DATE="2025-03-25T12:00:00" \
GIT_COMMITTER_NAME="Setup" GIT_COMMITTER_EMAIL="setup@example.com" GIT_COMMITTER_DATE="2025-03-25T12:00:00" \
git commit -m "Initial commit"

# Week 1: March 30 - April 5, 2025
make_commit "Alice" "alice@example.com" "2025-03-31T10:00:00" "Alice's first commit" "file1.txt" "Alice added this line\nAnd another line\nAnd one more line"
make_commit "Bob" "bob@example.com" "2025-04-01T14:30:00" "Bob's first commit" "file2.txt" "Bob added this line\nAnd another line"
make_commit "Alice" "alice@example.com" "2025-04-02T09:15:00" "Alice's second commit" "file1.txt" "Alice added this line\nAnd another line\nAnd one more line\nAdded a fourth line\nAnd a fifth line"
make_commit "Charlie" "charlie@example.com" "2025-04-03T16:45:00" "Charlie's first commit" "file1.txt" "Alice added this line\nAnd another line\nAnd one more line\nAdded a fourth line\nAnd a fifth line\nCharlie added this line"

# Week 2: April 6 - April 12, 2025
make_commit "Bob" "bob@example.com" "2025-04-07T11:20:00" "Bob's second commit" "file2.txt" "Bob added this line\nAnd another line\nBob added a third line\nAnd a fourth line"
make_commit "Alice" "alice@example.com" "2025-04-08T13:10:00" "Alice's third commit" "file1.txt" "Alice modified everything completely\nThis is all new content\nWith multiple lines\nTo show significant changes"
make_commit "Charlie" "charlie@example.com" "2025-04-09T15:30:00" "Charlie's second commit" "file2.txt" "Bob added this line\nAnd another line\nBob added a third line\nAnd a fourth line\nCharlie modified Bob's file\nAdding two new lines\nTo demonstrate changes"
make_commit "Bob" "bob@example.com" "2025-04-10T09:45:00" "Bob's third commit" "file1.txt" "Alice modified everything completely\nThis is all new content\nWith multiple lines\nTo show significant changes\nBob added a line to Alice's file"

# Week 3: April 13 - April 19, 2025
make_commit "Alice" "alice@example.com" "2025-04-14T10:30:00" "Alice's fourth commit" "file1.txt" "Alice modified everything completely\nThis is all new content\nWith multiple lines\nTo show significant changes\nBob added a line to Alice's file\nAlice added another line\nAnd one more"
make_commit "Charlie" "charlie@example.com" "2025-04-15T14:15:00" "Charlie's third commit" "file2.txt" "Charlie completely rewrote this file\nWith all new content\nSpanning multiple lines\nTo show a major change\nBy a single author"

# Build the gitstics tool if it doesn't exist
cd -
if [ ! -f "./gitstics" ]; then
    echo "Building gitstics tool..."
    go build -o gitstics
fi

# Run gitstics with the weekly flag
echo -e "\nRunning gitstics with weekly flag:"
./gitstics -weekly "$TEST_REPO"

# Run regular gitstics for comparison
echo -e "\nRunning regular gitstics for comparison:"
./gitstics "$TEST_REPO"

echo -e "\nTest repository is at: $TEST_REPO"
echo "You can remove it when done with: rm -rf $TEST_REPO"
