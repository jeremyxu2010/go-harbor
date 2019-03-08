package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
	)

type ChartsService struct {
	client *Client
}

type ListChartsOptions struct {
	ListOptions
	Repo   string `url:"repo,omitempty" json:"repo,omitempty"`
}

type Chart struct {
	Name          string `json:"name"`
	TotalVersions int    `json:"total_versions"`
	LatestVersion string `json:"latest_version"`
	Created       string `json:"created"`
	Updated       string `json:"updated"`
	Icon          string `json:"icon"`
	Home          string `json:"home"`
	Deprecated    bool   `json:"deprecated"`
}

func (s *ChartsService) ListChart(opt *ListChartsOptions) ([]Chart, *gorequest.Response, []error) {
	var charts []Chart
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, fmt.Sprintf("chartrepo/%s/charts", opt.Repo)).
		Query(*opt).
		EndStruct(&charts)
	return charts, &resp, errs
}

type DeleteChartOptions struct {
	Repo   string `url:"repo,omitempty" json:"repo,omitempty"`
	Name          string `url:"name,omitempty" json:"name,omitempty"`
}

func (s *ChartsService) DeleteChart(opt *DeleteChartOptions) (*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("chartrepo/%s/charts/%s", opt.Repo, opt.Name)).
		End()
	return &resp, errs
}
