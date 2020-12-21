package isaac

import (
	"time"
)

const (
	searchPresetsBasePath = "/api/v1/logs/presets"
)

type SearchPresetsService interface {
	Add(NewSearchPreset) (SearchPreset, error)
	Get(ID) (SearchPreset, error)
	List() ([]SearchPreset, error)
	Remove(SearchPreset) error
	Update(SearchPreset) (SearchPreset, error)
}

type SearchPresetsServiceOp struct {
	client *Client
}

type NewSearchPreset struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Tags []string `json:"tags,omitempty"`
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
	UserID ID `json:"userId"`
	Filters map[string]string `json:"filters"`
	CurrentSearch string `json:"currentSearch"`
	SearchOptions []SearchPresetOption `json:"searchOptions"`
	BasicFilter string `json:"basicFilter"`
}

type SearchPresetOption struct {
	Name string `json:"name"`
	DisplayName string `json:"displayName"`
	Value string `json:"value"`
}

type SearchPreset struct {
	NewSearchPreset `json:",inline"`
	ID ID `json:"_id"`
	Global bool `json:"isGlobal"`
}

func (item SearchPreset) ref() string {
	return item.ID.String()
}

func (s *SearchPresetsServiceOp) Add(preset NewSearchPreset) (SearchPreset, error) {
	var respStruct SearchPreset
	err := s.client.genericPost(searchPresetsBasePath, preset, &respStruct)
	return respStruct, err
}

func (s *SearchPresetsServiceOp) Get(id ID) (SearchPreset, error) {
	var respStruct SearchPreset
	err := s.client.genericGetID(searchPresetsBasePath, id, &respStruct)
	return respStruct, err
}

func (s *SearchPresetsServiceOp) List() ([]SearchPreset, error) {
	var respStruct []SearchPreset
	err := s.client.genericGet(searchPresetsBasePath, &respStruct)
	return respStruct, err
}

func (s *SearchPresetsServiceOp) Remove(preset SearchPreset) error {
	err := s.client.genericDeleteID(searchPresetsBasePath, preset, nil)
	return err
}

func (s *SearchPresetsServiceOp) Update(preset SearchPreset) (SearchPreset, error) {
	var respStruct SearchPreset
	err := s.client.genericPutID(searchPresetsBasePath, preset, &respStruct)
	return respStruct, err
}
