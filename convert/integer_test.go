package convert_test

import (
	"github.com/gowizzard/compver/convert"
	"testing"
)

// TestInteger is to test the convert integer function
// We set a local variable & convert it to integer
// The return value is output for logging
func TestInteger(t *testing.T) {

	value := "5"
	integer := convert.Integer(value)

	t.Logf("The gotten value of the integer is \"%d\". The expected value was \"%s\".\n", integer, value)

}
