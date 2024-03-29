package lg

import (
	"log"
	"os"
)

func init() {
	Info = log.New(os.Stdout, "INFO  ", log.LstdFlags)
	Err = log.New(os.Stdout, "\u001b[31mERROR \u001b[0m", log.LstdFlags)
}

var (
	Info *log.Logger
	Err  *log.Logger
)
