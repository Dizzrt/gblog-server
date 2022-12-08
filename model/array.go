package model

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (this *Array) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), "|")
	*this = ss
	return nil
}

func (this Array) Value() (driver.Value, error) {
	str := strings.Join(this, "|")
	return str, nil
}
