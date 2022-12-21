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
			name:     "WITH_PREFIX",
			version:  "v1.2.4",
			prefix:   "v",
			expected: "1.2.4",
		},
		{
			name:     "WITHOUT_PREFIX",
			version:  "1.11.8-beta.1",
			prefix:   "",
			expected: "1.11.8-beta.1",
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

// BenchmarkPrefix is to test the Prefix function benchmark timing.
func BenchmarkPrefix(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_ = statement.Prefix("v3.7.0-alpha.2+testing-12345a", "v")
	}

}
