package cmd

import (
	"io"
	"log"
	"os"
)

func newLogger(verbose bool) *log.Logger {
	if verbose {
		return log.New(os.Stderr, "", log.LstdFlags)
	} else {
		return log.New(io.Discard, "", 0)
	}
}
