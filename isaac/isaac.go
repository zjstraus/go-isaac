package isaac

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	client *http.Client
	shared service

	BaseURL *url.URL
	Token string

	ActionLogs ActionLogsService
	Actions ActionLogsService
	ActivityLogs ActivityLogsService
	AvailablePanels AvailablePanelsService
	Events EventsService
	Playables PlayablesService
	RegisteredPanels RegisteredPanelsService
	SearchPresets SearchPresetsService
	Variables VariablesService
}

type service struct {
	client *Client
}

type ErrorMessage struct {
	Request *http.Request
	Response *http.Response
	Message string `json:"message"`
	Detail interface{} `json:"detail"`
}

func (e *ErrorMessage) Error() string {
	return e.Message
}

type ID struct {
	Str string
	Int int
}

func (id *ID) MarshalJSON() ([]byte, error) {
	if id.Int != 0 {
		return []byte(fmt.Sprintf("%d", id.Int)), nil
	}
	return []byte("\"" + id.Str + "\""), nil
}

func (id *ID) UnmarshalJSON(value []byte) error{
	err := json.Unmarshal(value, &id.Str)
	if err != nil {
		return err
	}
	i, err := strconv.Atoi(id.Str)
	if err != nil {
		return nil
	}
	id.Int = i
	return nil
}

func (id *ID) String() string {
	return id.Str
}

func NewClient(url *url.URL, token string) *Client {
	c := &Client{
		client: http.DefaultClient,
		BaseURL: url,
		Token: token,
	}

	c.ActionLogs = &ActionLogsServiceOp{client: c}
	c.Actions = &ActionLogsServiceOp{client: c}
	c.ActivityLogs = &ActivityLogsServiceOp{client: c}
	c.AvailablePanels = &AvailablePanelsServiceOp{client: c}
	c.Events = &EventsServiceOp{client: c}
	c.Playables = &PlayablesServiceOp{client: c}
	c.RegisteredPanels = &RegisteredPanelsServiceOp{client: c}
	c.SearchPresets = &SearchPresetsServiceOp{client: c}
	c.Variables = &VariablesServiceOp{client: c}

	return c
}

func (c *Client) NewRequest(method string, path string, body interface{}) (*http.Request, error) {
	reqUrl, urlErr := c.BaseURL.Parse(path)
	if urlErr != nil {
		return nil, urlErr
	}

	var bodyBuf io.ReadWriter
	if body != nil {
		bodyBuf = &bytes.Buffer{}
		encoder := json.NewEncoder(bodyBuf)
		encErr := encoder.Encode(body)
		if encErr != nil {
			return nil, encErr
		}
	}

	req, reqErr := http.NewRequest(method, reqUrl.String(), bodyBuf)
	if reqErr != nil {
		return nil, reqErr
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if c.Token != "" {
		req.Header.Set("isaac-token", c.Token)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "go-isaac API client")

	return req, nil
}

func (c *Client) DoRequest(req *http.Request, v interface{}) (*http.Response, error){
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode >= 300 {
		errorDetail := &ErrorMessage{
			Request:  req,
			Response: resp,
		}
		body, bodyErr := ioutil.ReadAll(resp.Body)
		if bodyErr == nil && len(body) > 0 {
			jsonErr := json.Unmarshal(body, errorDetail)
			if jsonErr != nil {
				errorDetail.Message = jsonErr.Error()
				errorDetail.Detail = body
			}
		}
		return resp, errorDetail
	}

	if v != nil {
		jsonErr := json.NewDecoder(resp.Body).Decode(v)
		if jsonErr != nil {
			return nil, jsonErr
		}
	}

	return resp, nil
}
