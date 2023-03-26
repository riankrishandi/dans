package job

import (
	"net/http"
	"time"
)

type Job struct {
	httpClient   *http.Client
	JobListURL   string
	JobDetailURL string
}

func New() *Job {
	return &Job{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		JobListURL:   "http://dev3.dansmultipro.co.id/api/recruitment/positions.json",
		JobDetailURL: "http://dev3.dansmultipro.co.id/api/recruitment/positions",
	}
}
