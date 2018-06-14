package product

import (
	"strconv"

	"github.com/5112100070/trek-mp/src/global"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetDetailProduct(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SaveNewProduct(c *gin.Context) {
	name := c.PostForm("name")
	priceBuy, errParse := strconv.ParseInt(c.PostForm("price_to_buy"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		return
	}

	priceSell, errParse := strconv.ParseInt(c.PostForm("price_to_sell"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		return
	}

	p := initProduct(name, priceBuy, priceSell)
	errSave := p.SaveNewData()
	if errSave != nil {
		global.Error.Println(errSave)
		return
	}

}
