#!/usr/bin/env bash

BASE_DIRECTORY="."
GIT_BRANCH_NAME="feature/SIMPL-update-version"
GIT_REMOTE_NAME="origin"
GIT_REMOTE_BRANCH="develop"

while true; do
	case "$1" in
		--branch-name)
			GIT_BRANCH_NAME="$2";
			shift;
			shift;
			;;
		--remote-name)
			GIT_REMOTE_NAME="$2";
			shift;
			shift;
			;;
		--remote-branch)
			GIT_REMOTE_BRANCH="$2";
			shift;
			shift;
			;;
		--pipeline-version)
			PIPELINE_VERSION="$2";
			shift;
			shift;
			;;
		*)
			break;
			;;
	esac
done;

if [ -z ${PIPELINE_VERSION+x} ]; then
	echo "--pipeline-version is mandatory"; 
	exit 2;
fi

function main() {
	find "$BASE_DIRECTORY" -name .git | (
		while read line; do 
			GIT_DIR="$line";
			PROJECT_DIR="$(dirname "$GIT_DIR")"
			git_fetch_remote "$GIT_DIR";
			git_checkout_new_branch "$GIT_DIR";
			update_version_pipeline "$PROJECT_DIR";
		done;
	)
}

function git_fetch_remote() {
	GIT_DIR="$1";
	echo "Fetch project. Git: $GIT_DIR, remote name: $GIT_REMOTE_NAME";
	GIT_DIR="$GIT_DIR" git fetch "$GIT_REMOTE_NAME";
}

function git_checkout_new_branch() {
	GIT_DIR="$1";
	echo "Create new branch. Git: $GIT_DIR, Branch name: $GIT_BRANCH_NAME";
	GIT_DIR="$GIT_DIR" git checkout -b "$GIT_BRANCH_NAME";
}

function update_version_pipeline() {
	PROJECT_DIR="$1";
	PIPELINE_FILE="$(find "$PROJECT_DIR" -name pipeline.variables.sh)";
	echo "Update versione pipeline. Project directory: $PROJECT_DIR, pipeline version: $PIPELINE_VERSION, pipeline file: $PIPELINE_FILE";
	sed -i 's/PROJECT_VERSION_NUMBER=.*/PROJECT_VERSION_NUMBER="'"$PIPELINE_VERSION"'"/' "$PIPELINE_FILE"
}

main "$@";
