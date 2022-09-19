package command_line_test

import (
	"github.com/gowizzard/compver/v3/command_line"
	"testing"
)

// TestOutput is to test the output function.
func TestOutput(t *testing.T) {

	tests := []struct {
		attributes map[string]any
	}{
		{
			attributes: map[string]any{
				"major_number": 1,
				"minor_number": 10,
				"patch_number": 1,
			},
		},
		{
			attributes: map[string]any{
				"compare": "major update",
			},
		},
	}

	for _, value := range tests {
		command_line.Output(value.attributes)
	}

}
