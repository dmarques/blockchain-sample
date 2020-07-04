package services

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/blockchain-sample/domain"
	"github.com/blockchain-sample/utils"
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

func AddNewBlock(block domain.Block) {

	//parser arquivo /tmp/json e add
	chain := utils.ReadFileAndParse()
	chain = append(chain, block)

	// gBlock := &domain.BlockChain{
	// 	services.CreateGenesisBlock(),
	// }

	//Convert Struct to JSON File
	b, err := json.Marshal(chain)
	if err != nil {
		fmt.Println(err, "Error convert to JSON file!")
		return
	}

	utils.CreateAndWriteToFile(string(b))
}

func GetLastestBlock() domain.Block {

	chain := utils.ReadFileAndParse()
	lastBlock := chain[len(chain)-1]

	return lastBlock

}
