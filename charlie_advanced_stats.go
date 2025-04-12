package main

import (
	"sort"
	"time"
)

// AdvancedStats provides more sophisticated statistics
// for deeper analysis of repository activity

// AuthorActivityPattern represents the activity pattern of an author
type AuthorActivityPattern struct {
	Author           string
	DayOfWeekCounts  map[time.Weekday]int // Commits per day of week
	HourOfDayCounts  map[int]int          // Commits per hour of day
	MonthCounts      map[time.Month]int   // Commits per month
	AverageCommitGap float64              // Average days between commits
}

// CalculateAuthorActivityPatterns analyzes when authors tend to commit
func CalculateAuthorActivityPatterns(stats *RepositoryStats) map[string]*AuthorActivityPattern {
	// This is a placeholder implementation
	// In a real implementation, we would analyze the commit timestamps
	
	// For demonstration purposes, return a sample pattern
	return map[string]*AuthorActivityPattern{
		"Charlie": {
			Author: "Charlie",
			DayOfWeekCounts: map[time.Weekday]int{
				time.Monday:    2,
				time.Tuesday:   1,
				time.Wednesday: 3,
				time.Thursday:  2,
				time.Friday:    4,
				time.Saturday:  1,
				time.Sunday:    0,
			},
			HourOfDayCounts: map[int]int{
				9:  2,
				10: 3,
				11: 1,
				14: 2,
				15: 3,
				16: 2,
			},
			MonthCounts: map[time.Month]int{
				time.January:   1,
				time.February:  2,
				time.March:     3,
				time.April:     4,
				time.May:       2,
				time.June:      1,
				time.July:      0,
				time.August:    1,
				time.September: 2,
				time.October:   3,
				time.November:  2,
				time.December:  1,
			},
			AverageCommitGap: 3.5, // days
		},
	}
}

// FileAgeStats represents statistics about file age and modification frequency
type FileAgeStats struct {
	FileName           string
	CreationDate       time.Time
	LastModified       time.Time
	Age                time.Duration
	ModificationCount  int
	ModsPerMonth       float64
	AuthorCount        int
	PrimaryAuthor      string
	PrimaryAuthorShare float64 // percentage of changes by primary author
}

// CalculateFileAgeStats analyzes file age and modification patterns
func CalculateFileAgeStats(stats *RepositoryStats) []*FileAgeStats {
	// This is a placeholder implementation
	// In a real implementation, we would analyze the commit history
	
	// For demonstration purposes, return sample stats
	return []*FileAgeStats{
		{
			FileName:           "main.go",
			CreationDate:       time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
			LastModified:       time.Date(2025, 4, 10, 0, 0, 0, 0, time.UTC),
			Age:                time.Hour * 24 * 85, // ~85 days
			ModificationCount:  12,
			ModsPerMonth:       4.2,
			AuthorCount:        3,
			PrimaryAuthor:      "Alice",
			PrimaryAuthorShare: 0.6, // 60%
		},
		{
			FileName:           "analyze.go",
			CreationDate:       time.Date(2025, 2, 20, 0, 0, 0, 0, time.UTC),
			LastModified:       time.Date(2025, 4, 5, 0, 0, 0, 0, time.UTC),
			Age:                time.Hour * 24 * 44, // ~44 days
			ModificationCount:  8,
			ModsPerMonth:       5.5,
			AuthorCount:        2,
			PrimaryAuthor:      "Bob",
			PrimaryAuthorShare: 0.75, // 75%
		},
	}
}

// TeamCollaboration represents collaboration metrics between team members
type TeamCollaboration struct {
	AuthorPair      [2]string // pair of authors
	SharedFiles     int       // number of files both authors modified
	SequentialEdits int       // number of times one author edited after the other
	SameWeekEdits   int       // number of weeks both authors were active
}

// CalculateTeamCollaboration analyzes how team members work together
func CalculateTeamCollaboration(stats *RepositoryStats) []*TeamCollaboration {
	// This is a placeholder implementation
	// In a real implementation, we would analyze the commit history
	
	// For demonstration purposes, return sample collaboration stats
	return []*TeamCollaboration{
		{
			AuthorPair:      [2]string{"Alice", "Bob"},
			SharedFiles:     5,
			SequentialEdits: 8,
			SameWeekEdits:   4,
		},
		{
			AuthorPair:      [2]string{"Alice", "Charlie"},
			SharedFiles:     3,
			SequentialEdits: 6,
			SameWeekEdits:   3,
		},
		{
			AuthorPair:      [2]string{"Bob", "Charlie"},
			SharedFiles:     4,
			SequentialEdits: 7,
			SameWeekEdits:   3,
		},
	}
}

// GetTopCollaborators returns the most collaborative author pairs
func GetTopCollaborators(collaborations []*TeamCollaboration) []*TeamCollaboration {
	// Sort by number of shared files (could use other metrics)
	sort.Slice(collaborations, func(i, j int) bool {
		return collaborations[i].SharedFiles > collaborations[j].SharedFiles
	})
	
	// Return top 3 or all if less than 3
	if len(collaborations) <= 3 {
		return collaborations
	}
	return collaborations[:3]
}
