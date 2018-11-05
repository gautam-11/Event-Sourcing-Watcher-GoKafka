# Golang-kafka-mongodb-watcher

A watcher file written  in golang that watches for new events in the file-system like creation of  files and produces their content to kafka for storage .Later consumers subscribed to a specific topic fetch data from kafka and store into mongodb.

Installation guide:

1. dep ensure -update
2. go build .
3. go run watcher.go
4. Run zookeeper
5. Run Kafka after creating 4 topics csv , txt , xlsx , pdf , docx 
6. Drop the files in landing zones