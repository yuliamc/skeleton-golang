package timeutil

import (
	"time"
)

type nowFuncT func() time.Time

var nowFunc nowFuncT

func init() {
	resetClockImplementation()
}

func resetClockImplementation() {
	nowFunc = func() time.Time {
		return time.Now()
	}
}

func now() time.Time {
	return nowFunc()
}

func Now() time.Time {
	return now()
}

func Yesterday() time.Time {
	return Now().AddDate(0, 0, -1)
}

func Tomorrow() time.Time {
	return Now().AddDate(0, 0, 1)
}

func DateAdd(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

func HoursAdd(t time.Time, hours int) time.Time {
	return t.Add(time.Hour * time.Duration(hours))
}

func MinutesAdd(t time.Time, minutes int) time.Time {
	return t.Add(time.Minute * time.Duration(minutes))
}

func NowStr(formats ...interface{}) string {
	return StrFormat(Now(), formats...)
}

func StrFormat(t time.Time, formats ...interface{}) string {
	var format string = ISO8601TimeWithoutZone
	if len(formats) > 0 {
		// First parameter is the format.
		var first = formats[0]
		if first != nil {
			var ok bool
			format, ok = formats[0].(string)
			if !ok {
				format = ISO8601TimeWithoutZone
			}
		}
	}

	return t.Format(format)
}

func DateDifferenceCounter(date1 time.Time, date2 time.Time) int {
	return int(date1.Sub(date2).Hours() / 24)
}

func Parse(t string, layout string) (time.Time, error) {
	return time.Parse(layout, t)
}

// today is also counted as loan date - hence the -1
func MaturityDate(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days-1)
}
