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
		error    bool
		expected string
	}{
		{
			name:     "MAJOR_UPDATE",
			version1: "2.0.0",
			version2: "1.3.15",
			error:    false,
			expected: "major update",
		},
		{
			name:     "MAJOR_DOWNGRADE",
			version1: "1.4.0",
			version2: "2.3.15",
			error:    false,
			expected: "major downgrade",
		},
		{
			name:     "MINOR_UPDATE",
			version1: "3.7.0-alpha.2+testing-12345a",
			version2: "3.6.35-beta.4",
			error:    false,
			expected: "minor update",
		},
		{
			name:     "MINOR_DOWNGRADE",
			version1: "5.64.4",
			version2: "5.70.10",
			error:    false,
			expected: "minor downgrade",
		},
		{
			name:     "PATCH_UPDATE",
			version1: "12.9.5",
			version2: "12.9.3",
			error:    false,
			expected: "patch update",
		},
		{
			name:     "PATCH_DOWNGRADE",
			version1: "4.0.10",
			version2: "4.0.15",
			error:    false,
			expected: "patch downgrade",
		},
		{
			name:     "VERSION1_PARSE_ERROR",
			version1: "v1.0.0",
			version2: "1.9.0",
			error:    true,
			expected: "",
		},
		{
			name:     "VERSION2_PARSE_ERROR",
			version1: "2.0.0",
			version2: "v2.0.1",
			error:    true,
			expected: "",
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			compare, err := statement.Compare(value.version1, value.version2)
			if err != nil && !value.error {
				t.Error(err)
			}

			if !reflect.DeepEqual(value.expected, compare) && !value.error {
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
