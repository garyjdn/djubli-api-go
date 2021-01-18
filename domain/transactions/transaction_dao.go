package transactions

import (
	"github.com/garyjdn/djubli-api-go/datasources/mysql/transaction_db"
	"github.com/garyjdn/djubli-api-go/utils/errors"
)

const (
	queryGetTransactions = "SELECT `id`, `brand`, `name`, `year`, `type`, `color`, `km`, `location`, `price`, `date` FROM `transactions`;"
)

func (transaction *Transaction) GetAll() ([]Transaction, *errors.RestErr) {
	stmt, err := transaction_db.Client.Prepare(queryGetTransactions)
	if err != nil {
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.InternalServerError("database error")
	}
	defer rows.Close()

	res := make([]Transaction, 0)
	for rows.Next() {
		var transaction Transaction
		if err := rows.Scan(&transaction.Id, &transaction.Brand, &transaction.Name, &transaction.Year, &transaction.Type, &transaction.Color, &transaction.Km, &transaction.Location, &transaction.Price, &transaction.Date); err != nil {
			return nil, errors.InternalServerError("database error")
		}
		res = append(res, transaction)
	}

	return res, nil
}
