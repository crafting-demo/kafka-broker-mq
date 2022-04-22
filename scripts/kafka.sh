#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

# Download and extract Kafka, version 3.1.0.
rm -rf broker && mkdir broker && cd broker
wget -c https://dlcdn.apache.org/kafka/3.1.0/kafka_2.13-3.1.0.tgz -O - | tar -xz --strip 1
