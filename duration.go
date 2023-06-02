package humanize

import (
	"fmt"
	"time"
)

// A Duration is a time.Duration that is more readable for long
// time. i.e. 103h03m15s will format to 4 days 7 hours
type Duration time.Duration

// Missing time package constants for larger times
const (
	Day   = 24 * time.Hour // a day
	Week  = 7 * Day        // a Week
	Month = 30 * Day       // a Month
	Year  = 365 * Day      // a Year
)

// String implements fmt.Stringer interface
func (d Duration) String() string {
	dur := time.Duration(d)
	switch {
	case dur > Year:
		dur = dur.Truncate(Day)
		return formatDuration(dur, Year, Month, "year", "month")
	case dur > Month:
		dur = dur.Truncate(Day)
		return formatDuration(dur, Month, Week, "month", "week")
	case dur > Week:
		dur = dur.Truncate(Day)
		return formatDuration(dur, Week, Day, "week", "day")
	case dur > Day:
		dur = dur.Truncate(time.Hour)
		return formatDuration(dur, Day, time.Hour, "day", "hour")
	case dur > time.Hour:
		dur = dur.Truncate(time.Minute)
		return formatDuration(dur, time.Hour, time.Minute, "h", "m")
	case dur > time.Minute:
		return fmt.Sprintf("%s", dur.Truncate(time.Second))
	case dur > time.Second:
		return fmt.Sprintf("%s", dur.Truncate(time.Millisecond))
	default:
		return fmt.Sprintf("%s", dur)
	}
}

func mayPluralize(units string, n int, onlyLeft bool) string {
	if len(units) == 1 {
		return units
	}
	if n > 1 {
		units += "s"
	}
	if onlyLeft == true {
		return " " + units
	}
	return " " + units + " "
}

func formatDuration(d, base, rem time.Duration, baseUnit, remUnit string) string {
	trunc := d.Truncate(base)
	remainder := int((d - trunc) / rem)
	value := int(trunc / base)
	if remainder <= 0 {
		return fmt.Sprintf("%d%s", value, mayPluralize(baseUnit, value, true))
	}
	return fmt.Sprintf("%d%s%d%s",
		value, mayPluralize(baseUnit, value, false),
		remainder, mayPluralize(remUnit, remainder, true))
}
