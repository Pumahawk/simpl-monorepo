# Update Repository
helm repo update

# Add repository
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo add simpl-cloud-gateway-charts https://code.europa.eu/api/v4/projects/772/packages/helm/stable
helm repo add users-roles-charts https://code.europa.eu/api/v4/projects/771/packages/helm/stable
helm repo add simpl-fe-charts https://code.europa.eu/api/v4/projects/769/packages/helm/stable

# Install redis
helm install redis bitnami/redis --version 19.6.0 --values redis/values.yaml

# Install postgresql
helm install postgresql bitnami/postgresql \
--version 15.2.5 \
--values postgres/values-participant.yaml

# Install keycloak
helm install keycloak bitnami/keycloak \
--version 22.2.5 \
--values keycloak/values-participant.yaml

# Install simpl-cloud-gateway
helm install simpl-cloud-gateway simpl-cloud-gateway-charts/simpl-cloud-gateway \
--version 0.6.1 \
--values simpl-cloud-gateway/values-participant.yaml

# Install users-roles
helm install users-roles users-roles-charts/users-roles \
--version 0.6.4 \
--values users-roles/values-participant.yaml

# Install simpl-fe
helm install simpl-fe simpl-fe-charts/simpl-fe \
--version 0.6.1 \
--values simpl-fe/values-participant.yaml


