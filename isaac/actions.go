package isaac

import (
	"time"
)

const (
	actionsBasePath = "/api/v1/logs/scenarios"
)

type ActionsService interface {
	Add(NewAction) (Action, error)
	Get(ID) (Action, error)
	List() ([]Action, error)
	Remove(Action) error
	Update(Action) (Action, error)
}

type ActionsServiceOp struct {
	client *Client
}

type NewAction struct {
	Origin string `json:"origin"`
	Action string `json:"action"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Log bool `json:"log"`
}

type Action struct {
	NewActionLog `json:",inline"`
	ID ID `json:"_id"`
	UpdatedAt time.Time `json:"time"`
}

func (item Action) ref() string {
	return item.ID.String()
}

func (s *ActionsServiceOp) Add(action NewAction) (Action, error) {
	var respStruct Action
	err := s.client.genericPost(actionsBasePath, action, &respStruct)
	return respStruct, err
}

func (s *ActionsServiceOp) Get(id ID) (Action, error) {
	var respStruct Action
	err := s.client.genericGetID(actionsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *ActionsServiceOp) List() ([]Action, error) {
	var respStruct []Action
	err := s.client.genericGet(actionsBasePath, &respStruct)
	return respStruct, err
}

func (s *ActionsServiceOp) Remove(action Action) error {
	err := s.client.genericDeleteID(actionsBasePath, action, nil)
	return err
}

func (s *ActionsServiceOp) Update(action Action) (Action, error) {
	var respStruct Action
	err := s.client.genericPutID(actionsBasePath, action, &respStruct)
	return respStruct, err
}
