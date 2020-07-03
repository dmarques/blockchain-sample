package domain

import "time"

type Block struct {
	Index        int64     `json:"id"`
	DateCreated  time.Time `json:"date_created"`
	Data         string    `json:"data"`
	PreviousHash string    `json:"previous_hash"`
	Hash         string    `json:"hash"`
}
