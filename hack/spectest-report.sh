#!/bin/bash

set -xe

# Constants
PROJECT_ROOT=$(pwd)
SILA_DIR="${PROJECT_ROOT%/hack}/testing/spectest"
EXCLUSION_LIST="$SILA_DIR/exclusions.txt"
BAZEL_DIR="/tmp/spectest_report"
SPEC_REPO="git@github.com:sila/consensus-spec-tests.git"
SPEC_DIR="/tmp/consensus-spec"

# Create directory if it doesn't already exist
mkdir -p "$BAZEL_DIR"

# Add any passed flags to BAZEL_FLAGS
BAZEL_FLAGS=""
for flag in "$@"
do
    BAZEL_FLAGS="$BAZEL_FLAGS $flag"
done

# Run spectests
bazel test //testing/spectest/... --test_env=SPEC_TEST_REPORT_OUTPUT_DIR="$BAZEL_DIR" $BAZEL_FLAGS

# Ensure the SPEC_DIR exists and is a git repository
if [ -d "$SPEC_DIR/.git" ]; then
    echo "Repository already exists. Pulling latest changes."
    (cd "$SPEC_DIR" && git pull) || exit 1
else
    echo "Cloning the GitHub repository."
    git clone "$SPEC_REPO" "$SPEC_DIR" || exit 1
fi

# Finding all *_tests.txt files in BAZEL_DIR and concatenating them into tests.txt
find "$BAZEL_DIR" -type f -name '*_tests.txt' -exec cat {} + > "$SILA_DIR/tests.txt"

# Generating spec.txt
(cd "$SPEC_DIR" && find tests -maxdepth 4 -mindepth 4 -type d > "$SILA_DIR/spec.txt") || exit 1

# Comparing spec.txt with tests.txt and generating report.txt
while IFS= read -r line; do
   if grep -Fxq "$line" "$EXCLUSION_LIST"; then
      # If it's excluded and we have a test for it flag as an error
      if grep -q "$line" "$SILA_DIR/tests.txt"; then
          echo "Error: Excluded item found in tests.txt: $line"
          exit 1 # Exit with an error status
      else
          echo "Skipping excluded item: $line"
      fi
          continue
    fi
    if grep -q "$line" "$SILA_DIR/tests.txt"; then
        echo "found $line"
    else
        echo "missing $line"
    fi
done < "$SILA_DIR/spec.txt" > "$SILA_DIR/report.txt"

# Formatting report.txt
{
    echo "Sila Spectest Report"
    echo ""
    echo "Tests Missing"
    grep '^missing' "$SILA_DIR/report.txt"
    echo ""
    echo "Tests Found"
    grep '^found' "$SILA_DIR/report.txt"
} > "$SILA_DIR/report_temp.txt" && mv "$SILA_DIR/report_temp.txt" "$SILA_DIR/report.txt"

# Check for the word "missing" in the report and exit with an error if present
if grep -q '^missing' "$SILA_DIR/report.txt"; then
    echo "Error: 'missing' tests found in report: $SILA_DIR/report.txt"
    exit 1
fi

# Clean up
rm -f "$SILA_DIR/tests.txt" "$SILA_DIR/spec.txt"
