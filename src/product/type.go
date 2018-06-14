package product

import "time"

type product struct {
	ID         int64     `json:"product_id"`
	Name       string    `json:"product_name"`
	Status     int       `json:"status"`
	Type       int       `json:"type"`
	PriceBuy   int64     `json:"price_to_buy"`
	PriceSell  int64     `json:"price_to_sell"`
	CreateTime time.Time `json:"create_time"`
	ImgUrl     string    `json:"img_url"`
	Domain     string    `json:"domain_name"`
}
