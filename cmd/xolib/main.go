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

	//sr := xoclient.SRRef("5fa59666-11cc-258e-aa26-99bd09c9dbef")

	//	vdi, err := client.CreateVDI("test", int64(10737418240), sr)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}

	//ref := xoclient.VDIRef("13d510c7-9952-456f-afc9-8ec11e32d47a")
	//if err = client.DeleteVDI(ref); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	vms, err := client.GetVMByName("mortis")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(vms)
}
