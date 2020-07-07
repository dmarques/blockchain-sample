package services

import (
	"encoding/json"
	"fmt"

	"github.com/blockchain-sample/domain"
	"github.com/blockchain-sample/utils"
)

func AddNewTransaction(transaction domain.Transaction) {
	//Parser arquivo /tmp/chain.json e add
	pendingTrans := utils.ReadPendingTransactionFile()
	pendingTrans = append(pendingTrans, transaction)

	//Convert Struct to JSON File
	b, err := json.Marshal(pendingTrans)
	if err != nil {
		fmt.Println(err, "Error convert to JSON file!")
		return
	}

	utils.CreatePendingTransactionFile(string(b))
}

//Process the pending transactions and send to add into chain
func ProcessPendingTransactions() {
	//Difficulty to mine
	difficulty := 2
	miningRewards := 100.0

	//Parser arquivo /tmp/pending_transaction.json e add
	pendingTrans := utils.ReadPendingTransactionFile()

	//Generate new block with pending transactions
	newBlock, _ := CreateBlock(pendingTrans)

	//Mining the new Block
	Mine(difficulty, &newBlock)

	//Send to add into chain
	AddBlockToChain(newBlock)

	//Rewards
	utils.CreatePendingTransactionFile("[]")
	rewardTrans := domain.Transaction{
		FromAddress: "",
		ToAddress:   "mining_rewards",
		Amount:      miningRewards,
	}
	AddNewTransaction(rewardTrans)

}
