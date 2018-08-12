package conf

import (
	"database/sql"
	"log"
	"time"

	"github.com/5112100070/trek-mp/src/global"
	redigo "github.com/5112100070/trek-mp/src/global/redis"
)

func InitDB(c Config) global.DBBundle {
	productDB, err := initDBSqlt(c.Product.MasterDBConnection, c.Product.SlaveDBConnection)
	if err != nil {
		log.Fatalf("db.Product not available with config %v", err)
	}

	userDB, err := initDBSqlt(c.User.MasterDBConnection, c.User.SlaveDBConnection)
	if err != nil {
		log.Fatalf("db.User not available with config %v", err)
	}

	redisUser := redigo.New(c.User.Redis, &c.RedigoDefault)

	DB := global.DBBundle{
		Product:   productDB,
		User:      userDB,
		RedisUser: redisUser,
	}

	DB.Product.SetMaxOpenConnections(c.DBConfig.DBMaxConn)
	DB.Product.SetMaxIdleConnections(c.DBConfig.DBMaxIdleConn)
	DB.Product.SetConnMaxLifetime(time.Second * time.Duration(c.DBConfig.DBMaxLifetime))

	DB.User.SetMaxOpenConnections(c.DBConfig.DBMaxConn)
	DB.User.SetMaxIdleConnections(c.DBConfig.DBMaxIdleConn)
	DB.User.SetConnMaxLifetime(time.Second * time.Duration(c.DBConfig.DBMaxLifetime))

	var rows *sql.Rows
	rows, err = DB.Product.Query("SELECT 1")
	if err != nil {
		log.Fatal("Product DB is not accessible, please check config")
	}
	defer rows.Close()

	rows, err = DB.User.Query("SELECT 1")
	if err != nil {
		log.Fatal("User DB is not accessible, please check config")
	}
	defer rows.Close()

	_, err = redisUser.PING()
	if err != nil {
		log.Fatal("Redis User data not accessible, please check config")
	}

	return DB
}
