package models

type BroadcastTransactionRequest struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

type BroadcastTransactionResponse struct {
	TXHash string `json:"tx_hash"`
}
