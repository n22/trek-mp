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
	ID                  int64     `json:"product_id"`
	Name                string    `json:"product_name"`
	Status              int       `json:"status"`
	Type                int       `json:"type"`
	PriceToRentStr      string    `json:"price_to_sell"`
	PriceToBuyStr       string    `json:"price_to_buy"`
	PriceRentDailyStr   string    `json:"price_rent_daily"`
	PriceRentWeeklyStr  string    `json:"price_rent_weekly"`
	PriceRentMonthlyStr string    `json:"price_rent_monthly"`
	CreateTime          time.Time `json:"create_time"`
	ImgUrl              string    `json:"img_url"`
	Path                string    `json:"path"`
	PriceBuy            int64     `json:"-,omitempty"`
	PriceSell           int64     `json:"-,omitempty"`
	PriceRentDaily      int64     `json:"-,omitempty"`
	PriceRentWeekly     int64     `json:"-,omitempty"`
	PriceRentMonthly    int64     `json:"-,omitempty"`
}

type productRepo struct {
	DB             *sqlt.DB
	queryDBTimeout time.Duration
}
