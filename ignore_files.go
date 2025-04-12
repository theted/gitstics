package main

// CommonIgnoreFiles is a list of common dependency files that should be ignored by default
var CommonIgnoreFiles = []string{
	"package-lock.json",
	"yarn.lock",
	"go.sum",
	"Cargo.lock",
	"Gemfile.lock",
}
