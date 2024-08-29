- [Clone project](#clone-project)
  - [Init submodules after clone](#init-submodules-after-clone)
- [Build project](#build-project)
  - [Build optimazed](#build-optimazed)
- [Update all submodules](#update-all-submodules)

# Clone project

Command to clone project with all submodule dependencies

`git clone --recurse-submodules https://github.com/Pumahawk/simpl-monorepo.git`

## Init submodules after clone

```bash
# Clone repository
git clone https://github.com/Pumahawk/simpl-monorepo.git

#Align submodules
git submodule init
git submodule update
```

# Build project

Build project using maven with default configuration

`./build`

## Build optimazed

- Build only one project with dependencies - `./build -pl <project-path>`
- Build also make dependents - `./build -pl <project-path> -amd`
- Build projects in parallel - `./build -T100`
- Build projects in parallel and skip tests - `./build -T100 -Dmaven.test.skip`

# Update all submodules

- Align submodules commit - `git submodule update`
- Update module with original develop branch - `git submodule update --remote`
