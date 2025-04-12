package main

import (
	"testing"
	"time"
)

func TestCalculateAuthorActivityPatterns(t *testing.T) {
	// Create a test repository stats
	stats := &RepositoryStats{}
	
	// Calculate author activity patterns
	result := CalculateAuthorActivityPatterns(stats)
	
	// Check that we get a non-nil result
	if result == nil {
		t.Errorf("CalculateAuthorActivityPatterns() returned nil")
	}
	
	// Check that Charlie's pattern is included
	charliePattern, exists := result["Charlie"]
	if !exists {
		t.Errorf("Charlie's activity pattern not found in results")
	}
	
	// Check that the pattern has the expected structure
	if charliePattern != nil {
		// Check day of week counts
		if len(charliePattern.DayOfWeekCounts) != 7 {
			t.Errorf("Expected 7 days in DayOfWeekCounts, got %d", len(charliePattern.DayOfWeekCounts))
		}
		
		// Check hour of day counts
		if len(charliePattern.HourOfDayCounts) == 0 {
			t.Errorf("HourOfDayCounts is empty")
		}
		
		// Check month counts
		if len(charliePattern.MonthCounts) != 12 {
			t.Errorf("Expected 12 months in MonthCounts, got %d", len(charliePattern.MonthCounts))
		}
		
		// Check average commit gap
		if charliePattern.AverageCommitGap <= 0 {
			t.Errorf("Expected positive AverageCommitGap, got %f", charliePattern.AverageCommitGap)
		}
	}
}

func TestCalculateFileAgeStats(t *testing.T) {
	// Create a test repository stats
	stats := &RepositoryStats{}
	
	// Current time for reference
	now := time.Now()
	t.Logf("Running test at: %s", now.Format(time.RFC3339))
	
	// Calculate file age stats
	result := CalculateFileAgeStats(stats)
	
	// Check that we get a non-nil result
	if result == nil {
		t.Errorf("CalculateFileAgeStats() returned nil")
	}
	
	// Check that we have at least one file
	if len(result) == 0 {
		t.Errorf("CalculateFileAgeStats() returned empty slice")
	}
	
	// Check that the stats have the expected structure
	for i, fileStat := range result {
		if fileStat.FileName == "" {
			t.Errorf("File %d has empty name", i)
		}
		
		if fileStat.CreationDate.IsZero() {
			t.Errorf("File %s has zero CreationDate", fileStat.FileName)
		}
		
		if fileStat.LastModified.IsZero() {
			t.Errorf("File %s has zero LastModified", fileStat.FileName)
		}
		
		if fileStat.Age <= 0 {
			t.Errorf("File %s has non-positive Age: %v", fileStat.FileName, fileStat.Age)
		}
		
		if fileStat.ModificationCount <= 0 {
			t.Errorf("File %s has non-positive ModificationCount: %d", fileStat.FileName, fileStat.ModificationCount)
		}
		
		if fileStat.ModsPerMonth <= 0 {
			t.Errorf("File %s has non-positive ModsPerMonth: %f", fileStat.FileName, fileStat.ModsPerMonth)
		}
		
		if fileStat.AuthorCount <= 0 {
			t.Errorf("File %s has non-positive AuthorCount: %d", fileStat.FileName, fileStat.AuthorCount)
		}
		
		if fileStat.PrimaryAuthor == "" {
			t.Errorf("File %s has empty PrimaryAuthor", fileStat.FileName)
		}
		
		if fileStat.PrimaryAuthorShare <= 0 || fileStat.PrimaryAuthorShare > 1.0 {
			t.Errorf("File %s has invalid PrimaryAuthorShare: %f", fileStat.FileName, fileStat.PrimaryAuthorShare)
		}
	}
}

func TestCalculateTeamCollaboration(t *testing.T) {
	// Create a test repository stats
	stats := &RepositoryStats{}
	
	// Calculate team collaboration
	result := CalculateTeamCollaboration(stats)
	
	// Check that we get a non-nil result
	if result == nil {
		t.Errorf("CalculateTeamCollaboration() returned nil")
	}
	
	// Check that we have at least one collaboration
	if len(result) == 0 {
		t.Errorf("CalculateTeamCollaboration() returned empty slice")
	}
	
	// Check that the collaborations have the expected structure
	for i, collab := range result {
		if collab.AuthorPair[0] == "" || collab.AuthorPair[1] == "" {
			t.Errorf("Collaboration %d has empty author(s)", i)
		}
		
		if collab.SharedFiles <= 0 {
			t.Errorf("Collaboration %s-%s has non-positive SharedFiles: %d", 
				collab.AuthorPair[0], collab.AuthorPair[1], collab.SharedFiles)
		}
		
		if collab.SequentialEdits <= 0 {
			t.Errorf("Collaboration %s-%s has non-positive SequentialEdits: %d", 
				collab.AuthorPair[0], collab.AuthorPair[1], collab.SequentialEdits)
		}
		
		if collab.SameWeekEdits <= 0 {
			t.Errorf("Collaboration %s-%s has non-positive SameWeekEdits: %d", 
				collab.AuthorPair[0], collab.AuthorPair[1], collab.SameWeekEdits)
		}
	}
}

func TestGetTopCollaborators(t *testing.T) {
	// Create test collaborations
	collaborations := []*TeamCollaboration{
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
		{
			AuthorPair:      [2]string{"Alice", "David"},
			SharedFiles:     2,
			SequentialEdits: 5,
			SameWeekEdits:   2,
		},
		{
			AuthorPair:      [2]string{"Bob", "David"},
			SharedFiles:     1,
			SequentialEdits: 3,
			SameWeekEdits:   1,
		},
	}
	
	// Get top collaborators
	result := GetTopCollaborators(collaborations)
	
	// Check that we get the expected number of results
	if len(result) != 3 {
		t.Errorf("Expected 3 top collaborators, got %d", len(result))
	}
	
	// Check that the results are sorted by SharedFiles
	for i := 0; i < len(result)-1; i++ {
		if result[i].SharedFiles < result[i+1].SharedFiles {
			t.Errorf("Results not sorted by SharedFiles: %d < %d", 
				result[i].SharedFiles, result[i+1].SharedFiles)
		}
	}
	
	// Check that the top collaborator is Alice-Bob
	if result[0].AuthorPair[0] != "Alice" || result[0].AuthorPair[1] != "Bob" {
		t.Errorf("Expected top collaborator to be Alice-Bob, got %s-%s", 
			result[0].AuthorPair[0], result[0].AuthorPair[1])
	}
}
