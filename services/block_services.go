package services

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/blockchain-sample/domain"
)

func CreateBlockServices(data string) (domain.Block, error) {

	// index int64, timestamp time.Time, data string, previousHash string
	var block domain.Block
	var previousHash string

	block.Index = 1
	block.DateCreated = time.Now()
	block.Hash = GenerateHash(1, time.Now(), data, previousHash)
	block.PreviousHash = previousHash
	block.Data = data

	return block, nil

}

func GenerateHash(index int64, timestamp time.Time, data string, previousHash string) string {

	date := timestamp.Format("2006-01-02 00:00:00")
	sha256 := sha256.Sum256([]byte(strconv.FormatInt(index, 10) + date + data + previousHash))

	return fmt.Sprintf("%x", sha256)
}
