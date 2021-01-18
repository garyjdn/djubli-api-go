package app

import "github.com/garyjdn/djubli-api-go/controllers/transactions"

func mapUrls() {
	router.GET("/transactions", transactions.GetAll)
}
