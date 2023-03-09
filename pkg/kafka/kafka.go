package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topic []string, servers string, msgChan chan *ckafka.Message) {
	kafkaConsumer, err := ckafka.NewConsumer(&ckafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "goapp",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = kafkaConsumer.SubscribeTopics(topic, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
