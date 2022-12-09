package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"
const timeZone = "Asia/Shanghai"

type Time time.Time

func (this Time) MarshalJSON() ([]byte, error) {
	res := make([]byte, 0, len(timeFormat)+2)
	res = append(res, '"')
	res = time.Time(this).AppendFormat(res, timeFormat)
	res = append(res, '"')
	return res, nil
}

func (this *Time) UnmarshalJSON(data []byte) (err error) {
	now, _ := time.ParseInLocation(`"`+timeFormat+`'`, string(data), time.Local)
	*this = Time(now)
	return
}

func (this Time) Local() time.Time {
	loc, _ := time.LoadLocation(timeZone)
	return time.Time(this).In(loc)
}

func (this Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(this)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (this *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*this = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
