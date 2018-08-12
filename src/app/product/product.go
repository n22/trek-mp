package product

import "github.com/5112100070/trek-mp/src/utils"

func InitNewProduct(name string) Product {
	return Product{
		Name:       name,
		Status:     1,
		Type:       1,
		CreateTime: utils.GetTimeWIB(),
	}
}

func (p *Product) SetPriceToRent() {
	p.PriceToRentStr = utils.IntToRp(p.PriceSell)
	p.PriceRentDailyStr = utils.IntToRp(p.PriceRentDaily)
	p.PriceRentWeeklyStr = utils.IntToRp(p.PriceRentWeekly)
	p.PriceRentMonthlyStr = utils.IntToRp(p.PriceRentMonthly)
}

func (p *Product) SetPriceToBuy() {
	p.PriceToBuyStr = utils.IntToRp(p.PriceBuy)
}
