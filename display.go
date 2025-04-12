package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/olekukonko/tablewriter"
)

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

// displayWeeklyStats displays weekly code frequency statistics in an ASCII table
func displayWeeklyStats(stats *RepositoryStats) {
	// Create a slice of weeks for sorting
	weeks := make([]*WeeklyStats, 0, len(stats.WeeklyStats))
	for _, week := range stats.WeeklyStats {
		weeks = append(weeks, week)
	}

	// Sort weeks by date (ascending)
	sort.Slice(weeks, func(i, j int) bool {
		return weeks[i].Week.Before(weeks[j].Week)
	})

	// Create and configure the table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Week", "Author", "Lines Changed", "Lines/Week", "Commits"})
	table.SetBorder(true)
	table.SetAutoFormatHeaders(false)

	// Add rows for each week and author
	for _, week := range weeks {
		// Create a slice of authors for this week
		authors := make([]*WeeklyAuthorStats, 0, len(week.Authors))
		for _, author := range week.Authors {
			authors = append(authors, author)
		}

		// Sort authors by lines changed (descending)
		sort.Slice(authors, func(i, j int) bool {
			return authors[i].LinesChanged > authors[j].LinesChanged
		})

		// Format the week as YYYY-MM-DD
		weekStr := week.Week.Format("2006-01-02")

		// Add rows for each author in this week
		for i, author := range authors {
			weekDisplay := ""
			if i == 0 {
				// Only show the week for the first author in each week
				weekDisplay = weekStr
			}

			table.Append([]string{
				weekDisplay,
				author.Name,
				fmt.Sprintf("%d", author.LinesChanged),
				fmt.Sprintf("%.1f", float64(author.LinesChanged)),
				fmt.Sprintf("%d", author.CommitCount),
			})
		}

		// Add a separator between weeks
		if len(authors) > 0 {
			table.Append([]string{"", "", "", "", ""})
		}
	}

	// Render the table
	table.Render()
}
