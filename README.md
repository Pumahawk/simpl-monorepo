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

## Requirements

Before setting up the development environment, make sure the following are installed and properly configured:

- **[mise](https://mise.jdx.dev/)** – used as the package manager to automatically install and manage the required development tools.  
- **Docker** – must be installed and running, as it is required for containerized services and the local Minikube cluster.

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

