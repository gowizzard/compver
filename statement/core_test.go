package statement_test

import (
	"github.com/gowizzard/compver/v3/statement"
	"reflect"
	"testing"
)

// TestCore is to test the Core function
// With table driven tests
func TestCore(t *testing.T) {

	tests := []struct {
		version  string
		block    string
		expected int
	}{
		{
			version:  "v1.12.4",
			block:    "major",
			expected: 1,
		},
		{
			version:  "4.7.0",
			block:    "minor",
			expected: 7,
		},
		{
			version:  "v3.0.11-alpha.2+testing-12345a",
			block:    "patch",
			expected: 11,
		},
	}

	for _, value := range tests {

		number, err := statement.Core(&value.version, &value.block)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(value.expected, number) {
			t.Fatalf("expected: \"%d\", got \"%d\"", value.expected, number)
		}

	}

}
