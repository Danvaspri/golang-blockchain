package blockchain
type BlockChain struct {
	Blocks []*Block
}
type Block struct{
	Hash 		[]byte
	Data  		[]byte
	PrevHash 	[]byte
	Nonce 		int
}
//function to derive the hash of the block based on the previous block hash
//and the current block data, joined together and passed through the hashing
//algoritm


func CreateBlock(data string, prevHash []byte) *Block  {	
	block := &Block{[]byte{},[]byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	
	return block
}

func (chain *BlockChain) AddBlock(data string){
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data,prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

func Genesis() *Block{
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain{
	return &BlockChain{[]*Block{Genesis()}}
}