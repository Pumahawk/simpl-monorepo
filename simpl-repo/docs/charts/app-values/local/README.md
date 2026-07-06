# SIMPL IAA Chart for local development

## Prerequisites

The following guide assumes that the local development machine is already configured to :

* run a local K8S cluster (via Rancher Desktop, or minikube, k3s etc.)
* deploy Helm charts to the local K8S cluster
* run docker compose files (optional, Kafka-only)

**Note**: the procedure has currently been tested only on Windows with Rancher Desktop. Currently, some Windows-only features are used (such as `host.docker.internal`), and this will need to be addressed to enable development on other platforms appropriately.

## Usage

**Note**: see [Performance optimization for local development](#Performance-optimization-for-local-development) for tips about how to obtain a more fluent experience on local machine

## Run infrastructure on local K8S cluster

**Note**: if there is the need to change some EJBCA configuration, refer to [ejbca-init/README.md](ejbca-init/README.md)

1. `cd local-authority`

2. `helm install local-authority --create-namespace -n local-authority .`

This will launch all dependencies, such as PostgreSQL, Keycloak, EJBCA etc.

Refer to the [architectural schema](./Architecture.WIP.svg) for an overview of the components involved

## Import IntelliJ IDEA Run configurations

Ready-to-use run configurations are available for IntelliJ idea in the `./intellij/.run` folder.

Refer to [./intellij/README.md](./intellij/README.md) for usage instructions and also an outline of required configurations if using another IDE.


## Import IntelliJ IDEA Run configurations

Ready-to-use run configurations are available for IntelliJ idea in the `./intellij/.run` folder.

Refer to [./intellij/README.md](./intellij/README.md) for usage instructions and also an outline of required configurations if using another IDE.

# Performance optimization for local development

## Windows

On windows, it can be useful to tune WSL resource in order to reduce eccessive memory allocation.

### Configure WSL resources

Create or update `%homepath%/.wslconfig` with the configuration suitable for your machine.

Example on this machine:

```
Windows 11 Enterprise
RAM 32 GB
13th Gen Intel(R) Core(TM) i7-1355U 
NumberOfCores 10
NumberOfLogicalProcessors 12
```

`%homepath%/.wslconfig`
```
[wsl2]
memory=3GB
processors=2
swap=2GB
```

More info about WSL configuration can be found [here](https://learn.microsoft.com/en-us/windows/wsl/wsl-config).

## JDK

It can be useful to switch to a JVM implementation with lighter memory footprint, such as [OpenJ9](https://eclipse.dev/openj9/)

This can be done by adopting the [Semeru](https://www.ibm.com/support/pages/semeru-runtimes-getting-started) JDK distribution, that ships with OpenJ9

### IDE Configuration

Configure the IDE to run the Semeru distribution.

In IntelliJ IDEA, go to `File > Project Structure` and as `SDK` choose `IBM Semeru (AdoptOpenJDK OpenJ9)` or similar

TODO Add JAVA_HOME and PATH configuration

## IntelliJ IDEA

### Configure IntelliJ IDEA Memory Settings

If you are using IntellJ IDEA, it can be preferable to limit the memory footprint of the IDE:

`Menu > Help > Change Memory Settings >` Enter the desired heap size (e.g. 2048)

