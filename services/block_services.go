package services

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/blockchain-sample/domain"
)

func CreateBlock(transactions domain.Transactions) (domain.Block, error) {

	var block domain.Block
	var previousHash string

	lastBlock := GetLastestBlock()

	block.DateCreated = time.Now()
	block.Hash = GenerateHash(block.DateCreated, transactions, previousHash, 0)
	block.PreviousHash = lastBlock.Hash
	block.Data = transactions
	block.Nonce = 0

	return block, nil

}

func GenerateHash(timestamp time.Time, txs domain.Transactions, previousHash string, nonce int64) string {

	date := timestamp.Format("2006-01-02")
	strTxs, _ := json.Marshal(txs)
	sha256 := sha256.Sum256([]byte(date + string(strTxs) + previousHash + strconv.FormatInt(nonce, 10)))

	return fmt.Sprintf("%x", sha256)
}

func Mine(difficulty int, newBlock *domain.Block) {

	validatorSize := make([]string, difficulty)

	for i := 0; i < len(validatorSize); i++ {
		validatorSize[i] = "0"
	}
	validator := strings.Join(validatorSize, "")

	for strings.TrimSpace(newBlock.Hash[:difficulty]) != strings.TrimSpace(validator) {
		newBlock.Nonce++
		newBlock.Hash = GenerateHash(newBlock.DateCreated, newBlock.Data, newBlock.PreviousHash, newBlock.Nonce)
	}
}
