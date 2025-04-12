package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func main() {
	// Define command-line flags
	ignoreFilesFlag := flag.String("ignore", "", "Comma-separated list of additional files to ignore")
	fileFilterFlag := flag.String("ext", "", "File extension filter (e.g., .js, .go)")
	weeklyFlag := flag.Bool("weekly", false, "Show weekly code frequency statistics")
	
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
		Authors:     make(map[string]*AuthorStats),
		WeeklyStats: make(map[string]*WeeklyStats),
		FileFilter:  fileFilter,
		IgnoreFiles: make(map[string]bool),
	}

	// Load .gitignore patterns
	loadGitignore(repoPath, stats)

	// Add common files to ignore
	for _, file := range CommonIgnoreFiles {
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
	if *weeklyFlag {
		displayWeeklyStats(stats)
	} else {
		displayStats(stats)
	}
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
