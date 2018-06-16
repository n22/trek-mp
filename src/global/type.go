package global

import (
	"github.com/5112100070/trek-mp/src/app/product"
	"github.com/tokopedia/sqlt"
)

// DBBundle : Data type of DB
type DBBundle struct {
	Product *sqlt.DB
}

type RepoBundle struct {
	Product ProductService
}

type ProductService interface {
	Save(p product.Product) error
	GetProduct(productID int64) (product.Product, error)
	GetListProduct(start int, rows int, sortType string) ([]product.Product, error)
}
