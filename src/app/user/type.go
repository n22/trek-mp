package user

import (
	"time"

	redigo "github.com/5112100070/trek-mp/src/global/redis"
	"github.com/tokopedia/sqlt"
)

const (
	redis_key_cookie = "cookie:c_"
	redis_key_user   = "class:user:u_"
	redis_timeout    = time.Duration((6 * time.Hour))

	USER_STATUS_ACTIVE  = 1
	USER_STATUS_DELETED = 0
)

type User struct {
	ID         int64     `json:"user_id"`
	Username   string    `json:"username"`
	Password   string    `json:"-,omitempty"`
	FullName   string    `json:"fullname"`
	Status     int       `json:"status"`
	Type       int       `json:"type"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	ImgUrl     string    `json:"img_url"`
}

type userRepo struct {
	DB             *sqlt.DB
	redis          redigo.Redis
	queryDBTimeout time.Duration
}
