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
- Execute `chmod +x github-backup*` and execute the `github-backup*` file
- Or run via Docker (see "Docker" below)

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

This separates the binary entrypoint from reusable code and container assets, making builds/tests and packaging easier.

## Docker

You can build and run github-backup inside a container, which is useful for CI or environments without Go installed.

### Required environment variables

| Variable                                 | Description                                                                                                                                                                                                                                                       |
| ---------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| GITHUB_ACCESS_TOKEN                      | Your token, you created for github. Here you can find, how you can create one: [https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/](https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/) |
| GITHUB_BACKUP_USERS                      | Back up all repositories from users (multiple users can be separaterd by ',')                                                                                                                                                                                     |
| GITHUB_BACKUP_ORGS                       | Back up all repositories from organizations (multiple orgs can be separaterd by ',')                                                                                                                                                                              |
| SLACK_NOTIFICATION_ENDPOINT _(Optional)_ | This entry is just required, if you'd like to send update-debug messages to slack.                                                                                                                                                                                |
| BACKUP_MAX_FILES _(Optional)_            | The backup script will delete older backups. With this entry you define how many backups should be stored before it deletes newers. Older files in the directory will be deleted. The **Default Value is 5**. If you'd like to disable this feature, use 0.       |

### Build docker

Build locally (working dir `docker`):

```
docker build -t github-backup:latest .
```

### Run docker

There are two ways to set environment variables. You can find examples for both below.

#### Without environment file

```bash
docker run -e "GITHUB_ACCESS_TOKEN=<access-token>" -e "GITHUB_BACKUP_USERS=<usernames>" -e "GITHUB_BACKUP_ORGS=<orgs>" -e "SLACK_NOTIFICATION_ENDPOINT=<slack>" -e "BACKUP_MAX_FILES=<number of files, that should be available>" -v /full/path/to/backup/dir:/backup --rm nextunit/github-backup:latest
```

### Using a environment file

Environment file content:

```text
GITHUB_ACCESS_TOKEN=token-of-username-1
GITHUB_BACKUP_USERS=username-1
GITHUB_BACKUP_ORGS=username-1
SLACK_NOTIFICATION_ENDPOINT=https://hooks.slack.com/services/XXXXXX/XXXXXXXX/XXxxXXxx
BACKUP_MAX_FILES=10
```

Afterwards just execute the following command:

```bash
docker run --rm -v /full/path/to/backup/dir:/backup --env-file "./path/to/environment-file.conf" nextunit/github-backup:latest
```

## Contributing

- Contributions are welcome. Please open issues for bugs/feature requests and follow standard PR workflow.
