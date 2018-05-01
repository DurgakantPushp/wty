package connutils

import (
	"bytes"
	"encoding/gob"
	"log"
	"net"

	"github.com/pkg/errors"
	"github.com/wty/bchain/cmd"
)

// SendCmdData sends to conn
func SendCmdData(addr, cmdtype string, data interface{}) error {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(data); err != nil {
		log.Println("error in encoding data, error: ", err)
		return err
	}
	payload := bytes.Join([][]byte{cmd.Command2Bytes(cmdtype),
		buf.Bytes()}, []byte{})

	err := Send(addr, payload)

	log.Println(cmdtype, addr, "sent")
	return err
}

// Send sends payload to given addr on tcp4
func Send(addr string, payload []byte) error {
	conn, err := net.Dial("tcp4", addr)
	if err != nil {
		return errors.WithMessage(err, "error in connecting to center ")
	}
	defer conn.Close()

	_, err = conn.Write(payload)
	return err
}
