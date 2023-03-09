package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/arieuqis/go_training/internal/infra/database"
	"github.com/arieuqis/go_training/internal/usecase"
	"github.com/arieuqis/go_training/pkg/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

	//sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./orders.db")

	if err != nil {
		panic(err)
	}
	defer db.Close() // executes everything and then closes the connection

	repository := database.NewOrderRepository(db)
	usecase := usecase.CalculateFinalPrice{OrderRepository: repository}

	msgChanKafka := make(chan *ckafka.Message)

	topics := []string{"orders"}
	servers := "host.docker.internal:9094"
	fmt.Println("Kafka consumer has been started")
	go kafka.Consume(topics, servers, msgChanKafka)
	kafkaWorker(msgChanKafka, usecase)
}

func kafkaWorker(msgChanKafka chan *ckafka.Message, uc usecase.CalculateFinalPrice) {
	fmt.Println("Kafka worker has started")
	for msg := range msgChanKafka {
		var OrderInputDTO usecase.OrderInputDTO
		err := json.Unmarshal(msg.Value, &OrderInputDTO)
		if err != nil {
			panic(err)
		}
		outputDto, err := uc.Execute(OrderInputDTO)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Kafka has processed orde %s\n", outputDto.ID)
	}
}
