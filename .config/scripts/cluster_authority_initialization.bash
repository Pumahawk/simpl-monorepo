#!/usr/bin/env sh

wait_all_services_authority_up() {
  local services='authentication-provider-authority:8105
identity-provider-authority:8103
onboarding-authority:8104
security-attributes-provider-authority:8102
tier1-gateway-authority:8100
tier2-gateway-authority:8142
users-roles-authority:8101'

  max_n=60
  n=1
  for line in $services; do
    echo "Check service $line"
    until status="$(curl -s http://$line/actuator/health | grep UP)" || [ "$n" -ge "$max_n" ]; do
       echo "($n/$max_n) status=[$status] Wait service $line become healthy."
       echo "Wait 15"
       sleep 15
       n=$(($n + 1))
    done
    if [ $n == $max_n ]; then echo "Unable to check health of service $line"; return 1; fi
  done
  echo "Check microservices done"
}

wait_all_services_authority_up

MICROSERVICE_AUTHENTICATION_PROVIDER_INTERNAL_URL=authentication-provider-authority:8105
MICROSERVICE_IDENTITY_PROVIDER_INTERNAL_URL=identity-provider-authority:8103
CSR_CN=tier2-gateway-authority

CURL_W_OUT=$(mktemp)

echo "Start initialization process"

########################################
# Check authority already initializated
########################################

if ! KEYPAIR_RESPONSE="$(curl -s -w "%{http_code}" -X HEAD "$MICROSERVICE_AUTHENTICATION_PROVIDER_INTERNAL_URL/tier1/v2/keypairs/active" -o /tmp/t > "$CURL_W_OUT" && cat /tmp/t
)"; then
  CURL_EXIT_CODE=$?
  echo "Error keypairs active call. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

CURL_RESPONSE_CODE="$(cat "$CURL_W_OUT")"

echo "
==========================
Description: Check authority already initializated
Http response code: [$CURL_RESPONSE_CODE]
Response Body:
$KEYPAIR_RESPONSE
==========================
"

if [[ "$CURL_RESPONSE_CODE" -ge  200 && "$CURL_RESPONSE_CODE" -lt 300 ]]; then
  echo "Authority already initialized"
  return 0
fi

########################################
# Generating keypair...
########################################

if ! KEYPAIR_RESPONSE="$(curl -s -w "%{http_code}" -X POST "$MICROSERVICE_AUTHENTICATION_PROVIDER_INTERNAL_URL/tier1/v2/keypairs" -o /tmp/t > "$CURL_W_OUT" \
  --header 'Content-Type: application/json' \
  --data-raw '{"name":"initialization-authority"}' && cat /tmp/t
)"; then
  CURL_EXIT_CODE=$?
  echo "Error keypairs call. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

CURL_RESPONSE_CODE="$(cat "$CURL_W_OUT")"

echo "
==========================
Description: Generating keypair
Http response code: [$CURL_RESPONSE_CODE]
Response Body:
$KEYPAIR_RESPONSE
==========================
"

if [[ "$CURL_RESPONSE_CODE" -lt 200 || "$CURL_RESPONSE_CODE" -ge 300 ]]; then
  exit 1
fi

KEYPAIR_ID="$(echo "$KEYPAIR_RESPONSE" | sed -n '/id"/{s/.*{"id":"\([^"]*\)".*/\1/;p}')"

if [ -z "$KEYPAIR_ID" ]; then
  echo "Error retrieving keypair id. keypair_id=[$KEYPAIR_ID]"
  exit 1
fi

echo "KEYPAIR_ID=[$KEYPAIR_ID]"

########################################
# CSR request
########################################

if ! CSR_RESPONSE="$(curl -s -w "%{http_code}" -X POST \
"$MICROSERVICE_AUTHENTICATION_PROVIDER_INTERNAL_URL/tier1/v2/keypairs/$KEYPAIR_ID/csr" -o /tmp/t > "$CURL_W_OUT" \
--header 'Content-Type: application/json' \
--data-raw '{
  "commonName": "'"$CSR_CN"'",
  "country": "'"$CSR_CN"'",
  "organization": "'"$CSR_CN"'",
  "organizationalUnit": "'"$CSR_CN"'"
}' && cat /tmp/t )"; then
  CURL_EXIT_CODE=$?
  echo "Error csr call. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

echo "$CSR_RESPONSE" > /tmp/csr.json

CURL_RESPONSE_CODE="$(cat "$CURL_W_OUT")"

echo "
==========================
Description: CSR request
Http response code: [$CURL_RESPONSE_CODE]
Response Body:
$CSR_RESPONSE
==========================
"

if [[ "$CURL_RESPONSE_CODE" -lt 200 || "$CURL_RESPONSE_CODE" -ge 300 ]]; then
  exit 1
fi

if [ ! -s /tmp/csr.json ]; then
  echo "CSR file not created or empty"
  exit 1
fi

########################################
# Creating Authority participant
########################################

if ! PARTICIPANT_RESPONSE="$(curl -s -w "%{http_code}" -X POST \
"$MICROSERVICE_IDENTITY_PROVIDER_INTERNAL_URL/tier1/v2/participants" -o /tmp/t > "$CURL_W_OUT" \
--header 'Content-Type: application/json' \
--data-raw '{
  "organization": "local-authority",
  "participantType": "GOVERNANCE_AUTHORITY",
  "isAuthority": true
}' && cat /tmp/t)"; then
  CURL_EXIT_CODE=$?
  echo "Error participant creation. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

CURL_RESPONSE_CODE="$(cat "$CURL_W_OUT")"

echo "
==========================
Description: Create participant
Http response code: [$CURL_RESPONSE_CODE]
Response Body:
$PARTICIPANT_RESPONSE
==========================
"

if [[ "$CURL_RESPONSE_CODE" -lt 200 || "$CURL_RESPONSE_CODE" -ge 300 ]]; then
  exit 1
fi

PARTICIPANT_ID="$(echo "$PARTICIPANT_RESPONSE" | sed 's/.*"id":"\([^"]*\)".*/\1/')"

if [ -z "$PARTICIPANT_ID" ]; then
  echo "Error retrieving participant id"
  exit 1
fi

echo "PARTICIPANT_ID=[$PARTICIPANT_ID]"

########################################
# Uploading CSR
########################################

if ! CSR_UPLOAD_RESPONSE="$(curl -s -w "%{http_code}" -X PUT \
"$MICROSERVICE_IDENTITY_PROVIDER_INTERNAL_URL/tier1/v2/participants/$PARTICIPANT_ID/csr" -o /tmp/t > "$CURL_W_OUT" \
--header 'Content-Type: application/json' \
-d @/tmp/csr.json && cat /tmp/t)"; then
  CURL_EXIT_CODE=$?
  echo "Error uploading CSR. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

CURL_RESPONSE_CODE="$(cat "$CURL_W_OUT")"

echo "
==========================
Description: Upload CSR
Http response code: [$CURL_RESPONSE_CODE]
Response Body:
$CSR_UPLOAD_RESPONSE
==========================
"

if [[ "$CURL_RESPONSE_CODE" -lt 200 || "$CURL_RESPONSE_CODE" -ge 300 ]]; then
  exit 1
fi

########################################
# Create credentials
########################################

if ! CERT_RESPONSE="$(curl -s -w "%{http_code}" -X POST \
"$MICROSERVICE_IDENTITY_PROVIDER_INTERNAL_URL/tier1/v2/participants/$PARTICIPANT_ID/credentials" -o /tmp/t > "$CURL_W_OUT" \
--header 'Content-Type: application/json' \
--data-raw '{"reason":"initialization-authority"}' && cat /tmp/t)"; then
  CURL_EXIT_CODE=$?
  echo "Error creating credentials. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

echo "$CERT_RESPONSE" > /tmp/cert.json

CURL_RESPONSE_CODE="$(cat "$CURL_W_OUT")"

echo "
==========================
Description: Create credentials
Http response code: [$CURL_RESPONSE_CODE]
Response Body:
$CERT_RESPONSE
==========================
"

if [[ "$CURL_RESPONSE_CODE" -lt 200 || "$CURL_RESPONSE_CODE" -ge 300 ]]; then
  exit 1
fi

if [ ! -s /tmp/cert.json ]; then
  echo "/tmp/cert.json not created or empty"
  exit 1
fi

CONTENT="$(tr -d '\r' < /tmp/cert.json | sed -n '/"content":/{s/.*"content":"\([^"]*\)".*/\1/;p}')"

if [ -z "$CONTENT" ]; then
  echo "Error extracting certificate content"
  exit 1
fi

########################################
# Uploading credentials
########################################

if ! FINAL_RESPONSE="$(curl -s -w "%{http_code}" -X POST \
"$MICROSERVICE_AUTHENTICATION_PROVIDER_INTERNAL_URL/tier1/v2/credentials" -o /tmp/t > "$CURL_W_OUT" \
--header 'Content-Type: application/json' \
--data-raw "$(cat << EOF
{
  "reason": "initialization-authority",
  "content": "$CONTENT"
}
EOF
)" &&  cat /tmp/t)"; then
  CURL_EXIT_CODE=$?
  echo "Error uploading credentials. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

CURL_RESPONSE_CODE="$(cat "$CURL_W_OUT")"

echo "
==========================
Description: Upload credentials
Http response code: [$CURL_RESPONSE_CODE]
Response Body:
$FINAL_RESPONSE
==========================
"

if [[ "$CURL_RESPONSE_CODE" -lt 200 || "$CURL_RESPONSE_CODE" -ge 300 ]]; then
  exit 1
fi

echo "Script completed successfully"
