# Update chart dependencies

The cronjob template provided in the `cronjob.yaml` file defines a Kubernetes CronJob that runs the script in the [update-chart-dependencies.bash](./update-chart-dependencies.bash) file. The job keeps the dependencies of the `Chart.yaml` for each agent in the `iaa-dsdev` environment in sync with the version published on the `develop` branch of the microservices.

## Bulding and publishing the image

⚠️ Currently the image is pushed manually to `code.europa.eu:4567/simpl/simpl-open/development/iaa/charts:update-chart-dependencies-v1`. Since the image is periodically deleted by a Gitlab cleanup job, this procedure must be performed manually each time the cronjob fails due to image pull failure.

```
cd update-chart-dependencies

docker build . -f .\Dockerfile -t code.europa.eu:4567/simpl/simpl-open/development/iaa/charts:update-chart-dependencies-v1

docker login code.europa.eu:4567 -u <USERNAME> -p <TOKEN>

docker push code.europa.eu:4567/simpl/simpl-open/development/iaa/charts:update-chart-dependencies-v1

```

**Note:** the token should have the following scope: `read_registry`, `write_registry`, `api`