package product

import (
	"time"

	"github.com/tokopedia/sqlt"
)

type sort string

const (
	SORT_ASC  sort = "ASC"
	SORT_DESC sort = "DESC"
)

type Product struct {
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

type productRepo struct {
	DB             *sqlt.DB
	queryDBTimeout time.Duration
}
