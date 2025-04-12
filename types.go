package main

import "time"

// AuthorStats holds statistics for a single author
type AuthorStats struct {
	Name         string
	CommitCount  int
	LinesChanged int
}

// WeeklyAuthorStats holds statistics for a single author for a specific week
type WeeklyAuthorStats struct {
	Name         string
	CommitCount  int
	LinesChanged int
	Week         time.Time // Start of the week (Sunday)
}

// WeeklyStats holds statistics for a specific week
type WeeklyStats struct {
	Week         time.Time // Start of the week (Sunday)
	Authors      map[string]*WeeklyAuthorStats
	TotalCommits int
	TotalLines   int
}

// RepositoryStats holds statistics for the entire repository
type RepositoryStats struct {
	Authors      map[string]*AuthorStats
	WeeklyStats  map[string]*WeeklyStats // Key is ISO week string "YYYY-WW"
	TotalCommits int
	TotalLines   int
	FileFilter   string
	IgnoreFiles  map[string]bool
}
