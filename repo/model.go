package repo

import "context"

type GetUserRepoParam struct {
	Ctx      context.Context
	Username string
}

type GetUserRepoRes struct {
	User User
}
