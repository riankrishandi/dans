package controller

import (
	"github.com/riankrishandi/dans/job"
	"github.com/riankrishandi/dans/render"
	"github.com/riankrishandi/dans/repo"
)

type Controller struct {
	repo     *repo.Repo
	job      *job.Job
	renderer *render.Renderer
}

func New(repo *repo.Repo, j *job.Job, r *render.Renderer) *Controller {
	return &Controller{
		repo:     repo,
		job:      j,
		renderer: r,
	}
}
