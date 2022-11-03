package util

import "log"

// HandleError must give a fatal error if "err" is different of null
func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
