#!/bin/bash

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
	log "Extraction data from securityattributesprovider to $CSV_PATH"
	local db="securityattributesprovider"
	PGDATABASE="${db}" \
		PGUSER="${db}" \
		PGPASSWORD="${db}" \
		psql "$@" -f <( echo "$SQL_SAP" )

	log "Load data to onboarding from $CSV_PATH"
	local db="onboarding"
	PGDATABASE="${db}" \
		PGUSER="${db}" \
		PGPASSWORD="${db}" \
		psql "$@" -f <(echo "$SQL_ONBOARDING")
}

function log() {
	>&2 echo "$(date -Iseconds) - ${LOG_CONTEXT-"NC"} " "$@"
}


main "$@"

