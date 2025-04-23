#! /bin/bash

KUBECTL_CMD=(kubectl)
if [[ -n $K_NAMESPACE ]]; then
	KUBECTL_CMD+=("--namespace" "${K_NAMESPACE}")
fi

ALL_DATABASE=(
	"authenticationprovider"
	"ejbca"
	"identityprovider"
	"keycloak"
	"onboarding"
	"securityattributesprovider"
	"usersroles"
)

ALL_DEPLOYMENTS=(
	"authentication-provider"
	"ejbca-community-helm"
	"identity-provider"
	"onboarding"
	"security-attributes-provider"
	"tier1-gateway"
	"tier2-gateway"
	"users-roles"
)


function main() {
	LOG_CONTEXT=main
	log  "Clean dataspace start."

	LOG_CONTEXT="verifyEnvironment" verifyEnvironment || {
		log "Invalid cluster status. Validation error"
		return 1
	}

	LOG_CONTEXT="scaleAllTo0" scaleAllTo0 || {
		log "Unable to scale all to 0"
		return 1
	}

	LOG_CONTEXT="waitAllScaleTo0" waitAllScaleTo0 || {
		log "Unable to waitAllScaleTo0"
		return 1
	}

	LOG_CONTEXT="cleanDB" cleanDB || {
		log "Unable to clean database"
		return 1
	}

	LOG_CONTEXT="deleteEJBCASecrets" deleteEJBCASecrets || {
		log "Unable to delate EJBCA Secrets"
		log "Skip condition secrets"
	}

	LOG_CONTEXT="scaleAllTo1" scaleAllTo1 || {
		log "Unable to scale all to 1"
		return 1
	}

	LOG_CONTEXT="waitAllScaleTo1" waitAllScaleTo1 || {
		log "Unable to waitAllScaleTo1"
		return 1
	}

}

# Delete all database
function cleanDB() {
	log "Start Clean DB"
	for db in "${ALL_DATABASE[@]}"; do
		( dropdb "$db" && createdb -O "$db" "$db" ) || { log "Unable drop or create db $db"; return 1; }
	done
	log "Success Clean Database"
}

# Delete all EJBCA secrets
function deleteEJBCASecrets() {
	log "Delete EJBCA secrets"
	if ! kubectl_cmd delete secret ejbca-rest-api-secret; then
		log "Unable to delete secret ejbca-rest-api-secret"
		return 1
	fi
}

# Verify environment status.
#   Is an Authority dataspace?
#   Is EJBCA reachable?
#   Contains all expected database?
function verifyEnvironment() {
	log "Verify environment"

	isAuthorityDataspace \
	&& hasAllDeployments \
	&& verifyEJBCAStatus \
	&& exitsAllExpectedDatabase
}

# Verify is current namespace is an Authority dataspace
function isAuthorityDataspace() {
	log "Verify isAuthorityDataspace"
	log "Retrieve current namespace from onboardin deployment"
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

# Verify if EJBCA exist and is receable in current dataspace
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
		hasDeployment  "$dep" || { log "ERROR - Not found deployment $dep"; return 1; }
	done
}

function hasDeployment() {
	local dep_name="$1"
	log "Find deployment $dep_name"
	kubectl_cmd get deployment "$dep_name" > /dev/null
}

function exitsAllExpectedDatabase() {
	log "Verify exitsAllExpectedDatabase"

	for db in "${ALL_DATABASE[@]}"; do
		if ! psql -tAc "select 1 from pg_database where datname = '$db'" | grep -q 1; then
			log "ERROR - Not found dabatase '$db'"
			return 1
		fi
	done
}

function scaleAllTo1() {
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		log "Scale to 1 $dep"
		kubectl_cmd scale --replicas 1 deployment "$dep" || { log "ERROR - Unable to scale project $dep"; return 1; }
	done
}

function scaleAllTo0() {
	local cond="0/0"
	local seconds="5"
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		log "Scale to 0 $dep"
		kubectl_cmd scale --replicas 0 deployment "$dep" || { log "ERROR - Unable to scale project $dep"; return 1; }
	done
}

function waitAllScaleTo1() {
	local cond="1/1"
	local seconds="5"
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		waitDeploymentStatus  "$dep" "$cond" "$seconds" || { log "ERROR - Unable to wait project $dep"; return 1; }
	done
}

function waitAllScaleTo0() {
	local cond="0/0"
	local seconds="5"
	for dep in "${ALL_DEPLOYMENTS[@]}"; do
		waitDeploymentStatus  "$dep" "$cond" "$seconds" || { log "ERROR - Unable to wait project $dep"; return 1; }
	done
}

function waitDeploymentStatus() {
	local dep_name="$1"
	local dep_exp_status="$2"
	local wait_seconds="$3"

	local dep_status=0
	while [[ "$dep_status" != 1 ]]; do
		local dep_output="$( kubectl_cmd get deployment --no-headers "$dep_name" )"
		if [[ $? != 0 ]]; then
			log "ERROR - unable to get deployment $dep_name"
			return 1
		fi
		dep_status="$(echo "$dep_output" | grep -w $dep_exp_status | wc -l )"
		if [[ $dep_status != 1 ]]; then
			log "Waiting deployment $dep_name"
			sleep "$wait_seconds"
		fi
	done
}

function kubectl_cmd() {
	"${KUBECTL_CMD[@]}" "$@"
}

function log() {
	>&2 echo "$(date -Iseconds) - ${LOG_CONTEXT-"NC"} " "$@"
}

main
