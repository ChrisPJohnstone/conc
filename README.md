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
  - If unset tries to determine
    - Git branch name regexp `^{type}` for all types in spec. So `feature/....` would be `feat`
    - Falls back to `feat`
- Accept scope as a flag
  ```sh
  conc -s main "subject"
  conc --scope TEST-123 "subject"
  ```
  - If unset tries to determine
    - Git branch name regexp `^[^/]*/([A-Z]*-\d*)`. Example: `feature/TEST-123...` scope would be `TEST-123`
  - Set as `-` to disable
    ```sh
    conc -s '-' "subject"
    ```
- **NOT IMPLEMENTED**
  - Support shorthand for types
    - `-t d` resolves to `docs`
    - `-t r` is ambiguous `refactor` / `revert`
  - Handle commit rather than just writing message to stdout
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
