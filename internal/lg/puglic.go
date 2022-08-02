package lg

import (
	"log"
	"os"
)

func init() {
	Info = log.New(os.Stdout, "INFO  ", log.LstdFlags)
	Err = log.New(os.Stdout, "ERROR ", log.LstdFlags)
}

var (
	Info *log.Logger
	Err  *log.Logger
)
