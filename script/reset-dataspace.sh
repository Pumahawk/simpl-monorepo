#! /bin/bash

CRITICAL_MESSAGE="WARNING! Critical script running. ALL DATABASES WILL BE DROPPED and the EJBCA SECRET WILL BE DELETED. Ensure you understand the implications."

function log() {
	>&2 echo "$(date -Iseconds) - ${LOG_CONTEXT-"NC"} " "$@"
}

if [[ "$RESET_AUTHORITY_ENVIRONMENT" != "1" ]]; then
	log "$CRITICAL_MESSAGE"
	log "Set environment variable RESET_AUTHORITY_ENVIRONMENT=1"
	exit 1
fi

if [[ "$RESET_AUTHORITY_ENVIRONMENT_SKIP" != "1" ]]; then
	sleep_sec="10"
	log "$CRITICAL_MESSAGE"
	log "System will pause for $sleep_sec seconds, providing an opportunity to abort."
	sleep "$sleep_sec"
fi

# TYPE="${1?Insert type [participant, authority]}"
# log "Type: $type"


KUBECTL_CMD=(kubectl)
if [[ -n $K_NAMESPACE ]]; then
	KUBECTL_CMD+=("--namespace" "${K_NAMESPACE}")
fi

ALL_DATABASE=(
	"authenticationprovider:authenticationprovider:authenticationprovider"
	"ejbca:ejbca:ejbca2123"
	"identityprovider:identityprovider:identityprovider"
	"keycloak:keycloak:keycloak"
	"onboarding:onboarding:onboarding"
	"securityattributesprovider:securityattributesprovider:securityattributesprovider"
	"usersroles:usersroles:usersroles"
)

ALL_DEPLOYMENTS=(
	"deployment/authentication-provider"
	"deployment/ejbca-community-helm"
	"deployment/identity-provider"
	"deployment/onboarding"
	"deployment/security-attributes-provider"
	"deployment/tier1-gateway"
	"deployment/tier2-gateway"
	"deployment/users-roles"
	"statefulset.apps/keycloak"
)


function main() {
	LOG_CONTEXT=main

	log  "Clean dataspace start."

	# LOG_CONTEXT="verifyEnvironment" verifyEnvironment || {
	# 	log "Invalid cluster status. Validation error"
	# 	return 1
	# }

	# [[ "$TYPE" == "authority" ]] && \
	LOG_CONTEXT="deleteEJBCASecrets" deleteEJBCASecrets || {
		log "Unable to delate EJBCA Secrets"
		log "Skip condition secrets"
	}

	LOG_CONTEXT="cleanDB" cleanDB || {
		log "Unable to clean database"
		# return 1
	}

	LOG_CONTEXT="scaleAllTo0" scaleAllTo0 || {
		log "Unable to scale all to 0"
		# return 1
	}

	log "Sleep 5 seconds"
	sleep 5

	# LOG_CONTEXT="waitAllScaleTo0" waitAllScaleTo0 || {
	# 	log "Unable to waitAllScaleTo0"
	# 	# return 1
	# }

	LOG_CONTEXT="scaleAllTo1" scaleAllTo1 || {
		log "Unable to scale all to 1"
		# return 1
	}

	# LOG_CONTEXT="waitAllScaleTo1" waitAllScaleTo1 || {
	# 	log "Unable to waitAllScaleTo1"
	# 	# return 1
	# 	log "MANUALLY DELETE INIT JOB!!"
	# }

	LOG_CONTEXT="deleteInitJob" deleteInitJob || {
		log "Unable to delete init job"
	}

}

function cleanDB() {
	log "Start Clean DB"
       for db_conn in "${ALL_DATABASE[@]}"; do
	       local db="$(echo $db_conn | cut -d: -f1)"
	       local user="$(echo $db_conn | cut -d: -f2)"
	       local password="$(echo $db_conn | cut -d: -f3)"
               ( PGPASSWORD="${password}" dropdb -f -U${user} "$db" && PGPASSWORD="${password}" createdb -U$user -O "$user" "$db" ) || { log "Unable drop or create db $db"; }
       done
	log "Success Clean Database"
}

function deleteEJBCASecrets() {
	log "Delete EJBCA secrets"
	if ! kubectl_cmd delete secret ejbca-rest-api-secret; then
		log "Unable to delete secret ejbca-rest-api-secret"
		return 1
	fi
}

function verifyEnvironment() {
	log "Verify environment"

	isAuthorityDataspace \
	&& hasAllDeployments \
	&& verifyEJBCAStatus \
	&& exitsAllExpectedDatabase
}

# Skip this part.
# Script must work in authority and participant
function isAuthorityDataspace() {
	log "Verify isAuthorityDataspace"
	log "Retrieve current namespace from onboarding deployment"
	local namespace="$(kubectl_cmd get deployment -ojsonpath='{.items[].metadata.namespace}' | head -n1)"
	if [[ $? -ne 0 ]]; then
		log. "ERROR Unable to retrieve namespace"
		return 1
	fi
	log "Namespace: $namespace"
	if [[ "$namespace" =~ authority ]]; then
		log "OK - Current namespace contains string authority."
		return 0
	else
		log "ERROR - Invalid actual namespace. Must contain authority."
		return 1
	fi
}

function verifyEJBCAStatus() {
	log "Verify verifyEJBCAStatus"
	log "Retrieve deployment.apps/ejbca-community-helm"
	if kubectl_cmd get deployment.apps/ejbca-community-helm; then
		log "OK - deployment.apps/ejbca-community-helm"
		return 0
	else
		log "ERROR - deployment ejbca-community-helm not found"
		return 1
	fi
}

function hasAllDeployments() {
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		hasDeployment  "$dep" || { log "ERROR - Not found deployment $dep"; }
	done
}

function hasDeployment() {
	local dep_name="$1"
	log "Find deployment $dep_name"
	kubectl_cmd get deployment "$dep_name" > /dev/null
}

function exitsAllExpectedDatabase() {
	log "Verify exitsAllExpectedDatabase"

	for db_conn in "${ALL_DATABASE[@]}"; do
	       local db="$(echo $db_conn | cut -d: -f1)"
	       local user="$(echo $db_conn | cut -d: -f2)"
	       local password="$(echo $db_conn | cut -d: -f3)"
		if ! PGPASSWORD=$password psql -U$user -tAc "select 1 from pg_database where datname = '$db'" | grep -q 1; then
			log "ERROR - Not found dabatase '$db'"
			# return 1
		fi
	done
}

function scaleAllTo1() {
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		log "Scale to 1 $dep"
		kubectl_cmd scale --replicas 1 "$dep" || { log "ERROR - Unable to scale project $dep"; }
	done
}

function scaleAllTo0() {
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		log "Scale to 0 $dep"
		kubectl_cmd scale --replicas 0 "$dep" || { log "ERROR - Unable to scale project $dep"; }
	done
}

function waitAllScaleTo1() {
	local cond="[1-9]/"
	local seconds="5"
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		waitDeploymentStatus "$dep" "$cond" "$seconds" || { log "ERROR - Unable to wait project $dep"; }
	done
}

function waitAllScaleTo0() {
	local cond="0/"
	local seconds="5"
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		waitDeploymentStatus "$dep" "$cond" "$seconds" || { log "ERROR - Unable to wait project $dep"; }
	done
}

function waitDeploymentStatus() {
	local dep_name="$1"
	local dep_exp_status="$2"
	local wait_seconds="$3"

	local dep_status=0
	while [[ "$dep_status" != 1 ]]; do
		local dep_output="$( kubectl_cmd get --no-headers "$dep_name" )"
		if [[ $? != 0 ]]; then
			log "ERROR - unable to get deployment $dep_name"
			# return 1
		fi
		dep_status="$(echo "$dep_output" | grep -E "$dep_exp_status" | wc -l )"
		if [[ $dep_status != 1 ]]; then
			log "Waiting deployment $dep_name"
			sleep "$wait_seconds"
		fi
	done
}

function deleteInitJob() {
	kubectl_cmd delete job init-authority-job || kubectl_cmd delete job init-participant-job
}

function kubectl_cmd() {
	"${KUBECTL_CMD[@]}" "$@"
}

main "$@"
