#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT"

# Update system
sudo apt-get update

# Install libgmp3
sudo apt-get install libgmp3-dev -y

# Install mysql client
sudo apt-get install mysql-client libmysqlclient-dev -y

./mysql.sh
