package bchain

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/wty/utils/errutils"
)

const (
	blocksBucket = "block"
)

// Block represents any data along with its fingerprint and timestamp
type Block struct {
	// block headers
	Timestamp     time.Time `json:"timestamp"`
	PrevBlockHash []byte    `json:"prevHash"`
	Hash          []byte    `json:"hash"`
	Nonce         int       `json:"nonce"`
	// transaction or any data
	Data []byte `json:"data"`
}

// NewBlock returns a Block pointer
func NewBlock(data, prevHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now(),
		PrevBlockHash: prevHash,
		Data:          data,
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// NewGenesisBlock returns the first block
func NewGenesisBlock() *Block {
	return NewBlock([]byte("Genesis block"), []byte{})
}

// Serialize serializes a block in a byte array
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	errutils.ChkErr(err)
	return result.Bytes()
}

// DeserializeBlock deserializes a byte array in a block
func DeserializeBlock(data []byte) *Block {
	var block Block

	err := gob.NewDecoder(bytes.NewReader(data)).Decode(&block)
	errutils.ChkErr(err)
	return &block
}

func (b *Block) String() string {
	return fmt.Sprintf("\nTimestamp: %v\nHash: %s\nPrevHash: %s\nData: %s\nNonce: %v",
		b.Timestamp, hex.EncodeToString(b.Hash), hex.EncodeToString(b.PrevBlockHash),
		hex.EncodeToString(b.Data), b.Nonce)
}
