package global

import (
	"io"
	"log"
	"time"
)

func InitLogError(errorHandle io.Writer) {
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitDefaultQueryTimeOut(queryTimeout int64) {
	QueryTimeout = time.Second * time.Duration(queryTimeout)
}
