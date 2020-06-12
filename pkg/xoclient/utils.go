package xoclient

import (
	"github.com/arturoguerra/go-xolib/pkg/xolib"
	"github.com/mitchellh/mapstructure"
)

// Filter : filters
func (c *client) Filter(resp interface{}, filters map[string]string, object interface{}) bool {
	switch val := resp.(type) {
	case map[string]interface{}:
		for k, v := range filters {
			if val[k] != v {
				return false
			}
		}

		mapstructure.Decode(val, object)
		return true
	}

	return false
}

// getAllObjects
func (c *client) getAllObjects() (*xolib.MessageResult, error) {
	request := &xolib.MessageRequest{
		Method: "xo.getAllObjects",
	}

	resp, err := c.Call(request)

	return resp, err
}

// GetAll is just like getAllObjects
func (c *client) GetAll() (*xolib.MessageResult, error) {
	return c.getAllObjects()
}
