package bchain

import (
	"errors"

	"github.com/boltdb/bolt"
)

// BlockchainIterator iterates over a blockchain
type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// Next returns current block and moves the current cursor to
// a previous block
func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		if encodedBlock == nil {
			return errors.New("no block")
		}
		block = DeserializeBlock(encodedBlock)
		return nil
	})
	if err != nil {
		return nil
	}
	i.currentHash = block.PrevBlockHash
	return block
}
