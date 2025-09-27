# Table of contents
- [Description](#description)
  - [Project structure](#project-structure)
- [Clone repository](#clone-repository)
  - [Init submodules after clone](#init-submodules-after-clone)
  - [Update all submodules](#update-all-submodules)
- [Simple Code](#simple-code)
  - [Build project](#build-project)
    - [Build arguments](#build-arguments)
    - [Build optimazed](#build-optimazed)
    - [Examples](#examples)
- [Helm Charts](#helm-charts)
- [Skaffold Support](#skaffold-support)
- [Clone project](#clone-project)

# Description

This repository is designed to support the development of simple projects.

# Clone repository

To clone the project along with all submodule dependencies, use:

`git clone --recurse-submodules https://github.com/Pumahawk/simpl-monorepo.git`

## Initializing Submodules After Cloning

To ensure all submodules are properly initialized after cloning, use the following commands:

```bash
# Clone repository
git clone https://github.com/Pumahawk/simpl-monorepo.git

#Align submodules
git submodule init
git submodule update
```

## Updating All Submodules

- Align submodule commits: `git submodule update`
- Update each submodule to the latest commit from the original develop branch: `git submodule update --remote`

# Building the Simpl Code

## Building the Project

Use mise

```
mise install
mise run build
```

# Helm Charts

Custom Helm charts for deployment.

- simpl-participant - Helm chart for Simpl Participant.

## Simpl Participant Chart

To deploy a complete configuration for the participant, use:

```bash
helm install simpl-participant helm/simpl-participant -f helm/custom-simpl-participant.yaml
```

> Note: The file **helm/custom-simpl-participant.yaml** is provided as an example of customizable
> properties. Make sure to update the file with your personal host domain.

> Additionally, ensure the host domain and JDBC URL for the database connection reflect your
> specific Helm Release name and Kubernetes Namespace.

