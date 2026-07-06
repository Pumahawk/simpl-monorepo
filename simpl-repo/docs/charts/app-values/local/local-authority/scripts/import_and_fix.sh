#!/bin/sh
set -e

echo "🔄 Using HTTP alpinelinux repository (⚠️ Use only in local environment)"
sed -i 's|https://dl-cdn.alpinelinux.org|http://dl-cdn.alpinelinux.org|g' /etc/apk/repositories
echo "⚠️ HTTPS disabled successfully!"

echo "Installing packages..."
apk add --no-cache curl jq python3

echo "Parsing REALMS_JSON..."
count=$(echo "$REALMS_JSON" | jq length)

for i in $(seq 0 $((count - 1))); do
  NAME=$(echo "$REALMS_JSON" | jq -r ".[$i].name")
  URI=$(echo "$REALMS_JSON" | jq -r ".[$i].uri")
  REDIRECT_URIS=$(echo "$REALMS_JSON" | jq -r ".[$i].redirectUris")

  echo "Downloading $NAME from $URI"

  curl -s "$URI" \
    | sed "s|\${REDIRECT_URIS}|$REDIRECT_URIS|g" \
    | jq --arg realmName "$NAME" '.realm = $realmName' \
    > /config/$NAME.json

  echo "Rewriting IDs in $NAME.json"
  python3 /scripts/rewrite_ids.py /config/$NAME.json
done
