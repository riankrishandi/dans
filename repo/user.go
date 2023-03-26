package repo

import (
	"database/sql"
	"log"
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

func (r *Repo) GetUser(param GetUserRepoParam) (*GetUserRepoRes, error) {
	var user User
	err := r.client.GetContext(param.Ctx, &user, `SELECT
		id,
		username,
		password
	FROM tbl_user WHERE username=? LIMIT 1`, param.Username)
	switch err {
	case nil:
		return &GetUserRepoRes{user}, nil
	case sql.ErrNoRows:
		return nil, nil
	default:
		log.Printf("failed to get user from db: %s\n", err.Error())
		return nil, err
	}
}
