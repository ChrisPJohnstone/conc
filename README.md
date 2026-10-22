# Go Conventional Commit

Typing convetional commit message is getting tedious. This binary is to speed that up for me

## Usage

- Positional arguments are concatenated together to create the subject of the message
    ```sh
    conc added some new functions
    ```
- Accepts type as a flag. Only accepts valus from [spec](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional)
    ```sh
    conc -t fix "fixed some defect"
    conc --type test "added test case"
    ```
- **NOT IMPLEMENTED**
    - Handle commit rather than just writing message to stdout
    - Read ticket number from branch `feat(TEST-123): ...`
    - Read type from branch? `feature/TEST-123`
    - Body & Footer support for commits
