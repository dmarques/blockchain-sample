package domain

import "time"

type Block struct {
	DateCreated  time.Time    `json:"date_created"`
	Data         Transactions `json:"data"`
	PreviousHash string       `json:"previous_hash"`
	Hash         string       `json:"hash"`
	Nonce        int64        `json:"nounce"`
}
