package models

type Block struct {
	BlockNumber       int64
	Timestamp         uint64
	Difficulty        uint64
	Hash              string
	TransactionsCount int
	Transactions      []Transaction
}

type Transaction struct {
	Hash string
}
