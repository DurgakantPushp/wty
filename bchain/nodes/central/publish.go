package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
	"github.com/wty/bchain"
	"github.com/wty/bchain/mqtt"
	"github.com/wty/models"
)

const (
	topicBlocksInit   = "centre/blocks/init"
	topicBlocksUpdate = "centre/blocks/update"
)

// initializes admin with current blockchain
func sendBlocks2Admin() {
	blocks := bc.GetAllBlocks()

	if len(blocks) == 1 {
		log.Println("No gratitude added, so nothing to publish...")
		return
	}

	for _, block := range blocks {
		data, err := gob2JSONByte(block.Data)
		if err != nil {
			log.Println("gob to json conversion error:", err)
			continue
		}

		block.Data = data
	}
	bdata, err := json.Marshal(&blocks)
	if err != nil {
		log.Println("Error in marshaling blocks", err)
		return
	}

	err = mqtt.Publish("c", string(bdata))
	if err != nil {
		log.Println("Error in updating blockchain", err)
	}
	log.Println("Admin initialized for current blockchain", blocks)
}

func updateBlock2Admin(block *bchain.Block) {
	data, err := gob2JSONByte(block.Data)
	if err != nil {
		log.Println("gob to json conversion error:", err)
		return
	}
	block.Data = data

	bdata, err := json.Marshal(block)
	if err != nil {
		log.Println("Error in marshaling block", err)
		return
	}

	err = mqtt.Publish(topicBlocksUpdate, string(bdata))
	if err != nil {
		log.Println("Error in updating blockchain", err)
	}
	log.Println("Admin updated for current block", block)
}

func gob2JSONByte(in []byte) ([]byte, error) {
	var grat models.Gratitude

	err := gob.NewDecoder(bytes.NewBuffer(in)).Decode(&grat)
	if err != nil {
		return nil, errors.WithMessage(err, "Error: gob decoding block data")
	}

	var buf bytes.Buffer

	err = json.NewEncoder(&buf).Encode(&grat)
	if err != nil {
		return nil, errors.WithMessage(err, "Error: json encoding block data")
	}
	return buf.Bytes(), nil
}
