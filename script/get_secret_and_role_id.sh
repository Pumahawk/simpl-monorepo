function get_secret_and_role_id() {
	NAMESPACE="$1"
	SPATH="$2"
	APPROLE_PATH="$NAMESPACE.authenticationprovider"
	ROLE_ID=$(vault read -field=role_id auth/$APPROLE_PATH/role/$SPATH/role-id)
	SECRET_ID=$(vault write -f -field=secret_id auth/$APPROLE_PATH/role/$SPATH/secret-id)
	echo
	echo "Namespace: $NAMESPACE"
	echo "ROLE_ID: $ROLE_ID"
	echo "SECRET_ID: $SECRET_ID"
	echo
}
