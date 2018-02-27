#!/bin/bash

# Requires watchdog
# pip install watchdog
watchmedo shell-command \
    --patterns="*.go" \
    --command="go test"
