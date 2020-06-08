package xoclient

import (
	"errors"

	"github.com/arturoguerra/go-xolib/pkg/xolib"
)

// GetVMByName
func (c *client) GetVMByName(name string) (*VM, error) {
	filters := map[string]string{
		"type":       "VM",
		"name_label": name,
	}

	request := &xolib.MessageRequest{
		Method: "xo.getAllObjects",
	}

	resp, err := c.Call(request)
	if err != nil {
		return nil, err
	}

	// resp should be map[string]interface{} so we will treat it a such

	vms := make([]*VM, 0)

	for _, v := range (*resp).(map[string]interface{}) {
		vm := new(VM)
		if valid := c.Filter(v, filters, vm); valid {
			vms = append(vms, vm)
		}
	}

	if len(vms) == 1 {
		return vms[0], nil
	}

	return nil, errors.New("No VM found with this name")
}

// GetVMByUUID
func (c *client) GetVMByUUID(ref VMRef) (*VM, error) {
	filters := map[string]string{
		"type": "VM",
		"uuid": string(ref),
	}

	request := &xolib.MessageRequest{
		Method: "xo.getAllObjects",
	}

	resp, err := c.Call(request)
	if err != nil {
		return nil, err
	}

	// resp should be map[string]interface{} so we will treat it a such

	vms := make([]*VM, 0)

	for _, v := range (*resp).(map[string]interface{}) {
		vm := new(VM)
		if valid := c.Filter(v, filters, vm); valid {
			vms = append(vms, vm)
		}
	}

	if len(vms) == 1 {
		return vms[0], nil
	}

	return nil, errors.New("No VM found with this uuid")
}
