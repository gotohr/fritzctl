package fritz

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSwitchAndThermostatFiltering tests on the correctness of the switch/thermostat separation of a given device list.
func TestSwitchAndThermostatFiltering(t *testing.T) {
	f, err := os.Open("../testdata/devicelist_fritzos06.83.xml")
	assert.NoError(t, err)
	defer f.Close()

	var l Devicelist
	err = xml.NewDecoder(f).Decode(&l)
	assert.NoError(t, err)

	assert.Len(t, l.Thermostats(), 2)
	assert.Len(t, l.Switches(), 1)

	assert.Equal(t, len(l.Devices), len(l.Switches())+len(l.Thermostats()))
}

// TestSwitchAndThermostatFilteringIssue56 reproduces https://github.com/bpicode/fritzctl/issues/59.
func TestSwitchAndThermostatFilteringIssue56(t *testing.T) {
	f, err := os.Open("../testdata/devicelist_issue_59.xml")
	assert.NoError(t, err)
	defer f.Close()

	var l Devicelist
	err = xml.NewDecoder(f).Decode(&l)
	assert.NoError(t, err)

	assert.Len(t, l.Thermostats(), 4)
	assert.Len(t, l.Switches(), 8)

	assert.Equal(t, len(l.Devices), len(l.Switches())+len(l.Thermostats()))
}
