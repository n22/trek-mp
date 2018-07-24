package app

import (
	"strconv"

	"github.com/5112100070/trek-mp/src/app/product"
	"github.com/5112100070/trek-mp/src/global"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetProductPage(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, Content-Type")

	productService := global.GetServiceProduct()

	start, errParse := strconv.Atoi(c.DefaultQuery("start", "0"))
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	rows, errParse := strconv.Atoi(c.DefaultQuery("rows", "10"))
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	sortType := c.DefaultQuery("sort", "ASC")

	result, err := productService.GetListProduct(start, rows, sortType)
	if err != nil {
		global.Error.Println(err)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	global.OKResponse(c, result)
}

func GetDetailProduct(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, Content-Type")

	productService := global.GetServiceProduct()

	productPath := c.Query("path")

	// productID, errParse := strconv.ParseInt(c.Query("product_name"), 10, 64)
	// if errParse != nil {
	// 	global.Error.Println(errParse)
	// 	global.BadRequestResponse(c, nil)
	// 	return
	// }

	p, err := productService.GetProductByPath(productPath)
	if err != nil {
		global.Error.Println(err)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	global.OKResponse(c, p)
}

func SaveNewProduct(c *gin.Context) {
	productService := global.GetServiceProduct()

	name := c.PostForm("name")
	priceBuy, errParse := strconv.ParseInt(c.PostForm("price_to_buy"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	priceSell, errParse := strconv.ParseInt(c.PostForm("price_to_sell"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	p := product.InitNewProduct(name, priceBuy, priceSell)
	errSave := productService.Save(p)
	if errSave != nil {
		global.Error.Println(errSave)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	global.CreatedResponse(c, p)
}
