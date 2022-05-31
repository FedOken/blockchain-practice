#!/bin/sh

echo "Start run-container.sh"

# Move to work directory
cd ${ENV_WORK_DIR}

# Synchronizing code’s dependencies
go mod tidy

# Install binary
#go install

echo "Done run-container.sh"

# Leave the connection open
tail -F open-connection