package transactions

type Transaction struct {
	Id       int64  `json:"id"`
	Brand    string `json:"brand"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Type     string `json:"type"`
	Color    string `json:"color"`
	Km       int    `json:"km"`
	Location string `json:"location"`
	Price    int64  `json:"price"`
	Date     string `json:"date"`
}

type Transactions []Transaction
