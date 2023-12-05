package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CelsoTaliatelli/ms-walletconsumer/internal/usecase/create_transaction"
)

type WebAccountHandler struct {
	createTransactionUseCase create_transaction.createTransactionUseCase
}

func NewWebTransactionHandler(createTransactionUseCase create_transaction.createTransactionUseCase) *WebTransactionHandler {
	return &WebTransactionHandler{
		createTransactionUseCase: createTransactionUseCase,
	}
}

func (h *WebTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var dto create_transaction.CreateAccountInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	output, err := h.createTransactionUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
