## `keycloak-authenticator` local development

This folder contains a Dockerfile that can be use to build a `init-spi-local` embedding a local build of the `keycloak-authenticator` SPI.

This can be useful to easily debug a local build without the need to build the jar on the CI/CD infrastructure.

⚠️ This was tested using Rancher Desktop with `dockerd` (Moby) container runtime.

### Deployment

1. Change the following section of [../local-authority/values.yaml](../local-authority/values.yaml) to use `init-spi-local` image

```
keycloak:
  initContainers:
    - name: init-spi
      # Uncomment to allow running local builds of keycloak-authenticator, see /app-values/local/keycloak-authenticator/README.md
      image: init-spi-local:latest
      imagePullPolicy: Never
      #image: curlimages/curl
```

⚠️ Avoid committing this uncommented section of code since it is meant for local tests only, and might break other developer's keycloak deployment!

2. Manually build the `keycloak-authenticator` library from source code (from your local `keycloak-authenticator` project directory)
    * `mvn clean package`

3. Copy the jar from `target/deploy` folder into this folder

4. Manually build the docker image from this folder

    * `docker build -t init-spi-local .`

5. Restart the keycloak stateful set to ensure the new image is pulled from the local daemon

    * `kubectl rollout restart statefulset local-authority-keycloak -n local-authority`

### Debugging

Just port forward to the debug port of the keycloak container:

`kubectl port-forward local-authority-keycloak-0 5005:5005 -n local-authority`

When attempting to retrieve a token, the debugger should stop to any breakpoint inside `keycloak.authenticator.spi.AttributeAuthenticator.authenticate()` and subsequently invoked methods.
