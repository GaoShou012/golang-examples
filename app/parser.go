package app

import "time"

func ParserDate(val string) (*time.Time, error) {
	format := "2006-01-02"
	theTime, err := time.ParseInLocation(format, val, time.Local)
	return &theTime, err
}