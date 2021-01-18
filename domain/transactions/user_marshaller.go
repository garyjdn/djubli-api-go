package transactions

import "encoding/json"

func (transactions Transactions) Marshall() []interface{} {
	results := make([]interface{}, len(transactions))
	for index, user := range transactions {
		results[index] = user.Marshall()
	}
	return results
}

func (transaction *Transaction) Marshall() interface{} {
	transactionJson, _ := json.Marshal(transaction)
	var transactionRes Transaction
	json.Unmarshal(transactionJson, &transactionRes)
	return transactionRes
}
