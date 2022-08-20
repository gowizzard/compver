package convert_test

import (
	"github.com/gowizzard/compver/convert"
	"reflect"
	"testing"
)

// TestInteger is to test the convert integer function
// We set a local variable & convert it to integer
// The return value is output for logging
func TestInteger(t *testing.T) {

	tests := []struct {
		number   string
		expected int
	}{
		{
			number:   "5",
			expected: 5,
		},
		{
			number:   "265",
			expected: 265,
		},
		{
			number:   "78",
			expected: 78,
		},
		{
			number:   "132",
			expected: 132,
		},
		{
			number:   "25",
			expected: 25,
		},
	}

	for _, value := range tests {

		integer := convert.Integer(value.number)

		if !reflect.DeepEqual(value.expected, integer) {
			t.Fatalf("expected: %d, got %d", value.expected, integer)
		}

	}

}
