package main

import (
	"github.com/5112100070/trek-mp/src/product"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", product.Ping)
	r.GET("/product/detail", product.GetDetailProduct)
	r.Run(":3000") // listen and serve on 0.0.0.0:8080
}
