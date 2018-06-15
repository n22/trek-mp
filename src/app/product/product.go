package product

import "github.com/5112100070/trek-mp/src/utils"

func InitNewProduct(name string, priceToBuy int64, priceToSell int64) Product {
	return Product{
		Name:       name,
		PriceBuy:   priceToBuy,
		PriceSell:  priceToSell,
		Status:     1,
		Type:       1,
		CreateTime: utils.GetTimeWIB(),
	}
}
