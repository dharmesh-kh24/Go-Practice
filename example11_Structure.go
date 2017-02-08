package main
import "fmt"

type Blockchain struct {
	username string
	txId string
	blockNo int
}

func printChain (chain Blockchain) {
	fmt.Printf("User name = %s\n",chain.username)
	fmt.Printf("txId = %s\n",chain.txId)
	fmt.Printf("blockNo = %d",chain.blockNo)
}

func main(){

	var example Blockchain
	example.username = "mohsin"
	example.txId = "lsjiuen46556@cvnjkh"
	example.blockNo = 7
	printChain(example)

}