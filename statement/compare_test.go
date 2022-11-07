// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement_test

import (
	"github.com/gowizzard/compver/v5/statement"
	"reflect"
	"testing"
)

// TestCompare is to test the Compare function.
func TestCompare(t *testing.T) {

	tests := []struct {
		name     string
		version1 string
		version2 string
		expected string
	}{
		{
			name:     "MAJOR=update",
			version1: "2.0.0",
			version2: "1.3.15",
			expected: "major update",
		},
		{
			name:     "MAJOR=downgrade",
			version1: "1.4.0",
			version2: "2.3.15",
			expected: "major downgrade",
		},
		{
			name:     "MINOR=update",
			version1: "3.7.0-alpha.2+testing-12345a",
			version2: "3.6.35-beta.4",
			expected: "minor update",
		},
		{
			name:     "MINOR=downgrade",
			version1: "5.64.4",
			version2: "5.70.10",
			expected: "minor downgrade",
		},
		{
			name:     "PATCH=update",
			version1: "12.9.5",
			version2: "12.9.3",
			expected: "patch update",
		},
		{
			name:     "PATCH=downgrade",
			version1: "4.0.10",
			version2: "4.0.15",
			expected: "patch downgrade",
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			compare, err := statement.Compare(value.version1, value.version2)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(value.expected, compare) {
				t.Errorf("expected: \"%s\", got \"%s\"", value.expected, compare)
			}

		})

	}

}

// BenchmarkCompare is to test the Compare function benchmark timing.
func BenchmarkCompare(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := statement.Compare("1.5.0", "2.4.4")
		if err != nil {
			b.Error(err)
		}
	}

}

// FuzzCompare is to test the Compare function with fuzz testing.
func FuzzCompare(f *testing.F) {

	f.Add("2.2.0", "2.3.0")
	f.Fuzz(func(t *testing.T, s1, s2 string) {
		_, err := statement.Compare(s1, s2)
		if err != nil {
			f.Fuzz(err)
		}
	})

}
