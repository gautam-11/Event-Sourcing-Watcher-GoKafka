package KafkaProducer

import (
	"Event-Sourcing-Watcher-GoKafka/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

//Produce : Producer function to produce messages to kafka
func Produce(msg string, topic string) {
	// create producer
	producer, err := initProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}
	// publish without goroutene
	publish(msg, producer, topic)

	// publish with go routene
	// go publish(msg, producer)
}

func initProducer() (sarama.SyncProducer, error) {
	// setup sarama log to stdout
	//sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	configs := sarama.NewConfig()
	configs.Producer.Retry.Max = 5
	configs.Producer.RequiredAcks = sarama.WaitForAll
	configs.Producer.Return.Successes = true

	configuration, err := config.GetEnv()
	if err != nil {
		log.Println("Error in reading Kafka connection string")
	}

	kafkaConn := configuration.Constants.KAFKA_CONN

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kafkaConn}, configs)

	return prd, err
}

func publish(message string, producer sarama.SyncProducer, topic string) {
	// publish sync
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	p, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	} else {
		fmt.Println("Records sent")
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", p)
	fmt.Println("Offset: ", o)
}
