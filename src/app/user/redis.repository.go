package user

import (
	"encoding/json"
	"fmt"
)

func (repo userRepo) getUserByIdInRedis(userID int64) (User, error) {
	return repo.getUserInRedis(fmt.Sprintf("%v%v", redis_key_user, userID))
}

func (repo userRepo) getUserByCookieInRedis(cookie string) (User, error) {
	return repo.getUserInRedis(fmt.Sprintf("%v%v", redis_key_cookie, cookie))
}

func (repo userRepo) getUserInRedis(key string) (User, error) {
	var user User
	result, err := repo.redis.GET(key)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal([]byte(result), &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo userRepo) setUserByIdInRedis(user User) error {
	key := fmt.Sprintf("%v%v", redis_key_user, user.ID)

	err := repo.redis.SETEX(key, redis_timeout, user)
	if err != nil {
		return err
	}

	return nil

}

func (repo userRepo) setUserByCookieInRedis(nekot string, user User) error {
	key := fmt.Sprintf("%v%v", redis_key_cookie, nekot)

	data := struct {
		ID       int64  `json:"user_id"`
		Fullname string `json:"fullname"`
		Type     int    `json:"type"`
	}{
		ID:       user.ID,
		Fullname: user.FullName,
		Type:     user.Type,
	}
	m, _ := json.Marshal(data)

	err := repo.redis.SETEX(key, redis_timeout, string(m))
	if err != nil {
		return err
	}

	return nil
}
