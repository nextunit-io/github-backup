#/bin/bash

if [ -z "$GITHUB_ACCESS_TOKEN" ]; then
    echo "Please set the GITHUB_ACCESS_TOKEN environment variable."
    exit 1
fi
if [ -z "$GITHUB_BACKUP_USERS" ] && [ -z "$GITHUB_BACKUP_ORGS" ]; then
    echo "Please set the GITHUB_BACKUP_USERS or GITHUB_BACKUP_ORGS environment variable."
    exit 1
fi

# set Variables
DATE=`date '+%Y-%m-%d_%H-%M-%S'`
PRETTY_TIME_NOW=$(date "+%Y-%m-%d %H:%M:%S")
BACKUP_FILENAME="repository-backup-${DATE}.zip"
BACKUP_DIR="/backup"
ARCH=$(uname -m)

CMD=("-t" "$GITHUB_ACCESS_TOKEN" "-f" "${BACKUP_FILENAME}" "-i")

USERS=`echo $GITHUB_BACKUP_USERS | tr ',' '\n'`
ORGS=`echo $GITHUB_BACKUP_ORGS | tr ',' '\n'`
CONFIG="Users: ${GITHUB_BACKUP_USERS} // Orgs: ${GITHUB_BACKUP_ORGS}"

for user in $USERS; do
    CMD+=("--users=$user")
done

for org in $ORGS; do
    CMD+=("--orgs=$org")
done

if [ "$ARCH" == "x86_64" ] || [ "$ARCH" == "amd64" ]; then
    echo "Using github-backup-amd64 binary for architecture $ARCH"

    ./github-backup-amd64 ${CMD[@]}
elif [ "$ARCH" == "aarch64" ] || [ "$ARCH" == "arm64" ]; then
    echo "Using github-backup-arm64 binary for architecture $ARCH"

    ./github-backup-arm64 ${CMD[@]}
else
    echo "Unsupported architecture: $ARCH"

    if [ -n "$SLACK_NOTIFICATION_ENDPOINT" ]; then
      echo "ERROR: Doing slack notification."

      PAYLOAD=$(jq \
        --arg config "$CONFIG" \
        --arg time "$PRETTY_TIME_NOW" \
        --arg error "Unsupported architecture $ARCH" \
        '.attachments[0].fields[0].value=$config
        | .attachments[0].fields[1].value=$time
        | .attachments[0].fields[2].value=$error' \
        "slack/error.json"
      )

      curl -s -X POST -H 'Content-type: application/json' --data "$PAYLOAD" $SLACK_NOTIFICATION_ENDPOINT
    fi
    exit 1
fi

if [[ $? -ne 0 ]]; then
  if [ -n "$SLACK_NOTIFICATION_ENDPOINT" ]; then
      echo "ERROR: Doing slack notification."

      PAYLOAD=$(jq \
        --arg config "$CONFIG" \
        --arg time "$PRETTY_TIME_NOW" \
        --arg error "Error while executing github backup: ErrorCode: $?" \
        '.attachments[0].fields[0].value=$config
        | .attachments[0].fields[1].value=$time
        | .attachments[0].fields[2].value=$error' \
        "slack/error.json"
      )

      curl -s -X POST -H 'Content-type: application/json' --data "$PAYLOAD" $SLACK_NOTIFICATION_ENDPOINT
  fi

  exit 1
fi

mv $BACKUP_FILENAME $BACKUP_DIR

if [ -n "$BACKUP_MAX_FILES" ] && [ "$BACKUP_MAX_FILES" != "0" ]; then
    echo "Removing old backups. We just want to store ${BACKUP_MAX_FILES}"
    (find $BACKUP_DIR -type f|sort|tail -n $BACKUP_MAX_FILES;find $BACKUP_DIR -type f)|grep -v "^$BACKUP_DIR$"|sort|uniq -u|xargs rm -rfv
fi

if [ -n "$SLACK_NOTIFICATION_ENDPOINT" ]; then
    echo "Doing slack notification."
    
    PAYLOAD=$(jq \
      --arg config "$CONFIG" \
      --arg time "$PRETTY_TIME_NOW" \
      '.attachments[0].fields[0].value=$config
      | .attachments[0].fields[1].value=$time' \
      "slack/success.json"
    )

    curl -s -X POST -H 'Content-type: application/json' --data "$PAYLOAD" $SLACK_NOTIFICATION_ENDPOINT
fi