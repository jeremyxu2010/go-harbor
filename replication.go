package harbor

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type ReplicationsService struct {
	client *Client
}

type ListReplicationsOptions struct {
	ListOptions
	ProjectID    int64             `json:"project_id"`
	Name         string            `json:"name"`
}

type Replication struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Projects    []struct {
		ProjectID         int    `json:"project_id"`
		OwnerID           int    `json:"owner_id"`
		Name              string `json:"name"`
		CreationTime      string `json:"creation_time"`
		UpdateTime        string `json:"update_time"`
		Deleted           bool   `json:"deleted"`
		OwnerName         string `json:"owner_name"`
		Togglable         bool   `json:"togglable"`
		CurrentUserRoleID int    `json:"current_user_role_id"`
		RepoCount         int    `json:"repo_count"`
		ChartCount        int    `json:"chart_count"`
		Metadata          struct {
			Public             string `json:"public"`
			EnableContentTrust string `json:"enable_content_trust"`
			PreventVul         string `json:"prevent_vul"`
			Severity           string `json:"severity"`
			AutoScan           string `json:"auto_scan"`
		} `json:"metadata"`
	} `json:"projects"`
	Targets []struct {
		ID           int    `json:"id"`
		Endpoint     string `json:"endpoint"`
		Name         string `json:"name"`
		Username     string `json:"username"`
		Password     string `json:"password"`
		Type         int    `json:"type"`
		Insecure     bool   `json:"insecure"`
		CreationTime string `json:"creation_time"`
		UpdateTime   string `json:"update_time"`
	} `json:"targets"`
	Trigger struct {
		Kind          string `json:"kind"`
		ScheduleParam struct {
			Type    string `json:"type"`
			Weekday int    `json:"weekday"`
			Offtime int    `json:"offtime"`
		} `json:"schedule_param"`
	} `json:"trigger"`
	Filters []struct {
		Kind     string `json:"kind"`
		Value    string `json:"value"`
		Pattern  string `json:"pattern"`
		Metadata struct {
		} `json:"metadata"`
	} `json:"filters"`
	ReplicateExistingImageNow bool   `json:"replicate_existing_image_now"`
	ReplicateDeletion         bool   `json:"replicate_deletion"`
	CreationTime              string `json:"creation_time"`
	UpdateTime                string `json:"update_time"`
	ErrorJobCount             int    `json:"error_job_count"`
}


func (s *ReplicationsService) ListReplication(opt *ListReplicationsOptions) ([]Replication, *gorequest.Response, []error) {
	var replications []Replication
	resp, _, errs := s.client.
		NewRequest(gorequest.GET, "policies/replication").
		Query(*opt).
		EndStruct(&replications)
	return replications, &resp, errs
}

type DeleteReplicationOptions struct {
	Id      int               `json:"id"`
}

func (s *ReplicationsService) DeleteReplication(opt *DeleteReplicationOptions)(*gorequest.Response, []error) {
	resp, _, errs := s.client.
		NewRequest(gorequest.DELETE, fmt.Sprintf("policies/replication/%d", opt.Id)).
		End()
	return &resp, errs
}
