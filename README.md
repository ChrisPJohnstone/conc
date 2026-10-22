# Go Conventional Commit

Typing convetional commit message is getting tedious, quick helper to speed that up

## Spec

Scribblings of what I want, will clean up later

- Flag for type

    ```sh
    ${binary} -t fix
    ${binary} --type chore
    ```

    - Default to `feat`
    - Only accept types from [full spec](https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional)
        ```
        [
          'build',
          'chore',
          'ci',
          'docs',
          'feat',
          'fix',
          'perf',
          'refactor',
          'revert',
          'style',
          'test'
        ];
        ```

- Automatically gets ticket ID's from branch name
    - This will probably be hardcoded to what I work on
    - `{type}({branch}): {subject}`
- Subject will be positional argument, all I really use
- Might add support for body & footer as flags
- Haven't decided if it will actually commit, output message or open message in editor
    - Have it output the message and do git alias like `commit -m "$( cc $@ )"`?
