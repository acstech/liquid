package filters

import (
	"strings"
	"time"

	"github.com/acstech/liquid/core"
	"github.com/acstech/liquid/strftime"
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
	case 0:
		return Noop
	case 1:
		return (&DateFilter{format: parameters[0]}).ToString
	case 2:
		switch parameters[1].(type) {
		case *core.StaticIntValue, *core.StaticFloatValue, *core.DynamicValue:
			return (&DateFilter{format: parameters[0], offset: parameters[1]}).ToString
		}
		return (&DateFilter{format: parameters[0], inputFormat: parameters[1]}).ToString
	default:
		switch parameters[1].(type) {
		case *core.StaticIntValue, *core.StaticFloatValue, *core.DynamicValue:
			return (&DateFilter{format: parameters[0], offset: parameters[1], inputFormat: parameters[2]}).ToString
		}
		return (&DateFilter{format: parameters[0], inputFormat: parameters[1], offset: parameters[2]}).ToString
	}
}

// DateFilter ...
type DateFilter struct {
	format      core.Value
	offset      core.Value
	inputFormat core.Value
}

// ToString converts input to formatted date string
func (d *DateFilter) ToString(input interface{}, data map[string]interface{}) interface{} {
	inputFormat := ""
	if d.inputFormat != nil {
		inputFormat = core.ToString(d.inputFormat.Resolve(data))
	}

	t, ok := inputToTime(input, inputFormat)
	if ok == false {
		return input
	}

	if d.offset != nil {
		offset := d.offset.Resolve(data)

		switch typedOffset := offset.(type) {
		case int:
			t = t.Add(time.Duration(typedOffset) * time.Hour)
		case float64:
			t = t.Add(time.Duration(typedOffset*60) * time.Minute)
		}
	}

	format := processOrdinalDay(core.ToString(d.format.Resolve(data)), t)
	return strftime.Strftime(&t, format)
}

func inputToTime(input interface{}, inputFormat string) (time.Time, bool) {
	switch typed := input.(type) {
	case time.Time:
		return typed, true
	case string:
		return timeFromString(typed, inputFormat)
	case []byte:
		return timeFromString(string(typed), inputFormat)
	}
	if n, ok := core.ToInt(input); ok {
		return core.Now().UTC().Add(time.Minute * time.Duration(n)), true
	}
	return zeroTime, false
}

func timeFromString(s, inputFormat string) (time.Time, bool) {
	if s == "now" || s == "today" {
		return core.Now().UTC(), true
	}

	if inputFormat != "" {
		if t, err := time.Parse(inputFormat, s); err == nil {
			return t, true
		}
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
