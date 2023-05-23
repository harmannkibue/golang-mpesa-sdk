#!/usr/bin/env bash

set -e

# Get the current branch name
BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo "PUSHING TO BRANCH ${BRANCH}"

# Add all modified files to the Git index
git add .

# Check if there are any changes to commit
if git diff-index --quiet HEAD --; then
  echo "No changes to commit."
else
  # Prompt for a commit message
  read -rp "Enter a commit message: " COMMIT_MESSAGE

  # Commit changes with the provided message
  git commit -m "$COMMIT_MESSAGE"

  # Push the changes to the remote branch
  git push --verbose origin "$BRANCH"
fi