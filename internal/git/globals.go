package git

import (
	"io"
	"log"
)

var logger *log.Logger = log.New(io.Discard, "", 0)
