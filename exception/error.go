package exception

import "log"

func PanicIfNeeded(err interface{}) {
	if err != nil {
		log.Panic(err)
	}
}
