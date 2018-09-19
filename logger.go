package logger

import (
	"io"
	"log"

	"github.com/fatih/color"
)

var (
	// Info logger
	info *log.Logger
)

// Init the io.Writers
func Init(infoHandle io.Writer) {
	info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime)
}

// Info logs in color
func Info(format string, v ...interface{}) {

	color.Set(color.FgYellow)

	info.Printf(format, v...)

	color.Unset()
}

// Fatalf Printf
func Fatalf(format string, v ...interface{}) {

	color.Set(color.FgYellow)

	log.Printf(format, v...)

	color.Unset()
}

// Printf printf
func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
