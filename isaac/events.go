package isaac

const (
	eventsBasePath = "/api/v1/events"
)

type EventsService interface {
	Add(NewEvent) (Event, error)
	Get(ID) (Event, error)
	List() ([]Event, error)
	Remove(Event) error
	Update(Event) (Event, error)
}

type EventsServiceOp struct {
	client *Client
}

type NewEvent struct {
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	SubsystemExternalID string `json:"subsystemExternalId"`
	ExternalRef string `json:"externalRef"`
	Command string `json:"command"`
}


type Event struct {
	NewEvent `json:",inline"`
	ID ID `json:"_id"`
	SubsystemID ID `json:"subsystemId"`
}

func (item Event) ref() string {
	return item.ID.String()
}

func (s *EventsServiceOp) Add(alert NewEvent) (Event, error) {
	var respStruct Event
	err := s.client.genericPost(eventsBasePath, alert, &respStruct)
	return respStruct, err
}

func (s *EventsServiceOp) Get(id ID) (Event, error) {
	var respStruct Event
	err := s.client.genericGetID(eventsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *EventsServiceOp) List() ([]Event, error) {
	var respStruct []Event
	err := s.client.genericGet(eventsBasePath, &respStruct)
	return respStruct, err
}

func (s *EventsServiceOp) Remove(alert Event) error {
	err := s.client.genericDeleteID(eventsBasePath, alert, nil)
	return err
}

func (s *EventsServiceOp) Update(alert Event) (Event, error) {
	var respStruct Event
	err := s.client.genericPutID(eventsBasePath, alert, &respStruct)
	return respStruct, err
}
