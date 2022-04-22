#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/.."

# Download and extract Kafka, version 3.1.0.
rm -rf broker && mkdir broker && cd broker
wget -c https://dlcdn.apache.org/kafka/3.1.0/kafka_2.13-3.1.0.tgz -O - | tar -xz --strip 1

# Create queue topics.
bin/kafka-topics.sh --create --topic ts-react --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic go-gin --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic ts-express --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic ruby-rails --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic kotlin-spring --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic py-django --bootstrap-server localhost:9092
