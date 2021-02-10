echo "Running docker version of Elastic Logbeat Kibana (ELK) ..."
echo "Cleanup for elastic..."
docker kill elastic
docker rm -v elastic
echo "Cleanup for kibana..."
docker kill kibana
docker rm -v kibana
echo "Cleanup for filebeat..."
docker kill filebeat
docker rm -v filebeat

docker run \
    -p 9200:9200 \
    -p 9300:9300 \
    -e "discovery.type=single-node" \
    --name "elastic" \
    -d \
    --rm \
    docker.elastic.co/elasticsearch/elasticsearch:7.9.2

docker run \
    --name="kibana" \
    --link elastic:elasticsearch \
    -p 5601:5601 \
    -d \
    --rm \
    docker.elastic.co/kibana/kibana:7.9.2

echo "waiting for elasticsearch and kibana..."
sleep 10


export LOG_FILE="example-server"
docker run \
    --name="filebeat" \
    -d \
    --rm \
    --link kibana:kibana \
    --link elastic:elasticsearch \
    --volume="$(pwd)/filebeat.yml:/usr/share/filebeat/filebeat.yml:ro" \
    --volume="$(pwd)/../..:/example-server" \
    docker.elastic.co/beats/filebeat:7.9.2 

echo "To access kibana goto localhost:5601"