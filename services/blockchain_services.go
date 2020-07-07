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
		DateCreated: time.Now(),
		Data: domain.Transactions{
			domain.Transaction{
				FromAddress: "genesis_block",
				ToAddress:   "genesis_block",
				Amount:      0,
			},
		},
		PreviousHash: "",
		Hash:         GenerateHash(time.Now(), domain.Transactions{domain.Transaction{FromAddress: "genesis_block", ToAddress: "genesis_block", Amount: 0}}, "", 0),
		Nonce:        0,
	}
}

//Add processed block into Chain
func AddBlockToChain(block domain.Block) {

	//Parser arquivo /tmp/chain.json e add
	chain := utils.ReadChainFileAndParse()
	chain = append(chain, block)

	//Convert Struct to JSON File
	b, err := json.Marshal(chain)
	if err != nil {
		fmt.Println(err, "Error convert to JSON file!")
		return
	}

	utils.CreateChainFile(string(b))
}

func GetLastestBlock() domain.Block {

	chain := utils.ReadChainFileAndParse()
	lastBlock := chain[len(chain)-1]

	return lastBlock

}
