package controller

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/riankrishandi/dans/job"
)

func (c *Controller) HandleGetJobList() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get params.
		params := r.URL.Query()
		description := params.Get("description")
		location := params.Get("location")
		fullTime := params.Get("full_time")
		page := params.Get("page")
		pageInt, _ := strconv.Atoi(page)

		jobRes, err := c.job.GetJobList(job.GetJobListParam{
			Ctx:         r.Context(),
			Description: description,
			Location:    location,
			FullTime:    fullTime != "",
			Page:        pageInt,
		})
		if err != nil {
			c.renderer.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		c.renderer.RenderJSON(w, http.StatusOK, jobRes.List)
	})
}

func (c *Controller) HandleGetJobDetail() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		jobID := vars["jobID"]

		jobRes, err := c.job.GetJobDetail(job.GetJobDetailParam{
			Ctx: r.Context(),
			ID:  jobID,
		})
		if err != nil {
			c.renderer.RenderJSON(w, http.StatusBadRequest, err)
			return
		}

		c.renderer.RenderJSON(w, http.StatusOK, jobRes.Detail)
	})
}
