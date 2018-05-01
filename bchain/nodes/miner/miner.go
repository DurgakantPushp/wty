package main

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/wty/bchain"
	"github.com/wty/bchain/cmd"
	"github.com/wty/bchain/connutils"
	"github.com/wty/bchain/version"
	eu "github.com/wty/utils/errutils"
)

const (
	centreAddr = "localhost:3001"
	valAddr    = "localhost:3011"
)

var (
	myVersion   version.Version
	bc          *bchain.Blockchain
	chanUpdated = make(chan struct{}, 1)
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	port := os.Args[1]
	log.Println("port is", port)
	addr := "localhost:" + port

	myVersion.AddrFrom = addr

	dbFile := fmt.Sprintf(bchain.DBFileFmt, port)

	db, err := bolt.Open(dbFile, 0600, nil)
	eu.ChkErr(errors.WithMessage(err, "opening dbfile error"))
	defer db.Close()

	bc, err = bchain.NewBlockchainMiner(db)
	if err != nil || bc == nil {
		log.Fatalln("error in creating blockchain, error:", err)
	}

	myVersion.BestHeight = bc.GetHeight()
	log.Println("blockchain initialized, height", myVersion.BestHeight)

	ln, err := net.Listen("tcp4", addr)
	eu.ChkErr(err)
	defer ln.Close()

	defer close(chanUpdated)

	for {
		conn, err := ln.Accept()
		eu.ChkErr(err)
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	//	buf := make([]byte, 2560)

	//	_, err := conn.Read(buf)
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
	log.Println("command received:", cmd)

	switch cmd {
	case "version":
		version.Handle(data, myVersion, bc)
	case "block":
		//add received block
		err := addBlocks(data, bc)
		if err != nil {
			log.Println("blocks addition failed, error:", err)
			return
		}
		log.Println("Blockchain updated successfully")

		chanUpdated <- struct{}{}
	case "gratitude":
		update()
		<-chanUpdated
		mineBlock(data)
	}
}

func addBlocks(data []byte, bc *bchain.Blockchain) error {
	var blocks []*bchain.Block

	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&blocks)
	if err != nil {
		log.Println(hex.EncodeToString(data))
		return errors.WithMessage(err, "error in decoding blocks")
	}

	if err := bc.Update(blocks); err != nil {
		return err
	}
	myVersion.BestHeight += len(blocks)

	return nil
}

func mineBlock(data []byte) {
	log.Println("data received for mining")

	block, err := bc.Mine(data)
	if err != nil {
		log.Println("Mining failed. error: ", err)
		return
	}
	log.Println("data mined successfully. block", block)

	// send for validation
	connutils.SendCmdData(valAddr, "validate", block)
}

func update() {
	//update its blockchain from center
	err := connutils.SendCmdData(centreAddr, "version", myVersion)
	if err != nil {
		log.Println("error in updating version ", err)
	}

	log.Println("update request sent successfully")
}
