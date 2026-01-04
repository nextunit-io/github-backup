FROM golang:1.25-alpine AS builder

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app
COPY ../src .
RUN go build -o github-backup .



FROM alpine:latest

LABEL zero <zero@nextunit.io>

RUN \
    apk update && \
    apk add --update curl bash jq git tar

ENV GITHUB_BACKUP_USERS=""
ENV GITHUB_BACKUP_ORGS=""
ENV GITHUB_ACCESS_TOKEN=""
ENV BACKUP_MAX_FILES="5"
ENV SLACK_NOTIFICATION_ENDPOINT=""

RUN mkdir -p /backup

RUN adduser -S backup
RUN chown -R backup /backup

COPY run.sh /home/backup/run.sh
COPY --from=builder /app/github-backup /home/backup/github-backup
COPY slack /home/backup/slack

RUN chown backup /home/backup/run.sh
RUN chown backup /home/backup/github-backup

WORKDIR /home/backup
USER backup

RUN chmod +x run.sh
RUN chmod +x github-backup

ENTRYPOINT /bin/bash run.sh