#!/bin/bash

set -e

if [[ -z "${KAFKA_HOME}" ]]; then
    echo "KAFKA_HOME env variable undefined, exiting..." 
    exit 0
fi

# create topics
echo "Create topics..."
${KAFKA_HOME}/bin/kafka-topics.sh --create --topic order-received --bootstrap-server localhost:9092
${KAFKA_HOME}/bin/kafka-topics.sh --create --topic order-confirmed --bootstrap-server localhost:9092
${KAFKA_HOME}/bin/kafka-topics.sh --create --topic order-fulfilled --bootstrap-server localhost:9092
${KAFKA_HOME}/bin/kafka-topics.sh --create --topic notif-sent --bootstrap-server localhost:9092
${KAFKA_HOME}/bin/kafka-topics.sh --create --topic event-not-processed --bootstrap-server localhost:9092

# change retention for order-received topic to three days
echo "Alter 'order-received-topic..."
${KAFKA_HOME}/bin/kafka-configs.sh --alter --add-config retention.ms=259200000 --topic order-received --bootstrap-server localhost:9092

echo "List created topics..."
${KAFKA_HOME}/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list

