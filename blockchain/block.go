package blockchain


type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int

}

func CreateBlock(data string, prevHash[]byte) *Block {
	block := &Block{Data:  []byte(data),  PrevHash: prevHash}
	pow := NewProof(block)
	block.Nonce, block.Hash = pow.Run()
	return block
}

func (ch *BlockChain) AddBlock(data string)  {
	prevBlock := ch.Blocks[len(ch.Blocks) - 1]
	n := CreateBlock(data, prevBlock.Hash)
	ch.Blocks = append(ch.Blocks, n)
}

func Genesys() *Block  {
	return CreateBlock("Genesys", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesys()}}
}