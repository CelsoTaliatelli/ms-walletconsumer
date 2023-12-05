package handler

import (
	"fmt"
	"sync"

	"github.com/CelsoTaliatelli/ms-walletconsumer/pkg/events"
	"github.com/CelsoTaliatelli/ms-walletconsumer/pkg/kafka"
)

type UpdateBalanceKafkaHandler struct {
	Kafka *kafka.Consumer
}

func NewUpdateBalanceKafkaHandler(kafka *kafka.Consumer) *UpdateBalanceKafkaHandler {
	return &UpdateBalanceKafkaHandler{
		Kafka: kafka,
	}
}

func (h *UpdateBalanceKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Consume()
	fmt.Println("UpdateBalanceKafkaHandler called")
}
