package log

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout,"", log.Ldate | log.Ltime )

func Log(title interface{}, content interface{}) {
	logger.Printf("%s: %s", title, content)
}
