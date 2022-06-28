package queue

import (
	"github.com/Shopify/sarama"
)

type ConsumerGroup struct {
	ready   chan bool
	message chan []byte
}

func (cg *ConsumerGroup) Setup(sarama.ConsumerGroupSession) error {
	close(cg.ready)
	return nil
}

func (consumer *ConsumerGroup) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *ConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message := <-claim.Messages():
			consumer.message <- message.Value
			session.MarkMessage(message, "")

		case <-session.Context().Done():
			return nil
		}
	}
}
