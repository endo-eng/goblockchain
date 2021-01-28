package main

import (
	"fmt"
	"log"

	"github.com/umaibou1126/goblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println(w.PublicKeyStr())
	fmt.Println(w.BlockchainAddress())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.BlockchainAddress(), "B", 1.0)
	fmt.Printf("signature %s\n", t.GenerateSignature())
	// myBlockchainAddress := "my_blockchain_address"
	// blockChain := NewBlockchain(myBlockchainAddress)
	// blockChain.Print()

	// blockChain.AddTransaction("A", "B", 1.0)
	// blockChain.Mining()
	// // previousHash := blockChain.LastBlock().Hash()
	// // nonce := blockChain.ProofOfWork()
	// // blockChain.CreateBlock(nonce, previousHash)
	// blockChain.Print()

	// blockChain.AddTransaction("C", "D", 2.0)
	// blockChain.AddTransaction("X", "Y", 3.0)
	// // previousHash = blockChain.LastBlock().Hash()
	// // nonce = blockChain.ProofOfWork()
	// // blockChain.CreateBlock(nonce, previousHash)
	// blockChain.Mining()
	// blockChain.Print()

	// fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	// fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("C"))
	// fmt.Printf("D %.1f\n", blockChain.CalculateTotalAmount("D"))

}
