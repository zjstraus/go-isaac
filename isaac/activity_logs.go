package isaac

import (
	"time"
)

const (
	activityLogsBasePath = "/api/v1/logs"
)

type ActivityLogsService interface {
	Add(NewActivityLog) (ActivityLog, error)
	Get(ID) (ActivityLog, error)
	List() ([]ActivityLog, error)
}

type ActivityLogsServiceOp struct {
	client *Client
}

type NewActivityLog struct {
	Host string `json:"host"`
	Severity string `json:"severity"`
	Key string `json:"key"`
	Value string `json:"value"`
	Tags []string `json:"tags,omitempty"`
	CreatedBy ID `json:"createdBy"`
	CreatedByType string `json:"createdByType"`
}

type ActivityLog struct {
	NewActivityLog `json:",inline"`
	ID ID `json:"_id"`
	Time time.Time `json:"time"`
	SubsystemID string `json:"subsystemId"`
	SubsystemExternalID string `json:"subsystemExternalId"`
	DisplayName string `json:"displayName"`
}

func (item ActivityLog) ref() string {
	return item.ID.String()
}

type getApiV1LogsResponseWrapper struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
	TotalPages int `json:"totalPages"`
	TotalLogs int `json:"totalLogs"`
	Logs []ActivityLog `json:"logs"`
}

func (s *ActivityLogsServiceOp) Add(entry NewActivityLog) (ActivityLog, error) {
	var respStruct ActivityLog
	err := s.client.genericPost(activityLogsBasePath, entry, &respStruct)
	return respStruct, err
}

func (s *ActivityLogsServiceOp) Get(id ID) (ActivityLog, error) {
	var respStruct ActivityLog
	err := s.client.genericGetID(activityLogsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *ActivityLogsServiceOp) List() ([]ActivityLog, error) {
	var respStruct getApiV1LogsResponseWrapper
	err := s.client.genericGet(activityLogsBasePath, &respStruct)
	return respStruct.Logs, err
}
