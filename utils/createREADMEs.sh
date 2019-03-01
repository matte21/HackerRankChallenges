#!/bin/bash

# This script creates a README.md in directories under solutions/ which do not
# already have one. The created README.md is initialized with a template the
# user must manually fill.

cd `dirname "$0"`
template="# <PROBLEM_NAME>\n\nProblem statement is [here](<PROBLEM_STATEMENT_URL>)."
find ../solutions -maxdepth 1 -mindepth 1 -type d -exec bash -c "cd '{}' && [ ! -f README.md ] && echo -e \"$template\" > README.md" \;
