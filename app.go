package main

import (
	"fmt"
	"log"
	"os"

	"github.com/5112100070/trek-mp/src/conf"
	"github.com/5112100070/trek-mp/src/global"
	"github.com/5112100070/trek-mp/src/product"
	"github.com/gin-gonic/gin"
)

func init() {
	// init error logging
	global.InitLogError(os.Stderr)

	cfgenv := os.Getenv("TKPENV")
	network := os.Getenv("NETWORK")
	if cfgenv == "" {
		log.Println("[trek-mp] No environment set. Using 'development'.")
		log.Println("[trek-mp] Use 'export TKPENV=[development|alpha|staging|production]' to change.")
		cfgenv = "development"
	}

	fileLocation := fmt.Sprintf("/etc/trek-mp/sys-conf/%s.ini", cfgenv)
	log.Println(fmt.Sprintf("Using configuration : %s", fileLocation))
	log.Println(fmt.Sprintf("Running in network : %s", network))

	config, ok := conf.ReadConfig(fileLocation)
	if !ok {
		log.Fatal("Could not find configuration file")
	}

	global.DB = conf.InitDB(config)
	global.InitDefaultQueryTimeOut(config.DBConfig.QueryTimeout)

	product.InitProductService(global.DB.Product)
}

func main() {

	r := gin.Default()
	r.GET("/ping", product.Ping)
	r.GET("/product/detail", product.GetDetailProduct)

	r.POST("/product/save", product.SaveNewProduct)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
