package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/changjunpyo/nomadcoin/db"
	"github.com/changjunpyo/nomadcoin/utils"
	"strings"
)

const difficulty int = 2

type Block struct {
	Data       string `json:"data"`
	Hash       string `json:"hash"`
	PrevHash   string `json:"prevHash,omitempty"`
	Height     int    `json:"height"`
	Difficulty int    `json:"difficulty"`
	Nonce      int    `json:"nonce"`
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromByte(b, data)
}

var ErrNotFound = errors.New("block not fount")

func FindBlock(hash string) (*Block, error) {
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	return block, nil
}

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		blockAsString := fmt.Sprint(b)
		hash := fmt.Sprintf("%x",sha256.Sum256([]byte(fmt.Sprint(b))))
		fmt.Printf("Block as String:%s\nHash:%s\nTarget:%s\nNonce:%d\n\n\n",blockAsString, hash,target,b.Nonce)
		if strings.HasPrefix(hash, target){
			b.Hash = hash
			return
		} else{
			b.Nonce++
		}
	}
}


func createBlock(data string, prevHash string, height int) *Block {
	block := Block{
		Data:       data,
		Hash:       "",
		PrevHash:   prevHash,
		Height:     height,
		Difficulty: difficulty,
		Nonce:      0,
	}
	block.mine()
	block.persist()
	return &block
}
