package services

import (
	"github.com/garyjdn/djubli-api-go/domain/transactions"
	"github.com/garyjdn/djubli-api-go/utils/errors"
)

var (
	TransactionService transactionServiceInterface = &transactionService{}
)

type transactionService struct{}

type transactionServiceInterface interface {
	GetTransaction() (transactions.Transactions, *errors.RestErr)
}

func (s *transactionService) GetTransaction() (transactions.Transactions, *errors.RestErr) {
	dao := &transactions.Transaction{}
	return dao.GetAll()
}
