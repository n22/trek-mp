package global

import (
	"io"
	"log"
	"time"

	"github.com/5112100070/trek-mp/src/app/product"
)

func InitLogError(errorHandle io.Writer) {
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func InitDefaultQueryTimeOut(queryTimeout int64) time.Duration {
	return time.Second * time.Duration(queryTimeout)
}

func InitRepoBundle(dbBundle DBBundle, queryTimeout time.Duration) {
	Services = RepoBundle{
		Product: product.InitProductRepo(dbBundle.Product, queryTimeout),
	}
}
