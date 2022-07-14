package common

import (
	"encoding/json"
	"strconv"
	"strings"
)

type CustomFloat64 float64

func (fi *CustomFloat64) UnmarshalJSON(b []byte) error {

	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	i, err := strconv.ParseFloat(strings.ReplaceAll(s, ",", ""), 64)
	if err != nil {
		return err
	}
	*fi = CustomFloat64(i)
	return nil
}
