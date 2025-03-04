function init_vault_approle_and_kv() {

NAMESPACE="$1"

APPROLE_PATH="$NAMESPACE.authenticationprovider"
POLICY_NAME="$NAMESPACE.authenticationprovider"

SECRET_PATH="iaa.$NAMESPACE"
SECRET_SUBPATH="authenticationprovider"

vault secrets enable -path="$SECRET_PATH" kv
vault secrets list

vault auth enable -path="$APPROLE_PATH" approle
vault auth list

cat <<EOF > /tmp/my-policy.hcl
path "$SECRET_PATH/data/$SECRET_SUBPATH/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
path "$SECRET_PATH/metadata/$SECRET_SUBPATH/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
EOF

vault policy write $POLICY_NAME /tmp/my-policy.hcl
vault policy list

vault write auth/$APPROLE_PATH/role/myapp-role \
    token_policies="$POLICY_NAME" \
    token_ttl=1h \
    token_max_ttl=4h

ROLE_ID=$(vault read -field=role_id auth/$APPROLE_PATH/role/myapp-role/role-id)
SECRET_ID=$(vault write -f -field=secret_id auth/$APPROLE_PATH/role/myapp-role/secret-id)

vault write auth/$APPROLE_PATH/login role_id="$ROLE_ID" secret_id="$SECRET_ID"

echo "Namespace: $NAMESPACE"
echo "ROLE_ID: $ROLE_ID"
echo "SECRET_ID: $SECRET_ID"
}
