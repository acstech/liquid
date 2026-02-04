package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortsAnArrayOfInteger(t *testing.T) {
	filter := SortFactory(nil)
	values := filter([]int{3, 4, 1, 2, 3}, nil).([]int)
	assert.Equal(t, len(values), 5)
	assert.Equal(t, values[0], 1)
	assert.Equal(t, values[1], 2)
	assert.Equal(t, values[2], 3)
	assert.Equal(t, values[3], 3)
	assert.Equal(t, values[4], 4)
}

func TestSortsAnArrayOfStrings(t *testing.T) {
	filter := SortFactory(nil)
	values := filter([]string{"cc", "b", "aa", "g"}, nil).([]string)
	assert.Equal(t, len(values), 4)
	assert.Equal(t, values[0], "aa")
	assert.Equal(t, values[1], "b")
	assert.Equal(t, values[2], "cc")
	assert.Equal(t, values[3], "g")
}

func TestSortsAnArrayOfFloats(t *testing.T) {
	filter := SortFactory(nil)
	values := filter([]float64{1.1, 0.9, 1233.2, 21.994}, nil).([]float64)
	assert.Equal(t, len(values), 4)
	assert.Equal(t, values[0], 0.9)
	assert.Equal(t, values[1], 1.1)
	assert.Equal(t, values[2], 21.994)
	assert.Equal(t, values[3], 1233.2)
}

func TestSortsSortableData(t *testing.T) {
	filter := SortFactory(nil)
	values := filter(People{&Person{"Leto"}, &Person{"Paul"}, &Person{"Jessica"}}, nil).(People)
	assert.Equal(t, len(values), 3)
	assert.Equal(t, values[0].Name, "Jessica")
	assert.Equal(t, values[1].Name, "Leto")
	assert.Equal(t, values[2].Name, "Paul")
}

func TestSortsOtherValuesAsStrings(t *testing.T) {
	filter := SortFactory(nil)
	values := filter([]interface{}{933, "spice", true, 123.44, "123", false}, nil).([]interface{})
	assert.Equal(t, len(values), 6)
	assert.Equal(t, values[0].(string), "123")
	assert.Equal(t, values[1].(float64), 123.44)
	assert.Equal(t, values[2].(int), 933)
	assert.Equal(t, values[3].(bool), false)
	assert.Equal(t, values[4].(string), "spice")
	assert.Equal(t, values[5].(bool), true)
}

func TestSortSkipsNonArrays(t *testing.T) {
	filter := SortFactory(nil)
	assert.Equal(t, filter(1343, nil).(int), 1343)
}

type People []*Person

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Person struct {
	Name string
}

func (p *Person) String() string {
	return p.Name
}
