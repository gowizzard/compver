package version_core_test

import (
	"github.com/gowizzard/compver/version_core"
	"log"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {

	tests := []struct {
		version  string
		expected version_core.Core
	}{
		{
			version: "v1.2.4",
			expected: version_core.Core{
				Major: 1,
				Minor: 2,
				Patch: 4,
			},
		},
		{
			version: "v1.11.8-beta.1",
			expected: version_core.Core{
				Major: 1,
				Minor: 11,
				Patch: 8,
			},
		},
		{
			version: "v3.7.0-alpha.2+testing-12345a",
			expected: version_core.Core{
				Major: 3,
				Minor: 7,
				Patch: 0,
			},
		},
		{
			version: "3.34.7",
			expected: version_core.Core{
				Major: 3,
				Minor: 34,
				Patch: 7,
			},
		},
		{
			version: "0.12.0",
			expected: version_core.Core{
				Major: 0,
				Minor: 12,
				Patch: 0,
			},
		},
	}

	for _, value := range tests {

		core, err := version_core.Split(value.version)
		if err != nil {
			log.Fatalln(err)
		}

		numbers := []struct {
			Number1 int
			Number2 int
		}{
			{
				Number1: value.expected.Major,
				Number2: core.Major,
			},
			{
				Number1: value.expected.Minor,
				Number2: core.Minor,
			},
			{
				Number1: value.expected.Patch,
				Number2: core.Patch,
			},
		}

		for _, value := range numbers {
			if !reflect.DeepEqual(value.Number1, value.Number2) {
				t.Fatalf("expected: %d, got %d", value.Number1, value.Number2)
			}
		}

	}

}
