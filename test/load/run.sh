#!/usr/bin/env bash


echo "Running docker version of locust sharing host network..."
rm -rf __pycache__

docker run \
    --name locust \
    --network host \
    -p 8089:8089 \
    -v $(pwd):/mnt/locust \
    -d \
    locustio/locust \
    -f /mnt/locust/locustfile.py

echo "To access Locust goto http://localhost:8089"