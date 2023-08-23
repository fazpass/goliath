package queue

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func InitKafkaProducer(brokerList []string) (*sarama.AsyncProducer, error) {
	var config = sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Flush.Frequency = 500 * time.Millisecond

	var producer, err = sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return nil, err
	}

	go func() {
		for err = range producer.Errors() {
			log.Println(err)
		}
	}()

	return &producer, nil
}

func InitKafkaConsumer(brokerList []string) (*sarama.Consumer, error) {
	var config = sarama.NewConfig()
	var consumer, err = sarama.NewConsumer(brokerList, config)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func InitKafkaConsumerGroup(brokerList []string, groupId string) (*sarama.ConsumerGroup, error) {
	var config = sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	var consumer, err = sarama.NewConsumerGroup(brokerList, groupId, config)
	if err != nil {
		return nil, err
	}

	return &consumer, nil
}
