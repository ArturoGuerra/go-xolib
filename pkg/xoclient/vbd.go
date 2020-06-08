package xoclient

import (
	"errors"

	"github.com/arturoguerra/go-xolib/pkg/xolib"
)

// AttachVBD
func (c *client) AttachVBD(vdiref VDIRef, vmref VMRef) error {
	params := &xolib.Params{
		"mode": "RW",
		"vdi":  string(vdiref),
		"vm":   string(vmref),
	}

	request := &xolib.MessageRequest{
		Method: "vm.attachDisk",
		Params: params,
	}

	resp, err := c.Call(request)
	if err != nil {
		return err
	}

	if (*resp).(bool) {
		return nil
	}

	return errors.New("Unable to attach")
}

// GetVBDByUUID
func (c *client) GetVBDByUUID(ref VBDRef) (*VBD, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be a map[string]interface{} and will be treated as such

	vbds := make([]*VBD, 0)
	filters := map[string]string{
		"type": "VBD",
		"uuid": string(ref),
	}

	for _, v := range (*resp).(map[string]interface{}) {
		vbd := new(VBD)
		if valid := c.Filter(v, filters, vbd); valid {
			vbds = append(vbds, vbd)
		}
	}

	if len(vbds) == 1 {
		return vbds[0], nil
	}

	return nil, errors.New("VBD Not found")
}

// GetVBDByName
func (c *client) GetVBDByName(name string) (*VBD, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	// resp should be a map[string]interface{} and will be treated as such

	vbds := make([]*VBD, 0)
	filters := map[string]string{
		"type":       "VBD",
		"name_label": name,
	}

	for _, v := range (*resp).(map[string]interface{}) {
		vbd := new(VBD)
		if valid := c.Filter(v, filters, vbd); valid {
			vbds = append(vbds, vbd)
		}
	}

	if len(vbds) == 1 {
		return vbds[0], nil
	}

	return nil, errors.New("VBD Not found")
}

// DisconnectVBD
func (c *client) DisconnectVBD(ref VBDRef) error {
	params := &xolib.Params{
		"id": string(ref),
	}

	request := &xolib.MessageRequest{
		Method: "vbd.disconnect",
		Params: params,
	}

	resp, err := c.Call(request)
	if err != nil {
		return err
	}

	disconnected := (*resp).(bool)

	if !disconnected {
		return errors.New("Error disconneting VBD")
	}

	return nil
}

// DeleteVBD
func (c *client) DeleteVBD(ref VBDRef) error {
	params := &xolib.Params{
		"id": string(ref),
	}

	request := &xolib.MessageRequest{
		Method: "vbd.delete",
		Params: params,
	}

	resp, err := c.Call(request)
	if err != nil {
		return err
	}

	deleted := (*resp).(bool)

	if !deleted {
		return errors.New("Error deleting VBD")
	}

	return nil
}

// ConnectVBD
func (c *client) ConnectVBD(ref VBDRef) error {
	params := &xolib.Params{
		"id": string(ref),
	}

	request := &xolib.MessageRequest{
		Method: "vbd.connect",
		Params: params,
	}

	resp, err := c.Call(request)
	if err != nil {
		return err
	}

	connected := (*resp).(bool)

	if !connected {
		return errors.New("Error conneting VBD")
	}

	return nil
}

// GetVBDsFromVDI
func (c *client) GetVBDsFromVDI(ref VDIRef) ([]*VBD, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	filters := map[string]string{
		"type": "VBD",
		"VDI":  string(ref),
	}

	vbds := make([]*VBD, 0)

	for _, v := range (*resp).(map[string]interface{}) {
		vbd := new(VBD)
		if valid := c.Filter(v, filters, vbd); valid {
			vbds = append(vbds, vbd)
		}
	}

	if len(vbds) > 0 {
		return vbds, nil
	}

	return nil, errors.New("No VBD was found")
}

// GetVBDsFromVM
func (c *client) GetVBDsFromVM(ref VMRef) ([]*VBD, error) {
	resp, err := c.getAllObjects()
	if err != nil {
		return nil, err
	}

	filters := map[string]string{
		"type": "VBD",
		"VM":   string(ref),
	}

	vbds := make([]*VBD, 0)

	for _, v := range (*resp).(map[string]interface{}) {
		vbd := new(VBD)
		if valid := c.Filter(v, filters, vbd); valid {
			vbds = append(vbds, vbd)
		}
	}

	if len(vbds) > 0 {
		return vbds, nil
	}

	return nil, errors.New("No VBD was found")
}
