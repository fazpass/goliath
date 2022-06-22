package database

import (
	"log"
	"time"

	"github.com/Shopify/sarama"
)

func InitKafka(brokerList []string) *sarama.AsyncProducer {
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
