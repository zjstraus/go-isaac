package isaac

import "time"

const (
	variablesBasePath = "/api/v1/variables"
)

type VariablesService interface {
	Add(NewVariable) (Variable, error)
	Get(ID) (Variable, error)
	List() ([]Variable, error)
	Remove(Variable) error
	Update(Variable) (Variable, error)
	UpdateValue(Variable, string) (Variable, error)
}

type VariablesServiceOp struct {
	client *Client
}

type NewVariable struct {
	Value string `json:"lastValue"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	SubsystemExternalID string `json:"subsystemExternalId"`
	ExternalRef string `json:"externalRef"`
	Type string `json:"type"`
	Tags []string `json:"tags,omitempty"`
	Log bool `json:"log"`
	StoreHistory bool `json:"store_history"`
}

type Variable struct {
	NewVariable `json:",inline"`
	ID ID `json:"_id"`
	SubsystemID ID `json:"subsystemId"`
	SubsystemDisplayName string `json:"subsystemDisplayName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ValueUpdatedAt time.Time `json:"lastValueUpdatedAt"`
	Alert bool `json:"alert"`
}

func (item Variable) ref() string {
	return item.ID.String()
}

func (s *VariablesServiceOp) Add(alert NewVariable) (Variable, error) {
	var respStruct Variable
	err := s.client.genericPost(variablesBasePath, alert, &respStruct)
	return respStruct, err
}

func (s *VariablesServiceOp) Get(id ID) (Variable, error) {
	var respStruct Variable
	err := s.client.genericGetID(variablesBasePath, id, &respStruct)
	return respStruct, err
}

func (s *VariablesServiceOp) List() ([]Variable, error) {
	var respStruct []Variable
	err := s.client.genericGet(variablesBasePath, &respStruct)
	return respStruct, err
}

func (s *VariablesServiceOp) Remove(alert Variable) error {
	err := s.client.genericDeleteID(variablesBasePath, alert, nil)
	return err
}

func (s *VariablesServiceOp) Update(alert Variable) (Variable, error) {
	var respStruct Variable
	err := s.client.genericPutID(variablesBasePath, alert, &respStruct)
	return respStruct, err
}

type postApiV1VariablesValueWrapper struct {
	Value string `json:"value"`
}

func (s *VariablesServiceOp) UpdateValue(variable Variable, value string) (Variable, error) {
	var respStruct Variable
	reqStruct := postApiV1VariablesValueWrapper{Value: value}
	err := s.client.genericPost(variablesBasePath + "/" + variable.ID.String() + "/value", reqStruct, &respStruct)
	return respStruct, err
}
