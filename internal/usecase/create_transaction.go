package create_transaction

import (
	"github.com/CelsoTaliatelli/ms-walletconsumer/internal/database/transaction_db"
	"github.com/CelsoTaliatelli/ms-walletconsumer/internal/entity"
	"github.com/CelsoTaliatelli/ms-walletconsumer/pkg/events"
)

type CreateTransactionInputDTO struct {
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type BalanceUpdatedOutputDTO struct {
	ID            string  `json:"id"`
	AccountIDFrom string  `json:"account_id_from"`
	AccountIDTo   string  `json:"account_id_to"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionUseCase struct {
	EventDispatcher events.EventDispatcherInterface
	BalanceUpdated  events.EventInterface
}

func NewCreateTransactionUseCase(
	eventDispatcher events.EventDispatcherInterface,
	transactionCreated events.EventInterface,
	balanceUpdated events.EventInterface,
) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		EventDispatcher: eventDispatcher,
		BalanceUpdated:  balanceUpdated,
	}
}

func (uc *CreateTransactionUseCase) Execute(transaction_db transaction_db.TransactionDB, input CreateTransactionInputDTO) (*BalanceUpdatedOutputDTO, error) {
	output := &BalanceUpdatedOutputDTO{}
	balanceUpdatedOutput := &BalanceUpdatedOutputDTO{}

	transaction, err := entity.NewTransaction(input.AccountIDFrom, input.AccountIDTo, input.Amount)
	if err != nil {
		return err
	}

	err = transaction.Create(transaction)
	if err != nil {
		return err
	}
	output.ID = transaction.ID
	output.AccountIDFrom = input.AccountIDFrom
	output.AccountIDTo = input.AccountIDTo
	output.Amount = input.Amount

	balanceUpdatedOutput.AccountIDFrom = input.AccountIDFrom
	balanceUpdatedOutput.AccountIDTo = input.AccountIDTo
	balanceUpdatedOutput.Amount = input.Amount

	uc.BalanceUpdated.SetPayload(balanceUpdatedOutput)
	uc.EventDispatcher.Dispatch(uc.BalanceUpdated)
	return output, nil
}
