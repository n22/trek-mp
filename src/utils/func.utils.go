package utils

import (
	"fmt"
	"time"
)

func GetTimeWIB() time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(loc)
}

func ConvertTimeWIB(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	return t.In(loc)
}

func IntToRp(val int64) string {
	if val == 0 {
		return ""
	}

	var result string
	var counter int
	strVal := fmt.Sprintf("%d", val)
	for i := len(strVal) - 1; i >= 0; i-- {
		if counter%3 == 0 && counter != 0 {
			result = fmt.Sprintf(".%v", result)
		}
		result = fmt.Sprintf("%s%s", string(strVal[i]), result)
		counter++
	}
	return "Rp" + result + ",00"
}
