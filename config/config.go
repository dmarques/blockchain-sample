package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/blockchain-sample/domain"
	"github.com/blockchain-sample/services"
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

	//Create Temp File
	file, err := os.Create("/tmp/chain.json")
	if err != nil {
		fmt.Println(err, "Error creating file!")
		return
	}

	//Write to JSON file
	_, err = file.WriteString(string(b))
	if err != nil {
		fmt.Println("Error write on file")
		file.Close()
		return
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ReadFileAndParse() []domain.Block {

	file, _ := ioutil.ReadFile("/tmp/chain.json")

	data := []domain.Block{}
	_ = json.Unmarshal([]byte(file), &data)

	return data
}
