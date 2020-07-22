package utils

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	L *lumberjack.Logger
)

func init() {
	L = &lumberjack.Logger{
		Filename:   "lumberjack.log",
		MaxSize:    1, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
		// Compress:   true, // disabled by default
	}
	log.SetOutput(L)
	// l.Write([]byte("Herreee"))
}
