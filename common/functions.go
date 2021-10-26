package common

import (
	"fmt"
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetTimeString(currentTime *time.Time) string {
	year := fmt.Sprintf("%v", currentTime.Year())
	month := fmt.Sprintf("%v", currentTime.Month())
	day := fmt.Sprintf("%v", currentTime.Day())
	if len(day) < 2 {
		day = fmt.Sprintf("0%v", day)
	}
	hour := fmt.Sprintf("%v", currentTime.Hour())
	if len(hour) < 2 {
		hour = fmt.Sprintf("0%v", hour)
	}
	minute := fmt.Sprintf("%v", currentTime.Minute())
	if len(minute) < 2 {
		minute = fmt.Sprintf("0%v", minute)
	}
	second := fmt.Sprintf("%v", currentTime.Second())
	if len(second) < 2 {
		second = fmt.Sprintf("0%v", second)
	}

	return fmt.Sprintf("%v-%v-%v %v:%v:%v", year, month, day, hour, minute, second)
}