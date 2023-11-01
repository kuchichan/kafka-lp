#!/bin/bash

set -e

if [[ -z "${KAFKA_HOME}" ]]; then
    echo "KAFKA_HOME env variable undefined, exiting..." 
    exit 0
fi

${KAFKA_HOME}/bin/zookeeper-server-start.sh ${KAFKA_HOME}/config/zookeeper.properties
