package global

import (
	"github.com/gin-gonic/gin"
)

func GetServiceProduct() ProductService {
	return Services.Product
}

func GetServiceUser() UserService {
	return Services.User
}

func SetCookie(c *gin.Context, key string, value string) {
	c.SetCookie(key, value, (5 * 3600), "/", DetailServer.Domain, true, true)
}
