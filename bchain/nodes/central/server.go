package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/wty/bchain"
	"github.com/wty/bchain/cmd"
	"github.com/wty/bchain/version"
	"github.com/wty/models"
	eu "github.com/wty/utils/errutils"
)

const (
	port = "3001"
	host = "localhost"
)

type mineType int

const (
	unmined mineType = iota
	mining
	mined
	validated
	added
)

var (
	mapGratitude sync.Map
	myVersion    version.Version
	bc           *bchain.Blockchain
)

func init() {
	myVersion.AddrFrom = host + ":" + port
}

func main() {
	dbFile := fmt.Sprintf(bchain.DBFileFmt, port)

	db, err := bolt.Open(dbFile, 0600, nil)
	eu.ChkErr(errors.WithMessage(err, "opening dbfile error"))
	defer db.Close()

	bc = bchain.NewBlockchain(db)
	myVersion.BestHeight = bc.GetHeight()
	log.Println("blockchain initialized, height", myVersion.BestHeight)

	// update admin for current blockchain
	sendBlocks2Admin()

	ln, err := net.Listen("tcp4", host+":"+port)
	eu.ChkErr(err)

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		eu.ChkErr(err)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println("invalid data received, error: ", err)
		return
	}

	cmd, data, err := cmd.GetCmdData(data)
	if err != nil {
		log.Println("error in fetching command and data, error:", err)
		return
	}

	log.Println("command received", cmd)

	switch cmd {
	case "addGratitude":
		err := addGratitude(data)
		if err != nil {
			log.Println("Gratitude failed to add to mempool")
			return
		}
		log.Println("Gratitude added to mempool")
		//notify
	case "version":
		version.Handle(data, myVersion, bc)
	case "block":
		update(data)
	case "validated":
		addBlock(data)
	default:
		log.Println("invalid command received ", cmd)
	}
}

// decodes gratitude from data and stores it in mempool
func addGratitude(data []byte) error {
	var grat models.Gratitude

	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(grat)
	if err != nil {
		return err
	}
	mapGratitude.Store(grat, unmined)
	return nil
}

func addBlock(data []byte) {
	var block bchain.Block
	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&block)
	if err != nil {
		log.Println("error in decoding blocks ", err)
		return
	}

	if err := bc.AddBlock(&block); err != nil {
		log.Println("error in adding validated block to blockchain ", err)
		return
	}
	myVersion.BestHeight++

	go updateBlock2Admin(&block)

	log.Println("Validated block added successfully", myVersion.BestHeight)
}

func update(data []byte) {
	var blocks []*bchain.Block
	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&blocks)
	if err != nil {
		log.Println("error in decoding blocks ", err)
		return
	}

	err = bc.Update(blocks)
	if err != nil {
		log.Println("error in updating blockchain ", err)
	}
	log.Println("centre blockchain updated successfully")
}
