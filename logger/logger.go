package logger

import (
	"log"
	"os"
)

var Info *log.Logger
var Error *log.Logger
var Warn *log.Logger
var Debug *log.Logger

func init() {
	Info = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	Error = log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Llongfile)
	Warn = log.New(os.Stdout, "WARN: ", log.LstdFlags|log.Lshortfile)
	Debug = log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)
}
