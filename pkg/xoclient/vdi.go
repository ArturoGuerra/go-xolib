package xoclient

import (
	"errors"
	"fmt"

	"github.com/arturoguerra/go-xolib/pkg/xolib"
)

// GetVDIByUUID : ref VDIRef (string)
func (c *client) GetVDIByUUID(ref VDIRef) (*VDI, error) {
	request := xolib.MessageRequest{
		Method: "xo.getAllObjects",
	}

	resp, err := c.Call(&request)
	if err != nil {
		return nil, err
	}

	// resp should be a map[string]interface{} so we will treat it a such

	vdis := make([]*VDI, 0)

	filters := map[string]string{
		"type": "VDI",
		"uuid": string(ref),
	}

	for _, v := range (*resp).(map[string]interface{}) {
		vdi := new(VDI)
		if valid := c.Filter(v, filters, vdi); valid {
			vdis = append(vdis, vdi)
		}
	}

	if len(vdis) == 1 {
		return vdis[0], nil
	}

	return nil, errors.New("VDI Not found")
}

// CreateVDI : Name string, Size int64, SR SRRef
func (c *client) CreateVDI(name string, size int64, sr SRRef) (*VDI, error) {

	params := xolib.Params{
		"name": name,
		"size": size,
		"sr":   sr,
	}

	request := xolib.MessageRequest{
		Method: "disk.create",
		Params: &params,
	}

	resp, err := c.Call(&request)
	if err != nil {
		return nil, err
	}

	/*
		Resp here should be a VDIRef so we will treat it as such
	*/

	ref := VDIRef((*resp).(string))

	vdi, err := c.GetVDIByUUID(ref)
	return vdi, err
}

// DeleteVDI : deletes VDI
func (c *client) DeleteVDI(ref VDIRef) error {
	params := &xolib.Params{
		"id": string(ref),
	}

	request := &xolib.MessageRequest{
		Method: "vdi.delete",
		Params: params,
	}

	resp, err := c.Call(request)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}
