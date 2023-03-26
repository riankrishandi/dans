package job

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type JobDetail struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}

func (j *Job) GetJobDetail(param GetJobDetailParam) (*GetJobDetailRes, error) {
	// Prepare URL.
	url := fmt.Sprintf("%s/%s", j.JobDetailURL, param.ID)

	// Send GET request.
	res, err := http.Get(url)
	if err != nil {
		log.Printf("failed to send job detail request: %s\n", err.Error())
		return nil, err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("failed to read job detail response body: %s", err.Error())
		return nil, err
	}

	var detail JobDetail
	err = json.Unmarshal(bytes, &detail)
	if err != nil {
		log.Printf("failed to unmarshal json: %s\n", err.Error())
		return nil, err
	}

	return &GetJobDetailRes{detail}, nil
}
