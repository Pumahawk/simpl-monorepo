#!/usr/bin/env bash

set -o pipefail

wait_all_services_authority_up() {
  local services=()
  while read line; do services+=($line); done < <(yq '.services | to_entries | .[].key' .config/docker/simpl-services/docker-compose.yaml  | grep authority )
  echo "Start waiting for services: ${services[@]}"
  max_n=60
  n=1
  for line in "${services[@]}"; do
    until status="$(mise run simpl-services:compose ps -a "$line" --format json | jq -r .Health)" && [ "$status" == "healthy" ] || [ $n == $max_n ]; do
       if [ -z "$status" ]; then
         echo "Invalid health service=[$line], status=[$status]"
         return 1;
       fi
       echo "($n/$max_n) Wait service $line become healthy."
       echo "Wait 20s"
       sleep 20
       n=$(($n + 1))
    done
    if [ $n == $max_n ]; then echo "Unable to check health of service $line"; return 1; fi
  done
}

wait_all_services_authority_up || { echo "Unable to initializate authority"; exit 1; }

export AUTHORITY_AUTH_PROVIDER=localhost:8080
export AUTHORITY_IDENTITY_PROVIDER=localhost:8090

CURL_W_OUT=$(mktemp)

########################################
# Generating keypair...
########################################

if ! KEYPAIR_RESPONSE="$(curl -s -w "%output{$CURL_W_OUT}%{http_code}" -X POST "$AUTHORITY_AUTH_PROVIDER/tier1/v2/keypairs" \
  --header 'Content-Type: application/json' \
  --data-raw '{"name":"initialization-authority"}'
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

KEYPAIR_ID="$(sed -n '/id"/{s/.*{"id":"\([^"]*\)".*/\1/;p}' <<<"$KEYPAIR_RESPONSE")"

if [ -z "$KEYPAIR_ID" ]; then
  echo "Error retrieving keypair id. keypair_id=[$KEYPAIR_ID]"
  exit 1
fi

echo "KEYPAIR_ID=[$KEYPAIR_ID]"

########################################
# CSR request
########################################

if ! CSR_RESPONSE="$(curl -s -w "%output{$CURL_W_OUT}%{http_code}" -X POST \
"$AUTHORITY_AUTH_PROVIDER/tier1/v2/keypairs/$KEYPAIR_ID/csr" \
--header 'Content-Type: application/json' \
--data-raw '{
  "commonName": "tier2-gateway-authority",
  "country": "tier2-gateway-authority",
  "organization": "tier2-gateway-authority",
  "organizationalUnit": "tier2-gateway-authority"
}')"; then
  CURL_EXIT_CODE=$?
  echo "Error csr call. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

echo "$CSR_RESPONSE" > csr.json

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

if [ ! -s csr.json ]; then
  echo "CSR file not created or empty"
  exit 1
fi

########################################
# Creating Authority participant
########################################

if ! PARTICIPANT_RESPONSE="$(curl -s -w "%output{$CURL_W_OUT}%{http_code}" -X POST \
"$AUTHORITY_IDENTITY_PROVIDER/tier1/v2/participants" \
--header 'Content-Type: application/json' \
--data-raw '{
  "organization": "local-authority",
  "participantType": "GOVERNANCE_AUTHORITY",
  "isAuthority": true
}')"; then
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

PARTICIPANT_ID="$(sed 's/.*"id":"\([^"]*\)".*/\1/' <<<"$PARTICIPANT_RESPONSE")"

if [ -z "$PARTICIPANT_ID" ]; then
  echo "Error retrieving participant id"
  exit 1
fi

echo "PARTICIPANT_ID=[$PARTICIPANT_ID]"

########################################
# Uploading CSR
########################################

if ! CSR_UPLOAD_RESPONSE="$(curl -s -w "%output{$CURL_W_OUT}%{http_code}" -X PUT \
"$AUTHORITY_IDENTITY_PROVIDER/tier1/v2/participants/$PARTICIPANT_ID/csr" \
--header 'Content-Type: application/json' \
-d @csr.json)"; then
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

if ! CERT_RESPONSE="$(curl -s -w "%output{$CURL_W_OUT}%{http_code}" -X POST \
"$AUTHORITY_IDENTITY_PROVIDER/tier1/v2/participants/$PARTICIPANT_ID/credentials" \
--header 'Content-Type: application/json' \
--data-raw '{"reason":"initialization-authority"}')"; then
  CURL_EXIT_CODE=$?
  echo "Error creating credentials. curl_exit_code=[$CURL_EXIT_CODE]"
  exit 1
fi

echo "$CERT_RESPONSE" > cert.json

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

if [ ! -s cert.json ]; then
  echo "cert.json not created or empty"
  exit 1
fi

CONTENT="$(tr -d '\r' < cert.json | sed -n '/"content":/{s/.*"content":"\([^"]*\)".*/\1/;p}')"

if [ -z "$CONTENT" ]; then
  echo "Error extracting certificate content"
  exit 1
fi

########################################
# Uploading credentials
########################################

if ! FINAL_RESPONSE="$(curl -s -w "%output{$CURL_W_OUT}%{http_code}" -X POST \
"$AUTHORITY_AUTH_PROVIDER/tier1/v2/credentials" \
--header 'Content-Type: application/json' \
--data-raw "$(cat << EOF
{
  "reason": "initialization-authority",
  "content": "$CONTENT"
}
EOF
)")"; then
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

echo "✅ Script completed successfully"
