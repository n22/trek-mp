package product

import (
	"context"
	"time"

	"github.com/5112100070/trek-mp/src/global"
)

func initProduct(name string, priceToBuy int64, priceToSell int64) product {
	return product{
		Name:      name,
		PriceBuy:  priceToBuy,
		PriceSell: priceToSell,
	}
}

func (p *product) SaveNewData() error {
	p.CreateTime = time.Now()
	p.Status = 1
	p.Type = 1
	return p.Save()
}

func (p product) Save() error {
	query := `
		INSERT INTO ws_product (product_name, status, type, price_to_buy, price_to_sell, create_time, img_url, domain_name)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`

	ctx, cancel := context.WithTimeout(context.TODO(), global.QueryTimeout)
	defer cancel()

	dbProduct := global.DB.Product
	insertTalk, errPrepared := dbProduct.PreparexContext(ctx, query)
	if errPrepared != nil {
		global.Error.Printf(errPrepared.Error())
		return errPrepared
	}
	defer insertTalk.Close()

	_, errInsert := insertTalk.ExecContext(ctx, p.Name, p.Status, p.Type, p.PriceBuy, p.PriceSell, p.CreateTime, p.ImgUrl, p.Domain)
	if errInsert != nil {
		global.Error.Println(errInsert)
		return errInsert
	}

	return nil
}
