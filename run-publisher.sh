#!/bin/bash

set -e

if [[ -z "${KAFKA_HOME}" ]]; then
    echo "KAFKA_HOME env variable undefined, exiting..." 
    exit 0
fi

TOPIC_NAME=$1

${KAFKA_HOME}/bin/kafka-console-producer.sh --topic ${TOPIC_NAME} --bootstrap-server localhost:9092
