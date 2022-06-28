package queue

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/Shopify/sarama"
	helpermsg "github.com/fazpass/goliath/helper/queue/message"
)

func Consume(client sarama.Consumer, topic string, group string, message chan []byte) {
	var (
		wg              sync.WaitGroup
		partitions, err = client.Partitions(topic)
	)
	if err != nil {
		panic(err)
	}

	var partition int32
	for _, partition = range partitions {
		var partitionConsumer, err = client.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}

		defer partitionConsumer.AsyncClose()
		wg.Add(1)

		go func(partitionConsumer sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range partitionConsumer.Messages() {
				message <- msg.Value
			}
		}(partitionConsumer)
	}

	wg.Wait()
	client.Close()
}

func ConsumeGroup(ctx context.Context, client sarama.ConsumerGroup, topic string, group string, message chan []byte) {
	var (
		wg           sync.WaitGroup
		err          error
		groupHandler = ConsumerGroup{
			ready:   make(chan bool),
			message: message,
		}
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err = client.Consume(ctx, []string{topic}, &groupHandler); err != nil {
				panic(err)
			}

			if ctx.Err() != nil {
				return
			}

			groupHandler.ready = make(chan bool)
		}
	}()

	<-groupHandler.ready
	fmt.Println("Group consumer ready to consume message")

	<-ctx.Done()
	log.Println("terminating: context cancelled")

	wg.Wait()
	if err = client.Close(); err != nil {
		panic(err)
	}
}

func Produce(client sarama.AsyncProducer, topic string, msg *helpermsg.Message) {
	var producer = client
	producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: msg,
	}
}
