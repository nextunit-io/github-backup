# github-backup

A small CLI tool to back up GitHub repositories to a local directory. The tool queries GitHub (user or organization), then clones repositories locally (supports mirror-style clones) so you have a local copy for archival purposes.

## Features

- Backup all repos for a user or organization
- Support for private repositories via a GitHub token
- Mirror clone option to preserve refs
- Concurrent cloning to speed up backups

## Requirements

- git installed on your machine
- Golang v1.25.2
- A GitHub access token

## Installation (example)

- Download the version file in the releases
- Unzip the file
- Execute chmod +x `github-backup*` and execute the `github-backup*` file

## Usage example

- Common flags:
  ```
  --orgs <name>       Back up all repositories from organizations (use multiple --orgs / -o to download all in one)
  --users <name>      Back up all repositories from users (use multiple --users / -u to download all in one)
  --token <token>     GitHub token
  --file              Output file (e.g. `backup.zip`)
  --non-interactive   Deactivates the interactive mode
  --verbose           Enable verbose logging
  ```

## Output

- A directory tree under the destination folder containing clones of each repository.

## TODO

- Implement automated tests (unit and integration)
- Add function-level documentation and inline comments
- Improve error handling and retry logic
- Add packaging and release workflow

## Contributing

- Contributions are welcome. Please open issues for bugs/feature requests and follow standard PR workflow.
