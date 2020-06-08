package main

import (
	"fmt"

	"github.com/arturoguerra/go-xolib/pkg/xoclient"
	"github.com/arturoguerra/go-xolib/pkg/xolib"
	"github.com/mitchellh/mapstructure"
)

func parse(v interface{}) {
	for _, vv := range v.(map[string]interface{}) {
		switch val := vv.(type) {
		case nil:
		case int:
		case int64:
		case int32:
		case bool:
		case float64:
		case float32:
		case string:
		case []interface{}:
		case map[string]interface{}:
			parse(val)
		default:
			fmt.Println(val)
		}
	}
}

func filter(resp interface{}, filters map[string]string, object interface{}) bool {
	switch val := (resp).(type) {
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

func main() {
	cfg := xolib.LoadConfig()

	lib, err := xolib.NewXolib(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := xoclient.NewClient(lib)

	sr := xoclient.SRRef("5fa59666-11cc-258e-aa26-99bd09c9dbef")

	vdi, err := client.CreateVDI("test", int64(10737418240), sr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(vdi)
}
