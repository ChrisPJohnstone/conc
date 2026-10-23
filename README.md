# CONventional Commit (`conc`) Helper

Typing conventional commit messages is a slightly tedious and I'm extremely lazy.

This is a simple utility to ease the suffering.

## Usage

- Install with go
  ```sh
  go install github.com/ChrisPJohnstone/go-conventional-commit/cmd/conc@latest
  ```
- Get help using the tool
  ```sh
  conc --help
  ```
- Positional arguments are concatenated together to create the subject of the message
  ```sh
  conc added some new functions
  ```
- Accepts type as a flag.
  ```sh
  conc -t fix "fixed some defect"
  conc --type test "added test case"
  ```
  - Only accepts valus from [spec](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional)
  - Tries to determine default
    - Branch name regexp `^{type}` for all types in spec. So `feature/....` would be `feat`
    - Falls back to `feat`
- **NOT IMPLEMENTED**
  - Handle commit rather than just writing message to stdout
  - Read ticket number from branch `feat(TEST-123): ...`
  - Body & Footer support for commits

## Development

- The [`Makefile`](./Makefile) handles most things to do with compilation, installation & cleanup
  ```sh
  make              # Compile binary locally
  make clean        # Remove local binary
  make install      # Install binary to go path
  make uninstall    # Install binary to go path
  ```
- Run linter locally in a container
  ```sh
  make lint                         # Defaults to `podman`
  make lint CONTAINER_ENGINE=docker # Run linter in docker
  ```
