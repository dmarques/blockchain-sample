package config

import (
	"encoding/json"
	"fmt"

	"github.com/blockchain-sample/domain"
	"github.com/blockchain-sample/services"
	"github.com/blockchain-sample/utils"
)

func ConfigChainFile() {
	//Generate the chain, the genesis block and add to first position on chain
	gBlock := &domain.BlockChain{
		services.CreateGenesisBlock(),
	}

	//Convert Struct to JSON File
	b, err := json.Marshal(gBlock)
	if err != nil {
		fmt.Println(err, "Error convert to JSON file!")
		return
	}

	utils.CreateAndWriteToFile(string(b))

}
