package main

import (
	"fmt"
	"log"

	"github.com/umaibou1126/goblockchain/block"
	"github.com/umaibou1126/goblockchain/wallet"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	// w := wallet.NewWallet()
	// fmt.Println(w.PrivateKeyStr())
	// fmt.Println(w.PublicKeyStr())
	// fmt.Println(w.BlockchainAddress())

	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	blockchain := block.NewBlockchain(walletM.BlockchainAddress())
	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("Added?", isAdded)
	fmt.Printf("A %.1f\n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("M %.1f\n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
	// fmt.Printf("signature %s\n", t.GenerateSignature())
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
