package time

import (
	"time"

	"github.com/mergestat/timediff"
)

func GetCurrentTime() string {
	t := time.Now()

	formattedTime := t.Format(time.RFC3339)
	return formattedTime
}

func GetReadableTime(sorTime string) string {
	t, err := time.Parse(time.RFC3339, sorTime)
	if err != nil {
		panic(err)
	}
	return timediff.TimeDiff(time.Now().Add(time.Until(t)))
}
