# Event-Sourcing-Watcher using golang and kafka

#### A watcher file written in golang following the event-driven-architecture that watches for new events in a file directory like creation of files and produces their content to kafka for storage .Later consumers subscribed to a specific topic can fetch data from kafka and store into mongodb/cassandra.

##### Installation guide:

1. dep ensure -update <br />

2. Run kafka and zookeeper in Docker <br />
    #### docker run -d --name zookeeper -p 2181:2181 jplock/zookeeper <br />
    #### docker run -d --name kafka -p 7203:7203 -p 9092:9092 -e KAFKA_ADVERTISED_HOST_NAME=172.17.0.1  -e ZOOKEEPER_IP=172.17.0.1 ches/kafka <br />
3. Create 3 topics: csv , txt , xlsx  <br />
    #### docker run --rm ches/kafka kafka-topics.sh --create --topic csv --replication-factor 1 --partitions 1 --zookeeper 172.17.0.1:2181  <br />
    #### docker run --rm ches/kafka kafka-topics.sh --create --topic txt --replication-factor 1 --partitions 1 --zookeeper 172.17.0.1:2181  <br />
    #### docker run --rm ches/kafka kafka-topics.sh --create --topic xlsx --replication-factor 1 --partitions 1 --zookeeper 172.17.0.1:2181 <br />
4. Run Console consumer for each topic <br />
    #### docker run --rm ches/kafka kafka-console-consumer.sh --topic csv --zookeeper 172.17.0.1:2181 <br />
    #### docker run --rm ches/kafka kafka-console-consumer.sh --topic txt --zookeeper 172.17.0.1:2181 <br />
    #### docker run --rm ches/kafka kafka-console-consumer.sh --topic xlsx --zookeeper 172.17.0.1:2181 <br />
5. go build .   <br />
6. go run watcher.go  <br />
7. Drop the files in the landing zone(file directory you specify in watcher.go)  <br />
