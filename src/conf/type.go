package conf

import (
	"github.com/tokopedia/sqlt"
)

type Config struct {
	DBConfig DBConfig
	Product  ProductConfig
}

type DBConfig struct {
	DBMaxConn     int
	DBMaxIdleConn int
	DBMaxLifetime int64
	QueryTimeout  int64
}

type ProductConfig struct {
	SlaveDBConnection  string
	MasterDBConnection string
}

// DBBundle : Data type of DB
type DBBundle struct {
	Product *sqlt.DB
}
