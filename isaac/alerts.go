package isaac

const (
	alertsBasePath = "/api/v1/alerts"
)

type AlertsService interface {
	Add(NewAlert) (Alert, error)
	Get(ID) (Alert, error)
	List() ([]Alert, error)
	Remove(Alert) error
	Update(Alert) (Alert, error)
}

type AlertsServiceOp struct {
	client *Client
}

type NewAlert struct {
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	SubsystemExternalID string `json:"subsystemExternalId"`
	Variable string `json:"variable"`
	Value string `json:"value"`
	Operator string `json:"operator"`
	Message string `json:"message"`
	Criticality string `json:"criticality"`
	Lifecycle string `json:"lifeCycle"`
	Lock bool `json:"lock"`
	Recovery bool `json:"recovery"`
	Log bool `json:"Log"`
	Active bool `json:"Active"`
	Destinations []AlertDestination `json:"destinations"`
}

type AlertDestination struct {
	User string `json:"user"`
	Role string `json:"role"`
	Method []string `json:"method"`
}

type Alert struct {
	NewAlert `json:",inline"`
	ID ID `json:"_id"`
	VariableDisplayName string `json:"variableDisplayName"`
	SubsystemID ID `json:"subsystemId"`
	Locked bool `json:"isLocked"`
}

func (item Alert) ref() string {
	return item.ID.String()
}

func (s *AlertsServiceOp) Add(alert NewAlert) (Alert, error) {
	var respStruct Alert
	err := s.client.genericPost(alertsBasePath, alert, &respStruct)
	return respStruct, err
}

func (s *AlertsServiceOp) Get(id ID) (Alert, error) {
	var respStruct Alert
	err := s.client.genericGetID(alertsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *AlertsServiceOp) List() ([]Alert, error) {
	var respStruct []Alert
	err := s.client.genericGet(alertsBasePath, &respStruct)
	return respStruct, err
}

func (s *AlertsServiceOp) Remove(alert Alert) error {
	err := s.client.genericDeleteID(alertsBasePath, alert, nil)
	return err
}

func (s *AlertsServiceOp) Update(alert Alert) (Alert, error) {
	var respStruct Alert
	err := s.client.genericPutID(alertsBasePath, alert, &respStruct)
	return respStruct, err
}
