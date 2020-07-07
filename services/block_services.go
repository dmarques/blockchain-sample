package services

import (
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/blockchain-sample/domain"
)

func CreateBlockServices(data string) (domain.Block, error) {

	var block domain.Block
	var previousHash string

	lastBlock := GetLastestBlock()

	block.Index = lastBlock.Index + 1
	block.DateCreated = time.Now()
	block.Hash = GenerateHash(1, block.DateCreated, data, previousHash, 0)
	block.PreviousHash = lastBlock.Hash
	block.Data = data
	block.Nonce = 0

	return block, nil

}

func GenerateHash(index int64, timestamp time.Time, data string, previousHash string, nonce int64) string {

	date := timestamp.Format("2006-01-02")
	sha256 := sha256.Sum256([]byte(strconv.FormatInt(index, 10) + date + data + previousHash + strconv.FormatInt(nonce, 10)))

	return fmt.Sprintf("%x", sha256)
}

func Mine(difficulty int, newBlock *domain.Block) {

	log.Println("Mining the new Block...")

	validatorSize := make([]string, difficulty)

	for i := 0; i < len(validatorSize); i++ {
		validatorSize[i] = "0"
	}
	validator := strings.Join(validatorSize, "")

	for strings.TrimSpace(newBlock.Hash[:difficulty]) != strings.TrimSpace(validator) {
		newBlock.Nonce++
		newBlock.Hash = GenerateHash(newBlock.Index, newBlock.DateCreated, newBlock.Data, newBlock.PreviousHash, newBlock.Nonce)
	}
}
