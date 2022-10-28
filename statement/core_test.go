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
		expected any
	}{
		{
			name:     "MAJOR=1",
			version:  "1.12.4",
			block:    "major",
			expected: 1,
		},
		{
			name:     "MINOR=7",
			version:  "4.7.0",
			block:    "minor",
			expected: 7,
		},
		{
			name:     "PATCH=11",
			version:  "3.0.11+testing-12345a",
			block:    "patch",
			expected: 11,
		},
		{
			name:     "PRERELEASE=alpha.2",
			version:  "2.0.0-alpha.2",
			block:    "prerelease",
			expected: "alpha.2",
		},
		{
			name:     "BUILDMETADATA=meta",
			version:  "1.34.0-beta.1+meta",
			block:    "buildmetadata",
			expected: "meta",
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			result, err := statement.Core(value.version, value.block)
			if err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(value.expected, result) {
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
