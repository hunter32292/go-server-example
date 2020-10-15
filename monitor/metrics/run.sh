echo "Running docker version of promethues sharing host network..."

docker run \
    --network host \
    -v $(pwd):/etc/prometheus \
    prom/prometheus
