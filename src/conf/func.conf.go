package conf

import (
	"fmt"
	"log"

	"github.com/5112100070/trek-mp/src/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ruizu/gcfg"
	"github.com/tokopedia/sqlt"
)

func initDBSqlt(connSlave string, connMaster string) (*sqlt.DB, error) {
	dbUrl := fmt.Sprintf("%s;%s", connMaster, connSlave)
	db, err := sqlt.Open("mysql", dbUrl)
	if err != nil {
		return db, err
	}

	if err := db.Ping(); err != nil {
		return db, err
	}
	return db, nil
}

func InitGlobalVariable(config Config) {
	global.DetailServer = config.Server
}

func ReadConfig(filePath string) (Config, bool) {
	var c Config
	if err := gcfg.ReadFileInto(&c, filePath); err != nil {
		log.Printf("%s\n", err)
		return Config{}, false
	}
	return c, true
}
