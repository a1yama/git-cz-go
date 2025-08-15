#!/bin/bash

# Demo script for git-cz-go
# This script demonstrates the functionality in a test environment

echo "🚀 Starting git-cz-go demo..."
echo ""
echo "This demo will show all features of git-cz-go:"
echo "1. Commit type selection"
echo "2. Scope input"
echo "3. Subject/description"
echo "4. Body (detailed description)"
echo "5. Breaking changes"
echo "6. Footer (issue references)"
echo ""
echo "Setting up test repository..."

# Create test directory
TEST_DIR="/tmp/git-cz-go-demo-$(date +%s)"
mkdir -p "$TEST_DIR"
cd "$TEST_DIR"

# Initialize git repo
git init
git config user.email "test@example.com"
git config user.name "Test User"

# Create a test file
echo "# Demo Project" > README.md
echo "This is a test project for git-cz-go" >> README.md
git add README.md

echo ""
echo "✅ Test repository ready at: $TEST_DIR"
echo ""
echo "Now you can run git-cz-go to create a commit:"
echo ""
echo "  cd $TEST_DIR"
echo "  /Users/a1yama/ghq/git-cz-go/git-cz-go"
echo ""
echo "Or run it directly from here:"
echo "  /Users/a1yama/ghq/git-cz-go/git-cz-go"
echo ""
echo "Example flow:"
echo "  1. Type: feat (press 1)"
echo "  2. Scope: demo"
echo "  3. Subject: add initial demo setup"
echo "  4. Body: (optional, press Enter to skip or Enter twice to finish)"
echo "  5. Breaking: No (press N)"
echo "  6. Footer: (optional, press Enter to skip)"
echo "  7. Confirm: Yes"