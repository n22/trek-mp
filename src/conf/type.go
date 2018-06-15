package conf

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
