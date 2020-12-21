package isaac

const (
	availablePanelsBasePath = "/api/v1/controlpanels"
)

type AvailablePanelsService interface {
	Add(NewAvailablePanel) (AvailablePanel, error)
	Get(ID) (AvailablePanel, error)
	List() ([]AvailablePanel, error)
	Remove(AvailablePanel) error
	Update(AvailablePanel) (AvailablePanel, error)
}

type AvailablePanelsServiceOp struct {
	client *Client
}

type NewAvailablePanel struct {
	RegisteredPanelID ID `json:"registeredPanelId"`
	DisplayName string `json:"displayName"`
	Scale struct {
		Enabled bool `json:"enabled"`
		Width int `json:"width"`
		Height int `json:"height"`
	}
	OpenExternally bool `json:"openExternally"`
	Roles []ID `json:"roles"`
}

type AvailablePanel struct {
	NewAvailablePanel `json:",inline"`
	ID ID `json:"_id"`
	Location string `json:"location"`
	PermissionID ID `json:"permissionId"`
	Description string `json:"description"`
	SubsystemID ID `json:"subsystemId"`
	Type string `json:"type"`
	Connection map[string]interface{} `json:"connection"`
	RegisteredPanelName string `json:"displayNameRegisteredPanel"`
}

func (item AvailablePanel) ref() string {
	return item.ID.String()
}

func (s *AvailablePanelsServiceOp) Add(panel NewAvailablePanel) (AvailablePanel, error) {
	var respStruct AvailablePanel
	err := s.client.genericPost(availablePanelsBasePath, panel, &respStruct)
	return respStruct, err
}

func (s *AvailablePanelsServiceOp) Get(id ID) (AvailablePanel, error) {
	var respStruct AvailablePanel
	err := s.client.genericGetID(availablePanelsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *AvailablePanelsServiceOp) List() ([]AvailablePanel, error) {
	var respStruct []AvailablePanel
	err := s.client.genericGet(availablePanelsBasePath, &respStruct)
	return respStruct, err
}

func (s *AvailablePanelsServiceOp) Remove(panel AvailablePanel) error {
	err := s.client.genericDeleteID(availablePanelsBasePath, panel, nil)
	return err
}

func (s *AvailablePanelsServiceOp) Update(panel AvailablePanel) (AvailablePanel, error) {
	var respStruct AvailablePanel
	err := s.client.genericPutID(availablePanelsBasePath, panel, &respStruct)
	return respStruct, err
}
