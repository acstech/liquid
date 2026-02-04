package filters

import (
	"testing"
	"time"

	"github.com/acstech/liquid/core"
	"github.com/stretchr/testify/assert"
)

func init() {
	core.Now = func() time.Time {
		t, _ := time.Parse("Mon Jan 02 15:04:05 2006", "Mon Jan 02 15:04:05 2006")
		return t
	}
}

func TestDateNowWithBasicFormat(t *testing.T) {
	filter := DateFactory([]core.Value{stringValue("%Y %m %d")})
	assert.Equal(t, filter("now", nil).(string), "2006 01 02")
}

func TestDateTodayWithBasicFormat(t *testing.T) {
	filter := DateFactory([]core.Value{stringValue("%H:%M:%S%%")})
	assert.Equal(t, filter("today", nil).(string), "15:04:05%")
}

func TestDateWithSillyFormat(t *testing.T) {
	filter := DateFactory([]core.Value{stringValue("%w  %U  %j")})
	assert.Equal(t, filter("2014-01-10 21:31:28 +0800", nil).(string), "5  02  10")
}

func TestDateWithOffset(t *testing.T) {
	filter := DateFactory([]core.Value{stringValue("%Y %m %d %H %M %S"), intValue(5)})
	assert.Equal(t, filter("now", nil).(string), "2006 01 02 20 04 05")
}

func TestDateFormatting(t *testing.T) {
	input := "2013-02-06 09:37:42.753734088"
	expected := map[string]string{
		// literal %
		"%%":   "%",
		"%%%Y": "%2013",
		"%Y%%": "2013%",

		// year
		"%Y": "2013",
		"%G": "2013",
		"%C": "20",
		"%y": "13",
		"%g": "13",

		// month
		"%m":  "02",
		"%_m": " 2",
		"%-m": "2",
		"%B":  "February",
		"%^B": "FEBRUARY",
		"%b":  "Feb",
		"%^b": "FEB",
		"%h":  "Feb",

		// day
		"%d":  "06",
		"%-d": "6",
		"%e":  " 6",

		// day of the year
		"%j": "37",

		// Time
		// Hour
		"%H": "09",
		"%k": " 9",
		"%I": "09",
		"%l": " 9",
		"%P": "am",
		"%p": "AM",

		// Minute
		"%M": "37",

		// Second
		"%S": "42",

		// Millisecond of the second
		"%L": ".753",

		// Fractional seconds digits
		"%1N": ".7",
		"%2N": ".75",
		"%3N": ".753",
		"%4N": ".7537",
		"%5N": ".75373",
		"%6N": ".753734",
		"%7N": ".753734",
		"%8N": ".75373408",
		"%9N": ".753734088",

		// Time zone
		"%z":  "+0000",
		"%Z":  "UTC",
		"%A":  "Wednesday",
		"%^A": "WEDNESDAY",
		"%a":  "Wed",
		"%^a": "WED",
		"%w":  "3",
		"%u":  "3",
		"%V":  "06",
		"%s":  "1360143462",
		"%Q":  "1360143462753",
		"%t":  "\t",
		"%n":  "\n",
		"%c":  "Wed Feb  6 09:37:42 2013",
		"%D":  "02/06/13",
		"%v":  " 6-Feb-13",
		"%X":  "09:37:42",
		"%T":  "09:37:42",
		"%r":  "09:37:42 AM",
		"%R":  "09:37",

		"%Y-%m":  "2013-02",
		"%Y-%_m": "2013- 2",
		"%Y-%-m": "2013-2",

		"%y-%m":  "13-02",
		"%y-%_m": "13- 2",
		"%y-%-m": "13-2",

		"%Y-%m-%d":   "2013-02-06",
		"%Y-%_m-%d":  "2013- 2-06",
		"%Y-%-m-%d":  "2013-2-06",
		"%Y-%-m-%-d": "2013-2-6",
		"%O":         "6th",
	}

	for f, a := range expected {
		testFormat(t, input, f, a)
	}
}

func TestDateDates(t *testing.T) {
	testFormat := func(test *testing.T, input string, format string, assert string) {
		filter := DateFactory([]core.Value{&core.StaticStringValue{Value: format}})

		var data map[string]interface{} = make(map[string]interface{})
		data["p1"] = format
		if rs := filter(input, data); rs != assert {
			test.Errorf("Format: %s, Wanted: %s, Got: %s", format, assert, rs)
		}
	}

	input := []string{"2/6/2013", "2013-02-06"}
	expected := map[string]string{
		"%c": "Wed Feb  6 00:00:00 2013",
		"%D": "02/06/13",
		"%v": " 6-Feb-13",
		"%r": "12:00:00 AM",
		"%R": "00:00",

		"%Y-%m":  "2013-02",
		"%Y-%_m": "2013- 2",
		"%Y-%-m": "2013-2",

		"%y-%m":  "13-02",
		"%y-%_m": "13- 2",
		"%y-%-m": "13-2",

		"%Y-%m-%d":   "2013-02-06",
		"%Y-%_m-%d":  "2013- 2-06",
		"%Y-%-m-%d":  "2013-2-06",
		"%Y-%-m-%-d": "2013-2-6",
		"%O":         "6th",
	}

	for _, in := range input {
		for f, a := range expected {
			testFormat(t, in, f, a)
		}
	}
}

func TestDateDatesAndTimes(t *testing.T) {
	input := []string{"2/6/2013 13:15:00", "2013-02-06 13:15:00 -0700", "2013-02-06 13:15:00", "2013-02-06T13:15:00", "2013-02-06T13:15:00Z", "2/6/2013 1:15:00 PM", "2/6/2013 01:15:00 PM"}
	expected := map[string]string{
		"%c": "Wed Feb  6 13:15:00 2013",
		"%D": "02/06/13",
		"%v": " 6-Feb-13",
		"%r": "01:15:00 PM",
		"%R": "13:15",

		"%Y-%m":  "2013-02",
		"%Y-%_m": "2013- 2",
		"%Y-%-m": "2013-2",

		"%y-%m":  "13-02",
		"%y-%_m": "13- 2",
		"%y-%-m": "13-2",

		"%Y-%m-%d":   "2013-02-06",
		"%Y-%_m-%d":  "2013- 2-06",
		"%Y-%-m-%d":  "2013-2-06",
		"%Y-%-m-%-d": "2013-2-6",
		"%O":         "6th",
	}

	for _, in := range input {
		for f, a := range expected {
			testFormat(t, in, f, a)
		}
	}
}

func TestDateOrdinalDays(t *testing.T) {
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

func TestDateDatesAndTimesWithOffset(t *testing.T) {
	testFormat := func(test *testing.T, input string, format string, expected string) {
		filter := DateFactory([]core.Value{stringValue(format), intValue(4)})
		assert.Equal(t, filter(input, nil).(string), expected)
	}

	input := []string{"2/6/2013 13:15:00", "2013-02-06 13:15:00 -0000", "2013-02-06 13:15:00", "2013-02-06T13:15:00", "2013-02-06T13:15:00Z", "2/6/2013 1:15:00 PM", "2/6/2013 01:15:00 PM"}
	expected := map[string]string{
		"%c": "Wed Feb  6 17:15:00 2013",
		"%D": "02/06/13",
		"%v": " 6-Feb-13",
		"%r": "05:15:00 PM",
		"%R": "17:15",

		"%Y-%m":  "2013-02",
		"%Y-%_m": "2013- 2",
		"%Y-%-m": "2013-2",

		"%y-%m":  "13-02",
		"%y-%_m": "13- 2",
		"%y-%-m": "13-2",

		"%Y-%m-%d":   "2013-02-06",
		"%Y-%_m-%d":  "2013- 2-06",
		"%Y-%-m-%d":  "2013-2-06",
		"%Y-%-m-%-d": "2013-2-6",
		"%O":         "6th",
	}

	for _, in := range input {
		for f, a := range expected {
			testFormat(t, in, f, a)
		}
	}
}

func TestDateDatesAndTimesWithFloatOffset(t *testing.T) {
	testFormat := func(test *testing.T, input string, format string, expected string) {
		filter := DateFactory([]core.Value{stringValue(format), floatValue(4.5)})
		assert.Equal(t, filter(input, nil).(string), expected)
	}

	input := []string{"2/6/2013 13:15:00", "2013-02-06 13:15:00 -0000", "2013-02-06 13:15:00", "2013-02-06T13:15:00", "2013-02-06T13:15:00Z", "2/6/2013 1:15:00 PM", "2/6/2013 01:15:00 PM"}
	expected := map[string]string{
		"%c": "Wed Feb  6 17:45:00 2013",
		"%D": "02/06/13",
		"%v": " 6-Feb-13",
		"%r": "05:45:00 PM",
		"%R": "17:45",

		"%Y-%m":  "2013-02",
		"%Y-%_m": "2013- 2",
		"%Y-%-m": "2013-2",

		"%y-%m":  "13-02",
		"%y-%_m": "13- 2",
		"%y-%-m": "13-2",

		"%Y-%m-%d":   "2013-02-06",
		"%Y-%_m-%d":  "2013- 2-06",
		"%Y-%-m-%d":  "2013-2-06",
		"%Y-%-m-%-d": "2013-2-6",
		"%O":         "6th",
	}

	for _, in := range input {
		for f, a := range expected {
			testFormat(t, in, f, a)
		}
	}
}

func TestDateDatesAndTimesWithOffsetAndInputFormat(t *testing.T) {
	testFormat := func(test *testing.T, input string, format string, expected string, inputFormat string) {
		filter := DateFactory([]core.Value{stringValue(format), intValue(4), stringValue(inputFormat)})
		assert.Equal(t, filter(input, nil).(string), expected)
	}

	input := []string{"2-6-2013 17:15:00"}
	inputFormat := []string{"1-2-2006 15:04:05"}
	expected := map[string]string{
		"%c": "Wed Feb  6 21:15:00 2013",
		"%D": "02/06/13",
		"%v": " 6-Feb-13",
		"%r": "09:15:00 PM",
		"%R": "21:15",

		"%Y-%m":  "2013-02",
		"%Y-%_m": "2013- 2",
		"%Y-%-m": "2013-2",

		"%y-%m":  "13-02",
		"%y-%_m": "13- 2",
		"%y-%-m": "13-2",

		"%Y-%m-%d":   "2013-02-06",
		"%Y-%_m-%d":  "2013- 2-06",
		"%Y-%-m-%d":  "2013-2-06",
		"%Y-%-m-%-d": "2013-2-6",
		"%O":         "6th",
	}

	for idx, in := range input {
		for f, a := range expected {
			testFormat(t, in, f, a, inputFormat[idx])
		}
	}
}

func TestFormatDateDatesAndTimesWithInputFormat(t *testing.T) {
	testFormat := func(test *testing.T, input string, format string, expected string, inputFormat string) {
		filter := DateFactory([]core.Value{stringValue(format), stringValue(inputFormat)})
		assert.Equal(t, filter(input, nil).(string), expected)
	}

	input := []string{"2-6-2013 17:15:00"}
	inputFormat := []string{"1-2-2006 15:04:05"}
	expected := map[string]string{
		"%c": "Wed Feb  6 17:15:00 2013",
		"%D": "02/06/13",
		"%v": " 6-Feb-13",
		"%r": "05:15:00 PM",
		"%R": "17:15",

		"%Y-%m":  "2013-02",
		"%Y-%_m": "2013- 2",
		"%Y-%-m": "2013-2",

		"%y-%m":  "13-02",
		"%y-%_m": "13- 2",
		"%y-%-m": "13-2",

		"%Y-%m-%d":   "2013-02-06",
		"%Y-%_m-%d":  "2013- 2-06",
		"%Y-%-m-%d":  "2013-2-06",
		"%Y-%-m-%-d": "2013-2-6",
		"%O":         "6th",
	}

	for idx, in := range input {
		for f, a := range expected {
			testFormat(t, in, f, a, inputFormat[idx])
		}
	}
}

func testFormat(t *testing.T, input string, format string, expected string) {
	filter := DateFactory([]core.Value{stringValue(format)})
	assert.Equal(t, filter(input, nil).(string), expected)
}
