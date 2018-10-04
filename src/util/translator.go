package util

import (
	"strconv"
	"time"
)

func FromTime(time time.Time) string {
	return strconv.FormatInt(time.Unix(), 10)
}

func ToTime(text string) time.Time {
	intValue, _ := strconv.ParseInt(text, 10, 0)
	return time.Unix(intValue, 0)
}