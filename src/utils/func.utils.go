package utils

import "time"

func GetTimeWIB() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc)
}

func ConvertTimeWIB(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return t.In(loc)
}
