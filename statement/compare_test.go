// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement_test

import (
	"github.com/gowizzard/compver/v3/statement"
	"reflect"
	"testing"
)

// TestCompare is to test the Compare function.
func TestCompare(t *testing.T) {

	tests := []struct {
		version1 string
		version2 string
		expected string
	}{
		{
			version1: "v2.0.0",
			version2: "v1.3.15",
			expected: "major update",
		},
		{
			version1: "v1.4.0",
			version2: "v2.3.15",
			expected: "major downgrade",
		},
		{
			version1: "v3.7.0-alpha.2+testing-12345a",
			version2: "v3.6.35-beta.4",
			expected: "minor update",
		},
		{
			version1: "5.64.4",
			version2: "5.70.10",
			expected: "minor downgrade",
		},
		{
			version1: "v12.9.5",
			version2: "v12.9.3",
			expected: "patch update",
		},
		{
			version1: "v4.0.10",
			version2: "v4.0.15",
			expected: "patch downgrade",
		},
	}

	for _, value := range tests {

		compare, err := statement.Compare(value.version1, value.version2)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(value.expected, compare) {
			t.Fatalf("expected: \"%s\", got \"%s\"", value.expected, compare)
		}

	}

}
