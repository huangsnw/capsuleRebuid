package config

import (
	"log"
)

func LogConfig() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[SYSTEM LOG]")
}
