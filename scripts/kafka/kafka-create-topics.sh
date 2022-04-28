#!/bin/bash

. "${BASH_SOURCE[0]%/*}/config.sh"

ROOT="$(dirname "$0")"
cd "$ROOT/../../broker"

# Create queue topics.
bin/kafka-topics.sh --create --topic "$TOPIC_REACT" --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic "$TOPIC_GIN" --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic "$TOPIC_EXPRESS" --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic "$TOPIC_RAILS" --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic "$TOPIC_SPRING" --if-not-exists --bootstrap-server localhost:9092
bin/kafka-topics.sh --create --topic "$TOPIC_DJANGO" --if-not-exists --bootstrap-server localhost:9092
