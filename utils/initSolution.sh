#!/bin/bash

# This script creates the directories and files associated with a solution.
# After this script invocation, the user must provide the missing pieces, such
# as the solution implementation.

# Check invocation arguments
usage_msg="Usage: $0 <solution_name> <difficulty>\n\t<difficulty> = easy | medium | hard"
if [ "$#" -ne 2 ]; then
  echo -e $usage_msg >&2
  exit 1
fi
difficulty=`echo $2 | tr A-Z a-z`
if [ "$difficulty" != "easy" ] && [ "$difficulty" != "medium" ] && [ "$difficulty" != "hard" ]; then
  echo -e $usage_msg >&2
  exit 2
fi

# Init needed variables
hacker_rank_challenges_dir=`dirname "$0"`
solution_name="$1"
solution_name_lc=`echo $solution_name | tr A-Z a-z`

# Create dir hierarchy for the solution
cd `dirname "$0"`/..
mkdir -p "solutions/$solution_name/pkg/solution/"

# Create the README.md
readme_template="# PROBLEM_NAME\n\nProblem statement is [here](PROBLEM_STATEMENT_URL)."
echo -e $readme_template | sed 's/PROBLEM_NAME/'$solution_name'/g' > "solutions/$solution_name/README.md"

# Create the source file containing the solution
solution_template="package solution\n\n//Complete the SOLUTION_NAME function below\n//Time complexity:\n//Space complexity:\nfunc SOLUTION_NAME() {\n\n}"
echo -e $solution_template | sed 's/SOLUTION_NAME/'$solution_name'/g' > "solutions/$solution_name/pkg/solution/${solution_name_lc}.go"

# Insert link to solution in root README.md. Currently supported only on
# MAC OS X, because sed behaves differently on Linux.
if [ "$(uname)" == "Darwin" ]; then
  sed -i.bak '/\* '$difficulty'/a\
  \ \ \* \['$solution_name'\](solutions\/'$solution_name')\
  ' README.md
  if [[ $? -eq 0 ]]; then
    rm README.md.bak
  else
    echo "Updating the root README.md failed: restoring previous version from backup."
    rm README.md
    mv README.md.bak README.md
  fi
else
  echo "WARNING: A link to the solution sub directory has not been automatically added. This is currently only supported in MAC OS X. Add the link manually."
fi

echo "Don't forget to manually insert in solutions/$solution_name/README.md the URL to the problem statement!"
