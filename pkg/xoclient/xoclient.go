package xoclient

/*
This package is used as an interface for xolib, its mostly used for xcpng-csi so feature support will be limited.
however you can still used xolib since it has full capabilities thanks to its generic like interface

Pull requests are welcome to add more commands :)
*/

import "github.com/arturoguerra/go-xolib/pkg/xolib"

type (
	client struct {
		xolib.Xolib
	}

	// Client is the main interface used to interact with xo client
	Client interface {
		GetAllVDI() ([]*VDI, error)
		GetVDI(IDRef) (*VDI, error)
		CreateVDI(IDRef, int) (*VDI, error)

		GetAllVBD() ([]*VBD, error)
		GetVDB(IDRef) (*VBD, error)

		GetAllServers() ([]*Server, error)
		GetServer(IDRef) (*Server, error)

		GetAllObjects(*Filter) ([]*Generic, error)
	}
)

// NewClient returns a client
func NewClient(lib xolib.Xolib) Client {
	return &client{
		Xolib: lib,
	}
}
