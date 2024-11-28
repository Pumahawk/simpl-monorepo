#!/usr/bin/env bash

BASE_DIRECTORY="simpl-repo"
GIT_BRANCH_NAME="feature/SIMPL-update-version"
GIT_REMOTE_NAME="origin"
GIT_REMOTE_BRANCH="develop"

STASH_ENABLE="true"
OPENAPI_HOST="https://t1.authority.dev.aruba-simpl.cloud"
FORMAT_JSON="true"
SKIP_FETCH="false"
FORMAT_JSON_CMD="jq ."

function print_help() {
	echo "Update parent pom and script version"
	echo;
column -ts \; <<EOF
     --pom-version VERSION;Mandatory;New pom version
     --pipeline-version VERSION;Default: like pom version;New pipeline version
     --openapi-version VERSION;Default: like pom version;New openapi version
     --dir PATH;Default: $BASE_DIRECTORY;Base directory
     --branch-name NAME;Default: $GIT_BRANCH_NAME;Create new branch with given name
     --remote-name NAME;Default: $GIT_REMOTE_NAME;Remote repository name
     --remote-branch NAME;Default: $GIT_REMOTE_BRANCH;Remote branch name
     --git-stash;Default: true;Enable stash mode
     --no-git-stash;;Disable stash mode
     --no-fetch;;Disable stash mode
     --fetch --no-fetch;Default: true;Enable or disable git fetch
EOF
}

while true; do
	case "$1" in
		--dir)
			BASE_DIRECTORY="$2";
			shift;
			shift;
			;;
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
		--pom-version)
			POM_VERSION="$2";
			shift;
			shift;
			;;
		--pipeline-version)
			PIPELINE_VERSION="$2";
			shift;
			shift;
			;;
		--openapi-version)
			OPENAPI_VERSION="$2";
			shift;
			shift;
			;;
		--git-stash)
			STASH_ENABLE="true"
			shift;
			;;
		--no-git-stash)
			STASH_ENABLE="false"
			shift;
			;;
		--format-json)
			FORMAT_JSON="true"
			shift;
			;;
		--no-format-json)
			FORMAT_JSON="false"
			shift;
			;;
		--fetch)
			SKIP_FETCH="false"
			shift;
			;;
		--no-fetch)
			SKIP_FETCH="true"
			shift;
			;;
		--help)
			print_help;
			exit 0;
			;;
		*)
			break;
			;;
	esac
done;

if [ -z ${POM_VERSION+x} ]; then
	echo "--pom-version is mandatory"; 
	exit 2;
fi

if [ -z ${PIPELINE_VERSION+x} ]; then
	PIPELINE_VERSION="$POM_VERSION"
fi

if [ -z ${OPENAPI_VERSION+x} ]; then
	OPENAPI_VERSION="$POM_VERSION"
fi

if [ $FORMAT_JSON == "false" ]; then
	FORMAT_JSON_CMD="cat"
fi


function main() {
	find "$BASE_DIRECTORY" -name .git | (
		while read line; do 
			GIT_DIR="$line";
			PROJECT_DIR="$(dirname "$GIT_DIR")"
			if [ $SKIP_FETCH == "false" ]; then
				git_fetch_remote "$GIT_DIR";
			fi
			git_stash_all "$GIT_DIR";
			git_checkout_new_branch "$GIT_DIR";
			update_version_pipeline "$PROJECT_DIR";
			update_version_pom "$PROJECT_DIR";
			git_commit "$GIT_DIR";
		done;
	);
	update_all_openapi
}

function git_fetch_remote() {
	GIT_DIR="$1";
	echo "Fetch project. Git: $GIT_DIR, remote name: $GIT_REMOTE_NAME";
	GIT_DIR="$GIT_DIR" git fetch "$GIT_REMOTE_NAME";
}

function git_checkout_new_branch() {
	GIT_DIR="$1";
	echo "Create new branch. Git: $GIT_DIR, Branch name: $GIT_BRANCH_NAME";
	GIT_DIR="$GIT_DIR" git checkout --detach
	GIT_DIR="$GIT_DIR" git branch -D "$GIT_BRANCH_NAME"
	GIT_DIR="$GIT_DIR" git checkout -b "$GIT_BRANCH_NAME" origin/develop;
}

function git_stash_all() {
	GIT_DIR="$1";
	if [ $STASH_ENABLE == "true" ]; then
		GIT_DIR="$GIT_DIR" git add -A;
		GIT_DIR="$GIT_DIR" git stash;
	else
		echo "Stash not enabled";
	fi
}

function update_version_pipeline() {
	PROJECT_DIR="$1";
	PIPELINE_FILE="$(find "$PROJECT_DIR" -name pipeline.variables.sh)";
	echo "Update versione pipeline. Project directory: $PROJECT_DIR, pipeline version: $PIPELINE_VERSION, pipeline file: $PIPELINE_FILE";
	sed -i 's/PROJECT_VERSION_NUMBER=.*/PROJECT_VERSION_NUMBER="'"$PIPELINE_VERSION"'"/' "$PIPELINE_FILE"
}

function update_version_pom() {
	PROJECT_DIR="$1";
	echo "Update pom version"
	sed -i '/simpl-.*parent/,/<\/version>/s/<version>.*<\/version>/<version>'"$POM_VERSION"'<\/version>/' $(find "$PROJECT_DIR" -name pom.xml);
	sed -i 's!<simpl.httpclient.version>.*</simpl.httpclient.version>!<simpl.httpclient.version>'$POM_VERSION'</simpl.httpclient.version>!' $(find "$PROJECT_DIR" -name pom.xml);
}

function git_commit() {
	GIT_DIR="$1";
	GIT_MESSAGE="Update pipeline version";
	echo "Git add all and commit. Message: $GIT_MESSAGE";
	GIT_DIR="$GIT_DIR" git add -A;
	GIT_DIR="$GIT_DIR" git commit -m "$GIT_MESSAGE";
}

function update_all_openapi() {
	update_openapi "$BASE_DIRECTORY/agent-service"                     "$OPENAPI_HOST/public/auth-api/v3/api-docs"
	update_openapi "$BASE_DIRECTORY/identity-provider"                 "$OPENAPI_HOST/public/identity-api/v3/api-docs"
	update_openapi "$BASE_DIRECTORY/onboarding"                        "$OPENAPI_HOST/public/onboarding-api/v3/api-docs";
	update_openapi "$BASE_DIRECTORY/security-attributes-provider"      "$OPENAPI_HOST/public/sap-api/v3/api-docs"
	update_openapi "$BASE_DIRECTORY/users-roles"                       "$OPENAPI_HOST/public/user-api/v3/api-docs"
}

function update_openapi() {
	PROJECT_DIR="$1";
	URL="$2";
	echo "Update openapi $PROJECT_DIR"
	curl -s "$URL" | $FORMAT_JSON_CMD > "$PROJECT_DIR/openapi/openApi-doc-$OPENAPI_VERSION-release.json"
	GIT_DIR="$PROJECT_DIR/.git" git add openapi;
	GIT_DIR="$PROJECT_DIR/.git" git commit -m "Update opena Api, version: $OPENAPI_VERSION"
}

main "$@";
