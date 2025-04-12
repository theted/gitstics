package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// TestAliceMainFunctionality is a basic test for the main functionality
func TestAliceMainFunctionality(t *testing.T) {
	// Skip this test if we're running in CI or don't want to create temp dirs
	if os.Getenv("SKIP_REPO_TESTS") != "" {
		t.Skip("Skipping test that requires creating a git repository")
	}

	// Create a temporary directory for the test repository
	tempDir, err := os.MkdirTemp("", "gitstics-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Initialize git repository
	repo, err := git.PlainInit(tempDir, false)
	if err != nil {
		t.Fatalf("Failed to initialize git repo: %v", err)
	}

	// Create a test file
	testFilePath := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFilePath, []byte("Initial content"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Add the file to git
	w, err := repo.Worktree()
	if err != nil {
		t.Fatalf("Failed to get worktree: %v", err)
	}

	_, err = w.Add("test.txt")
	if err != nil {
		t.Fatalf("Failed to add file: %v", err)
	}

	// Make a commit as Test User
	_, err = w.Commit("Initial commit", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Test User",
			Email: "test@example.com",
		},
	})
	if err != nil {
		t.Fatalf("Failed to commit: %v", err)
	}

	// Modify the file
	if err := os.WriteFile(testFilePath, []byte("Modified content\nWith multiple lines\nFor testing"), 0644); err != nil {
		t.Fatalf("Failed to modify test file: %v", err)
	}

	// Commit the changes as Charlie
	_, err = w.Add("test.txt")
	if err != nil {
		t.Fatalf("Failed to add modified file: %v", err)
	}

	_, err = w.Commit("Update test file", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Charlie",
			Email: "charlie@example.com",
		},
	})
	if err != nil {
		t.Fatalf("Failed to commit as Charlie: %v", err)
	}

	// Create a second test file
	testFile2Path := filepath.Join(tempDir, "test2.txt")
	if err := os.WriteFile(testFile2Path, []byte("Second file content"), 0644); err != nil {
		t.Fatalf("Failed to write second test file: %v", err)
	}

	// Add and commit the second file as Alice
	_, err = w.Add("test2.txt")
	if err != nil {
		t.Fatalf("Failed to add second file: %v", err)
	}

	_, err = w.Commit("Add second test file", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "Alice Developer",
			Email: "alice@example.com",
		},
	})
	if err != nil {
		t.Fatalf("Failed to commit as Alice: %v", err)
	}

	// Now test the repository stats functionality
	// Initialize repository stats
	stats := &RepositoryStats{
		Authors:     make(map[string]*AuthorStats),
		WeeklyStats: make(map[string]*WeeklyStats),
		FileFilter:  "",
		IgnoreFiles: make(map[string]bool),
	}

	// Analyze the repository
	err = AnalyzeRepository(repo, stats)
	if err != nil {
		t.Fatalf("Failed to analyze repository: %v", err)
	}

	// Check that we have the expected authors
	expectedAuthors := []string{"Test User", "Charlie", "Alice Developer"}
	for _, author := range expectedAuthors {
		if _, ok := stats.Authors[author]; !ok {
			t.Errorf("Expected author %s not found in stats", author)
		}
	}

	// Check that the total commit count is correct
	if stats.TotalCommits != 3 {
		t.Errorf("Expected 3 total commits, got %d", stats.TotalCommits)
	}

	// Check that Charlie has 1 commit
	if stats.Authors["Charlie"].CommitCount != 1 {
		t.Errorf("Expected Charlie to have 1 commit, got %d", stats.Authors["Charlie"].CommitCount)
	}

	// Test the display functionality by capturing stdout
	originalStdout := os.Stdout
	r, pipeWriter, _ := os.Pipe()
	os.Stdout = pipeWriter

	// Display the stats
	DisplayStats(stats)

	// Restore stdout
	pipeWriter.Close()
	os.Stdout = originalStdout

	// Read the captured output
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Check that the output contains the expected authors
	for _, author := range expectedAuthors {
		if !bytes.Contains(buf.Bytes(), []byte(author)) {
			t.Errorf("Expected output to contain author %s", author)
		}
	}

	// Check that the output contains the expected headers
	expectedHeaders := []string{"Author", "Commits", "Lines Changed", "Lines Changed %", "Commits %"}
	for _, header := range expectedHeaders {
		if !bytes.Contains(buf.Bytes(), []byte(header)) {
			t.Errorf("Expected output to contain header %s", header)
		}
	}

	t.Logf("Test repository successfully analyzed and displayed")
}
