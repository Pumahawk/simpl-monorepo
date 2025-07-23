function lookup() {
	local file="$1"
	local properties="$(cat "$file" |
	yq '.properties | to_entries[] | select(.value.type == "string" and .value.maxLength == null and  .value.format == null) | .key')";
		if [ -n "$properties" ]; then
			# echo -e "$file\\n[\\n $properties \\n]\\n" 
			echo "$properties" | while read line; do
				grep -Hn "${line}:" "$file" | head -n1
			done;

		fi
}

function main() {
	while read line; do
		log "process: $line"
		lookup "$line"
	done

}

function log() {
	>&2 echo "$@"
}

main
