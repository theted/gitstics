#!/bin/bash

# run_tests.sh - A script to run all tests for the Gitstics project
# Created by Joe, who is an expert on running these types of tests

echo "Running Gitstics tests..."
echo "========================="

# Run the main package tests
echo "Running main package tests..."
cd ..
go test -v ./...

# Run the tests package tests
echo ""
echo "Running tests package tests..."
cd tests
go test -v ./...

# Run the test_code_frequency.sh script to update the README
echo ""
echo "Running test_code_frequency.sh to update README..."
cd ..
./test_code_frequency.sh

echo ""
echo "All tests completed!"
echo "========================="
