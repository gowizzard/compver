// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package version_core_test

import (
	"github.com/gowizzard/compver/v5/version_core"
	"log"
	"reflect"
	"testing"
)

// TestSplit is to check the version core split function.
// We simulate the version input and check the result.
func TestSplit(t *testing.T) {

	tests := []struct {
		name     string
		version  string
		error    bool
		expected version_core.Core
	}{
		{
			name:    "1.2.4",
			version: "1.2.4",
			error:   false,
			expected: version_core.Core{
				Major:         1,
				Minor:         2,
				Patch:         4,
				PreRelease:    "",
				BuildMetadata: "",
			},
		},
		{
			name:    "1.11.8-beta.1",
			version: "1.11.8-beta.1",
			error:   false,
			expected: version_core.Core{
				Major:         1,
				Minor:         11,
				Patch:         8,
				PreRelease:    "beta.1",
				BuildMetadata: "",
			},
		},
		{
			name:    "3.7.0-alpha.2+testing-12345a",
			version: "3.7.0-alpha.2+testing-12345a",
			error:   false,
			expected: version_core.Core{
				Major:         3,
				Minor:         7,
				Patch:         0,
				PreRelease:    "alpha.2",
				BuildMetadata: "testing-12345a",
			},
		},
		{
			name:    "PARSE_ERROR",
			version: "v1.0.0",
			error:   true,
			expected: version_core.Core{
				Major:         0,
				Minor:         0,
				Patch:         0,
				PreRelease:    "",
				BuildMetadata: "",
			},
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			core, err := version_core.Split(value.version)
			if err != nil && !value.error {
				log.Fatalln(err)
			}

			if !value.error {

				numbers := []struct {
					expected any
					got      any
				}{
					{
						expected: value.expected.Major,
						got:      core.Major,
					},
					{
						expected: value.expected.Minor,
						got:      core.Minor,
					},
					{
						expected: value.expected.Patch,
						got:      core.Patch,
					},
					{
						expected: value.expected.PreRelease,
						got:      core.PreRelease,
					},
					{
						expected: value.expected.BuildMetadata,
						got:      core.BuildMetadata,
					},
				}

				for _, value := range numbers {
					if !reflect.DeepEqual(value.got, value.expected) {
						t.Errorf("expected: \"%d\", got \"%d\"", value.got, value.expected)
					}
				}

			}

		})

	}

}

// BenchmarkSplit is to test the Split function benchmark timing.
func BenchmarkSplit(b *testing.B) {

	for i := 0; i < b.N; i++ {
		_, err := version_core.Split("3.7.0-alpha.2+testing-12345a")
		if err != nil {
			b.Error(err)
		}
	}

}

// FuzzSplit is to test the Split function with fuzz testing.
func FuzzSplit(f *testing.F) {

	f.Add("2.2.0+testing-67890b")
	f.Fuzz(func(t *testing.T, s string) {
		_, err := version_core.Split(s)
		if err != nil {
			f.Error(err)
		}
	})

}
