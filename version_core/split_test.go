// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package version_core_test

import (
	"github.com/gowizzard/compver/v3/version_core"
	"log"
	"reflect"
	"testing"
)

// TestSplit is to check the version core split function.
// We simulate the version input and check the result.
func TestSplit(t *testing.T) {

	tests := []struct {
		version  string
		expected version_core.Core
	}{
		{
			version: "v1.2.4",
			expected: version_core.Core{
				Major:         1,
				Minor:         2,
				Patch:         4,
				Prerelease:    "",
				Buildmetadata: "",
			},
		},
		{
			version: "v1.11.8-beta.1",
			expected: version_core.Core{
				Major:         1,
				Minor:         11,
				Patch:         8,
				Prerelease:    "beta.1",
				Buildmetadata: "",
			},
		},
		{
			version: "v3.7.0-alpha.2+testing-12345a",
			expected: version_core.Core{
				Major:         3,
				Minor:         7,
				Patch:         0,
				Prerelease:    "alpha.2",
				Buildmetadata: "testing-12345a",
			},
		},
		{
			version: "3.34.7",
			expected: version_core.Core{
				Major:         3,
				Minor:         34,
				Patch:         7,
				Prerelease:    "",
				Buildmetadata: "",
			},
		},
		{
			version: "0.12.0+meta",
			expected: version_core.Core{
				Major:         0,
				Minor:         12,
				Patch:         0,
				Prerelease:    "",
				Buildmetadata: "meta",
			},
		},
	}

	for _, value := range tests {

		core, err := version_core.Split(value.version)
		if err != nil {
			log.Fatalln(err)
		}

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
				expected: value.expected.Prerelease,
				got:      core.Prerelease,
			},
			{
				expected: value.expected.Buildmetadata,
				got:      core.Buildmetadata,
			},
		}

		for _, value := range numbers {
			if !reflect.DeepEqual(value.got, value.expected) {
				t.Fatalf("expected: \"%d\", got \"%d\"", value.got, value.expected)
			}
		}

	}

}
