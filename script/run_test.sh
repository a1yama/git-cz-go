#!/bin/bash
set -e

# Get directory of this script
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# Get root directory of the project
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

# Go to the root directory
cd "$ROOT_DIR"

# Run tests with verbose output
echo "Running tests..."
go test -v ./internal/...

# Run tests with coverage
echo "Running tests with coverage..."
go test -cover ./internal/...

echo "Tests completed successfully!"