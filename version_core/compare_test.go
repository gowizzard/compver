package version_core_test

import (
	"github.com/gowizzard/compver/v3/version_core"
	"reflect"
	"testing"
)

// TestCompare is to check the version core compare function.
// We simulate the major, minor & patch values and check the result.
func TestCompare(t *testing.T) {

	tests := []struct {
		version1 version_core.Core
		version2 version_core.Core
		expected string
	}{
		{
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
	}

	for _, value := range tests {

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
			t.Fatalf("expected: \"%s\", got \"%s\"", value.expected, compare)
		}

	}

}
