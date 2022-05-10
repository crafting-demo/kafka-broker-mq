#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT"

# Update system
sudo apt-get update

# Install libgmp3
sudo apt-get install libgmp3-dev -y

# Install mysql client
sudo apt-get install mysql-client libmysqlclient-dev -y

# Install postgres client
sudo apt-get install postgresql-client -y

# Setup/prepare all databases.
./db/setup-db.sh

# Download/extract kafka.
./kafka/get-kafka.sh

# Build go server.
./go-server.sh
