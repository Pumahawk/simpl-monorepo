function main() {
	while true; do
		case "$1" in
			--chart-file)
				CHART_FILE=${2?Add chart file path}
				shift
				shift
				;;
			--git-url)
				REPO_URL=${2?Add path to config file}
				shift
				shift
				;;
			--path)
				REPO_DIR=${2?Add path to config file}
				shift
				shift
				;;
			--git-credential)
				GIT_CREDENTIAL=${2?Add git credential}
				shift
				shift
				;;
			--git-branch)
				GIT_BRANCH=${2?Add git branch}
				shift
				shift
				;;
			--version-filter)
				HELM_VERSION_FILTER=${2?Add version filter}
				shift
				shift
				;;
			*)
				break;
				;;
		esac
	done
	log "Start update project."
	log "Chart file: $CHART_FILE"
	git_init_credential "$GIT_CREDENTIAL"
	git_clone_pull "$REPO_URL" "$REPO_DIR"
	CHART_DATA="$(cat "$CHART_FILE" | yq -p yaml -o json)"
	DEP_LIST="$(dep_list "$CHART_FILE")"
	UPDATE_LIST="$(HELM_VERSION_FILTER="$HELM_VERSION_FILTER" update_list "$DEP_LIST")"
	CHART_DATA="$(update_chart_data "$CHART_DATA" "$UPDATE_LIST")"
	write_chart_data "$CHART_FILE" "$CHART_DATA"
	# helm_dependency_update "$CHART_FILE"
	git_commit_and_push "$CHART_FILE"
	log "Pop directory"
	popd
}

function dep_list() {
	CHART_FILE="${1?Missin chart file parameter}"
	cat "$CHART_FILE" | yq -p yaml -o json | jq -c '.dependencies[] | select(.repository | test("code.europa.eu"))'
}

function update_list() {
	DEP_LIST="${1?"Missing parameter DEP_LIST"}"
	echo "$DEP_LIST" | while read line; do
		REPO=$(echo "$line" | jq -r .repository)
		NAME=$(echo "$line" | jq -r .name)
		local version_current=$(echo "$line" | jq -r .version)
		VERSION="$(helm_get_repo_last_version "$REPO" "$NAME" "$version_current")"
		echo "$line" | jq -c ".version = \"$VERSION\""
	done
}

function update_chart_data() {
	CHART_DATA="${1?"Missing parameter CHART_DATA"}"
	UPDATE_LIST="${2?"Missing parameter UPDATE_LIST"}"
	while read line; do
		NAME=$(echo "$line" | jq -r .name)
		VERSION=$(echo "$line" | jq -r .version)
		log Name: $NAME, Version: $VERSION
		CHART_DATA="$(echo "$CHART_DATA" | jq '.dependencies = [(.dependencies[] | if .name == "'"$NAME"'" then .version = "'"$VERSION"'" else . end)]')"
	done < <(echo "$UPDATE_LIST");
	echo "$CHART_DATA"
}

function write_chart_data() {
	CHART_FILE="${1?"Missing parameter CHART_FILE"}"
	CHART_DATA="${2?"Missing parameter CHART_DATA"}"
	log "Save $CHART_FILE"
	echo "$CHART_DATA" | json_to_yaml > "$CHART_FILE"
}

function helm_get_repo_last_version() {
	HELM_REPO_URL="${1?Missing helm repo url parameter}"
	HELM_REPO_NAME="${2?Missing helm repo name parameter}"
	local version_current="${3?Missing current version}"
	log "Get last version. Name: $HELM_REPO_NAME, Url: $HELM_REPO_URL, Filter: $HELM_VERSION_FILTER"
	local version="$(helm_client_index_json "$HELM_REPO_URL" | jq -r '.entries."'"$HELM_REPO_NAME"'" | sort_by(.created) '"$( [ -n "$HELM_VERSION_FILTER" ] && echo "| [ .[] | select($HELM_VERSION_FILTER)]")"' | last | .version')"
	if [[ "$version" == "null" || "$version" == "" ]]; then
		echo "$version_current"
	else
		echo "$version"
	fi
}

function helm_client_index_json() {
	HELM_REPO_URL="${1?Missing helm repo url parameter}"
	http_c "$HELM_REPO_URL/index.yaml" | yq -p yaml -o json
}

function git_clone_pull() {
	REPO_URL="${1?Repo url}"
	REPO_DIR="${2?Repo dir}"
	if [[ ! -d "$REPO_DIR" ]] ; then
		log "Create clone repository. Dir: $REPO_DIR, Url: $REPO_URL"
		git -c credential.helper="store --file /tmp/gitcredential" clone "$REPO_URL" "$REPO_DIR"
	fi
	log "Move to repo dir $REPO_DIR"
	pushd "$REPO_DIR"
	log "Git checkout branch $GIT_BRANCH"
	git checkout "$GIT_BRANCH"
	log "Git pull"
	git -c credential.helper="store --file /tmp/gitcredential" pull
}

function helm_dependency_update() {
	CHART_FILE="${1?"Missing parameter CHART_FILE"}"
	log "Helm update dependencies"
	pushd $(dirname "$CHART_FILE")
	helm dep update
	popd
}
function git_commit_and_push() {
	CHART_FILE="${1?"Missing parameter CHART_FILE"}"
	log "Git add, commit and push"
	git status
	git add $(dirname "$CHART_FILE") \
		&& git commit -m "Update dependencies version" \
		&& git -c credential.helper="store --file /tmp/gitcredential" push
}

function git_init_credential() {
	GIT_CREDENTIAL=${1?Add git credential}
	log "Set git credential"
	echo "$GIT_CREDENTIAL" > /tmp/gitcredential
}

function http_c() {
	curl -s "$@"
}

function json_to_yaml {
	yq -p json -o yaml
}

function yaml_to_json() {
	yq -p yaml -o json
}

function err() {
	log "$@"
}

function log() {
	>&2 echo "$(date -Iseconds) - ${LOG_CONTEXT-"NC"} " "$@"
}

main "$@"
