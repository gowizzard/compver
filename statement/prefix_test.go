// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package statement_test

import (
	"github.com/gowizzard/compver/v5/statement"
	"reflect"
	"testing"
)

// TestPrefix is to test the prefix statement function.
func TestPrefix(t *testing.T) {

	tests := []struct {
		name     string
		version  string
		prefix   string
		expected string
	}{
		{
			name:     "VERSION=v1.2.4",
			version:  "v1.2.4",
			prefix:   "v",
			expected: "1.2.4",
		},
		{
			name:     "VERSION=1.11.8-beta.1",
			version:  "1.11.8-beta.1",
			prefix:   "",
			expected: "1.11.8-beta.1",
		},
		{
			name:     "VERSION=v3.7.0-alpha.2+testing-12345a",
			version:  "v3.7.0-alpha.2+testing-12345a",
			prefix:   "v",
			expected: "3.7.0-alpha.2+testing-12345a",
		},
		{
			name:     "VERSION=version3.34.7",
			version:  "version3.34.7",
			prefix:   "version",
			expected: "3.34.7",
		},
		{
			name:     "VERSION=v0.12.0",
			version:  "v0.12.0",
			prefix:   "v",
			expected: "0.12.0",
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			result := statement.Prefix(value.version, value.prefix)

			if !reflect.DeepEqual(value.expected, result) {
				t.Errorf("expected: \"%s\", got \"%s\"", value.expected, result)
			}

		})

	}

}
