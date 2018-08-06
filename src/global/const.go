package global

import (
	"log"
)

//Error Logger
var Error *log.Logger

//All DB Access
var DB DBBundle

//All Repository Access
var Services RepoBundle

//Detail Server
var DetailServer ServerConfig

var UserCookie = map[string]string{
	"production":  "_TREK_",
	"staging":     "_TREK_TEST_",
	"development": "_TREK_DEV_",
}
