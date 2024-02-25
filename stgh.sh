#!/bin/bash

REPO_NAME="NihonGo"
FOSSIL_DIR="${HOME}/Projects/fossilized"
FOSSIL_FILE="${FOSSIL_DIR}/${REPO_NAME}.fossil"
echo "Fossil file: ${FOSSIL_FILE}"
GITHUB_REPO_URL="https://github.com/crgimenes/$REPO_NAME.git"
TEMP_DIR="temp_git_repo"

if [ -d "$TEMP_DIR" ]; then
    echo "Removing old $TEMP_DIR"
    rm -rf "$TEMP_DIR"
fi

mkdir "$TEMP_DIR"
cd "$TEMP_DIR"


echo "creating git repository..."
git init

echo "exporting from fossil..."
fossil export --git $FOSSIL_FILE | git fast-import

git checkout trunk

echo "setting remote... $GITHUB_REPO_URL"
git remote add origin "$GITHUB_REPO_URL"

# rename trunk to master
git branch -m trunk master

echo "pushing to github..."
git push -u origin master -f

echo "removing temp dir..."
cd ..
rm -rf "$TEMP_DIR"

echo "fim."

