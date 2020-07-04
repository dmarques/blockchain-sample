package services

import (
	"time"

	"github.com/blockchain-sample/domain"
)

func CreateGenesisBlock() domain.Block {
	return domain.Block{
		Index:        0,
		DateCreated:  time.Now(),
		Data:         "genesis_block",
		PreviousHash: "",
		Hash:         GenerateHash(0, time.Now(), "genesis_block", ""),
	}
}

func AddNewBlock() {

	//parser arquivo /tmp/json e add
}

func GetLastestBlock() {

}
