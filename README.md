# Table of contents
- [SIMPL - Monorepo Development Environment](#simpl---monorepo-development-environment)
  - [Project Overview](#project-overview)
  - [Requirements](#requirements)
  - [Architecture](#architecture)
  - [Getting Started](#getting-started)
    - [Useful Commands](#useful-commands)

# SIMPL - Monorepo Development Environment

This repository provides a unified setup for initializing and managing the development environment of a project composed of multiple microservices.  

The goal of this monorepo is to **automate the entire workflow**, including:
- Downloading all related projects
- Keeping dependencies up to date
- Building services
- Initializing the local development cluster

To achieve this, the monorepo leverages **[mise](https://mise.jdx.dev/)**, a package manager that:
- Automatically installs the required development tools
- Provides custom scripts created by developers to streamline repetitive or complex tasks

## Project Overview

The system is built as a collection of **microservices** developed in **Java** using the **Spring Boot** framework.  

- **Build & Execution**:  
  Each microservice is built with **Maven** and can be run using the `spring-boot-maven-plugin`.

- **Third-Party Dependencies**:  
  The project relies on several external services that are deployed locally on a **Minikube cluster**, including:
  - **[EJBCA](https://www.ejbca.org/)** (certificate authority)
  - **[Keycloak](https://www.keycloak.org/)** (identity and access management)

The monorepo provides automation for managing both the internal microservices and the required third-party services, ensuring a consistent and reproducible development environment.

## Architecture

The development environment combines **locally executed microservices** with a **Kubernetes-based cluster** for third-party dependencies.  

- **Microservices execution**  
  Each service is executed directly on the host machine using the developer’s preferred toolchain, such as:
  - **IntelliJ IDEA** (for development and debugging)  
  - **Maven** (for building and running via `spring-boot-maven-plugin`)  
  - **mise** (for installing and managing development tools)  

- **Kubernetes cluster**  
  A **Minikube node** is configured to run inside Docker, providing a lightweight local Kubernetes environment.  

- **Third-party services**  
  Dependencies required for development are deployed into the Kubernetes cluster using **Helm charts**, including:
  - Database  
  - **EJBCA** (certificate authority)  
  - **Redpanda** (streaming platform)  
  - **Keycloak** (identity and access management)  

- **Service integration**  
  When launched with the provided pre-configurations, each microservice automatically connects to its dependencies within the cluster, ensuring a consistent and reproducible setup.

## Requirements

Before setting up the development environment, make sure the following are installed and properly configured:

- **[Git](https://git-scm.com/)** – used to clone the repository and manage submodules. Git is required for:
  - Cloning the monorepo
  - Initializing submodules (`git submodule init`)
  - Updating all submodules recursively (`git submodule update --remote --recursive`)

- **[mise](https://mise.jdx.dev/)** – a package manager that automatically installs and manages all development tools and scripts needed to build, run, and manage the microservices and cluster.

- **Docker** – must be installed and running. Docker is required to:
  - Run Minikube inside a Docker container
  - Execute containerized services (databases, Keycloak, EJBCA, etc.)
  - Handle port forwarding between the host and the cluster

## Getting Started

Once the required dependencies are installed, you can set up the monorepo and initialize the development environment.

```bash
# Clone the repository
git clone https://github.com/Pumahawk/simpl-monorepo.git

# Move into the project directory
cd simpl-monorepo

# Initialize the project
mise run initialization:project
```

> **Important:**  
> If your environment requires Z-Scaler certificates, it is recommended to run:
> 
>    ```mise run initialization:zscaler```
> 
> Upon successful execution, the cluster will be available with all necessary port forwards active.

### Useful Commands

```bash
# Stop Kubernetes cluster
mise run cluster:stop

# Start Kubernetes cluster
mise run cluster:start

# Destroy the cluster (if needed)
mise run destruction:all
```

## Task Descriptions

### Project Initialization
- **`initialization:project`**  
  Initializes the entire project environment by executing the following subtasks:
  - `initialization:code` – initializes git submodules and updates all projects.
  - `initialization:cluster` – creates the Minikube cluster, configures port forwarding, and installs or upgrades the authority chart.
  - `initialization:build:code` – builds all microservices without running tests.

- **`initialization:code`**  
  Sets up the project source code:
  - `git:init` – initializes git submodules.
  - `git:update-all` – updates all submodules recursively to the latest remote commits.

- **`initialization:cluster`**  
  Sets up the local Kubernetes environment:
  - `cluster:create` – starts Minikube with a dedicated Docker network and profile.
  - `cluster:forward-node-up` – sets up port forwarding for local access.
  - `cluster:authority-install-or-upgrade` – installs or upgrades the local authority Helm chart and switches namespace.

- **`initialization:build:code`**  
  Builds all projects without executing tests using Maven (`mvnd`).

- **`initialization:zscaler`**  
  Configures the environment for Z-Scaler certificates:
  - `zscaler:install-jdk` – imports Z-Scaler CA into Java keystore.
  - `zscaler:install-cluster` – installs Z-Scaler certificates inside the Minikube cluster.

### Project Destruction
- **`destruction:all`**  
  Completely removes the development environment:
  - `cluster:destroy` – deletes the Minikube cluster.
  - `cluster:forward-node-down` – shuts down all port forwards.
  - `zscaler:remove-jdk` – removes the Z-Scaler CA from the Java keystore.

### Cluster Management
- **`cluster:start`** – starts the Minikube cluster and sets up port forwarding.
- **`cluster:stop`** – stops the Minikube cluster and disables port forwarding.
- **`cluster:status`** – shows the status of the Minikube cluster.
- **`cluster:bash`** – opens a shell inside the Minikube control plane container.
- **`cluster:forward-node-compose`** – runs the Docker Compose file for port forwarding.
- **`cluster:forward-node-up`** – starts port forwarding services.
- **`cluster:forward-node-down`** – stops port forwarding services.
- **`cluster:authority-install-or-upgrade`** – installs or upgrades the Helm chart for the authority.
- **`cluster:authority-uninstall`** – uninstalls the authority Helm chart.

### Microservices Execution
Each microservice can be started individually using `micro:run:<service>` tasks. The tasks configure Spring profiles and secrets appropriately:

- **Authority Services:** `authenticationprovider`, `identityprovider`, `onboarding`, `securityattributesprovider`, `tierone`, `tiertwo`, `usersroles`
- **Consumer Services:** `authenticationprovider`, `tierone`, `tiertwo`, `usersroles`

All services are executed using Maven’s `spring-boot:run` plugin with `mvnd` for optimized builds.

### Code Management
- **`code:fmt`** – formats the code using Spotless.
- **`code:build`** – builds all projects with Maven, including tests.
- **`code:build:no-test`** – builds all projects but skips tests and license downloads.
- **`code:test-automation:run-by-tag`** – runs automation tests filtered by Cucumber tags.

### Logging
- **`jqlog`** – formats JSON log output to a readable line-by-line format using `jq`.

### Z-Scaler Integration
- **`zscaler:install-jdk`** – imports Z-Scaler CA into the Java keystore.
- **`zscaler:remove-jdk`** – removes the Z-Scaler CA from the Java keystore.

