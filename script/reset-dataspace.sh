#! /bin/bash

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
	( dropdb authenticationprovider && createdb -O authenticationprovider authenticationprovider ) || { log "Unable drop or create db authenticationprovider"; return 1; }
	( dropdb ejbca && createdb -O ejbca ejbca ) || { log "Unable drop or create db ejbca"; return 1; }
	( dropdb identityprovider && createdb -O identityprovider identityprovider ) || { log "Unable drop or create db identityprovider"; return 1; }
	( dropdb keycloak && createdb -O keycloak keycloak ) || { log "Unable drop or create db keycloak"; return 1; }
	( dropdb onboarding && createdb -O onboarding onboarding ) || { log "Unable drop or create db onboarding"; return 1; }
	( dropdb securityattributesprovider && createdb -O securityattributesprovider securityattributesprovider ) || { log "Unable drop or create db securityattributesprovider"; return 1; }
	( dropdb usersroles && createdb -O usersroles usersroles ) || { log "Unable drop or create db usersroles"; return 1; }
	log "Success Clean Database"
}

# Delete all EJBCA secrets
function deleteEJBCASecrets() {
	log "Delete EJBCA secrets"
	if ! kubectl delete secret ejbca-rest-api-secret; then
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
	&& verifyEJBCAStatus \
	&& exitsAllExpectedDatabase
}

function log() {
	>&2 echo "$(date -Iseconds) - ${LOG_CONTEXT-"NC"} " "$@"
}


# Verify is current namespace is an Authority dataspace
function isAuthorityDataspace() {
	log "Verify isAuthorityDataspace"
	log "Retrieve current namespace from onboardin deployment"
	local namespace="$(kubectl get deployment -ojsonpath='{.items[].metadata.namespace}' | head -n1)"
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
	if kubectl get deployment.apps/ejbca-community-helm; then
		log "OK - deployment.apps/ejbca-community-helm"
		return 0
	else
		log "ERROR - deployment ejbca-community-helm not found"
		return 1
	fi
}

function exitsAllExpectedDatabase() {
	log "Verify exitsAllExpectedDatabase"
	log "Find database:"
	log "- authenticationprovider"
	log "- ejbca"
	log "- identityprovider"
	log "- keycloak"
	log "- onboarding"
	log "- securityattributesprovider"
	log "- usersroles"

 	if ! psql -tAc "select 1 from pg_database where datname = 'authenticationprovider'" | grep -q 1; then
		log "ERROR - Not found dabatase 'authenticationprovider'"
		return 1
	fi
 	if ! psql -tAc "select 1 from pg_database where datname = 'ejbca'" | grep -q 1; then
		log "ERROR - Not found dabatase 'ejbca'"
		return 1
	fi
 	if ! psql -tAc "select 1 from pg_database where datname = 'identityprovider'" | grep -q 1; then
		log "ERROR - Not found dabatase 'identityprovider'"
		return 1
	fi
 	if ! psql -tAc "select 1 from pg_database where datname = 'keycloak'" | grep -q 1; then
		log "ERROR - Not found dabatase 'keycloak'"
		return 1
	fi
 	if ! psql -tAc "select 1 from pg_database where datname = 'onboarding'" | grep -q 1; then
		log "ERROR - Not found dabatase 'onboarding'"
		return 1
	fi
 	if ! psql -tAc "select 1 from pg_database where datname = 'securityattributesprovider'" | grep -q 1; then
		log "ERROR - Not found dabatase 'securityattributesprovider'"
		return 1
	fi
 	if ! psql -tAc "select 1 from pg_database where datname = 'usersroles'" | grep -q 1; then
		log "ERROR - Not found dabatase 'usersroles'"
		return 1
	fi
}

function scaleAllTo1() {
	log "Scale to 1 authentication-provider"
	kubectl scale --replicas 1 deployment "authentication-provider" || { log "ERROR - Unable to scale project authentication-provider"; return 1; }
	log "Scale to 1 ejbca-community-helm"
	kubectl scale --replicas 1 deployment "ejbca-community-helm" || { log "ERROR - Unable to scale project ejbca-community-helm"; return 1; }
	log "Scale to 1 identity-provider"
	kubectl scale --replicas 1 deployment "identity-provider" || { log "ERROR - Unable to scale project identity-provider"; return 1; }
	log "Scale to 1 onboarding"
	kubectl scale --replicas 1 deployment "onboarding" || { log "ERROR - Unable to scale project onboarding"; return 1; }
	log "Scale to 1 security-attributes-provider"
	kubectl scale --replicas 1 deployment "security-attributes-provider" || { log "ERROR - Unable to scale project security-attributes-provider"; return 1; }
	log "Scale to 1 tier1-gateway"
	kubectl scale --replicas 1 deployment "tier1-gateway" || { log "ERROR - Unable to scale project tier1-gateway"; return 1; }
	log "Scale to 1 tier2-gateway"
	kubectl scale --replicas 1 deployment "tier2-gateway" || { log "ERROR - Unable to scale project tier2-gateway"; return 1; }
	log "Scale to 1 users-roles"
	kubectl scale --replicas 1 deployment "users-roles" || { log "ERROR - Unable to scale project users-roles"; return 1; }
}

function scaleAllTo0() {
	log "Scale to 0 authentication-provider"
	kubectl scale --replicas 0 deployment "authentication-provider" || { log "ERROR - Unable to scale project authentication-provider"; return 1; }
	log "Scale to 0 ejbca-community-helm"
	kubectl scale --replicas 0 deployment "ejbca-community-helm" || { log "ERROR - Unable to scale project ejbca-community-helm"; return 1; }
	log "Scale to 0 identity-provider"
	kubectl scale --replicas 0 deployment "identity-provider" || { log "ERROR - Unable to scale project identity-provider"; return 1; }
	log "Scale to 0 onboarding"
	kubectl scale --replicas 0 deployment "onboarding" || { log "ERROR - Unable to scale project onboarding"; return 1; }
	log "Scale to 0 security-attributes-provider"
	kubectl scale --replicas 0 deployment "security-attributes-provider" || { log "ERROR - Unable to scale project security-attributes-provider"; return 1; }
	log "Scale to 0 tier1-gateway"
	kubectl scale --replicas 0 deployment "tier1-gateway" || { log "ERROR - Unable to scale project tier1-gateway"; return 1; }
	log "Scale to 0 tier2-gateway"
	kubectl scale --replicas 0 deployment "tier2-gateway" || { log "ERROR - Unable to scale project tier2-gateway"; return 1; }
	log "Scale to 0 users-roles"
	kubectl scale --replicas 0 deployment "users-roles" || { log "ERROR - Unable to scale project users-roles"; return 1; }
}

function waitAllScaleTo1() {
	local cond="1/1"
	local seconds="5"
	waitDeploymentStatus  "authentication-provider" "$cond" "$seconds" || { log "ERROR - Unable to wait project authentication-provider"; return 1; }
	waitDeploymentStatus  "ejbca-community-helm" "$cond" "$seconds" || { log "ERROR - Unable to wait project ejbca-community-helm"; return 1; }
	waitDeploymentStatus  "identity-provider" "$cond" "$seconds" || { log "ERROR - Unable to wait project identity-provider"; return 1; }
	waitDeploymentStatus  "onboarding" "$cond" "$seconds" || { log "ERROR - Unable to wait project onboarding"; return 1; }
	waitDeploymentStatus  "security-attributes-provider" "$cond" "$seconds" || { log "ERROR - Unable to wait project security-attributes-provider"; return 1; }
	waitDeploymentStatus  "tier1-gateway" "$cond" "$seconds" || { log "ERROR - Unable to wait project tier1-gateway"; return 1; }
	waitDeploymentStatus  "tier2-gateway" "$cond" "$seconds" || { log "ERROR - Unable to wait project tier2-gateway"; return 1; }
	waitDeploymentStatus  "users-roles" "$cond" "$seconds" || { log "ERROR - Unable to wait project users-roles"; return 1; }
}

function waitAllScaleTo0() {
	local cond="0/0"
	local seconds="5"
	waitDeploymentStatus  "authentication-provider" "$cond" "$seconds" || { log "ERROR - Unable to wait project authentication-provider"; return 1; }
	waitDeploymentStatus  "ejbca-community-helm" "$cond" "$seconds" || { log "ERROR - Unable to wait project ejbca-community-helm"; return 1; }
	waitDeploymentStatus  "identity-provider" "$cond" "$seconds" || { log "ERROR - Unable to wait project identity-provider"; return 1; }
	waitDeploymentStatus  "onboarding" "$cond" "$seconds" || { log "ERROR - Unable to wait project onboarding"; return 1; }
	waitDeploymentStatus  "security-attributes-provider" "$cond" "$seconds" || { log "ERROR - Unable to wait project security-attributes-provider"; return 1; }
	waitDeploymentStatus  "tier1-gateway" "$cond" "$seconds" || { log "ERROR - Unable to wait project tier1-gateway"; return 1; }
	waitDeploymentStatus  "tier2-gateway" "$cond" "$seconds" || { log "ERROR - Unable to wait project tier2-gateway"; return 1; }
	waitDeploymentStatus  "users-roles" "$cond" "$seconds" || { log "ERROR - Unable to wait project users-roles"; return 1; }
}

function waitDeploymentStatus() {
	local dep_name="$1"
	local dep_exp_status="$2"
	local wait_seconds="$3"

	local dep_status=0
	while [[ "$dep_status" != 1 ]]; do
		local dep_output="$( kubectl get deployment --no-headers "$dep_name" )"
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

main
