// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement_test

import (
	"github.com/gowizzard/compver/v4/statement"
	"reflect"
	"testing"
)

// TestCore is to test the Core function with table driven tests.
func TestCore(t *testing.T) {

	tests := []struct {
		version  string
		block    string
		expected any
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
			version:  "v3.0.11+testing-12345a",
			block:    "patch",
			expected: 11,
		},
		{
			version:  "v2.0.0-alpha.2",
			block:    "prerelease",
			expected: "alpha.2",
		},
		{
			version:  "v1.34.0-beta.1+meta",
			block:    "buildmetadata",
			expected: "meta",
		},
	}

	for _, value := range tests {

		result, err := statement.Core(value.version, value.block)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(value.expected, result) {
			t.Fatalf("expected: \"%d\", got \"%d\"", value.expected, result)
		}

	}

}
