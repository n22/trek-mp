package app

import (
	"fmt"
	"strconv"

	"github.com/5112100070/trek-mp/src/app/product"
	"github.com/5112100070/trek-mp/src/global"
	"github.com/5112100070/trek-mp/src/utils"
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

func GetDetailProductById(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, Content-Type")

	productService := global.GetServiceProduct()

	productID, errParse := strconv.ParseInt(c.Query("product-id"), 10, 64)
	if errParse != nil {
		global.Error.Println(errParse)
		global.BadRequestResponse(c, nil)
		return
	}

	p, err := productService.GetProduct(productID)
	if err != nil {
		global.Error.Println(err)
		global.InternalServerErrorResponse(c, nil)
		return
	}

	global.OKResponse(c, p)
}

func SaveNewProduct(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	productService := global.GetServiceProduct()

	var p product.Product

	productID, errParse := strconv.ParseInt(c.PostForm("product_id"), 10, 64)
	if errParse != nil {
		productID = 0
		return
	}
	productName := c.PostForm("product_name")

	if productID != 0 {
		var errGetProduct error
		p, errGetProduct = productService.GetProduct(productID)
		if errGetProduct != nil {
			global.Error.Println(errGetProduct)
			global.BadRequestResponse(c, nil)
			return
		}
	} else {
		p = product.InitNewProduct(productName)
	}

	if c.PostForm("price_rent_daily") != "" {
		priceRentDaily, errParse := strconv.ParseInt(c.PostForm("price_rent_daily"), 10, 64)
		if errParse != nil {
			global.Error.Println(errParse)
			global.BadRequestResponse(c, nil)
			return
		}
		if priceRentDaily != 0 {
			p.PriceRentDaily = priceRentDaily
		}
		fmt.Println(priceRentDaily)
	}

	if c.PostForm("price_rent_weekly") != "" {
		priceRentWeekly, errParse := strconv.ParseInt(c.PostForm("price_rent_weekly"), 10, 64)
		if errParse != nil {
			global.Error.Println(errParse)
			global.BadRequestResponse(c, nil)
			return
		}
		if priceRentWeekly != 0 {
			p.PriceRentWeekly = priceRentWeekly
		}
		p.PriceSell = p.PriceRentWeekly
	}

	if c.PostForm("price_rent_monthly") != "" {
		priceRentMonthly, errParse := strconv.ParseInt(c.PostForm("price_rent_monthly"), 10, 64)
		if errParse != nil {
			global.Error.Println(errParse)
			global.BadRequestResponse(c, nil)
			return
		}
		if priceRentMonthly != 0 {
			p.PriceRentMonthly = priceRentMonthly
		}
	}
	p.Path = c.PostForm("path")
	fmt.Println(c.PostForm("path"))

	if productID == 0 {
		errSave := productService.Save(p)
		if errSave != nil {
			global.Error.Println(errSave)
			global.InternalServerErrorResponse(c, nil)
			return
		}
	} else {
		errUpdate := productService.Update(p)
		if errUpdate != nil {
			global.Error.Println(errUpdate)
			global.InternalServerErrorResponse(c, nil)
			return
		}

		fmt.Println(p)
	}

	global.CreatedResponse(c, p)
}

func MakeLogin(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	username := c.PostForm("username")
	password := c.PostForm("secret")

	userService := global.GetServiceUser()

	dataResponse := map[string]interface{}{
		"is_success": false,
	}

	result, nekot, err := userService.MakeLogin(username, password)
	if err != nil {
		global.Error.Println(err)
		global.InternalServerErrorResponse(c, dataResponse)
		return
	}

	if result {
		global.OKResponse(c, map[string]interface{}{
			"env":        global.UserCookie[utils.GetEnv()],
			"nekot":      nekot,
			"is_success": result,
		})
	} else {
		global.UnAuthorizeResponse(c, dataResponse)
	}
}
