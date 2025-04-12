package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

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

			// Get the week start date (Sunday)
			commitTime := c.Author.When
			year, week := commitTime.ISOWeek()
			weekStart := getWeekStart(year, week)
			weekKey := fmt.Sprintf("%d-W%02d", year, week)

			// Get or create weekly stats
			weeklyStats, ok := stats.WeeklyStats[weekKey]
			if !ok {
				weeklyStats = &WeeklyStats{
					Week:    weekStart,
					Authors: make(map[string]*WeeklyAuthorStats),
				}
				stats.WeeklyStats[weekKey] = weeklyStats
			}

			// Get or create weekly author stats
			weeklyAuthorStats, ok := weeklyStats.Authors[authorName]
			if !ok {
				weeklyAuthorStats = &WeeklyAuthorStats{
					Name: authorName,
					Week: weekStart,
				}
				weeklyStats.Authors[authorName] = weeklyAuthorStats
			}

			// Update weekly stats
			weeklyAuthorStats.CommitCount++
			weeklyAuthorStats.LinesChanged += linesChanged
			weeklyStats.TotalCommits++
			weeklyStats.TotalLines += linesChanged
		}

		return nil
	})

	return err
}

// getWeekStart returns the start date (Sunday) of the given ISO week
func getWeekStart(year, week int) time.Time {
	// Get the date of the first day of the year
	jan1 := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	
	// Get the day of the week for January 1st
	jan1Weekday := int(jan1.Weekday())
	
	// Calculate days to add to get to the first day (Monday) of week 1
	daysToFirstMonday := (8 - jan1Weekday) % 7
	if jan1Weekday == 0 { // Sunday
		daysToFirstMonday = 1
	}
	
	// Calculate the date of the first Monday of week 1
	firstMonday := jan1.AddDate(0, 0, daysToFirstMonday)
	
	// Calculate the date of the first day (Sunday) of the requested week
	// Week 1 starts with the Monday closest to January 1st
	weekStart := firstMonday.AddDate(0, 0, (week-1)*7-1)
	
	return weekStart
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
