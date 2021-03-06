package rest

import (
	"net/url"
	"testing"

	"github.com/couchbaselabs/go.assert"
)

func TestGetRestrictedIntQuery(t *testing.T) {

	defaultValue := uint64(42)
	minValue := uint64(20)
	maxValue := uint64(100)

	// make sure it returns default value when passed empty Values
	values := make(url.Values)
	restricted := getRestrictedIntQuery(
		values,
		"foo",
		defaultValue,
		minValue,
		maxValue,
		false,
	)
	assert.Equals(t, restricted, defaultValue)

	// make sure it returns default value when passed Values that doesn't contain key
	values.Set("bar", "99")
	restricted = getRestrictedIntQuery(
		values,
		"foo",
		defaultValue,
		minValue,
		maxValue,
		false,
	)
	assert.Equals(t, restricted, defaultValue)

	// make sure it returns appropriate value from Values
	values.Set("foo", "99")
	restricted = getRestrictedIntQuery(
		values,
		"foo",
		defaultValue,
		minValue,
		maxValue,
		false,
	)
	assert.Equals(t, restricted, uint64(99))

	// make sure it is limited to max when value value is over max
	values.Set("foo", "200")
	restricted = getRestrictedIntQuery(
		values,
		"foo",
		defaultValue,
		minValue,
		maxValue,
		false,
	)
	assert.Equals(t, restricted, maxValue)

	// make sure it is limited to min when value value is under min
	values.Set("foo", "1")
	restricted = getRestrictedIntQuery(
		values,
		"foo",
		defaultValue,
		minValue,
		maxValue,
		false,
	)
	assert.Equals(t, restricted, minValue)

	// Return zero when allowZero=true
	values.Set("foo", "0")
	restricted = getRestrictedIntQuery(
		values,
		"foo",
		defaultValue,
		minValue,
		maxValue,
		true,
	)
	assert.Equals(t, restricted, uint64(0))

	// Return minValue when allowZero=false
	values.Set("foo", "0")
	restricted = getRestrictedIntQuery(
		values,
		"foo",
		defaultValue,
		minValue,
		maxValue,
		false,
	)
	assert.Equals(t, restricted, minValue)
}
