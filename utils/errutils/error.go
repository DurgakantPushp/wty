package errutils

import "log"

func ChkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
