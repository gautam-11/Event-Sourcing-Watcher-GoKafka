# Golang-kafka-mongodb-watcher

A watcher file written  in golang that watches for new events in the file-system like creation of  files and produces their content to kafka for storage .Later consumers subscribed to a specific topic fetch data from kafka and store into mongodb/cassandra.

Installation guide:

1. dep ensure -update

2. Run kafka and zookeeper in Docker
    Run Zookeeper :  sudo docker run -d --name zookeeper -p 2181:2181 jplock/zookeeper
    Run Kafka: sudo docker run -d --name kafka -p 7203:7203 -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=172.17.0.1  -e ZOOKEEPER_IP=172.17.0.1 ches/kafka
3. Create 3 topics: csv , txt , xlsx
    sudo docker run --rm ches/kafka kafka-topics.sh --create --topic csv --replication-factor 1 --partitions 1 --zookeeper 172.17.0.1:2181
    sudo docker run --rm ches/kafka kafka-topics.sh --create --topic txt --replication-factor 1 --partitions 1 --zookeeper 172.17.0.1:2181
    sudo docker run --rm ches/kafka kafka-topics.sh --create --topic xlsx --replication-factor 1 --partitions 1 --zookeeper 172.17.0.1:2181
4. Run Console consumer for each topic
    sudo docker run --rm ches/kafka kafka-console-consumer.sh --topic csv --zookeeper 172.17.0.1:2181
    sudo docker run --rm ches/kafka kafka-console-consumer.sh --topic txt --zookeeper 172.17.0.1:2181
    sudo docker run --rm ches/kafka kafka-console-consumer.sh --topic xlsx --zookeeper 172.17.0.1:2181
5. go build .
6. go run watcher.go
7. Drop the files in landing zone