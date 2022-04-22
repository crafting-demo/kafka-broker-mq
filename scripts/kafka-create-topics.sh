#!/bin/bash

ROOT="$(dirname "$0")"
cd "$ROOT/../broker"

# Create queue topics.
bin/kafka-topics.sh --create --topic ts-react --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic go-gin --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic ts-express --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic ruby-rails --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic kotlin-spring --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic py-django --if-not-exists --bootstrap-server localhost:9092
