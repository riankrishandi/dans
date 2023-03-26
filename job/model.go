package job

import "context"

type GetJobListParam struct {
	Ctx         context.Context
	Description string
	Location    string
	FullTime    bool
	Page        int
}

type GetJobListRes struct {
	List []JobDetail
}

type GetJobDetailParam struct {
	Ctx context.Context
	ID  string
}

type GetJobDetailRes struct {
	Detail JobDetail
}
