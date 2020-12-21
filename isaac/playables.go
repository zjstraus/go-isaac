package isaac

const (
	playablesBasePath = "/api/v1/playables"
)

type PlayablesService interface {
	Add(NewPlayable) (Playable, error)
	Get(ID) (Playable, error)
	List() ([]Playable, error)
	Remove(Playable) error
	Update(Playable) (Playable, error)
}

type PlayablesServiceOp struct {
	client *Client
}

type NewPlayable struct {
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	Duration int `json:"duration"`
	Group string `json:"group"`
	SubsystemExternalID string `json:"subsystemExternalId"`
	ExternalRef string `json:"externalRef"`
	Command string `json:"command"`
	ForceCached bool `json:"forceCached"`
	IsCached bool `json:"isCached"`
	Type string `json:"type"`
	Metadata interface{} `json:"metadata"`
	DurationType string `json:"durationType"`
}

type TemplatePath struct {
	Group string `json:"group"`
	Name string `json:"name"`
	ID ID `json:"_id"`
}

func (item TemplatePath) ref() string {
	return item.ID.String()
}

type Playable struct {
	NewPlayable `json:",inline"`
	ID ID `json:"_id"`
	SubsystemID ID `json:"subsystemId"`
}

func (item Playable) ref() string {
	return item.ID.String()
}

func (s *PlayablesServiceOp) Add(alert NewPlayable) (Playable, error) {
	var respStruct Playable
	err := s.client.genericPost(playablesBasePath, alert, &respStruct)
	return respStruct, err
}

func (s *PlayablesServiceOp) Get(id ID) (Playable, error) {
	var respStruct Playable
	err := s.client.genericGetID(playablesBasePath, id, &respStruct)
	return respStruct, err
}

func (s *PlayablesServiceOp) List() ([]Playable, error) {
	var respStruct []Playable
	err := s.client.genericGet(playablesBasePath, &respStruct)
	return respStruct, err
}

func (s *PlayablesServiceOp) Remove(alert Playable) error {
	err := s.client.genericDeleteID(playablesBasePath, alert, nil)
	return err
}

func (s *PlayablesServiceOp) Update(alert Playable) (Playable, error) {
	var respStruct Playable
	err := s.client.genericPutID(playablesBasePath, alert, &respStruct)
	return respStruct, err
}
