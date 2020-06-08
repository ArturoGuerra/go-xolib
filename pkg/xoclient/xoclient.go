package xoclient

/*
This package is used as an interface for xolib, its mostly used for xcpng-csi so feature support will be limited.
however you can still used xolib since it has full capabilities thanks to its generic like interface

Pull requests are welcome to add more commands :)
*/

import (
	"github.com/arturoguerra/go-xolib/pkg/xolib"
)

type (
	client struct {
		xolib.Xolib
	}

	// VDIMethods are all methods that create/modify/get/delete a VDI from an SR
	VDIMethods interface {
		// CreateVDI : Name string, Size int64 (bytes), sr string (SR ID)
		CreateVDI(string, int64, SRRef) (*VDI, error)
		// DeleteVDI : VDI ID IDRef
		//DeleteVDI(VDIRef) error
		// ForceDeleteVDI : VDI ID IDRef
		//ForceDeleteVDI(VDIRef) error

		//GetVDIByName(string) ([]*VDI, error)
		GetVDIByUUID(VDIRef) (*VDI, error)
	}

	// VBDMethods are all methods that create/modify/get/delete a VBD from an VM
	VBDMethods interface {
		// Attach : Attaches a VDI to a VM by creating a VBD
		Attach(VDIRef, VMRef) (VBDRef, error)
		// Disconnect : Disconnects a VBD from a VM
		Disconnect(VBDRef) error
		// Delete : Deletes a VBD
		Delete(VBDRef) error

		GetVBDByName(string) ([]*VBD, error)
		GetVBDByUUID(VBDRef) (*VBD, error)
	}

	// VMMethods are methods used to create/modify/get/delete a VM from a HOST
	VMMethods interface {
		GetVMByName(string) ([]*VM, error)
		GetVMByUUID(VMRef) (*VM, error)
	}

	// SRMethods are methods used to create/modify/get/delete an SR from a HOST
	SRMethods interface {
		GetSRByName(string) ([]*SR, error)
		GetSRByUUID(string) (*SR, error)
	}

	// HostMethods are used to modify/add/remove/get a Host from Xen Orchestra
	HostMethods interface {
		GetHostByName(string) ([]*Host, error)
		GetHostByTag(string) ([]*Host, error)
		GetHostByUUID(string) (*Host, error)
	}

	// Client is the main interface used to interact with xo client
	Client interface {
		VDIMethods
		//VBDMethods
		//VMMethods
		//SRMethods
		//HostMethods
	}
)

// NewClient returns a client
func NewClient(lib xolib.Xolib) Client {
	return &client{
		Xolib: lib,
	}
}
