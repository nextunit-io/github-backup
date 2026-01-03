#/bin/bash

if [ -z "$GITHUB_ACCESS_TOKEN" ]; then
    echo "Please set the GITHUB_ACCESS_TOKEN environment variable."
    exit 1
fi
if [ -z "$GITHUB_BACKUP_USERS" ] && [ -z "$GITHUB_BACKUP_ORGS" ]; then
    echo "Please set the GITHUB_BACKUP_USERS or GITHUB_BACKUP_ORGS environment variable."
    exit 1
fi

DATE=`date '+%Y-%m-%d_%H-%M-%S'`
BACKUP_FILENAME="repository-backup-${DATE}.zip"
BACKUP_DIR="/backup"

CMD=("-t" "$GITHUB_ACCESS_TOKEN" "-f" "${BACKUP_FILENAME}" "-i")

USERS=`echo $GITHUB_BACKUP_USERS | tr ',' '\n'`
ORGS=`echo $GITHUB_BACKUP_ORGS | tr ',' '\n'`

for user in $USERS; do
    CMD+=("--users=$user")
done

for org in $ORGS; do
    CMD+=("--orgs=$org")
done

./github-backup ${CMD[@]}

if [[ $? -ne 0 ]]; then
  if [ -n "$SLACK_NOTIFICATION_ENDPOINT" ]; then
      echo "ERROR: Doing slack notification."
      curl -s -X POST -H 'Content-type: application/json' --data '{"text":"GITHUB Backup failed!"}' $SLACK_NOTIFICATION_ENDPOINT
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
    curl -s -X POST -H 'Content-type: application/json' --data '{"text":"GITHUB Backup successfully done!"}' $SLACK_NOTIFICATION_ENDPOINT
fi