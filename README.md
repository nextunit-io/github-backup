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
- Clone the repo and install dependencies:
  ```bash
  git clone <repo-url>
  cd github-backup
  npm install
  ```

## Usage example
- Example using an environment token and a Node entrypoint:
  ```bash
  GITHUB_TOKEN=your_token_here node ./index.js --org nextunit-io --dest ./backups --mirror
  ```

- Common flags:
  ```
  --org <name>       Back up all repositories from an organization
  --user <name>      Back up all repositories from a user
  --dest <dir>       Destination directory for backups
  --token <token>    GitHub token (can also be set via GITHUB_TOKEN env var)
  --mirror           Use git --mirror clones
  --include-forks    Include forked repositories
  --concurrency N    Number of concurrent clones
  --verbose          Enable verbose logging
  ```

## Output
- A directory tree under the destination folder containing clones of each repository (bare or mirror clones depending on flags).

## TODO
- Implement automated tests (unit and integration)
- Add function-level documentation and inline comments
- Add CI (GitHub Actions) for linting and tests
- Add configuration file support and examples
- Improve error handling and retry logic
- Add packaging and release workflow

## Contributing
- Contributions are welcome. Please open issues for bugs/feature requests and follow standard PR workflow.



