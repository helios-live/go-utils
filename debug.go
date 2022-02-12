package utils

import (
	"io"
	"log"

	"github.com/davecgh/go-spew/spew"
)

// DebugLevel is the level of debug verbosity
var DebugLevel int = 99

// Debug writes the p params if level <= utils.DebugLevel
func Debug(level int, p ...interface{}) {
	if level <= DebugLevel {
		log.Println(p...)
	}
}

// Debugf writes the formated format with p params if level <= utils.DebugLevel
func Debugf(level int, format string, p ...interface{}) {
	if level <= DebugLevel {
		log.Printf(format, p...)
	}
}

// PrintReader Prints whatever it reads
type PrintReader struct {
	io.Reader
	Prefix string
}

func (p PrintReader) Read(b []byte) (n int, err error) {
	n, err = p.Reader.Read(b)
	log.Println(p.Prefix, "PrintReader", n, err)
	if n > 0 {
		spew.Dump(b[0:n])
	}
	return n, err
}
