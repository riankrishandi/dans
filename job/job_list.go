package job

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func (j *Job) GetJobList(param GetJobListParam) (*GetJobListRes, error) {
	// Prepare params.
	params := url.Values{}
	if param.Description != "" {
		params.Add("description", param.Description)
	}
	if param.Location != "" {
		params.Add("location", param.Location)
	}
	if param.FullTime {
		params.Add("full_time", "true")
	}
	if param.Page != 0 {
		params.Add("page", strconv.Itoa(param.Page))
	}

	// Prepare URL.
	u, _ := url.ParseRequestURI(j.JobListURL)
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)

	// Send GET request.
	res, err := http.Get(urlStr)
	if err != nil {
		log.Printf("failed to send job list request: %s\n", err.Error())
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read job list response body: %s", err.Error())
		return nil, err
	}

	var jobList []JobDetail
	err = json.Unmarshal(bytes, &jobList)
	if err != nil {
		log.Printf("failed to unmarshal json: %s\n", err.Error())
		return nil, err
	}

	return &GetJobListRes{jobList}, nil
}
