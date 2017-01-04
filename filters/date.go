package filters

import (
	"strings"
	"time"

	"github.com/acstech/go-strftime"
	"github.com/acstech/liquid/core"
)

var (
	zeroTime         = time.Time{}
	timeInputFormats = []string{
		"2006-01-02 15:04:05 -0700",
		"2006-01-02 15:04:05",
		"2006-01-02",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z",
		"1/2/2006 03:04:05 PM",
		"1/2/2006 3:04:05 PM",
		"1/2/2006",
		"1/2/2006 15:04:05",
	}
)

// DateFactory creates an date filter
func DateFactory(parameters []core.Value) core.Filter {
	switch len(parameters) {
	case 1:
		return (&DateFilter{parameters[0], nil}).ToString
	case 2:
		return (&DateFilter{parameters[0], parameters[1]}).ToString
	default:
		return Noop
	}
}

// DateFilter ...
type DateFilter struct {
	format core.Value
	offset core.Value
}

// ToString converts input to formatted date string
func (d *DateFilter) ToString(input interface{}, data map[string]interface{}) interface{} {
	t, ok := inputToTime(input)
	if ok == false {
		return input
	}

	if d.offset != nil {
		if offset, ok := core.ToInt(d.offset.Resolve(data)); ok {
			t = t.Add(time.Duration(offset) * time.Hour)
		}
	}

	format := processOrdinalDay(core.ToString(d.format.Resolve(data)), t)
	return strftime.Strftime(&t, format)
}

func inputToTime(input interface{}) (time.Time, bool) {
	switch typed := input.(type) {
	case time.Time:
		return typed, true
	case string:
		return timeFromString(typed)
	case []byte:
		return timeFromString(string(typed))
	}
	if n, ok := core.ToInt(input); ok {
		return core.Now().Add(time.Minute * time.Duration(n)), true
	}
	return zeroTime, false
}

func timeFromString(s string) (time.Time, bool) {
	if s == "now" || s == "today" {
		return core.Now(), true
	}

	for _, f := range timeInputFormats {
		t, err := time.Parse(f, s)
		if err == nil {
			return t, true
		}

	}

	return zeroTime, false
}

func processOrdinalDay(format string, time time.Time) string {
	if strings.Contains(format, "%O") {
		ordinal := "%-dth"
		day := time.Day()
		if !(day > 10 && day < 14) {
			switch day % 10 {
			case 1:
				ordinal = "%-dst"
				break
			case 2:
				ordinal = "%-dnd"
				break
			case 3:
				ordinal = "%-drd"
				break
			}
		}
		return strings.Replace(format, "%O", ordinal, -1)
	}

	return format
}
