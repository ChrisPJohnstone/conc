# CONventional Commit (`conc`) Helper

Typing conventional commit messages is a slightly tedious and I'm extremely lazy.

This is a simple utility to ease the suffering.

## Usage

- Install with go
  ```sh
  go install github.com/ChrisPJohnstone/conc@latest
  ```
- Create git alias
  ```sh
  git config --global alias.conc '!f() { git commit -m "$(conc "$@")"; }; f'
  ```
- Get help using the tool
  ```sh
  git conc --help
  ```
- Positional arguments are concatenated together to create the subject of the message
  ```sh
  git conc added some new functions
  ```
- Accepts type as a flag.
  ```sh
  git conc -t fix "fixed some defect"
  git conc --type test "added test case"
  ```
  - Only accepts valus from [spec](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional)
  - You only need to type enough for it to be unique
    ```sh
    git conc -t t "message"     # Resolves to 'test'
    git conc -t f "message"     # Ambiguous as could be feat or fix
    git conc -t fe "message"    # Resolves to 'feat'
    ```
  - If unset tries to determine
    - Git branch name regexp `^{type}` for all types in spec. So `feature/....` would be `feat`
    - Falls back to `feat`
- Accepts scope as a flag
  ```sh
  git conc -s main "subject"
  git conc --scope TEST-123 "subject"
  ```
  - If unset tries to determine
    - Git branch name regexp `^[^/]*/([A-Z]*-\d*)`. Example: `feature/TEST-123...` scope would be `TEST-123`
  - Set as `-` to disable
    ```sh
    git conc -s '-' "subject"
    ```
- Accepts body and footer flags
  ```sh
  git conc "subject" -b "body" -f "footer"
  git conc "subject" --body "body" --footer "footer"
  ```

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
