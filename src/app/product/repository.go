package product

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/5112100070/trek-mp/src/utils"
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
			price_to_rent_daily,
			price_to_rent_weekly,
			price_to_rent_monthly,
			create_time,
			img_url,
			path)
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`

	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	dbProduct := repo.DB
	insertProduct, errPrepared := dbProduct.PreparexContext(ctx, query)
	if errPrepared != nil {
		return errPrepared
	}
	defer insertProduct.Close()

	_, errInsert := insertProduct.ExecContext(ctx, p.Name, p.Status, p.Type, p.PriceBuy, p.PriceSell, p.PriceRentDaily, p.PriceRentWeekly, p.PriceRentMonthly, p.CreateTime, p.ImgUrl, p.Path)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (repo productRepo) Update(p Product) error {
	query := `
		UPDATE 
				ws_product 
		SET
				product_name = ?,
				status = ?,
				type = ?,
				price_to_sell = ?,
				price_to_rent_daily = ?,
				price_to_rent_weekly = ?,
				price_to_rent_monthly = ?,
				img_url = ?,
				path = ?
		WHERE
				product_id = ?
		`

	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	dbProduct := repo.DB
	insertProduct, errPrepared := dbProduct.PreparexContext(ctx, query)
	if errPrepared != nil {
		return errPrepared
	}
	defer insertProduct.Close()

	_, errInsert := insertProduct.ExecContext(ctx, p.Name, p.Status, p.Type, p.PriceSell, p.PriceRentDaily, p.PriceRentWeekly, p.PriceRentMonthly, p.ImgUrl, p.Path, p.ID)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (repo productRepo) GetProduct(productID int64) (Product, error) {
	var p Product
	query := `
		SELECT 
			product_id,
			product_name,
			price_to_buy,
			price_to_sell,
			price_to_rent_daily,
			price_to_rent_weekly,
			price_to_rent_monthly,			
			status,
			type,
			create_time,
			img_url,
			path
		FROM
			ws_product
		WHERE
			product_id=?
		LIMIT 1
	`

	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	selectQuery, errPrepare := repo.DB.PreparexContext(ctx, query)
	if errPrepare != nil {
		return p, errPrepare
	}

	var rawTime time.Time
	errScan := selectQuery.QueryRowxContext(ctx, productID).Scan(&p.ID,
		&p.Name,
		&p.PriceBuy,
		&p.PriceSell,
		&p.PriceRentDaily,
		&p.PriceRentWeekly,
		&p.PriceRentMonthly,
		&p.Status,
		&p.Type,
		&rawTime,
		&p.ImgUrl,
		&p.Path)
	if errScan != nil {
		return p, errScan
	}
	p.CreateTime = utils.ConvertTimeWIB(rawTime)
	p.SetPriceToRent()
	p.SetPriceToBuy()

	return p, nil
}

func (repo productRepo) GetProductByName(productName string) (Product, error) {
	var p Product
	query := `
		SELECT 
			product_id,
			product_name,
			price_to_buy,
			price_to_sell,
			price_to_rent_daily,
			price_to_rent_weekly,
			price_to_rent_monthly,			
			status,
			type,
			create_time,
			img_url,
			path
		FROM
			ws_product
		WHERE
			product_name=?
		LIMIT 1
	`

	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	selectQuery, errPrepare := repo.DB.PreparexContext(ctx, query)
	if errPrepare != nil {
		return p, errPrepare
	}

	var rawTime time.Time
	errScan := selectQuery.QueryRowxContext(ctx, productName).Scan(&p.ID,
		&p.Name,
		&p.PriceBuy,
		&p.PriceSell,
		&p.PriceRentDaily,
		&p.PriceRentWeekly,
		&p.PriceRentMonthly,
		&p.Status,
		&p.Type,
		&rawTime,
		&p.ImgUrl,
		&p.Path)
	if errScan != nil && errScan != sql.ErrNoRows {
		return p, errScan
	}
	p.CreateTime = utils.ConvertTimeWIB(rawTime)
	p.SetPriceToRent()
	p.SetPriceToBuy()

	return p, nil
}

func (repo productRepo) GetProductByPath(path string) (Product, error) {
	var p Product
	query := `
		SELECT 
			product_id,
			product_name,
			price_to_buy,
			price_to_sell,
			price_to_rent_daily,
			price_to_rent_weekly,
			price_to_rent_monthly,			
			status,
			type,
			create_time,
			img_url,
			path
		FROM
			ws_product
		WHERE
			path=?
		LIMIT 1
	`

	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	selectQuery, errPrepare := repo.DB.PreparexContext(ctx, query)
	if errPrepare != nil {
		return p, errPrepare
	}

	var rawTime time.Time
	errScan := selectQuery.QueryRowxContext(ctx, path).Scan(&p.ID,
		&p.Name,
		&p.PriceBuy,
		&p.PriceSell,
		&p.PriceRentDaily,
		&p.PriceRentWeekly,
		&p.PriceRentMonthly,
		&p.Status,
		&p.Type,
		&rawTime,
		&p.ImgUrl,
		&p.Path)
	if errScan != nil && errScan != sql.ErrNoRows {
		return p, errScan
	}
	p.CreateTime = utils.ConvertTimeWIB(rawTime)
	p.SetPriceToRent()
	p.SetPriceToBuy()

	return p, nil
}

func (repo productRepo) GetListProduct(start int, rows int, sortType string) ([]Product, error) {
	if start < 0 || rows <= 0 {
		start = 0
		rows = 10
	}
	if sort(sortType) != SORT_ASC && sort(sortType) != SORT_DESC {
		sortType = string(SORT_ASC)
	}

	query := fmt.Sprintf(`
		SELECT 
			product_id,
			product_name,
			price_to_buy,
			price_to_sell,
			status,
			type,
			create_time,
			img_url,
			path
		FROM ws_product
		WHERE
			status = 1 AND
			type = 1
		ORDER BY product_id %s	
		LIMIT %v,%v		
	`, sortType, start, rows)
	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	dbProduct := repo.DB
	selectProduct, errPrepared := dbProduct.PreparexContext(ctx, query)
	if errPrepared != nil {
		return nil, errPrepared
	}
	defer selectProduct.Close()

	resultQuery, errQuery := selectProduct.QueryContext(ctx)
	if errQuery != nil {
		return nil, errQuery
	}

	var result []Product
	for resultQuery.Next() {
		var p Product
		var rawTime time.Time
		errScan := resultQuery.Scan(&p.ID,
			&p.Name,
			&p.PriceBuy,
			&p.PriceSell,
			&p.Status,
			&p.Type,
			&rawTime,
			&p.ImgUrl,
			&p.Path)
		if errScan != nil {
			return nil, errScan
		}

		p.CreateTime = utils.ConvertTimeWIB(rawTime)
		p.SetPriceToRent()
		p.SetPriceToBuy()
		result = append(result, p)
	}

	return result, nil
}
