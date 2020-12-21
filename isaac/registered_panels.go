package isaac

import "time"

const (
	registeredPanelsBasePath = "/api/v1/controlpanels"
)

type RegisteredPanelsService interface {
	Add(NewRegisteredPanel) (RegisteredPanel, error)
	Get(ID) (RegisteredPanel, error)
	List() ([]RegisteredPanel, error)
	Remove(RegisteredPanel) error
	Update(RegisteredPanel) (RegisteredPanel, error)
}

type RegisteredPanelsServiceOp struct {
	client *Client
}

type NewRegisteredPanel struct {
	DisplayName string `json:"displayName"`
	SubsystemID ID `json:"subystemId"`
	SubsystemExternalID string `json:"subsystemExternalId"`
	ExternalRef string `json:"externalRef"`
	Type string `json:"type"`
	Active bool `json:"active"`
	Description string `json:"description"`
	Connection map[string]interface{} `json:"connection"`
}

type RegisteredPanel struct {
	NewRegisteredPanel `json:",inline"`
	ID ID `json:"_id"`
	Location string `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
}

func (item RegisteredPanel) ref() string {
	return item.ID.String()
}

func (s *RegisteredPanelsServiceOp) Add(panel NewRegisteredPanel) (RegisteredPanel, error) {
	var respStruct RegisteredPanel
	err := s.client.genericPost(registeredPanelsBasePath, panel, &respStruct)
	return respStruct, err
}

func (s *RegisteredPanelsServiceOp) Get(id ID) (RegisteredPanel, error) {
	var respStruct RegisteredPanel
	err := s.client.genericGetID(registeredPanelsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *RegisteredPanelsServiceOp) List() ([]RegisteredPanel, error) {
	var respStruct []RegisteredPanel
	err := s.client.genericGet(registeredPanelsBasePath, &respStruct)
	return respStruct, err
}

func (s *RegisteredPanelsServiceOp) Remove(panel RegisteredPanel) error {
	err := s.client.genericDeleteID(registeredPanelsBasePath, panel, nil)
	return err
}

func (s *RegisteredPanelsServiceOp) Update(panel RegisteredPanel) (RegisteredPanel, error) {
	var respStruct RegisteredPanel
	err := s.client.genericPutID(registeredPanelsBasePath, panel, &respStruct)
	return respStruct, err
}
