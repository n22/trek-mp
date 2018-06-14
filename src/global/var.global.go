package global

import (
	"log"
	"time"

	"github.com/5112100070/trek-mp/src/conf"
)

//Error Logger
var Error *log.Logger

//All initial Configuration
// var Config conf.Config

//All DB Access
var DB conf.DBBundle

//Default query timeout
var QueryTimeout time.Duration
