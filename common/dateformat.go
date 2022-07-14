package common

import (
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "02-Jan-2006"

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.ParseInLocation(ctLayout, string(b), loc)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(ctLayout)), nil
}

type CustomTimeStamp struct {
	time.Time
}

const ctsLayout = "Jan 02, 2006 15:04:05"

func (ct *CustomTimeStamp) UnmarshalJSON(b []byte) (err error) {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	ct.Time, err = time.ParseInLocation(ctsLayout, string(b), loc)
	return
}

func (ct *CustomTime) CustomTimeStamp() ([]byte, error) {
	return []byte(ct.Time.Format(ctsLayout)), nil
}
