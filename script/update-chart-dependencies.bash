function main() {
	CHART_FILE="${1?Missing chart file parameter}"
	log "Start update project."
	log "Chart file: $CHART_FILE"
	CHART_DATA="$(cat "$CHART_FILE" | yq -p yaml -o json)"
	DEP_LIST="$(dep_list "$CHART_FILE")"
	UPDATE_LIST="$(update_list "$DEP_LIST")"
	CHART_DATA="$(update_chart_data "$CHART_DATA" "$UPDATE_LIST")"
	write_chart_data "$CHART_FILE" "$CHART_DATA"
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
		VERSION="$(helm_get_repo_last_version "$REPO" "$NAME")"
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
	log "Get last version. Name: $HELM_REPO_NAME, Url: $HELM_REPO_URL, Filter: $HELM_VERSION_FILTER"
	helm_client_index_json "$HELM_REPO_URL" | jq -r '.entries."'"$HELM_REPO_NAME"'" | sort_by(.created) '"$( [ -n "$HELM_VERSION_FILTER" ] && echo "| [ .[] | select($HELM_VERSION_FILTER)]")"' | last | .version'
}

function helm_client_index_json() {
	HELM_REPO_URL="${1?Missing helm repo url parameter}"
	http_c "$HELM_REPO_URL/index.yaml" | yq -p yaml -o json
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
