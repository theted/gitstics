package main

// Export types and functions for testing
var (
	// Functions
	ExportedAnalyzeRepository = AnalyzeRepository
	ExportedDisplayStats      = DisplayStats
	ExportedDisplayWeeklyStats = DisplayWeeklyStats
)

// Export types for testing
type (
	ExportedAuthorStats       = AuthorStats
	ExportedWeeklyAuthorStats = WeeklyAuthorStats
	ExportedWeeklyStats       = WeeklyStats
	ExportedRepositoryStats   = RepositoryStats
)
