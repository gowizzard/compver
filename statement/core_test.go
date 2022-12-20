// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement_test

import (
	"github.com/gowizzard/compver/v5/statement"
	"reflect"
	"testing"
)

// TestCore is to test the Core function with table driven tests.
func TestCore(t *testing.T) {

	tests := []struct {
		name     string
		version  string
		block    string
		error    bool
		expected any
	}{
		{
			name:     "MAJOR",
			version:  "1.12.4",
			block:    "major",
			error:    false,
			expected: 1,
		},
		{
			name:     "MINOR",
			version:  "4.7.0",
			block:    "minor",
			error:    false,
			expected: 7,
		},
		{
			name:     "PATCH",
			version:  "3.0.11+testing-12345a",
			block:    "patch",
			error:    false,
			expected: 11,
		},
		{
			name:     "PRERELEASE",
			version:  "2.0.0-alpha.2",
			block:    "prerelease",
			error:    false,
			expected: "alpha.2",
		},
		{
			name:     "BUILDMETADATA",
			version:  "1.34.0-beta.1+meta",
			block:    "buildmetadata",
			error:    false,
			expected: "meta",
		},
		{
			name:     "CORE_ERROR",
			version:  "v1.0.0",
			block:    "",
			error:    true,
			expected: nil,
		},
		{
			name:     "BLOCK_ERROR",
			version:  "1.2.3",
			block:    "error",
			error:    true,
			expected: nil,
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			result, err := statement.Core(value.version, value.block)
			if err != nil && !value.error {
				t.Error(err)
			}

			if !reflect.DeepEqual(value.expected, result) && !value.error {
				t.Errorf("expected: \"%d\", got \"%d\"", value.expected, result)
			}

		})

	}

}

// BenchmarkCore is to test the Core function benchmark timing.
func BenchmarkCore(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := statement.Core("3.5.0", "major")
		if err != nil {
			b.Error(err)
		}
	}

}

// FuzzCore is to test the Core function with fuzz testing.
func FuzzCore(f *testing.F) {

	f.Add("1.20.4", "patch")
	f.Fuzz(func(t *testing.T, s1, s2 string) {
		_, err := statement.Core(s1, s2)
		if err != nil {
			f.Error(err)
		}
	})

}
