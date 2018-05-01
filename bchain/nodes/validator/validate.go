package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"net"

	"github.com/wty/bchain/connutils"

	"github.com/wty/bchain"
	"github.com/wty/bchain/cmd"
	"github.com/wty/bchain/version"
	"github.com/wty/utils/errutils"
)

const (
	centreAddr = "localhost:3001"
	addr       = "localhost:3011"
)

var (
	myVersion version.Version
	bc        *bchain.Blockchain
)

func main() {
	ln, err := net.Listen("tcp4", addr)
	errutils.ChkErr(err)
	defer ln.Close()

	log.Println("validator node started on ", addr)
	for {
		conn, err := ln.Accept()
		errutils.ChkErr(err)

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
	case "version":
		version.Handle(data, myVersion, bc)
	case "block":
		//add received block
		addBlocks(data, bc)
	case "validate":
		log.Println("block received for validation")
		validate(data)
	}
}

func addBlocks(data []byte, bc *bchain.Blockchain) {
	var blocks []*bchain.Block

	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&blocks)
	errutils.ChkErr(err)

	if err := bc.Update(blocks); err != nil {
		log.Println("blocks addition failed")
		return
	}
	log.Println("Blockchain updated successfully")
}

func validate(data []byte) {
	var block bchain.Block

	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&block)
	if err != nil {
		log.Println("Validation failed. Error: ", err)
		return
	}

	pow := bchain.NewProofOfWork(&block)

	if pow.Validate() {
		log.Println("Block validated")
		// add to itself and send to server
		if err := connutils.SendCmdData(centreAddr, "validated", block); err != nil {
			log.Println("validated block sent to central node", block)
		}
	} else {
		log.Println("Block validation failed")
	}
}
