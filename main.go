package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/olekukonko/tablewriter"
)

// AuthorStats holds statistics for a single author
type AuthorStats struct {
	Name         string
	CommitCount  int
	LinesChanged int
}

// RepositoryStats holds statistics for the entire repository
type RepositoryStats struct {
	Authors     map[string]*AuthorStats
	TotalCommits int
	TotalLines   int
	FileFilter   string
	IgnoreFiles  map[string]bool
}

func main() {
	// Define command-line flags
	ignoreFilesFlag := flag.String("ignore", "", "Comma-separated list of additional files to ignore")
	fileFilterFlag := flag.String("ext", "", "File extension filter (e.g., .js, .go)")
	
	// Parse command-line arguments
	flag.Parse()
	args := flag.Args()

	repoPath := "."
	fileFilter := *fileFilterFlag

	if len(args) > 0 {
		// Check if the first argument is a file extension (starts with a dot)
		if strings.HasPrefix(args[0], ".") {
			fileFilter = args[0]
		} else {
			repoPath = args[0]
			
			// Check if there's a second argument for file extension
			if len(args) > 1 && strings.HasPrefix(args[1], ".") {
				fileFilter = args[1]
			}
		}
	}

	// Open the repository
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		fmt.Printf("Error opening repository: %s\n", err)
		os.Exit(1)
	}

	// Initialize repository stats
	stats := &RepositoryStats{
		Authors:    make(map[string]*AuthorStats),
		FileFilter: fileFilter,
		IgnoreFiles: make(map[string]bool),
	}

	// Load .gitignore patterns
	loadGitignore(repoPath, stats)

	// Add common files to ignore
	commonIgnoreFiles := []string{
		"package-lock.json",
		"yarn.lock",
		"go.sum",
		"Cargo.lock",
		"Gemfile.lock",
	}
	for _, file := range commonIgnoreFiles {
		stats.IgnoreFiles[file] = true
	}
	
	// Add user-specified files to ignore
	if *ignoreFilesFlag != "" {
		userIgnoreFiles := strings.Split(*ignoreFilesFlag, ",")
		for _, file := range userIgnoreFiles {
			stats.IgnoreFiles[strings.TrimSpace(file)] = true
		}
	}

	// Get repository statistics
	err = analyzeRepository(repo, stats)
	if err != nil {
		fmt.Printf("Error analyzing repository: %s\n", err)
		os.Exit(1)
	}

	// Display statistics
	displayStats(stats)
}

// loadGitignore loads patterns from .gitignore file
func loadGitignore(repoPath string, stats *RepositoryStats) {
	gitignorePath := filepath.Join(repoPath, ".gitignore")
	content, err := os.ReadFile(gitignorePath)
	if err != nil {
		// .gitignore file might not exist, which is fine
		return
	}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.HasPrefix(line, "#") {
			stats.IgnoreFiles[line] = true
		}
	}
}

// analyzeRepository analyzes the Git repository and collects statistics
func analyzeRepository(repo *git.Repository, stats *RepositoryStats) error {
	// Get the HEAD reference
	ref, err := repo.Head()
	if err != nil {
		return err
	}

	// Get the commit object
	commit, err := repo.CommitObject(ref.Hash())
	if err != nil {
		return err
	}

	// Create a commit iterator
	commitIter, err := repo.Log(&git.LogOptions{From: commit.Hash})
	if err != nil {
		return err
	}
	defer commitIter.Close()

	// Iterate through commits
	err = commitIter.ForEach(func(c *object.Commit) error {
		// Get author name
		authorName := c.Author.Name

		// Variables to track if this commit should be counted
		commitAffectsFilteredFiles := false
		linesChanged := 0

		// Get commit stats
		if c.NumParents() > 0 {
			// For non-initial commits, compare with parent
			parent, err := c.Parent(0)
			if err == nil {
				patch, err := parent.Patch(c)
				if err == nil {
					for _, fileStat := range patch.Stats() {
						// Check if file should be included based on filter and ignore rules
						if shouldIncludeFile(fileStat.Name, stats.FileFilter, stats.IgnoreFiles) {
							commitAffectsFilteredFiles = true
							linesChanged += fileStat.Addition + fileStat.Deletion
						}
					}
				}
			}
		} else {
			// For initial commit, count all lines as additions
			files, err := c.Files()
			if err == nil {
				err = files.ForEach(func(f *object.File) error {
					if shouldIncludeFile(f.Name, stats.FileFilter, stats.IgnoreFiles) {
						commitAffectsFilteredFiles = true
						content, err := f.Contents()
						if err == nil {
							lineCount := len(strings.Split(content, "\n"))
							linesChanged += lineCount
						}
					}
					return nil
				})
			}
		}

		// Only count this commit if it affects files matching our filter
		if commitAffectsFilteredFiles {
			// Get or create author stats
			authorStats, ok := stats.Authors[authorName]
			if !ok {
				authorStats = &AuthorStats{
					Name: authorName,
				}
				stats.Authors[authorName] = authorStats
			}

			// Increment commit count
			authorStats.CommitCount++
			stats.TotalCommits++

			// Add lines changed
			authorStats.LinesChanged += linesChanged
			stats.TotalLines += linesChanged
		}

		return nil
	})

	return err
}

// shouldIncludeFile checks if a file should be included in statistics
func shouldIncludeFile(filename string, fileFilter string, ignoreFiles map[string]bool) bool {
	// Check if file is in ignore list
	if _, ok := ignoreFiles[filename]; ok {
		return false
	}

	// Check if file matches the extension filter
	if fileFilter != "" && !strings.HasSuffix(filename, fileFilter) {
		return false
	}

	return true
}

// displayStats displays repository statistics in an ASCII table
func displayStats(stats *RepositoryStats) {
	// Create a slice of authors for sorting
	authors := make([]*AuthorStats, 0, len(stats.Authors))
	for _, author := range stats.Authors {
		authors = append(authors, author)
	}

	// Sort authors by commit count (descending)
	sort.Slice(authors, func(i, j int) bool {
		return authors[i].CommitCount > authors[j].CommitCount
	})

	// Create and configure the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Author", "Commits", "Lines Changed", "Lines Changed %", "Commits %"})
	table.SetBorder(true)
	table.SetAutoFormatHeaders(false)

	// Add author rows
	for _, author := range authors {
		linesPercent := 0.0
		if stats.TotalLines > 0 {
			linesPercent = float64(author.LinesChanged) / float64(stats.TotalLines) * 100
		}

		commitsPercent := 0.0
		if stats.TotalCommits > 0 {
			commitsPercent = float64(author.CommitCount) / float64(stats.TotalCommits) * 100
		}

		table.Append([]string{
			author.Name,
			fmt.Sprintf("%d", author.CommitCount),
			fmt.Sprintf("%d", author.LinesChanged),
			fmt.Sprintf("%.1f%%", linesPercent),
			fmt.Sprintf("%.1f%%", commitsPercent),
		})
	}

	// Add total row
	table.Append([]string{
		"TOTAL",
		fmt.Sprintf("%d", stats.TotalCommits),
		fmt.Sprintf("%d", stats.TotalLines),
		"100%",
		"100%",
	})

	// Render the table
	table.Render()
}
