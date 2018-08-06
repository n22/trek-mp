package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/5112100070/trek-mp/src/utils"
	"github.com/tokopedia/sqlt"
)

func InitUserRepo(userDB *sqlt.DB, queryTimeout time.Duration) *userRepo {
	return &userRepo{
		DB:             userDB,
		queryDBTimeout: queryTimeout,
	}
}

func (repo userRepo) GetUser(userID int64) (User, error) {
	var u User
	query := `
		SELECT
			user_id,
			username,
			fullname,
			status,
			type,
			create_time,
			update_time,
			img_url
		FROM
			ws_user
		WHERE
			user_id=?
		LIMIT 1
	`

	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	selectQuery, errPrepare := repo.DB.PreparexContext(ctx, query)
	if errPrepare != nil {
		return u, errPrepare
	}

	var rawCreateTime time.Time
	var rawUpdateTime time.Time
	errScan := selectQuery.QueryRowxContext(ctx, userID).Scan(&u.ID,
		&u.Username,
		&u.FullName,
		&u.Status,
		&u.Type,
		&rawCreateTime,
		&rawUpdateTime,
		&u.ImgUrl,
	)
	if errScan != nil {
		return u, errScan
	}
	u.CreateTime = utils.ConvertTimeWIB(rawCreateTime)
	u.UpdateTime = utils.ConvertTimeWIB(rawUpdateTime)

	return u, nil
}

func (repo userRepo) MakeLogin(username string, password string) (bool, string, error) {
	var userID, status int64
	var nekot string
	query := `
		SELECT
			user_id,
			status
		FROM
			ws_user
		WHERE
			username=? AND
			password=?
		LIMIT 1
	`
	ctx, cancel := context.WithTimeout(context.TODO(), repo.queryDBTimeout)
	defer cancel()

	selectQuery, errPrepare := repo.DB.PreparexContext(ctx, query)
	if errPrepare != nil {
		return false, nekot, errPrepare
	}

	errScan := selectQuery.QueryRowxContext(ctx, username, password).Scan(&userID, &status)
	if errScan != nil && errScan != sql.ErrNoRows {
		return false, nekot, errScan
	}

	if userID != 0 && status == USER_STATUS_ACTIVE {
		return true, utils.GenerateMD5(fmt.Sprintf("%v%v", userID, time.Now())), nil
	} else {
		return false, nekot, nil
	}
}
