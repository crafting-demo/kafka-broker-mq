#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT"

# Update system
sudo apt-get update

# Download/extract kafka.
./kafka/get-kafka.sh

# Build go server.
./go-server.sh

# Setup/prepare all databases.
./db/setup-db.sh
