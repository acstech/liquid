package filters

import (
	"testing"
	"time"

	"github.com/acstech/liquid/core"
	"github.com/karlseguin/gspec"
)

func init() {
	core.Now = func() time.Time {
		t, _ := time.Parse("Mon Jan 02 15:04:05 2006", "Mon Jan 02 15:04:05 2006")
		return t
	}
}

func TestDateNowWithBasicFormat(t *testing.T) {
	spec := gspec.New(t)
	filter := DateFactory([]core.Value{stringValue("%Y %m %d")})
	spec.Expect(filter("now", nil).(string)).ToEqual("2006 01 02")
}

func TestDateTodayWithBasicFormat(t *testing.T) {
	spec := gspec.New(t)
	filter := DateFactory([]core.Value{stringValue("%H:%M:%S%%")})
	spec.Expect(filter("today", nil).(string)).ToEqual("15:04:05%")
}

func TestDateWithSillyFormat(t *testing.T) {
	spec := gspec.New(t)
	filter := DateFactory([]core.Value{stringValue("%w  %U  %j")})
	spec.Expect(filter("2014-01-10 21:31:28 +0800", nil).(string)).ToEqual("5  02  10")
}

func TestDateWithOffset(t *testing.T) {
	spec := gspec.New(t)
	filter := DateFactory([]core.Value{stringValue("%Y %m %d %H %M %S"), intValue(5)})
	spec.Expect(filter("now", nil).(string)).ToEqual("2006 01 02 20 04 05")
}

func TestDateOrdinalDays(t *testing.T) {
	spec := gspec.New(t)
	testFormat := func(test *testing.T, input string, format string, assert string) {
		filter := DateFactory([]core.Value{stringValue(format)})

		spec.Expect(filter(input, nil).(string)).ToEqual(assert)
	}

	testFormat(t, "1/1/2016", "%O", "1st")
	testFormat(t, "1/2/2016", "%O", "2nd")
	testFormat(t, "1/3/2016", "%O", "3rd")
	testFormat(t, "1/4/2016", "%O", "4th")
	testFormat(t, "1/5/2016", "%O", "5th")
	testFormat(t, "1/6/2016", "%O", "6th")
	testFormat(t, "1/7/2016", "%O", "7th")
	testFormat(t, "1/8/2016", "%O", "8th")
	testFormat(t, "1/9/2016", "%O", "9th")
	testFormat(t, "1/10/2016", "%O", "10th")
	testFormat(t, "1/11/2016", "%O", "11th")
	testFormat(t, "1/12/2016", "%O", "12th")
	testFormat(t, "1/13/2016", "%O", "13th")
	testFormat(t, "1/14/2016", "%O", "14th")
	testFormat(t, "1/15/2016", "%O", "15th")
	testFormat(t, "1/16/2016", "%O", "16th")
	testFormat(t, "1/17/2016", "%O", "17th")
	testFormat(t, "1/18/2016", "%O", "18th")
	testFormat(t, "1/19/2016", "%O", "19th")
	testFormat(t, "1/20/2016", "%O", "20th")
	testFormat(t, "1/21/2016", "%O", "21st")
	testFormat(t, "1/22/2016", "%O", "22nd")
	testFormat(t, "1/23/2016", "%O", "23rd")
	testFormat(t, "1/24/2016", "%O", "24th")
	testFormat(t, "1/25/2016", "%O", "25th")
	testFormat(t, "1/26/2016", "%O", "26th")
	testFormat(t, "1/27/2016", "%O", "27th")
	testFormat(t, "1/28/2016", "%O", "28th")
	testFormat(t, "1/29/2016", "%O", "29th")
	testFormat(t, "1/30/2016", "%O", "30th")
	testFormat(t, "1/31/2016", "%O", "31st")
}
