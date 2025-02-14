#/bin/bash

pushd common; git fetch; git checkout origin/develop; popd;

mvn clean install -Dmaven.test.skip -f common/simpl-api-iaa

cp -v -t agent-service/openapi/ common/simpl-api-iaa/src/main/resources/static/openapi/authenticationprovider-v1.yaml
cp -v -t identity-provider/openapi/ common/simpl-api-iaa/src/main/resources/static/openapi/identityprovider-v1.yaml
cp -v -t onboarding/openapi/ common/simpl-api-iaa/src/main/resources/static/openapi/onboarding-v1.yaml
cp -v -t security-attributes-provider/openapi/ common/simpl-api-iaa/src/main/resources/static/openapi/securityattributesprovider-v1.yaml
cp -v -t users-roles/openapi/ common/simpl-api-iaa/src/main/resources/static/openapi/usersroles-v1.yaml

pushd agent-service; git fetch; git checkout -b feature/openapi-align origin/develop; git add -A; git commit -m "Update openapi"; git push -o ci -u origin feature/openapi-align; popd;
pushd identity-provider; git fetch; git checkout -b feature/openapi-align origin/develop; git add -A; git commit -m "Update openapi"; git push -o ci -u origin feature/openapi-align; popd;
pushd onboarding; git fetch; git checkout -b feature/openapi-align origin/develop; git add -A; git commit -m "Update openapi"; git push -o ci -u origin feature/openapi-align; popd;
pushd security-attributes-provider; git fetch; git checkout -b feature/openapi-align origin/develop; git add -A; git commit -m "Update openapi"; git push -o ci -u origin feature/openapi-align; popd;
pushd users-roles; git fetch; git checkout -b feature/openapi-align origin/develop; git add -A; git commit -m "Update openapi"; git push -o ci -u origin feature/openapi-align; popd;

