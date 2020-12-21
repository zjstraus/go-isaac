package isaac

import "net/http"

type isaacItem interface {
	ref() string
}

func (c *Client) genericPost(path string, item interface{}, outval interface{}) error {
	req, reqErr := c.NewRequest(http.MethodPost, path, item)
	if reqErr != nil {
		return reqErr
	}

	_, respErr := c.DoRequest(req, outval)

	return respErr
}


func (c *Client) genericPutID(basePath string, item isaacItem, outval interface{}) error {
	req, reqErr := c.NewRequest(http.MethodPut, basePath + "/" + item.ref(), item)
	if reqErr != nil {
		return reqErr
	}

	_, respErr := c.DoRequest(req, outval)

	return respErr
}

func (c *Client) genericGetID(basePath string, id ID, outval interface{}) error {
	req, reqErr := c.NewRequest(http.MethodGet, basePath + "/" + id.String(), nil)
	if reqErr != nil {
		return reqErr
	}

	_, respErr := c.DoRequest(req, outval)

	return respErr
}

func (c *Client) genericGet(basePath string, outval interface{}) error {
	req, reqErr := c.NewRequest(http.MethodGet, basePath, nil)
	if reqErr != nil {
		return reqErr
	}

	_, respErr := c.DoRequest(req, outval)

	return respErr
}

func (c *Client) genericDeleteID(basePath string, item isaacItem, outval interface{}) error {
	req, reqErr := c.NewRequest(http.MethodDelete, basePath + "/" + item.ref(), nil)
	if reqErr != nil {
		return reqErr
	}

	_, respErr := c.DoRequest(req, outval)

	return respErr
}
