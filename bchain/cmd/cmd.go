package cmd

import (
	"errors"
	"fmt"
)

const (
	// CommandLen is max length of command
	CommandLen = 20
)

var (
	errLessData = errors.New("error: data length less than command length")
)

// Bytes2Command return command from byte slice
func Bytes2Command(data []byte) (string, error) {
	if len(data) < CommandLen {
		return "", errLessData
	}

	var cmd []byte
	for i := 0; i < CommandLen; i++ {
		if data[i] != 0x0 {
			cmd = append(cmd, data[i])
		}
	}
	return fmt.Sprintf("%s", cmd), nil
}

// Command2Bytes returns []byte from string command
func Command2Bytes(cmd string) []byte {
	var bcmd [CommandLen]byte

	for i, c := range cmd {
		bcmd[i] = byte(c)
	}
	return bcmd[:]
}

// GetCmdData returns cmd and data from input data
func GetCmdData(data []byte) (command string, out []byte, err error) {
	if len(data) < CommandLen {
		err = errLessData
		return
	}
	command, err = Bytes2Command(data[:CommandLen])

	if err != nil {
		return
	}
	out = data[CommandLen:]

	return
}
