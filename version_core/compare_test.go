// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package version_core_test

import (
	"github.com/gowizzard/compver/v5/version_core"
	"reflect"
	"testing"
)

// TestCompare is to check the version core compare function.
// We simulate the major, minor & patch values and check the result.
func TestCompare(t *testing.T) {

	tests := []struct {
		name     string
		version1 version_core.Core
		version2 version_core.Core
		expected string
	}{
		{
			name: "MAJOR_UPDATE",
			version1: version_core.Core{
				Major: 12,
				Minor: 0,
				Patch: 2,
			},
			version2: version_core.Core{
				Major: 11,
				Minor: 35,
				Patch: 9,
			},
			expected: "major update",
		},
		{
			name: "MAJOR_DOWNGRADE",
			version1: version_core.Core{
				Major: 3,
				Minor: 34,
				Patch: 2,
			},
			version2: version_core.Core{
				Major: 4,
				Minor: 0,
				Patch: 0,
			},
			expected: "major downgrade",
		},
		{
			name: "MINOR_UPDATE",
			version1: version_core.Core{
				Major: 1,
				Minor: 3,
				Patch: 0,
			},
			version2: version_core.Core{
				Major: 1,
				Minor: 2,
				Patch: 5,
			},
			expected: "minor update",
		},
		{
			name: "MINOR_DOWNGRADE",
			version1: version_core.Core{
				Major: 4,
				Minor: 3,
				Patch: 5,
			},
			version2: version_core.Core{
				Major: 4,
				Minor: 12,
				Patch: 0,
			},
			expected: "minor downgrade",
		},
		{
			name: "PATCH_UPDATE",
			version1: version_core.Core{
				Major: 1,
				Minor: 19,
				Patch: 8,
			},
			version2: version_core.Core{
				Major: 1,
				Minor: 19,
				Patch: 7,
			},
			expected: "patch update",
		},
		{
			name: "PATCH_DOWNGRADE",
			version1: version_core.Core{
				Major: 6,
				Minor: 9,
				Patch: 8,
			},
			version2: version_core.Core{
				Major: 6,
				Minor: 9,
				Patch: 11,
			},
			expected: "patch downgrade",
		},
		{
			name: "NO_CHANGES",
			version1: version_core.Core{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
			version2: version_core.Core{
				Major: 1,
				Minor: 0,
				Patch: 0,
			},
			expected: "no changes",
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {

			var blocks []version_core.Block
			blocks = append(
				blocks, version_core.Block{
					Name:    "major",
					Number1: value.version1.Major,
					Number2: value.version2.Major,
				}, version_core.Block{
					Name:    "minor",
					Number1: value.version1.Minor,
					Number2: value.version2.Minor,
				}, version_core.Block{
					Name:    "patch",
					Number1: value.version1.Patch,
					Number2: value.version2.Patch,
				},
			)

			compare := version_core.Compare(blocks)

			if !reflect.DeepEqual(value.expected, compare) {
				t.Errorf("expected: \"%s\", got \"%s\"", value.expected, compare)
			}

		})

	}

}

// BenchmarkCompare is to test the Compare function benchmark timing.
func BenchmarkCompare(b *testing.B) {

	blocks := []version_core.Block{
		{
			Name:    "major",
			Number1: 0,
			Number2: 2,
		},
		{
			Name:    "minor",
			Number1: 13,
			Number2: 0,
		},
		{
			Name:    "patch",
			Number1: 1,
			Number2: 4,
		},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = version_core.Compare(blocks)
	}

}
