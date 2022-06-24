package queue

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func InitKafkaProducer(brokerList []string) *sarama.AsyncProducer {
	var config = sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	var producer, err = sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		panic(err)
	}

	go func() {
		for err = range producer.Errors() {
			log.Println(err)
		}
	}()

	return &producer
}

func InitKafkaConsumer(brokerList []string) *sarama.Consumer {
	var config = sarama.NewConfig()
	var consumer, err = sarama.NewConsumer(brokerList, config)
	if err != nil {
		panic(err)
	}

	return &consumer
}
