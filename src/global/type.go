package global

import (
	"github.com/5112100070/trek-mp/src/app/product"
	"github.com/5112100070/trek-mp/src/app/user"
	redigo "github.com/5112100070/trek-mp/src/global/redis"
	"github.com/tokopedia/sqlt"
)

type ServerConfig struct {
	Domain  string
	AppPort string
}

// DBBundle : Data type of DB
type DBBundle struct {
	Product   *sqlt.DB
	User      *sqlt.DB
	RedisUser redigo.Redis
}

type RepoBundle struct {
	Product ProductService
	User    UserService
}

type ProductService interface {
	Save(p product.Product) error
	GetProduct(productID int64) (product.Product, error)
	GetProductByName(productName string) (product.Product, error)
	GetProductByPath(path string) (product.Product, error)
	GetListProduct(start int, rows int, sortType string) ([]product.Product, error)
}

type UserService interface {
	GetUser(userID int64) (user.User, error)
	MakeLogin(username string, password string) (bool, string, error)
}
