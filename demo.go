package main

import (
	"encoding/json"
	"github/nsego/common"
)

type Foo struct {
	Value common.CustomFloat64 `json:"openInterest"`
}

func Demo() float64 {
	var f Foo

	if err := json.Unmarshal([]byte(`{ "openInterest": "1,12,300.50" }`), &f); err != nil {
		panic(err)
	}

	return float64(f.Value)
}
