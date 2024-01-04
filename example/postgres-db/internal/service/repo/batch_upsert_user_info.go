package repo

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type UserInfoRecord struct {
	ID string `db:"id" json:"id"`
}

type UserList []*UserInfoRecord

func (r *Repo) BatchUpsertUserInfo(ctx context.Context, ul UserList) error {
	batch := &pgx.Batch{}
	for _, u := range ul {
		sql := `
    insert into
        user_info (
            id,
			created_at
        )
    values
        (
            $1,
            $2
        ) on conflict (id) do
    update
    set
		updated_at = now() returning *;
    `
		batch.Queue(sql,
			u.ID,
			time.Now(),
		)
	}

	br := r.db.SendBatch(ctx, batch)
	defer br.Close()

	for i := 0; i < len(ul); i++ {
		_, err := br.Exec()
		if err != nil {
			return err
		}
	}

	return nil
}
