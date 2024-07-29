package utils

import "time"

const TIMEZONE = "Asia/Bangkok"

func GetTimeLocation() (*time.Location, error) {
	loc, err := time.LoadLocation(TIMEZONE)
	if err != nil {
		return &time.Location{}, err
	}
	return loc, nil
}

func GetNow() (time.Time, error) {
	loc, err := GetTimeLocation()
	if err != nil {
		return time.Time{}, err
	}
	now := time.Now().In(loc)
	return now, nil
}
