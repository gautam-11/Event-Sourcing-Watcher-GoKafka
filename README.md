# Golang-kafka-mongodb-watcher

A watcher file written  in golang that watches for new events in the file-system like creation of  files and produces their content to kafka for storage .Later consumers subscribed to a specific topic fetch data from kafka and store into mongodb.