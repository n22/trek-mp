package conf

import (
	"database/sql"
	"log"
	"time"
)

func InitDB(c Config) DBBundle {
	productDB, err := initDBSqlt(c.Product.MasterDBConnection, c.Product.SlaveDBConnection)
	if err != nil {
		log.Fatalf("db.Product not available with config %v", err)
	}

	DB := DBBundle{
		productDB,
	}

	DB.Product.SetMaxOpenConnections(c.DBConfig.DBMaxConn)
	DB.Product.SetMaxIdleConnections(c.DBConfig.DBMaxIdleConn)
	DB.Product.SetConnMaxLifetime(time.Second * time.Duration(c.DBConfig.DBMaxLifetime))

	var rows *sql.Rows
	rows, err = DB.Product.Query("SELECT 1")
	if err != nil {
		log.Fatal("Product DB is not accessible, please check config")
	}
	defer rows.Close()

	return DB
}