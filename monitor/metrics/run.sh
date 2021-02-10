echo "Running docker version of prometheus sharing host network..."

docker run \
    --network host \
    -v $(pwd):/etc/prometheus \
    prom/prometheus

echo "To access prometheus goto http://localhost:8080"