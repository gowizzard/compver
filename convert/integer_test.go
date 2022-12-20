// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package convert_test

import (
	"github.com/gowizzard/compver/v5/convert"
	"reflect"
	"testing"
)

// TestInteger is to test the convert integer function and
// convert a number it to an integer and check the result.
func TestInteger(t *testing.T) {

	tests := []struct {
		name     string
		number   string
		expected int
	}{
		{
			name:     "5",
			number:   "5",
			expected: 5,
		},
		{
			name:     "265",
			number:   "265",
			expected: 265,
		},
		{
			name:     "78",
			number:   "78",
			expected: 78,
		},
		{
			name:     "132",
			number:   "132",
			expected: 132,
		},
		{
			name:     "25",
			number:   "25",
			expected: 25,
		},
		{
			name:     "PARSE_ERROR",
			number:   "error",
			expected: 0,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			integer := convert.Integer(value.number)

			if !reflect.DeepEqual(value.expected, integer) {
				t.Errorf("expected: \"%d\", got \"%d\"", value.expected, integer)
			}

		})

	}

}

// BenchmarkInteger is to test the Integer function benchmark timing.
func BenchmarkInteger(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = convert.Integer("14")
	}

}

// FuzzInteger is to test the Integer function with fuzz testing.
func FuzzInteger(f *testing.F) {

	f.Add("42")
	f.Fuzz(func(t *testing.T, s string) {
		_ = convert.Integer(s)
	})

}
