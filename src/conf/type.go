package conf

import (
	"github.com/5112100070/trek-mp/src/global"
	redigo "github.com/5112100070/trek-mp/src/global/redis"
)

type Config struct {
	DBConfig      DBConfig
	Server        global.ServerConfig
	Product       ProductConfig
	User          UserConfig
	RedigoDefault redigo.RedisConfig
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

type UserConfig struct {
	Redis              string
	SlaveDBConnection  string
	MasterDBConnection string
}
