package bchain

import (
	"bytes"
	"errors"

	"github.com/boltdb/bolt"
	eu "github.com/wty/utils/errutils"
)

// Blockchain stores blocks consisting of gratitude
type Blockchain struct {
	DB  *bolt.DB
	tip []byte
}

var (
	errInvalidBlock = errors.New("Invalid Block")
)

// AddBlock adds a block at the end of chain
func (bc *Blockchain) AddBlock(block *Block) error {
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if bytes.Equal(bc.tip, block.PrevBlockHash) {
			if err := b.Put(block.Hash, block.Serialize()); err != nil {
				return err
			}

			if err := b.Put([]byte("l"), block.Hash); err != nil {
				return err
			}

			bc.tip = block.Hash
		} else {
			return errInvalidBlock
		}
		return nil
	})
	return err
}

// Mine mines a block at the end of chain
// also check max transactions only then mine
func (bc *Blockchain) Mine(data []byte) (*Block, error) {
	var lastHash []byte

	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		return nil, err
	}

	block := NewBlock(data, lastHash)

	return block, nil
}

// Iterator returns blockchain iterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.DB}
}

// NewBlockchain creates and returns a blockchain
func NewBlockchain(db *bolt.DB) *Blockchain {
	var tip []byte

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			eu.ChkErr(err)
			err = b.Put(genesis.Hash, genesis.Serialize())
			eu.ChkErr(err)

			err = b.Put([]byte("l"), genesis.Hash)
			eu.ChkErr(err)

			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	eu.ChkErr(err)

	return &Blockchain{db, tip}
}

// NewBlockchainMiner creates and returns a blockchain
func NewBlockchainMiner(db *bolt.DB) (*Blockchain, error) {
	var tip []byte

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			btmp, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			b = btmp
		} else {
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Blockchain{db, tip}, nil
}

// GetAllBlocks returns all blocks
func (bc *Blockchain) GetAllBlocks() []*Block {
	var bl []*Block

	bci := bc.Iterator()
	for {
		block := bci.Next()
		bl = append(bl, block)

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return bl
}

// GetHeight returns number of blocks
func (bc *Blockchain) GetHeight() (height int) {
	bci := bc.Iterator()
	for {
		block := bci.Next()
		if block == nil {
			return
		}
		height++

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return
}

// GetBlocks get block upto height lenght
func (bc *Blockchain) GetBlocks(height int) []*Block {
	if height < 1 {
		return nil
	}

	var bl []*Block
	h := 0
	bci := bc.Iterator()

	for {
		block := bci.Next()
		bl = append(bl, block)
		h++
		if h == height {
			break
		}
		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
	return bl
}

// Update updates current blockchain
func (bc *Blockchain) Update(blocks []*Block) error {
	// check nil blockchain
	err := bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			btmp, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				return err
			}
			b = btmp
		}

		for _, block := range blocks {
			err := b.Put(block.Hash, block.Serialize())
			if err != nil {
				return err
			}
			err = b.Put([]byte("l"), block.Hash)
			if err != nil {
				return err
			}

			bc.tip = block.Hash
		}
		return nil
	})
	return err
}
