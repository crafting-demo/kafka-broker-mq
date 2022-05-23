#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT"

# Update system
sudo apt-get update

# Setup/prepare all databases.
# Note: use snapshots if available in current org,
#       or create them for subsequent use.
# ./db/setup-db.sh

# Download/extract kafka.
./kafka/get-kafka.sh

# Build go server.
./go-server.sh
