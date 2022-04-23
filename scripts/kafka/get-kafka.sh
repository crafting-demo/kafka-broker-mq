#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/../.."

# Download and extract Kafka, version 3.1.0.
KAFKA_DOWNLOAD_URL="https://dlcdn.apache.org/kafka/3.1.0/kafka_2.13-3.1.0.tgz"
rm -rf broker && mkdir broker && cd broker
wget -c "$KAFKA_DOWNLOAD_URL" -O - | tar -xz --strip 1
