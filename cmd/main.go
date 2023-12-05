package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/CelsoTaliatelli/wallet-consumer/internal/entity/transaction"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		println(err)
		panic(err)
	}
	defer db.Close()

	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	topics := []string{"balances"}

	/*kafkaConsumer := kafka.NewConsumer(configMap, topics)

	eventDispatcher := events.NewEventDispatcher()
	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispatcher.Register("UpdatedBalance", handler.NewUpdateBalanceKafkaHandler(kafkaConsumer))
	//database := transaction_db.NewTransactionDB(db)

	transaction, err := create_transaction.NewCreateTransactionUseCase(eventDispatcher, balanceUpdatedEvent)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}*/

	fmt.Println("app is running")
	c, err := ckafka.NewConsumer(configMap)
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition)
			json := json.Unmarshal(msg.Value, *transaction.Transaction)
			fmt.Println(json)
		}
	}
}
