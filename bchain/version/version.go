package version

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/wty/bchain"
	"github.com/wty/bchain/cmd"
	"github.com/wty/bchain/connutils"
	"github.com/wty/utils/errutils"
)

// Version represents each node blockchain version
type Version struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

// Decode decode version from data
func Decode(data []byte) (Version, error) {
	var vers Version

	err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&vers)
	return vers, err
}

// Handle handles version payload
func Handle(data []byte, myVersion Version, bc *bchain.Blockchain) {
	version, err := Decode(data)
	errutils.ChkErr(err)

	diffHeight := myVersion.BestHeight - version.BestHeight
	if diffHeight >= 0 {
		err := sendBlocks(version.AddrFrom, diffHeight, bc)
		if err != nil {
			log.Println("block send failed, error: ", err)
		} else {
			log.Println("blocks sent successfully, #blocks", diffHeight)
		}
	} else if diffHeight < 0 {
		err := sendVersion(version.AddrFrom, myVersion)
		if err != nil {
			log.Println("version send failed, error: ", err)
		} else {
			log.Println("version sent successfully from", myVersion.AddrFrom, "to", version.AddrFrom)
		}
	}
}

func sendBlocks(addr string, diffHeight int, bc *bchain.Blockchain) (err error) {
	// get diff blocks
	blocks := bc.GetBlocks(diffHeight)

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(blocks); err != nil {
		log.Println("error in encoding blocks, error: ", err)
		return err
	}
	payload := bytes.Join([][]byte{cmd.Command2Bytes("block"),
		buf.Bytes()}, []byte{})

	return connutils.Send(addr, payload)
}

// sendVersion sends version
func sendVersion(addr string, vers Version) (err error) {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(vers); err != nil {
		log.Println("error in encoding version, error: ", err)
		return err
	}
	payload := bytes.Join([][]byte{cmd.Command2Bytes("version"),
		buf.Bytes()}, []byte{})

	return connutils.Send(addr, payload)
}
