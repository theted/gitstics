package main

import (
	"reflect"
	"testing"
)

func TestCalculateAverageCommitSize(t *testing.T) {
	// Create a test repository stats
	stats := &RepositoryStats{
		Authors: map[string]*AuthorStats{
			"Alice": {
				Name:         "Alice",
				CommitCount:  4,
				LinesChanged: 20,
			},
			"Bob": {
				Name:         "Bob",
				CommitCount:  2,
				LinesChanged: 10,
			},
			"Charlie": {
				Name:         "Charlie",
				CommitCount:  3,
				LinesChanged: 15,
			},
			"EmptyUser": {
				Name:         "EmptyUser",
				CommitCount:  0,
				LinesChanged: 0,
			},
		},
	}

	// Calculate average commit size
	result := CalculateAverageCommitSize(stats)

	// Expected results
	expected := map[string]float64{
		"Alice":     5.0,  // 20 lines / 4 commits
		"Bob":       5.0,  // 10 lines / 2 commits
		"Charlie":   5.0,  // 15 lines / 3 commits
		"EmptyUser": 0.0,  // 0 lines / 0 commits
	}

	// Compare results
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("CalculateAverageCommitSize() = %v, want %v", result, expected)
	}
}

func TestCalculateCodeChurn(t *testing.T) {
	// This is a placeholder test for now
	// In a real implementation, we would:
	// 1. Create a test repository stats
	// 2. Call CalculateCodeChurn
	// 3. Compare the results with expected values
	
	stats := &RepositoryStats{}
	result := CalculateCodeChurn(stats)
	
	// For now, just check that we get a non-nil result
	if result == nil {
		t.Errorf("CalculateCodeChurn() returned nil")
	}
}

func TestCalculateContributorDiversity(t *testing.T) {
	// This is a placeholder test for now
	// In a real implementation, we would:
	// 1. Create a test repository stats
	// 2. Call CalculateContributorDiversity
	// 3. Compare the results with expected values
	
	stats := &RepositoryStats{}
	result := CalculateContributorDiversity(stats)
	
	// For now, just check that we get a non-nil result
	if result == nil {
		t.Errorf("CalculateContributorDiversity() returned nil")
	}
}
