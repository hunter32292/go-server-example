#!/usr/bin/env bash

echo "Running docker version of prometheus sharing host network..."

docker run \
    --network host \
    -v $(pwd):/etc/prometheus \
    -d \
    --rm \
    prom/prometheus

echo "To access prometheus goto http://localhost:9090"
