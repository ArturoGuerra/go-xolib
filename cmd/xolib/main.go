package main

import (
	"fmt"

	"github.com/arturoguerra/go-xolib/pkg/xoclient"
	"github.com/arturoguerra/go-xolib/pkg/xolib"
)

func main() {
	cfg := xolib.LoadConfig()

	lib, err := xolib.NewXolib(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := xoclient.NewClient(lib)

	hostRef := xoclient.HostRef("8564fc43-26c8-49e5-b26a-f3f0e77db1fc")
	srRef := xoclient.SRRef("5fa59666-11cc-258e-aa26-99bd09c9dbef")

	vm, err := client.GetVMByName("mortis")
	if err != nil {
		fmt.Println(err)
		return
	}

	host, err := client.GetHostByUUID(hostRef)
	if err != nil {
		fmt.Println(err)
		return
	}

	if vm.PoolID != host.PoolID {
		fmt.Println("PoolIDs dont match")
		return
	}

	vdiRef, err := client.CreateVDI("test", int64(10737418240), srRef)
	if err != nil {
		fmt.Println(err)
		return
	}

	vdi, err := client.GetVDIByUUID(*vdiRef)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, ref := range vdi.VBDs {
		if err = client.DeleteVBD(ref); err != nil {
			fmt.Println(err)
		}
	}

	if err = client.AttachVBD(*vdiRef, vm.UUID); err != nil {
		fmt.Println(err)
		return
	}

	vdi, err = client.GetVDIByUUID(*vdiRef)
	if err != nil {
		fmt.Println(err)
		return
	}

	var vbd *xoclient.VBD

	for _, ref := range vdi.VBDs {
		if v, err := client.GetVBDByUUID(ref); err == nil {
			if v.VM == vm.UUID {
				vbd = v
				break
			}
		}
	}

	fmt.Println(vbd.Device)

	for _, ref := range vdi.VBDs {
		client.DisconnectVBD(ref)
		client.DeleteVBD(ref)
	}

	if err = client.DeleteVDI(*vdiRef); err != nil {
		fmt.Println(err)
		return
	}
}
