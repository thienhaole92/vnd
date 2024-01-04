package repo

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/thienhaole92/vnd/logger"
)

func (r *Repo) GetUserById(ctx context.Context, userId string) (*UserInfoRecord, error) {
	log := logger.GetLogger("GetUserById")

	var user UserInfoRecord
	sql := `
		select
			*
		from
			user_info
		where
			id = $1
		`

	log.Debugw("log data", "sql", sql, "id", userId)
	err := pgxscan.Get(ctx, r.db, &user, sql, userId)
	return &user, err
}
