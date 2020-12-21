package isaac

import (
	"time"
)

const (
	actionLogsBasePath = "/api/v1/logs/scenariosLog"
)

type ActionLogsService interface {
	Add(NewActionLog) (ActionLog, error)
	Get(ID) (ActionLog, error)
	List() ([]ActionLog, error)
}

type ActionLogsServiceOp struct {
	client *Client
}

type NewActionLog struct {
	Origin string `json:"origin"`
	Action string `json:"action"`
	Text string `json:"text"`
	DisplayName string `json:"displayName"`
}

type ActionLog struct {
	NewActionLog `json:",inline"`
	ID ID `json:"_id"`
	Time time.Time `json:"time"`
}

func (item ActionLog) ref() string {
	return item.ID.String()
}

type getApiV1LogsActionLogsResponseWrapper struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
	TotalPages int `json:"totalPages"`
	TotalLogs int `json:"totalLogs"`
	Logs []ActionLog `json:"logs"`
}

func (s *ActionLogsServiceOp) Add(entry NewActionLog) (ActionLog, error) {
	var respStruct ActionLog
	err := s.client.genericPost(actionLogsBasePath, entry, &respStruct)
	return respStruct, err
}

func (s *ActionLogsServiceOp) Get(id ID) (ActionLog, error) {
	var respStruct ActionLog
	err := s.client.genericGetID(actionLogsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *ActionLogsServiceOp) List() ([]ActionLog, error) {
	var respStruct getApiV1LogsActionLogsResponseWrapper
	err := s.client.genericGet(actionLogsBasePath, &respStruct)
	return respStruct.Logs, err
}
