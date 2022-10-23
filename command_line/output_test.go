// Copyright 2022 Jonas Kwiedor. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package command_line_test

import (
	"github.com/gowizzard/compver/v4/command_line"
	"testing"
)

// TestOutput is to test the output function.
func TestOutput(t *testing.T) {

	tests := []struct {
		name       string
		attributes map[string]any
	}{
		{
			name: "ATTRIBUTES=3",
			attributes: map[string]any{
				"major_number": 1,
				"minor_number": 10,
				"patch_number": 1,
			},
		},
		{
			name: "ATTRIBUTES=1",
			attributes: map[string]any{
				"compare": "major update",
			},
		},
	}

	for _, value := range tests {

		t.Run(value.name, func(t *testing.T) {
			command_line.Output(value.attributes)
		})

	}

}
