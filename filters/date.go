package filters

import (
	"context"
	"time"

	"github.com/acstech/go-strftime"
	"github.com/acstech/liquid/core"
)

var (
	zeroTime = time.Time{}
)

// Creates an date filter
func DateFactory(ctx context.Context, parameters []core.Value) core.Filter {
	if len(parameters) == 0 {
		return Noop
	}
	return (&DateFilter{parameters[0]}).ToString
}

type DateFilter struct {
	format core.Value
}

func (d *DateFilter) ToString(input interface{}, data map[string]interface{}) interface{} {
	time, ok := inputToTime(input)
	if ok == false {
		return input
	}
	return strftime.Strftime(&time, core.ToString(d.format.Resolve(data)))
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
	t, err := time.Parse("2006-01-02 15:04:05 -0700", s)
	if err != nil {
		return zeroTime, false
	}
	return t, true
}
