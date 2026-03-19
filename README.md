
- [Simpl IAA Dev Environment Automation Toolkit](#simpl-iaa-dev-environment-automation-toolkit)
  - [Overview](#overview)
  - [Goals](#goals)
  - [Prerequisites](#prerequisites)
    - [Supported Environment](#supported-environment)
    - [Required Tools](#required-tools)
  - [Key Features](#key-features)
    - [Source Code Management](#source-code-management)
  - [Download code](#download-code)
    - [Build & Compilation](#build--compilation)
  - [Build backend services skips tests](#build-backend-services-skips-tests)
  - [Build backend services](#build-backend-services)
    - [Automated Testing](#automated-testing)
  - [Running tests tasks](#running-tests-tasks)
    - [Local Execution](#local-execution)
  - [Run backend service commands](#run-backend-service-commands)
    - [Environment Initialization](#environment-initialization)
  - [Initialization Code and Cluster](#initialization-code-and-cluster)
  - [Initialization code only](#initialization-code-only)
  - [Initialization keycloak](#initialization-keycloak)
  - [Initialization ejbca](#initialization-ejbca)
    - [Docker Compose Management](#docker-compose-management)
    - [Global Task Orchestration](#global-task-orchestration)
  - [Cluster Initialization](#cluster-initialization)
    - [Entry Point](#entry-point)
    - [Execution Flow](#execution-flow)
    - [Result](#result)

# Simpl IAA Dev Environment Automation Toolkit

## TL;DR

```bash
git clone https://github.com/Pumahawk/simpl-monorepo.git
cd simpl-monorepo
mise trust -E zscaler,wait
mise run -E zscaler,wait initialization:project
```

## Overview

This repository provides a set of tools designed to simplify the management of the development environment for the **Simpl IAA** project.

Its main goal is to streamline the onboarding process for new developers by minimizing the manual steps required to properly set up a local development environment.

To achieve this, the repository leverages [Mise](https://mise.jdx.dev/), a tool for managing development dependencies and tasks. It enables automatic installation and configuration of the main tools required for development, including:

- gomplate
- helm
- java
- jq
- kubectl
- kubectx
- kubens
- minikube
- mvnd
- yq

In addition to dependency management, the repository includes a collection of scripts that automate the core operational tasks involved in developing and running the project.

## Goals

- Reduce development environment setup time
- Standardize workflows across developers
- Automate repetitive and error-prone tasks
- Simplify local execution of the microservices ecosystem

## Prerequisites

Before using this repository, ensure that your environment meets the following requirements.

### Supported Environment

> **Important**
> This setup is actively maintained and tested using **WSL (Windows Subsystem for Linux) with Ubuntu**.

### Required Tools

The following tools must be available on your system:

- [Docker Desktop](https://docs.docker.com/desktop/) (required to run containers and Kubernetes locally)
- [Mise](https://mise.jdx.dev/) (used to manage development dependencies and tasks)

> All additional development tools (e.g. kubectl, JDK, mvnd, minikube) are automatically installed and managed via Mise.

> **Advanced Configuration (Without Docker Desktop)**
> It is possible to run the environment without using Docker Desktop, but this requires additional manual configuration.
> In particular, you must update the **CoreDNS ConfigMap** in your Kubernetes cluster to resolve host.docker.internal to the WSL host.
> This setup is intended for advanced users and is not officially supported by default.

## Key Features

This repository provides a set of automation tools to simplify the setup and management of the **Simpl IAA** development environment.

### Source Code Management

- Automated download of all microservices source code (both **frontend** and **backend**)
- Support for **multi-module repositories**

```bash
## Download code
mise run initialization:code 
```

### Build & Compilation

- Backend microservices compilation using **Maven multi-module** structure
- Optimized build process through **mvnd**

```bash
## Build backend services skips tests
mise run code:build:no-test 

## Build backend services
mise run code:build

```

### Automated Testing

- Execution of backend automated tests using **Cucumber**
- Support for:
  - Feature file execution
- Tag-based test selection
- Integration with **Surefire Cucumber plugin**

```bash
## Running tests tasks

mise run code:test-automation:run-all
mise run code:test-automation:run-by-feature feature/file/path.feature
mise run code:test-automation:run-by-tag "@SIMPL=xxx and @issue=yyy"
mise run code:test-automation:allure-report-serve
```

### Local Execution

- Startup of all microservices using the **local profile**

```bash
## Run backend service commands
mise run micro:run:authority-authenticationprovider
mise run micro:run:authority-identityprovider
mise run micro:run:authority-onboarding
mise run micro:run:authority-securityattributesprovider
mise run micro:run:authority-tierone
mise run micro:run:authority-tiertwo
mise run micro:run:authority-usersroles
mise run micro:run:consumer-authenticationprovider
mise run micro:run:consumer-tierone
mise run micro:run:consumer-tiertwo
mise run micro:run:consumer-usersroles
```

### Environment Initialization

- Automated setup of the full development environment, including:
  - Kubernetes cluster initialization (via minikube)
  - **Keycloak** setup and realm initialization
  - **EJBCA** initialization
  - Configuration of microservices public keys
  - Management of separate **Docker Compose** configurations for:
    - Frontend
    - Backend

```bash
## Initialization Code and Cluster
mise run initialization:project

## Initialization code only
mise run initialization:cluster

## Initialization keycloak
mise run cluster:keycloak-autoconfigure

## Initialization ejbca
mise run cluster:initialization_ejbca
```


### Docker Compose Management

- Unified wrapper around **Docker Compose** configurations for:
  - Frontend microservices
  - Backend microservices
- Simplified commands to manage containers without directly interacting with Docker Compose files
- Consistent interface for starting, stopping, and monitoring services

```bash
mise run cluster:redpanda:compose
mise run simpl-services-fe:compose
mise run simpl-services:compose
```

### Global Task Orchestration

- Centralized task management via [Mise](https://mise.jdx.dev/).
- Ability to control the entire microservices ecosystem with a single command

```bash
mise run simpl:start
mise run simpl:restart
mise run simpl:stop
```

## Cluster Initialization

The `initialization:cluster` task is responsible for provisioning and configuring the entire runtime environment, including Kubernetes, infrastructure services, and microservices.

It executes a sequence of steps that must run in a specific order to ensure a fully working system.

### Entry Point

```
mise run initialization:cluster
```

### Execution Flow

The process performs the following steps:

1. **Cluster Creation**

- Creates a local Kubernetes cluster using **Minikube**
- Uses a dedicated profile and Docker network
- Prepares the namespace `local-authority`

2. **ZScaler Configuration (Optional)**

- Installs ZScaler certificates:
  - Into the JDK keystore
  - Inside the Kubernetes cluster
- This step is executed only if ZScaler support is enabled

3. **Port Forwarding Setup**

- Starts a dedicated Docker Compose stack for **port forwarding**
- Exposes internal Kubernetes services to the local machine

4. **Redpanda Startup**

- Starts **Redpanda** (Kafka-compatible streaming platform) via Docker Compose
- Provides messaging infrastructure required by the system

5. **Authority Deployment (Helm)**

- Installs or upgrades the **local-authority Helm chart**
- Deploys core infrastructure components inside Kubernetes, including:
  - Keycloak
  - PostgreSQL
  - EJBCA

6. **Test Automation Configuration**

- Waits for Keycloak to become available
- Retrieves client secrets dynamically
- Injects secrets into test configuration files
- Enables execution of automated tests without manual setup

7. **EJBCA Initialization**

- Waits for the EJBCA pod to be ready
- Uploads configuration files into the container
- Executes initialization EJBCA
- Extracts:
  - PKCS#12 keystores
  - Truststores
- Makes credentials available to dependent services

8. **EJBCA Certificate Retrieval**

- Downloads the EJBCA public certificate (`.pem`)
- Stores it locally for use by microservices

9. **Tier1 Authority Bootstrap**

Starts the **tier1 authority gateway** service
Enables initial interactions with identity and certificate services

10. **Keycloak Auto-Configuration**

- Configures realms:
  - `authority`
  - `participant`
  - `onboarding`
- Applies base configuration (e.g. frontend URLs)
- Generates and exports service account secrets

11. **Backend Services Startup**

- Starts all backend microservices via Docker Compose
- Pulls latest images if required

12. **Frontend Services Startup**

- Starts frontend applications via a separate Docker Compose stack

### Result

After this process completes:

- Kubernetes cluster is fully operational
- Core infrastructure (Keycloak, EJBCA, Redpanda) is configured
- Certificates and secrets are generated and propagated
- Backend and frontend microservices are running
- The system is ready for development and testing
