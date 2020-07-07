package domain

import "time"

type BlockChain []struct {
	Index        int64     `json:"index"`
	DateCreated  time.Time `json:"date_created"`
	Data         string    `json:"data"`
	PreviousHash string    `json:"previous_hash"`
	Hash         string    `json:"hash"`
	Nonce        int64     `json:"nounce"`
}
