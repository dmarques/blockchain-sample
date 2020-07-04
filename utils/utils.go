package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/blockchain-sample/domain"
)

func ReadFileAndParse() []domain.Block {

	file, _ := ioutil.ReadFile("/tmp/chain.json")

	data := []domain.Block{}
	_ = json.Unmarshal([]byte(file), &data)

	return data
}

func CreateAndWriteToFile(data string) {
	//Create Temp File
	file, err := os.Create("/tmp/chain.json")
	if err != nil {
		fmt.Println(err, "Error creating file!")
		return
	}

	//Write to JSON file
	_, err = file.WriteString(data)
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
