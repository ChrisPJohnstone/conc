# Contributing

- All commits to must follow the [Conventional Commits](https://www.conventionalcommits.org/) format
- Pull requests run checks before merging is possible
  - [golangci-lint-test](https://github.com/golangci/golangci-lint-action) Does linting & formatting of go files
  - `go test` is run
- This project uses [Release Please](https://github.com/googleapis/release-please) to automate versioning and changelog generation.

## Development

- The [`Makefile`](./Makefile) handles most things to do with compilation, installation & cleanup
  ```sh
  make              # Compile binary locally
  make clean        # Remove local binary
  make install      # Install binary to go path
  make uninstall    # Remove binary from go path
  ```
- Run linter locally in a container

  ```sh
  make lint                         # Defaults to `podman`
  make lint CONTAINER_ENGINE=docker # Run linter in docker
  ```

  - A container engine such as [podman](https://podman.io/) or [docker](https://www.docker.com/) is required

- Run tests or checks
  ```sh
  make test   # Runs go tests
  make check  # Runs CI checks (linter & tests)
  ```
