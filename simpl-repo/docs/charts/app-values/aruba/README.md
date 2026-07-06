# Demo

## Installing an agent as namespace

### _new_consumer

```sh
  cd _new_consumer
  helm dependency update
  k create ns iaa-demo-new-consumer
  helm install iaa-demo-new-consumer . --namespace iaa-demo-new-consumer --values ./values.yaml
```

### Setup Authority

```sh
  curl --location 'https://t1.iaa-ds-authority.demo.aruba-simpl.cloud/auth/realms/authority/protocol/openid-connect/token' \
    --header 'Content-Type: application/x-www-form-urlencoded' \
    --data-urlencode 'grant_type=password' \
    --data-urlencode 'client_id=frontend-cli' \
    --data-urlencode 'password=password' \
    --data-urlencode 'username=e.j' | jq
```

```sh
  k -n iaa-demo-authority  port-forward svc/identity-provider 8080:8080

  curl --location 'http://localhost:8080/participant/initialize' \
    -X POST \
    --header 'Authorization: Bearer <JWT>' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "distinguishedName" : {
            "commonName" : "t2.iaa-ds-authority.demo.aruba-simpl.cloud",
            "organization": "Aruba SpA",
            "organizationalUnit" : "SIMPL Open",
            "country": "IT"
        },
        "identityAttributes" : []
    }'
```

## New cluster on KAAS

### Nginx ingress controller

```sh
  helm upgrade --install ingress-nginx ingress-nginx \
    --repo https://kubernetes.github.io/ingress-nginx \
    --namespace ingress-nginx \
    --create-namespace \
    --values ./nginx-ingress-controller.yaml
```

### CertManager

```sh
  helm repo add jetstack https://charts.jetstack.io --force-update

  helm upgrade --install cert-manager jetstack/cert-manager \
    --create-namespace \
    --namespace cert-manager \
    --version v1.16.2 \
    --set crds.enabled=true
```

#### Cluster Issuer

```sh
  kubectl apply -f ./cluster-issuer.yaml
```
