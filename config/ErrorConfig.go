package config

import "log"

func PrintError(e error, message string) {
	if e != nil {
		log.Println(message, e)
	}
}

func PanicError(e error, message string) {
	if e != nil {
		log.Panic(message)
	}
}
