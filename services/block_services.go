package services

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"

	"github.com/blockchain-sample/domain"
)

func CreateBlockServices(index int64, timestamp time.Time, data string, previousHash string) (domain.Block, error) {

	var block domain.Block

	return block, nil

}

func GenerateHash(blockData domain.Block) string {

	index := strconv.FormatInt(blockData.Index, 10)
	date := blockData.DateCreated.Format("2006-01-02 00:00:00")

	sha256 := sha256.Sum256([]byte(index + date + blockData.Data + blockData.PreviousHash))

	return fmt.Sprintf("%x", sha256)
}
