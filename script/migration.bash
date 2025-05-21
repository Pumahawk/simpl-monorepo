#!/bin/bash

# From securityattributesprovider
# To onboarding

TMP_TABLE="tmp_stage";
CSV_PATH="/tmp/identity_attributes_participant_types.csv";

SQL_SAP="$(cat << EOF

BEGIN;

\\copy ( \
	select \
	ia.code as code, \
	iapt.participant_type as participant_type \
	from identity_attribute ia join identity_attribute_participant_type iapt on ia.id = iapt.identity_attribute_id \
) TO '$CSV_PATH' CSV HEADER;

COMMIT;

EOF
)";

SQL_ONBOARDING="$(cat << EOF
BEGIN;

CREATE TEMP TABLE $TMP_TABLE (
	code character varying(255),
	participant_type character varying(255)
);

\\copy tmp_stage FROM '$CSV_PATH' CSV HEADER;

INSERT into onboarding_procedure_template_identity_attribute (
	select opt.id as onboarding_procedure_template_id, tmp.code as identity_attribute_code
	from
		onboarding_procedure_template opt
		join participant_type pt on opt.participant_type_id = pt.id
		join $TMP_TABLE tmp on tmp.participant_type = pt.value
) on CONFLICT ( onboarding_procedure_template_id, identity_attribute_code) DO NOTHING;

COMMIT;

EOF
)";

function main() {

	while [ -n "$1" ]; do
		case "$1" in
			--db-host)
				local db_host="$2"
				shift;
				;;
			--from-db)
				local from_db="$2"
				shift;
				;;
			--from-user)
				local from_user="$2"
				shift;
				;;
			--from-password)
				local from_password="$2"
				shift;
				;;
			--to-db)
				local to_db="$2"
				shift;
				;;
			--to-user)
				local to_user="$2"
				shift;
				;;
			--to-password)
				local to_password="$2"
				shift;
				;;
			*)
				;;
		esac;
		shift;
	done;

	local from_db="${from_db-"securityattributesprovider"}"
	local to_db="${to_db-"onboarding"}"

	[[ -z "${db_host}" ]] && { log "Missing parameter --db-host"; exit 1; }
	[[ -z "${from_user}" ]] && { log "Missing parameter --from-user"; exit 1; }
	[[ -z "${from_password}" ]] && { log "Missing parameter --from-password"; exit 1; }
	[[ -z "${to_user}" ]] && { log "Missing parameter --to-user"; exit 1; }
	[[ -z "${to_password}" ]] && { log "Missing parameter --to-password"; exit 1; }

	log "Extraction data from securityattributesprovider to $CSV_PATH"
	PGHOST="${db_host}" \
		PGDATABASE="${from_db}" \
		PGUSER="${from_user}" \
		PGPASSWORD="${from_password}" \
		psql "$@" -f <( echo "$SQL_SAP" ) || { log "Unable to retrieve data from sap"; exit 1; }

	log "Load data to onboarding from $CSV_PATH"
	PGHOST="${db_host}" \
		PGDATABASE="${to_db}" \
		PGUSER="${to_user}" \
		PGPASSWORD="${to_password}" \
		psql "$@" -f <(echo "$SQL_ONBOARDING") || { log "Unable to write data to onboarding"; exit 1; }
}

function log() {
	>&2 echo "$(date -Iseconds) - ${LOG_CONTEXT-"NC"} " "$@"
}


main "$@"

