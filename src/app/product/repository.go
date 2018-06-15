package product

import (
	"context"
	"time"

	"github.com/tokopedia/sqlt"
)

func InitProductRepo(productDB *sqlt.DB, queryTimeout time.Duration) *productRepo {
	return &productRepo{
		DB:             productDB,
		queryDBTimeout: queryTimeout,
	}
}

func (repo productRepo) Save(p Product) error {
	query := `
		INSERT INTO ws_product (
			product_name,
			status,
			type,
			price_to_buy,
			price_to_sell,
			create_time,
			img_url,
			domain_name)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`

	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	dbProduct := repo.DB
	insertTalk, errPrepared := dbProduct.PreparexContext(ctx, query)
	if errPrepared != nil {
		return errPrepared
	}
	defer insertTalk.Close()

	_, errInsert := insertTalk.ExecContext(ctx, p.Name, p.Status, p.Type, p.PriceBuy, p.PriceSell, p.CreateTime, p.ImgUrl, p.Domain)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (repo productRepo) GetProduct() {

}
